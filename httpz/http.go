package httpz

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func WriteOk(resp http.ResponseWriter, data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		internal(err, resp)
	}
	resp.Header().Add("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)
	resp.Write(bytes)
}

func WriteError(resp http.ResponseWriter, err error) {
	error, ok := err.(Error)
	if !ok {
		internal(errors.New("given error is not right interface!"), resp)
	} else {
		resp.WriteHeader(error.Code())
		resp.Header().Add("Content-Type", "application/json")
		bytes, err := json.Marshal(error)
		if err != nil {
			internal(errors.New("can't marshall err: !"+err.Error()), resp)
		} else {
			resp.Write(bytes)
		}
	}
}

func internal(err error, resp http.ResponseWriter) {
	resp.WriteHeader(http.StatusInternalServerError)
	resp.Header().Add("Content-Type", "application/json")
	bytes, err := json.Marshal(microError{
		Msg:        err.Error(),
		StatusCode: http.StatusInternalServerError,
	})
	if err != nil {
		log.Panic(err)
	}
	resp.Write(bytes)
}

type Error interface {
	Message()   string
	Code() 		int
	Error()		string
}

type microError struct {
	Msg 		string  `json:"message"`
	StatusCode 	int 	`json:"status_code"`
}

func (m microError) Message() string {
	return m.Msg
}

func (m microError) Code() int {
	return m.StatusCode
}

func (m microError) Error() string {
	return "microservice-kit-error: "+m.Msg
}

func NewError(err error, statusCode int) error {
	return microError{
		Msg:        err.Error(),
		StatusCode: statusCode,
	}
}

func NewMsgError(error string, statusCode int) error {
	return microError{
		Msg:        error,
		StatusCode: statusCode,
	}
}