package email

import (
	"errors"
	"log"
)

type mock struct {
	failureCount 		int
}

func (m *mock) Send(mail Mail) error {
	log.Print("sending email..")
	if m.failureCount > 0 {
		m.failureCount--
		return errors.New("email-client: error during sending")
	}
	log.Printf("from: %v", mail.From)
	log.Printf("to: %v", mail.To)
	log.Printf("subject: %v", mail.Subject)
	log.Printf("body: %v", mail.Body)
	return nil
}

func (m mock) Provider() string {
	return "Mock"
}

func NewEmailMock(failureCount int) EmailClient {
	m := &mock{failureCount: failureCount}
	return m
}