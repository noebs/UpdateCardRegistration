// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-03-26 07:55:28.2925086 +0200 CAT m=+2.172010301

package docs

import (
	"bytes"

	"github.com/alecthomas/template"
)

var doc = `{
    "info": {
        "contact": {},
        "license": {}
    },
    "paths": {
        "/billInquiry": {
            "post": {
                "description": "get accounts",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get all transactions made by a specific terminal ID",
                "parameters": [
                    {
                        "description": "Bill Inquiry Request Fields",
                        "name": "billInquiry",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/ebs_fields.BillInquiryFields"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/main.SuccessfulResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "/billPayment": {
            "post": {
                "description": "get accounts",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get all transactions made by a specific terminal ID",
                "parameters": [
                    {
                        "description": "Bill Payment Request Fields",
                        "name": "billPayment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/ebs_fields.BillPaymentFields"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/main.SuccessfulResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "/cardTransfer": {
            "post": {
                "description": "get accounts",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get all transactions made by a specific terminal ID",
                "parameters": [
                    {
                        "description": "Card Transfer Request Fields",
                        "name": "cardTransfer",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/ebs_fields.CardTransferFields"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/main.SuccessfulResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "/cashIn": {
            "post": {
                "description": "get accounts",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get all transactions made by a specific terminal ID",
                "parameters": [
                    {
                        "description": "Cash In Request Fields",
                        "name": "cashOut",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/ebs_fields.CashInFields"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/main.SuccessfulResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "/cashOut": {
            "post": {
                "description": "get accounts",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get all transactions made by a specific terminal ID",
                "parameters": [
                    {
                        "description": "Cash Out Request Fields",
                        "name": "cashOut",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/ebs_fields.CashOutFields"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/main.SuccessfulResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "/changePin": {
            "post": {
                "description": "get accounts",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get all transactions made by a specific terminal ID",
                "parameters": [
                    {
                        "description": "Change PIN Request Fields",
                        "name": "changePIN",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/ebs_fields.ChangePINFields"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/main.SuccessfulResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "/dashboard/get_tid": {
            "get": {
                "description": "get accounts",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get all transactions made by a specific terminal ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "search list transactions by terminal ID",
                        "name": "id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "/miniStatement": {
            "post": {
                "description": "get accounts",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get all transactions made by a specific terminal ID",
                "parameters": [
                    {
                        "description": "Mini Statement Request Fields",
                        "name": "miniStatement",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/ebs_fields.MiniStatementFields"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/main.SuccessfulResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "/purchase": {
            "post": {
                "description": "get accounts",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get all transactions made by a specific terminal ID",
                "parameters": [
                    {
                        "description": "Purchase Request Fields",
                        "name": "purchase",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/ebs_fields.PurchaseFields"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/main.SuccessfulResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "/workingKey": {
            "post": {
                "description": "get accounts",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get all transactions made by a specific terminal ID",
                "parameters": [
                    {
                        "description": "Working Key Request Fields",
                        "name": "workingKey",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/ebs_fields.WorkingKeyFields"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/main.SuccessfulResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "ebs_fields.BillInquiryFields": {
            "type": "object",
            "required": [
                "PAN",
                "PIN",
                "clientId",
                "expDate",
                "payeeId",
                "personalPaymentInfo",
                "systemTraceAuditNumber",
                "terminalId",
                "tranAmount",
                "tranDateTime"
            ],
            "properties": {
                "PAN": {
                    "type": "string"
                },
                "PIN": {
                    "type": "string"
                },
                "clientId": {
                    "type": "string"
                },
                "expDate": {
                    "type": "string"
                },
                "payeeId": {
                    "type": "string"
                },
                "personalPaymentInfo": {
                    "type": "string"
                },
                "systemTraceAuditNumber": {
                    "type": "integer"
                },
                "terminalId": {
                    "type": "string"
                },
                "tranAmount": {
                    "type": "number"
                },
                "tranCurrencyCode": {
                    "type": "string"
                },
                "tranDateTime": {
                    "type": "string"
                }
            }
        },
        "ebs_fields.BillPaymentFields": {
            "type": "object",
            "required": [
                "PAN",
                "PIN",
                "clientId",
                "expDate",
                "payeeId",
                "personalPaymentInfo",
                "systemTraceAuditNumber",
                "terminalId",
                "tranAmount",
                "tranDateTime"
            ],
            "properties": {
                "PAN": {
                    "type": "string"
                },
                "PIN": {
                    "type": "string"
                },
                "clientId": {
                    "type": "string"
                },
                "expDate": {
                    "type": "string"
                },
                "payeeId": {
                    "type": "string"
                },
                "personalPaymentInfo": {
                    "type": "string"
                },
                "systemTraceAuditNumber": {
                    "type": "integer"
                },
                "terminalId": {
                    "type": "string"
                },
                "tranAmount": {
                    "type": "number"
                },
                "tranCurrencyCode": {
                    "type": "string"
                },
                "tranDateTime": {
                    "type": "string"
                }
            }
        },
        "ebs_fields.CardTransferFields": {
            "type": "object",
            "required": [
                "PAN",
                "PIN",
                "clientId",
                "expDate",
                "systemTraceAuditNumber",
                "terminalId",
                "toCard",
                "tranAmount",
                "tranDateTime"
            ],
            "properties": {
                "PAN": {
                    "type": "string"
                },
                "PIN": {
                    "type": "string"
                },
                "clientId": {
                    "type": "string"
                },
                "expDate": {
                    "type": "string"
                },
                "systemTraceAuditNumber": {
                    "type": "integer"
                },
                "terminalId": {
                    "type": "string"
                },
                "toCard": {
                    "type": "string"
                },
                "tranAmount": {
                    "type": "number"
                },
                "tranCurrencyCode": {
                    "type": "string"
                },
                "tranDateTime": {
                    "type": "string"
                }
            }
        },
        "ebs_fields.CashInFields": {
            "type": "object"
        },
        "ebs_fields.CashOutFields": {
            "type": "object"
        },
        "ebs_fields.ChangePINFields": {
            "type": "object",
            "required": [
                "PAN",
                "PIN",
                "clientId",
                "expDate",
                "newPin",
                "systemTraceAuditNumber",
                "terminalId",
                "tranDateTime"
            ],
            "properties": {
                "PAN": {
                    "type": "string"
                },
                "PIN": {
                    "type": "string"
                },
                "clientId": {
                    "type": "string"
                },
                "expDate": {
                    "type": "string"
                },
                "newPin": {
                    "type": "string"
                },
                "systemTraceAuditNumber": {
                    "type": "integer"
                },
                "terminalId": {
                    "type": "string"
                },
                "tranDateTime": {
                    "type": "string"
                }
            }
        },
        "ebs_fields.GenericEBSResponseFields": {
            "type": "object",
            "properties": {
                "DisputeRRN": {
                    "type": "string"
                },
                "PAN": {
                    "type": "string"
                },
                "additionalAmount": {
                    "type": "number"
                },
                "additionalData": {
                    "type": "string"
                },
                "approvalCode": {
                    "type": "integer"
                },
                "clientId": {
                    "type": "string"
                },
                "ebsserviceName": {
                    "type": "string"
                },
                "fromAccount": {
                    "type": "string"
                },
                "fromCard": {
                    "type": "string"
                },
                "miniStatementRecords": {
                    "type": "string"
                },
                "otp": {
                    "type": "string"
                },
                "otpId": {
                    "type": "string"
                },
                "phoneNumber": {
                    "type": "string"
                },
                "referenceNumber": {
                    "type": "integer"
                },
                "responseCode": {
                    "type": "integer"
                },
                "responseMessage": {
                    "type": "string"
                },
                "responseStatus": {
                    "type": "string"
                },
                "serviceId": {
                    "type": "string"
                },
                "systemTraceAuditNumber": {
                    "type": "integer"
                },
                "terminalId": {
                    "type": "string"
                },
                "toAccount": {
                    "type": "string"
                },
                "toCard": {
                    "type": "string"
                },
                "tranAmount": {
                    "type": "number"
                },
                "tranCurrencyCode": {
                    "type": "string"
                },
                "tranDateTime": {
                    "type": "string"
                },
                "tranFee": {
                    "type": "number"
                },
                "voucherNumber": {
                    "type": "integer"
                },
                "workingKey": {
                    "type": "string"
                }
            }
        },
        "ebs_fields.MiniStatementFields": {
            "type": "object",
            "required": [
                "PAN",
                "PIN",
                "clientId",
                "expDate",
                "systemTraceAuditNumber",
                "terminalId",
                "tranDateTime"
            ],
            "properties": {
                "PAN": {
                    "type": "string"
                },
                "PIN": {
                    "type": "string"
                },
                "clientId": {
                    "type": "string"
                },
                "expDate": {
                    "type": "string"
                },
                "systemTraceAuditNumber": {
                    "type": "integer"
                },
                "terminalId": {
                    "type": "string"
                },
                "tranDateTime": {
                    "type": "string"
                }
            }
        },
        "ebs_fields.PurchaseFields": {
            "type": "object",
            "required": [
                "PAN",
                "PIN",
                "clientId",
                "expDate",
                "systemTraceAuditNumber",
                "terminalId",
                "tranAmount",
                "tranDateTime"
            ],
            "properties": {
                "PAN": {
                    "type": "string"
                },
                "PIN": {
                    "type": "string"
                },
                "clientId": {
                    "type": "string"
                },
                "expDate": {
                    "type": "string"
                },
                "systemTraceAuditNumber": {
                    "type": "integer"
                },
                "terminalId": {
                    "type": "string"
                },
                "tranAmount": {
                    "type": "number"
                },
                "tranCurrencyCode": {
                    "type": "string"
                },
                "tranDateTime": {
                    "type": "string"
                }
            }
        },
        "ebs_fields.WorkingKeyFields": {
            "type": "object",
            "required": [
                "clientId",
                "systemTraceAuditNumber",
                "terminalId",
                "tranDateTime"
            ],
            "properties": {
                "clientId": {
                    "type": "string"
                },
                "systemTraceAuditNumber": {
                    "type": "integer"
                },
                "terminalId": {
                    "type": "string"
                },
                "tranDateTime": {
                    "type": "string"
                }
            }
        },
        "main.SuccessfulResponse": {
            "type": "object",
            "properties": {
                "ebs_response": {
                    "type": "object",
                    "$ref": "#/definitions/ebs_fields.GenericEBSResponseFields"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo swaggerInfo

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}
