package consumer

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	noebsCrypto "github.com/adonese/crypto"
	gateway "github.com/adonese/noebs/apigateway"
	"github.com/adonese/noebs/ebs_fields"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-redis/redis/v7"
	"github.com/golang-jwt/jwt"
	"github.com/pquerna/otp/totp"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type State struct {
	Db    *gorm.DB
	Redis *redis.Client
	Auth  Auther
}

type Auther interface {
	VerifyJWT(token string) (*gateway.TokenClaims, error)
	GenerateJWT(token string) (string, error)
}

//GenerateAPIKey An Admin-only endpoint that is used to generate api key for our clients
// the user must submit their email to generate a unique token per email.
// FIXME #59 #58 #61 api generation should be decoupled from apigateway package
func (s *State) GenerateAPIKey(c *gin.Context) {
	var m map[string]string
	if err := c.ShouldBindJSON(&m); err != nil {
		if _, ok := m["email"]; ok {
			k, _ := gateway.GenerateAPIKey()
			s.Redis.SAdd("apikeys", k)
			c.JSON(http.StatusOK, gin.H{"result": k})
			return
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"message": "missing_field"})
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error in email"})

	}

}

//ApiKeyMiddleware used to authenticate clients using X-Email and X-API-Key headers
//FIXME issue #58 #61
func (s *State) ApiKeyMiddleware(c *gin.Context) {
	email := c.GetHeader("X-Email")
	key := c.GetHeader("X-API-Key")
	if email == "" || key == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "unauthorized"})
		return
	}

	res, err := s.Redis.HGet("api_keys", email).Result()
	if err != redis.Nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "unauthorized"})
		return
	}
	if key == res {
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "unauthorized"})
		return
	}
}

//FIXME issue #58 #61
func (s *State) IpFilterMiddleware(c *gin.Context) {
	ip := c.ClientIP()
	if u := c.GetString("username"); u != "" {
		s.Redis.HIncrBy(u+":ips_count", ip, 1)
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "unauthorized_access"})
	}
}

func (s *State) LoginHandler(c *gin.Context) {

	var req ebs_fields.User
	if err := c.ShouldBindWith(&req, binding.JSON); err != nil {
		// The request is wrong
		log.Printf("The request is wrong. %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "code": "bad_request"})
		return
	}
	log.Printf("the processed request is: %v\n", req)
	if req.Mobile == "" {
		req.Mobile = req.Username
	}
	req.Username = req.Mobile // making username a mobile
	u := ebs_fields.User{}
	if notFound := s.Db.Where("username = ? or email = ? or mobile = ?", strings.ToLower(req.Username), strings.ToLower(req.Username), strings.ToLower(req.Username)).First(&u).Error; errors.Is(notFound, gorm.ErrRecordNotFound) {
		// service id is not found
		log.Printf("User with service_id %s is not found.", req.Username)
		c.JSON(http.StatusBadRequest, gin.H{"message": notFound.Error(), "code": "not_found"})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "wrong password entered", "code": "wrong_password"})
		return
	}

	token, err := s.Auth.GenerateJWT(u.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.Writer.Header().Set("Authorization", token)
	c.JSON(http.StatusOK, gin.H{"authorization": token, "user": u})
}

//SingleLoginHandler is used for one-time authentications. It checks a signed entered otp keys against
// the user's credentials (user's stored public key)
func (s *State) SingleLoginHandler(c *gin.Context) {

	var req gateway.Token
	c.ShouldBindWith(&req, binding.JSON)
	log.Printf("the processed request is: %v\n", req)

	u := ebs_fields.User{}
	if notFound := s.Db.Where("username = ? or email = ? or mobile = ?", strings.ToLower(req.Mobile), strings.ToLower(req.Mobile), strings.ToLower(req.Mobile)).First(&u).Error; errors.Is(notFound, gorm.ErrRecordNotFound) {
		log.Printf("User with service_id %s is not found.", req.Mobile)
		c.JSON(http.StatusBadRequest, gin.H{"message": notFound.Error(), "code": "not_found"})
		return
	}

	if _, encErr := noebsCrypto.VerifyWithHeaders(u.PublicKey, req.Signature, req.Message); encErr != nil {
		log.Printf("invalid signature in refresh: %v", encErr)
		c.JSON(http.StatusBadRequest, gin.H{"message": encErr.Error(), "code": "bad_request"})
		return
	}

	// Validate the otp using user's stored public key
	if totp.Validate(req.Message, u.EncodePublickey()) == false {
		c.JSON(http.StatusBadRequest, gin.H{"message": "wrong otp entered", "code": "wrong_otp"})
		return
	}
	token, err := s.Auth.GenerateJWT(u.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.Writer.Header().Set("Authorization", token)
	c.JSON(http.StatusOK, gin.H{"authorization": token, "user": u})
}

//RefreshHandler generates a new access token to the user using
// their signed public key.
// the user will sign their username with their private key, and noebs will verify
// the signature using the stored public key for the user
func (s *State) RefreshHandler(c *gin.Context) {
	var req gateway.Token
	if err := c.ShouldBindWith(&req, binding.JSON); err != nil {
		// The request is wrong
		log.Printf("The request is wrong. %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "code": "bad_request"})
		return
	}
	claims, err := s.Auth.VerifyJWT(req.JWT)
	if e, ok := err.(*jwt.ValidationError); ok {
		if e.Errors&jwt.ValidationErrorExpired != 0 {
			log.Info("refresh: auth username is: %s", claims.Username)
			user, _ := ebs_fields.NewUserByMobile(claims.Username, s.Db)
			// should verify signature here...
			if user.PublicKey == "" {
				log.Printf("user: %s has no registered pubkey", user.Mobile)
			}
			log.Printf("grabbed user is: %#v", user.Mobile)
			if _, encErr := noebsCrypto.VerifyWithHeaders(user.PublicKey, req.Signature, req.Message); encErr != nil {
				log.Printf("invalid signature in refresh: %v", encErr)
				c.JSON(http.StatusBadRequest, gin.H{"message": encErr.Error(), "code": "bad_request"})
				return
			}
			auth, _ := s.Auth.GenerateJWT(claims.Username)
			c.Writer.Header().Set("Authorization", auth)
			c.JSON(http.StatusOK, gin.H{"authorization": auth})

		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Malformed token", "code": "jwt_malformed"})
			return
		}
	} else if err == nil {
		// FIXME it is better to let the endpoint explicitly Get the claim off the user
		//  as we will assume the auth server will reside in a different domain!
		log.Printf("the username is: %s", claims.Username)

		auth, _ := s.Auth.GenerateJWT(claims.Username)
		c.Writer.Header().Set("Authorization", auth)
		c.JSON(http.StatusOK, gin.H{"authorization": auth})
	}
}

//CreateUser to register a new user to noebs
func (s *State) CreateUser(c *gin.Context) {
	u := ebs_fields.User{}
	if s.Db == nil {
		panic("wtf")
	}
	if err := c.ShouldBindBodyWith(&u, binding.JSON); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	// default username to mobile, in case username was not provided
	if u.Username == "" {
		u.Username = u.Mobile
	}
	// validate u.Password to include at least one capital letter, one symbol and one number
	// and that it is at least 8 characters long
	if !validatePassword(u.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Password must be at least 8 characters long, and must include at least one capital letter, one symbol and one number", "code": "password_invalid"})
		return
	}

	// make sure that the user doesn't exist in the database

	if err := u.HashPassword(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	// make the user capital - small
	u.SanitizeName()
	if err := s.Db.Create(&u).Error; err != nil {
		// unable to create this user; see possible reasons
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error(), "code": "duplicate_username"})
		return
	}

	s.Redis.Set(u.Mobile, u.Username, 0)
	ip := c.ClientIP()
	s.Redis.HSet(u.Username+":ips_count", "first_ip", ip)

	c.JSON(http.StatusCreated, gin.H{"ok": "object was successfully created", "details": u})
}

func (s *State) GenerateSignInCode(c *gin.Context) {

	var req gateway.Token
	c.ShouldBindWith(&req, binding.JSON)
	// default username to mobile, in case username was not provided
	if req.Mobile == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Mobile number was not sent", "code": "bad_request"})
		return
	}
	user, _ := ebs_fields.NewUserByMobile(req.Mobile, s.Db)
	if user.PublicKey == "" {
		// user has no public key
		c.JSON(http.StatusBadRequest, gin.H{"message": "Mobile number was not sent", "code": "bad_request"})
		return
	}
	key, err := generateOtp(user.EncodePublickey())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Mobile number was not sent", "code": "bad_request"})
		return
	}
	// this function doesn't have to be blocking.
	go sendSMS(SMS{Mobile: req.Mobile, Message: fmt.Sprintf("Your one-time access code is: %s. DON'T share it with anyone.", key)})
	c.JSON(http.StatusCreated, gin.H{"status": "ok", "message": "Password reset link has been sent to your mobile number. Use the info to login in to your account."})
}

//FIXME issue #61
func GetServiceID(c *gin.Context) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
	}

	id := c.Query("id")
	if id == "" {
		c.AbortWithStatusJSON(400, gin.H{"message": errNoServiceID.Error()})
	}

	fmt.Printf("the qparam is: %v\n", id)
	var res Service

	if err := db.Where("username = ?", id).First(&res).Error; err != nil {
		c.AbortWithStatusJSON(404, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ok": "this object is available"})
}

//APIAuth API-Key middleware. Currently is used by consumer services
//FIXME issue #61
func (s *State) APIAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if key := c.GetHeader("api-key"); key != "" {
			if !isMember("apikeys", key, s.Redis) {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": "wrong_api_key",
					"message": "visit https://soluspay.net/contact for a key"})
				return
			}
		}
		c.Next()
	}

}
