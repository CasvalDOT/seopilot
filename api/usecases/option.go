package usecases

import (
	"api/models"
	"api/repositories"
	"strconv"
)

type optionUsecase struct {
	roleRepo  repositories.RoleRepository
	alertRepo repositories.AlertRepository
}

type OptionUsecase interface {
	ExtractRoleOptions() []models.Option
	ExtractAlertOptions() []models.Option
}

func (o *optionUsecase) ExtractRoleOptions() []models.Option {
	roles, _ := o.roleRepo.GetMany(nil)
	var data []models.Option = []models.Option{}
	for _, role := range roles {
		data = append(data, models.Option{
			Label: role.Name,
			Value: strconv.Itoa(role.ID),
		})
	}

	return data
}

func (o *optionUsecase) ExtractAlertOptions() []models.Option {
	alerts, _ := o.alertRepo.GetMany(nil)
	var data []models.Option = []models.Option{}
	for _, alert := range alerts {
		data = append(data, models.Option{
			Label: alert.Name,
			Value: strconv.Itoa(alert.ID),
		})
	}

	return data
}

func NewOptionUsecase(roleRepo repositories.RoleRepository, alertRepo repositories.AlertRepository) OptionUsecase {
	return &optionUsecase{
		roleRepo:  roleRepo,
		alertRepo: alertRepo,
	}
}
