package presenters

import (
	"api/models"
)

type alertPresenter struct{}

type AlertPresenter interface {
	View(models.Alert) models.AlertViewResponse
	ViewSimple(models.Alert) models.AlertViewSimpleResponse
	ViewAny([]models.Alert) models.AlertViewAnyResponse
}

func (p *alertPresenter) View(alert models.Alert) models.AlertViewResponse {
	var data models.AlertViewResponse = models.AlertViewResponse{
		ID:          alert.ID,
		Name:        alert.Name,
		Description: alert.Description,
		CreatedAt:   alert.CreatedAt,
		UpdatedAt:   alert.UpdatedAt,
	}

	return data
}

func (p *alertPresenter) ViewSimple(alert models.Alert) models.AlertViewSimpleResponse {
	var data models.AlertViewSimpleResponse = models.AlertViewSimpleResponse{
		ID:   alert.ID,
		Name: alert.Name,
	}

	return data
}

func (p *alertPresenter) ViewAny(alerts []models.Alert) models.AlertViewAnyResponse {
	var data []models.AlertViewResponse = []models.AlertViewResponse{}
	for _, alert := range alerts {
		data = append(data, p.View(alert))
	}

	return models.AlertViewAnyResponse{
		Data: data,
	}
}

func NewAlertPresenter() AlertPresenter {
	return &alertPresenter{}
}
