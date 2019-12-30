package mail

import (
	"fmt"
	"log"
)

type mock struct {
	mail 		Mail
	Fail 		bool
	count 		int
	retryMax   	int
}

func (mock *mock) Mail() Mail {
	return mock.mail
}

func (mock *mock) Send() error {
	log.Print("sending email..")
	if mock.count > mock.retryMax {
		mock.Fail = false
	}
	if mock.Fail {
		return fmt.Errorf("failure during sending mail")
	}
	log.Printf("from: %v", mock.mail.From)
	log.Printf("to: %v", mock.mail.To)
	log.Printf("subject: %v", mock.mail.Subject)
	log.Printf("body: %v", mock.mail.Body)
	return nil
}