package mail

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/smtp"
)

const (
	HotmailServer = "smtp.office365.com"
	HotmailPort = "587"
	GMailServer = "smtp-relay.gmail.com"
	GMailPort = "587"
)

type Mail struct {
	From 		From 		// PIIIA.com <info@piiia.com>
	To 			[]string 	// Pepe Unlimited Oy <info@pepeunlimited.com>
	Subject 	string
	Body 		string
}

type From struct {
	email string
	name string
}

type Client interface {
	Mail() Mail
	Send() error
}

type client struct {
	auth smtp.Auth
	server string
	port string
	mail Mail
}

func (client client) Mail() Mail {
	return client.mail
}

func (client client) Send() error {
	log.Printf("sending using smtp %v", client.serverport())
	for _, to := range client.mail.To {
		log.Printf("to: %v", to)
		err := smtp.SendMail(client.serverport(), client.auth, client.mail.From.email, []string{to}, client.body(client.mail.From, to, client.mail.Subject, client.mail.Body))
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
	return []byte("From: " + fmt.Sprintf("%v <%v>", from.name, from.email) + "\r\n" +
		"To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"Content-Type: " + "text/html; charset=\"UTF-8\"" + "\r\n" +
		"\r\n" +
		msg)
}

type loginAuth struct {
	username, password string
}

func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte(a.username), nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("Unkown fromServer")
		}
	}
	return nil, nil
}

func NewSmtpClient(mail Mail, port string, server string, auth smtp.Auth) Client {
	return client{mail:mail, port:port, server:server, auth: auth}
}