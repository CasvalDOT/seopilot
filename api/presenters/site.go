package presenters

import (
	"api/models"
)

type sitePresenter struct {
	alertPresenter  AlertPresenter
	sourcePresenter SourcePresenter
}

type SitePresenter interface {
	View(models.Site) models.SiteViewResponse
	ViewAny([]models.Site) models.SiteViewAnyResponse
}

func (p *sitePresenter) getSource(source models.Source) models.SourceViewResponse {
	return models.SourceViewResponse{
		ID:        source.ID,
		Name:      source.Name,
		Code:      source.Code,
		UpdatedAt: source.UpdatedAt,
		CreatedAt: source.CreatedAt,
	}
}

func (p *sitePresenter) getAlert(alert models.Alert) models.AlertViewResponse {
	return models.AlertViewResponse{
		ID:          alert.ID,
		Name:        alert.Name,
		Description: alert.Description,
		CreatedAt:   alert.CreatedAt,
		UpdatedAt:   alert.UpdatedAt,
	}
}

func (p *sitePresenter) getAttributes(attributes []models.SiteAttribute) []models.SiteAttributeViewResponse {
	var data []models.SiteAttributeViewResponse = []models.SiteAttributeViewResponse{}
	for _, attribute := range attributes {
		data = append(data, models.SiteAttributeViewResponse{
			ID:    attribute.ID,
			Name:  attribute.Name,
			Value: attribute.Value,
		})
	}

	return data
}

func (p *sitePresenter) View(site models.Site) models.SiteViewResponse {
	var data models.SiteViewResponse = models.SiteViewResponse{
		ID:         site.ID,
		URL:        site.URL,
		Status:     site.Status,
		Source:     p.sourcePresenter.View(site.Source),
		Alert:      p.alertPresenter.View(site.Alert),
		Attributes: p.getAttributes(site.Attributes),
		CreatedAt:  site.CreatedAt,
		UpdatedAt:  site.UpdatedAt,
	}

	return data
}

func (p *sitePresenter) ViewSimple(site models.Site) models.SiteSimpleViewResponse {
	return models.SiteSimpleViewResponse{
		ID:        site.ID,
		URL:       site.URL,
		Status:    site.Status,
		Source:    p.sourcePresenter.ViewSimple(site.Source),
		Alert:     p.alertPresenter.ViewSimple(site.Alert),
		CreatedAt: site.CreatedAt,
		UpdatedAt: site.UpdatedAt,
	}
}

func (p *sitePresenter) ViewAny(sites []models.Site) models.SiteViewAnyResponse {
	var data []models.SiteSimpleViewResponse = []models.SiteSimpleViewResponse{}
	for _, site := range sites {
		data = append(data, p.ViewSimple(site))
	}

	return models.SiteViewAnyResponse{
		Data: data,
	}
}

func NewSitePresenter(alertPresenter AlertPresenter, sourcePresenter SourcePresenter) SitePresenter {
	return &sitePresenter{
		alertPresenter:  alertPresenter,
		sourcePresenter: sourcePresenter,
	}
}
