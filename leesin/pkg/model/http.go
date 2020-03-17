package model

type Http struct {
	StatusCode int         `json:"statuscode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}
