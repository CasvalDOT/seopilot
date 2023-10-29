package presenters

import (
	"api/models"
)

type userPresenter struct{}

type UserPresenter interface {
	View(models.User) models.UserViewResponse
	ViewAny([]models.User) models.UserViewAnyResponse
}

func (p *userPresenter) View(user models.User) models.UserViewResponse {
	userRoles := []models.RoleSimple{}
	for _, role := range user.Roles {
		userRoles = append(userRoles, models.RoleSimple{
			ID:   role.ID,
			Name: role.Name,
		})
	}

	var data models.UserViewResponse = models.UserViewResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Roles:     userRoles,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return data
}

func (p *userPresenter) ViewAny(users []models.User) models.UserViewAnyResponse {
	var data []models.UserViewResponse = []models.UserViewResponse{}
	for _, user := range users {
		data = append(data, p.View(user))
	}

	return models.UserViewAnyResponse{
		Data: data,
	}
}

func NewUserPresenter() UserPresenter {
	return &userPresenter{}
}
