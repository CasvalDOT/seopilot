package emails

import "api/models"

type userActivateAccountEmail struct {
	*email
}

type UserActivateAccountTemplateData struct {
	Subject string
	User    models.User
	Url     string
}

func NewUserActivateAccountEmail(data UserActivateAccountTemplateData) Email {
	data.Subject = "Activate your account"

	email := &email{
		to:       []string{data.User.Email},
		template: "templates/user-activate-account.tmpl",
		subject:  data.Subject,
		data:     data,
	}

	return &userActivateAccountEmail{
		email,
	}
}
