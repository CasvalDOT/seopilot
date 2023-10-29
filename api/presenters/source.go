package presenters

import (
	"api/models"
)

type sourcePresenter struct{}

type SourcePresenter interface {
	View(models.Source) models.SourceViewResponse
	ViewSimple(models.Source) models.SourceViewSimpleResponse
}

func (p *sourcePresenter) View(source models.Source) models.SourceViewResponse {
	var data models.SourceViewResponse = models.SourceViewResponse{
		ID:        source.ID,
		Name:      source.Name,
		Code:      source.Code,
		CreatedAt: source.CreatedAt,
		UpdatedAt: source.UpdatedAt,
	}

	return data
}

func (p *sourcePresenter) ViewSimple(source models.Source) models.SourceViewSimpleResponse {
	return models.SourceViewSimpleResponse{
		ID:   source.ID,
		Name: source.Name,
	}
}

func NewSourcePresenter() SourcePresenter {
	return &sourcePresenter{}
}
