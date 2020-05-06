package email

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

type builder struct {
	from		From
	to 			[]string
	subject 	string
	content 	string
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

func (builder builder) Build() Mail {
	return Mail{From: builder.from, To: builder.to, Body:builder.content, Subject:builder.subject}
}

type Build interface {
	Build() Mail
}

func NewBuilder() FromStep {
	return builder{from:From{}}
}