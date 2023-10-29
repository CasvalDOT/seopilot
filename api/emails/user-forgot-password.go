package emails

import (
	"time"
)

type userForgotPasswordEmail struct {
	*email
}

type UserForgotPasswordTemplateData struct {
	Subject   string
	Email     string
	Url       string
	Name      string
	CreatedAt time.Time
}

func NewUserForgotPasswordEmail(data UserForgotPasswordTemplateData) Email {
	data.Subject = "Forgot password request"
	email := &email{
		to:       []string{data.Email},
		template: "templates/user-forgot-password.tmpl",
		subject:  data.Subject,
		data:     data,
	}

	return &userForgotPasswordEmail{
		email,
	}
}
