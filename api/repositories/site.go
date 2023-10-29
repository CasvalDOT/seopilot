package repositories

import (
	"api/database"
	"api/models"
	"time"
)

type siteRepository struct {
	database database.Database
}

type SiteRepository interface {
	GetOne(id int) (models.Site, error)
	GetMany(map[string]string) ([]models.Site, error)
	GetByURL(string) (models.Site, error)
	Insert(models.Site) (models.Site, error)
	Save(models.Site) (models.Site, error)
	Delete(int) (models.Site, error)
}

func (r *siteRepository) GetOne(id int) (models.Site, error) {
	var entity models.Site

	result := r.database.
		Preload("Alert").
		Preload("Source").
		Preload("Attributes").
		First(&entity, id)

	return entity, result.Error
}

func (r *siteRepository) GetMany(filters map[string]string) ([]models.Site, error) {
	var entities []models.Site = []models.Site{}

	query := r.database.Preload("Alert").Preload("Source")

	if value, ok := filters["search"]; ok {
		if value != "" {
			query = query.Where("url ilike ?", "%"+value+"%")
		}
	}

	res := query.Find(&entities)

	return entities, res.Error
}

func (c *siteRepository) Insert(data models.Site) (models.Site, error) {
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()

	res := c.database.Create(&data)
	if res.Error != nil {
		return data, res.Error
	}

	return c.GetOne(data.ID)
}

func (r *siteRepository) GetByURL(url string) (models.Site, error) {
	var entity models.Site

	result := r.database.Where("url = ?", url).First(&entity)

	return entity, result.Error
}

func (c *siteRepository) Save(data models.Site) (models.Site, error) {
	data.UpdatedAt = time.Now()

	res := c.database.Save(&data)

	return data, res.Error
}

func (c *siteRepository) Delete(id int) (models.Site, error) {
	var model models.Site

	result := c.database.First(&model, id)
	if result.Error != nil {
		return model, result.Error
	}

	res := c.database.Delete(&model)

	return model, res.Error
}

func NewSiteRepository(database database.Database) SiteRepository {
	return &siteRepository{
		database: database,
	}
}
