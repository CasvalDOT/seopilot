package events

import (
	"api/emails"
)

type onForgotPassword struct {
	*event
}

func (e *onForgotPassword) Listen() {
	e.listen(func(event any) {
		data := event.(emails.UserForgotPasswordTemplateData)
		email := emails.NewUserForgotPasswordEmail(data)
		email.Send()
	})
}

func NewOnForgotPassword() Event {
	l := onForgotPassword{
		event: &event{
			ch: make(chan any),
		},
	}

	go l.Listen()

	return &l
}
