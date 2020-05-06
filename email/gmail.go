package email

import "net/smtp"

const (
	GMailServer = "smtp-relay.gmail.com"
	GMailPort = "587"
)

func (g gmail) Send(mail Mail) error {
	return g.client.Send(mail)
}

func (g gmail) Provider() string {
	return "GMail"
}

type gmail struct {
	client *client
}

func NewGMailClient(username string, password string) EmailClient {
	return gmail{client: &client{
		auth:   smtp.PlainAuth("", username, password, GMailServer),
		server: GMailServer,
		port:   GMailPort,
	}}
}