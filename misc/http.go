package misc

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func WriteOK() {

}

func WriteError(resp http.ResponseWriter, err error) {
	error, ok := err.(Error)
	if !ok {
		internal(errors.New("is not Error-interface!"), resp)
	} else {
		resp.WriteHeader(error.Code())
		resp.Header().Add("Content-Type", "application/json")
		em, err := json.Marshal(error)
		if err != nil {
			internal(errors.New("can't marshall err: !"+err.Error()), resp)
		} else {
			resp.Write(em)
		}
	}
}

func internal(err error, response http.ResponseWriter) {
	response.WriteHeader(500)
	response.Header().Add("Content-Type", "application/json")
	em, err := json.Marshal(MicroError{
		Msg:        err.Error(),
		StatusCode: 500,
	})
	if err != nil {
		log.Panic(err)
	}
	response.Write(em)
}

type Error interface {
	Message()   string
	Code() 		int
	Error()		string
}

type MicroError struct {
	Msg 		string  `json:"message"`
	StatusCode 	int 	`json:"status_code"`
}

func (m MicroError) Message() string {
	return m.Msg
}

func (m MicroError) Code() int {
	return m.StatusCode
}

func (m MicroError) Error() string {
	return "microservice-kit-error: "+m.Msg
}