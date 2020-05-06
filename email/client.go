package email

import (
	"fmt"
	"log"
	"net"
	"net/smtp"
)

type EmailClient interface {
	Send(mail Mail)  error
	Provider() string
}

type client struct {
	auth smtp.Auth
	server string
	port string
}

func (client client) Send(mail Mail) error {
	log.Printf("sending using smtp %v", client.serverport())
	for _, to := range mail.To {
		log.Printf("to: %v", to)
		err := smtp.SendMail(client.serverport(), client.auth, mail.From.Email, []string{to}, client.body(mail.From, to, mail.Subject, mail.Body))
		if err != nil {
			return err
		}
	}
	log.Print("..finished!")
	return nil
}

func (client client) serverport() string {
	return net.JoinHostPort(client.server, client.port)
}

func (client) body(from From, to string, subject string, msg string) []byte {
	return []byte("From: " + fmt.Sprintf("%v <%v>", from.Name, from.Email) + "\r\n" +
		"To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"Content-Type: " + "text/html; charset=\"UTF-8\"" + "\r\n" +
		"\r\n" +
		msg)
}