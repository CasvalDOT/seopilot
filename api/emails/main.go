package emails

import (
	"embed"
	"fmt"
	"html/template"
	"strings"
	"time"

	"gopkg.in/gomail.v2"
)

//go:embed templates/*
var emailTemplates embed.FS

type Email interface {
	Send() error
	AddReceiver(string)
	AddData(any)
}

type email struct {
	template string
	subject  string
	to       []string
	body     string
	data     any
}

type EmailPayload struct {
	Data    any
	AppName string
	Year    int
}

func (e *email) compileBody(emailTemplate string, data any) (string, error) {
	body := template.Must(template.ParseFS(emailTemplates, emailTemplate, "templates/base.tmpl"))
	out := &strings.Builder{}
	err := body.ExecuteTemplate(out, "base", EmailPayload{
		Data:    data,
		AppName: "Seopilot",
		Year:    time.Now().Year(),
	})
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	e.body = out.String()

	return e.body, nil
}

func (e *email) AddReceiver(email string) {
	e.to = append(e.to, email)
}

func (e *email) Send() error {
	body, err := e.compileBody(e.template, e.data)
	if err != nil {
		return err
	}

	for _, to := range e.to {
		err = e.send(to, e.subject, body)
		if err != nil {
			return err
		}
	}

	e.to = []string{}

	return nil
}

func (e *email) AddData(data any) {
	e.data = data
}

func (e *email) send(to string, subject string, message string) error {
	m := gomail.NewMessage()
	m.SetHeader("Subject", subject)
	m.SetHeader("From", "l5Z0U@example.com")
	m.SetHeader("To", to)
	m.SetBody("text/html", message)

	err := gomail.NewDialer("mailpit", 1025, "l5Z0U@example.com", "password").DialAndSend(m)
	if err != nil {
		return err
	}

	return nil
}
