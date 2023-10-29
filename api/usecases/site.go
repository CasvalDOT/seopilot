package usecases

import (
	"api/constants"
	ae "api/errors"
	"api/models"
	"api/repositories"
	"errors"
)

type siteUsecase struct {
	siteRepo   repositories.SiteRepository
	sourceRepo repositories.SourceRepository
	alertRepo  repositories.AlertRepository
}

type SiteUsecase interface {
	GetSite(int) (models.Site, error)
	GetSites(map[string]string) ([]models.Site, error)
	CreateSite(models.SiteCreateRequest) (models.Site, error)
	UpdateSite(int, models.SiteUpdateRequest) (models.Site, error)
	DeleteSite(int) (models.Site, error)
}

func (u *siteUsecase) GetSite(id int) (models.Site, error) {
	site, err := u.siteRepo.GetOne(id)
	if err != nil {
		return site, err
	}

	return site, nil
}

func (u *siteUsecase) GetSites(filters map[string]string) ([]models.Site, error) {
	sites, err := u.siteRepo.GetMany(filters)
	if err != nil {
		return sites, err
	}

	return sites, nil
}

func (u *siteUsecase) CreateSite(payload models.SiteCreateRequest) (models.Site, error) {
	// Take the source manual
	source, _ := u.sourceRepo.GetByCode(models.SourceCodeManual)

	site := models.Site{
		URL:      payload.URL,
		AlertID:  payload.AlertID,
		SourceID: source.ID,
		Status:   constants.SiteStatusPending,
	}

	oldSite, _ := u.siteRepo.GetByURL(site.URL)
	if oldSite.ID != 0 {
		return site, errors.New(ae.ErrorKeySiteURLAlreadyExists)
	}

	// Take the alert to associate
	if payload.AlertID != nil {
		alert, err := u.alertRepo.GetOne(*payload.AlertID)
		if err != nil {
			return site, errors.New(ae.ErrorKeyAlertNotFound)
		}

		site.Alert = alert
	}

	newSite, err := u.siteRepo.Insert(site)
	if err != nil {
		return newSite, errors.New(ae.ErrorKeySiteCannotCreate)
	}

	return newSite, nil
}

func (u *siteUsecase) UpdateSite(id int, payload models.SiteUpdateRequest) (models.Site, error) {
	site, err := u.siteRepo.GetOne(id)
	if err != nil || site.ID == 0 {
		return site, errors.New(ae.ErrorKeySiteNotFound)
	}

	// Check if email already exists
	siteByURL, _ := u.siteRepo.GetByURL(payload.URL)
	if siteByURL.ID != 0 && siteByURL.ID != site.ID {
		return site, errors.New(ae.ErrorKeySiteURLAlreadyExists)
	}

	site.URL = payload.URL

	// Associate Alert here
	if payload.AlertID != nil && *payload.AlertID != 0 {
		alert, err := u.alertRepo.GetOne(*payload.AlertID)
		if err != nil {
			return site, errors.New(ae.ErrorKeyAlertNotFound)
		}

		site.Alert = alert
	}

	site, err = u.siteRepo.Save(site)
	if err != nil {
		return site, errors.New(ae.ErrorKeySiteCannotUpdate)
	}

	return site, nil
}

func (u *siteUsecase) DeleteSite(id int) (models.Site, error) {
	site, err := u.siteRepo.GetOne(id)
	if err != nil || site.ID == 0 {
		return site, errors.New(ae.ErrorKeySiteNotFound)
	}

	// Cannot delete site with source different from manula
	// because the system can re-take the source at next sync
	// with cloud provider
	if site.Source.Code != models.SourceCodeManual {
		return site, errors.New(ae.ErrorKeySiteCannotDeleteForWrongSource)
	}

	site, err = u.siteRepo.Delete(id)
	if err != nil {
		return site, errors.New(ae.ErrorKeySiteCannotDelete)
	}

	return site, nil
}

func NewSiteUsecase(
	siteRepo repositories.SiteRepository,
	sourceRepo repositories.SourceRepository,
	alertRepo repositories.AlertRepository,
) SiteUsecase {
	return &siteUsecase{
		siteRepo:   siteRepo,
		sourceRepo: sourceRepo,
		alertRepo:  alertRepo,
	}
}
