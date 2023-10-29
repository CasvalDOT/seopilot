package usecases

import (
	ae "api/errors"
	"api/models"
	"api/repositories"
	"errors"
	"fmt"
)

type alertUsecase struct {
	alertRepo repositories.AlertRepository
}

type AlertUsecase interface {
	GetAlert(int) (models.Alert, error)
	GetAlerts(map[string]string) ([]models.Alert, error)
	CreateAlert(models.AlertCreateRequest) (models.Alert, error)
	UpdateAlert(int, models.AlertUpdateRequest) (models.Alert, error)
	DeleteAlert(int) (models.Alert, error)
}

func (u *alertUsecase) GetAlert(id int) (models.Alert, error) {
	alert, err := u.alertRepo.GetOne(id)
	if err != nil {
		return alert, err
	}

	return alert, nil
}

func (u *alertUsecase) GetAlerts(filters map[string]string) ([]models.Alert, error) {
	alerts, err := u.alertRepo.GetMany(filters)
	if err != nil {
		return alerts, err
	}

	return alerts, nil
}

func (u *alertUsecase) CreateAlert(payload models.AlertCreateRequest) (models.Alert, error) {
	alert := models.Alert{
		Name:        payload.Name,
		Description: payload.Description,
	}

	newAlert, err := u.alertRepo.Insert(alert)
	if err != nil {
		return newAlert, errors.New(ae.ErrorKeyAlertCannotCreate)
	}

	// Send email

	return newAlert, nil
}

func (u *alertUsecase) UpdateAlert(id int, payload models.AlertUpdateRequest) (models.Alert, error) {
	alert, err := u.alertRepo.GetOne(id)
	if err != nil || alert.ID == 0 {
		return alert, errors.New(ae.ErrorKeyAlertNotFound)
	}

	// Check if roles are valid
	alert.Name = payload.Name
	alert.Description = payload.Description

	alert, err = u.alertRepo.Save(alert)
	if err != nil {
		fmt.Println(err)
		return alert, errors.New(ae.ErrorKeyAlertCannotUpdate)
	}

	return alert, nil
}

func (u *alertUsecase) DeleteAlert(id int) (models.Alert, error) {
	alert, err := u.alertRepo.GetOne(id)
	if err != nil || alert.ID == 0 {
		return alert, errors.New(ae.ErrorKeyAlertNotFound)
	}

	alert, err = u.alertRepo.Delete(id)
	if err != nil {
		return alert, errors.New(ae.ErrorKeyAlertCannotDelete)
	}

	return alert, nil
}

func NewAlertUsecase(alertRepo repositories.AlertRepository) AlertUsecase {
	return &alertUsecase{
		alertRepo: alertRepo,
	}
}
