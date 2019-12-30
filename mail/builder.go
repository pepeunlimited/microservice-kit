package mail

import (
	"log"
	"net/smtp"
)

type Provider string

const (
	GMail Provider = "GMail"
	Hotmail        = "Hotmail"
	Mock           = "Mock"
)

type builder struct {
	from		From
	to 			[]string
	subject 	string
	content 	string
	password 	string
	username	string
}

type ContentStep interface {
	Content(content string) Build
}

func (builder builder) Content(content string) Build {
	builder.content = content
	return builder
}

func (builder builder) Subject(subject string) ContentStep {
	builder.subject = subject
	return builder
}

type SubjectStep interface {
	Subject(subject string) ContentStep
}

func (builder builder) From(email string, name string) ToStep {
	builder.from.email = email
	builder.from.name = name
	return builder
}

type FromStep interface {
	From(email string, name string) ToStep
}

func (builder builder) To(address []string) SubjectStep {
	builder.to = address
	return builder
}

type ToStep interface {
	To(address []string) SubjectStep
}

func (builder builder) Build(provider Provider) Client {
	mail := Mail{From: builder.from, To: builder.to, Body:builder.content, Subject:builder.subject}
	var client Client
	switch provider {
	case GMail:
		client = NewSmtpClient(mail, GMailPort, GMailServer, smtp.PlainAuth("", builder.username, builder.password, GMailServer))
	case Hotmail:
		client = NewSmtpClient(mail, HotmailPort, HotmailServer, LoginAuth(builder.username, builder.password))
	case Mock:
		client = &mock{mail: mail, Fail: false}
	default:
		log.Panic("not supported smtp client")
	}
	return client
}

type Build interface {
	Build(provider Provider) Client
}

func NewBuilder(username string, password string) FromStep {
	return builder{username:username, password:password, from:From{}}
}