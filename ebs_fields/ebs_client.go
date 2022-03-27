package ebs_fields

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

//EBSHttpClient the client to interact with EBS
func EBSHttpClient(url string, req []byte) (int, EBSParserFields, error) {

	verifyTLS := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	ebsClient := http.Client{
		Timeout:   30 * time.Second,
		Transport: verifyTLS,
	}

	log.Printf("our request to EBS: %v", string(req))
	reqBuffer := bytes.NewBuffer(req)

	var ebsGenericResponse EBSParserFields

	reqHandler, err := http.NewRequest(http.MethodPost, url, reqBuffer)

	if err != nil {
		fmt.Println(err.Error())
		log.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Error in establishing connection to the host")
		return 500, ebsGenericResponse, err
	}
	reqHandler.Header.Set("Content-Type", "application/json")

	ebsResponse, err := ebsClient.Do(reqHandler)
	if err != nil {
		log.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Error in establishing connection to the host")
		return http.StatusGatewayTimeout, ebsGenericResponse, EbsGatewayConnectivityErr
	}

	defer ebsResponse.Body.Close()

	responseBody, err := ioutil.ReadAll(ebsResponse.Body)
	if err != nil {
		log.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Error reading ebs response")
		return http.StatusInternalServerError, ebsGenericResponse, EbsGatewayConnectivityErr
	}

	log.Printf("ebs_raw: %s", string(responseBody))
	// check Content-type is application json, if not, panic!
	// check if content type includes application/json
	if !strings.Contains(ebsResponse.Header.Get("Content-Type"), "application/json") {
		log.WithFields(logrus.Fields{
			"error":   "wrong content type parsed",
			"details": ebsResponse.Header.Get("Content-Type"),
		}).Error("ebs response content type is not application/json")
		return http.StatusInternalServerError, ebsGenericResponse, ContentTypeErr
	}

	var tmpRes IPINResponse

	if err := json.Unmarshal(responseBody, &ebsGenericResponse); err == nil {
		// there's no problem in Unmarshalling
		if ebsGenericResponse.ResponseCode == 0 || strings.Contains(ebsGenericResponse.ResponseMessage, "Success") {
			return http.StatusOK, ebsGenericResponse, nil
		} else {
			// if ebsGenericResponse.ResponseCode == 51 {
			// 	// this is a special case just to mock ebs
			// 	data := tmpRes.newResponse()
			// 	data.GenericEBSResponseFields.PAN = "7222061020105617937"
			// 	data.GenericEBSResponseFields.ExpDate = "2308"
			// 	data.GenericEBSResponseFields.FinancialInstitutionID = "ESSD"
			// 	data.GenericEBSResponseFields.EntityID = "249912141679"
			// 	data.GenericEBSResponseFields.EntityType = "Phone No"
			// 	data.GenericEBSResponseFields.PhoneNumber = "0912141679"
			// 	data.GenericEBSResponseFields.CreationDate = "0714"
			// 	data.GenericEBSResponseFields.EntityGroup = "0"
			// 	data.GenericEBSResponseFields.PanCategory = "Standard"
			// 	return http.StatusOK, data, nil
			// } else
			// the error here should be nil!
			// we don't actually have any errors!
			return http.StatusBadGateway, ebsGenericResponse, errors.New(ebsGenericResponse.ResponseMessage)
		}

	} else {
		// there is an error in handling the incoming EBS's ebsResponse
		// log the err here please
		log.WithFields(logrus.Fields{
			"error":        err.Error(),
			"all_response": string(responseBody),
			"ebs_fields":   ebsGenericResponse,
		}).Info("ebs response transaction")
		if strings.Contains(err.Error(), " EBSParserFields.tranDateTime of type string") { // fuck me
			// we have TWO cases here:
			// fuck ebs
			json.Unmarshal(responseBody, &tmpRes)
			if tmpRes.ResponseCode == 0 || strings.Contains(tmpRes.ResponseMessage, "Success") {
				return http.StatusOK, tmpRes.newResponse(), nil
			} else {
				return http.StatusBadGateway, ebsGenericResponse, errors.New(ebsGenericResponse.ResponseMessage)
			}
		}

		return http.StatusInternalServerError, ebsGenericResponse, err
	}

}

type IPINResponse struct {
	UUID            string `json:"UUID"`
	TranDateTime    int    `json:"tranDateTime"`
	ResponseMessage string `json:"responseMessage"`
	ResponseStatus  string `json:"responseStatus"`
	PubKeyValue     string `json:"pubKeyValue"`
	ResponseCode    int64  `json:"responseCode"`
	Pan             string `json:"pan"`
	ExpDate         string `json:"expDate"`
	Username        string `json:"userName"`
}

//newResponse the
func (i IPINResponse) newResponse() EBSParserFields {
	var res GenericEBSResponseFields
	res.ResponseCode = int(i.ResponseCode)
	res.ResponseMessage = i.ResponseMessage
	res.PubKeyValue = i.PubKeyValue
	res.TranDateTime = strconv.Itoa(i.TranDateTime)
	res.UUID = i.UUID
	res.PAN = i.Pan
	res.ExpDate = i.ExpDate
	return EBSParserFields{GenericEBSResponseFields: res}
}
