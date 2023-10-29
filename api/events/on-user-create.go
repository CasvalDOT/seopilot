package events

import (
	"api/emails"
)

type onUserCreate struct {
	*event
}

func (e *onUserCreate) Listen() {
	e.listen(func(event any) {
		data := event.(emails.UserActivateAccountTemplateData)
		email := emails.NewUserActivateAccountEmail(data)
		email.Send()
	})
}

func NewOnUserCreate() Event {
	l := onUserCreate{
		event: &event{
			ch: make(chan any),
		},
	}

	go l.Listen()

	return &l
}
