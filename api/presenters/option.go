package presenters

import (
	"api/models"
)

type optionPresenter struct{}

type OptionPresenter interface {
	ViewAny([]models.Option) models.OptionData
}

func (p *optionPresenter) ViewAny(options []models.Option) models.OptionData {
	return models.OptionData{
		Data: options,
	}
}

func NewOptionPresenter() OptionPresenter {
	return &optionPresenter{}
}
