package email

import (
	"errors"
	"net/smtp"
)

const (
	HotmailServer = "smtp.office365.com"
	HotmailPort   = "587"
)

type hotmail struct {
	client *client
}

func (h hotmail) Send(mail Mail) error {
	return h.client.Send(mail)
}

func (h hotmail) Provider() string {
	return "Hotmail"
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
			return nil, errors.New("unknown error during sign-in")
		}
	}
	return nil, nil
}

func NewHotmailClient(username string, password string) EmailClient {
	return hotmail{client: &client{
		auth:   LoginAuth(username, password),
		server: HotmailServer,
		port:   HotmailPort,
	}}
}