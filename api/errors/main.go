package errors

import "net/http"

type ApiError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

const (
	ErrorKeyAuthInvalidAuthHeader        = "auth.invalid_auth_header"
	ErrorKeyAuthEmailOrPasswordIncorrent = "auth.email_or_password_incorrent"
	ErrorKeyAuthPasswordNotMatch         = "auth.password_not_match"
	ErrorKeyAuthInvalidToken             = "auth.invalid_token"
	ErrorKeyAuthUnauthorized             = "auth.unauthorized"
)

const (
	ErrorKeyUserMailAlreadyExists = "users.mail_already_exists"
	ErrorKeyUserNotFound          = "users.not_found"
	ErrorKeyUserCannotCreate      = "users.cannot_create"
	ErrorKeyUserCannotUpdate      = "users.cannot_update"
	ErrorKeyUserCannotDelete      = "users.cannot_delete"
)

const (
	ErrorKeySiteURLAlreadyExists           = "sites.url_already_exists"
	ErrorKeySiteNotFound                   = "sites.not_found"
	ErrorKeySiteCannotCreate               = "sites.cannot_create"
	ErrorKeySiteCannotUpdate               = "sites.cannot_update"
	ErrorKeySiteCannotDelete               = "sites.cannot_delete"
	ErrorKeySiteCannotDeleteForWrongSource = "sites.cannot_delete_for_wrong_source"
)

const (
	ErrorKeyAlertNotFound     = "alert.not_found"
	ErrorKeyAlertCannotCreate = "alert.cannot_create"
	ErrorKeyAlertCannotUpdate = "alert.cannot_update"
	ErrorKeyAlertCannotDelete = "alert.cannot_delete"
)

var errorMap = map[string]ApiError{
	"generic": {
		Message: "Something went wrong",
		Status:  http.StatusInternalServerError,
	},
	ErrorKeyUserNotFound: {
		Message: "Users not found",
		Status:  http.StatusNotFound,
	},
	ErrorKeyUserCannotCreate: {
		Message: "Cannot create user",
		Status:  http.StatusInternalServerError,
	},
	ErrorKeyUserCannotUpdate: {
		Message: "Cannot update user",
		Status:  http.StatusInternalServerError,
	},
	ErrorKeyUserCannotDelete: {
		Message: "Cannot delete user",
		Status:  http.StatusInternalServerError,
	},
	ErrorKeyUserMailAlreadyExists: {
		Message: "Email already exists",
		Status:  http.StatusBadRequest,
	},
	ErrorKeySiteURLAlreadyExists: {
		Message: "URL already exists",
		Status:  http.StatusBadRequest,
	},
	ErrorKeySiteNotFound: {
		Message: "Sites not found",
		Status:  http.StatusNotFound,
	},
	ErrorKeySiteCannotCreate: {
		Message: "Cannot create site",
		Status:  http.StatusInternalServerError,
	},
	ErrorKeySiteCannotUpdate: {
		Message: "Cannot update site",
		Status:  http.StatusInternalServerError,
	},
	ErrorKeySiteCannotDelete: {
		Message: "Cannot delete site",
		Status:  http.StatusInternalServerError,
	},
	ErrorKeySiteCannotDeleteForWrongSource: {
		Message: "Cannot delete site with this type of source",
		Status:  http.StatusBadRequest,
	},
	ErrorKeyAlertNotFound: {
		Message: "Alert not found",
		Status:  http.StatusNotFound,
	},
	ErrorKeyAlertCannotCreate: {
		Message: "Cannot create alert",
		Status:  http.StatusInternalServerError,
	},
	ErrorKeyAlertCannotUpdate: {
		Message: "Cannot update alert",
		Status:  http.StatusInternalServerError,
	},
	ErrorKeyAlertCannotDelete: {
		Message: "Cannot delete alert",
		Status:  http.StatusInternalServerError,
	},
	ErrorKeyAuthInvalidAuthHeader: {
		Message: "Invalid authorization header",
		Status:  http.StatusUnauthorized,
	},
	ErrorKeyAuthEmailOrPasswordIncorrent: {
		Message: "Email or password incorrent",
		Status:  http.StatusUnauthorized,
	},
	ErrorKeyAuthPasswordNotMatch: {
		Message: "Password not match",
		Status:  http.StatusBadRequest,
	},
	ErrorKeyAuthInvalidToken: {
		Message: "Invalid token",
		Status:  http.StatusBadRequest,
	},
	ErrorKeyAuthUnauthorized: {
		Message: "Unauthorized",
		Status:  http.StatusUnauthorized,
	},
}

func NewApiError(key string) ApiError {
	value, ok := errorMap[key]
	if ok {
		return value
	}

	err := errorMap["generic"]
	err.Message = key

	return err
}
