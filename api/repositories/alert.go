package repositories

import (
	"api/database"
	"api/models"
	"time"
)

type alertRepository struct {
	database database.Database
}

type AlertRepository interface {
	GetOne(id int) (models.Alert, error)
	GetMany(map[string]string) ([]models.Alert, error)
	Insert(models.Alert) (models.Alert, error)
	Save(models.Alert) (models.Alert, error)
	Delete(int) (models.Alert, error)
}

func (r *alertRepository) GetOne(id int) (models.Alert, error) {
	var entity models.Alert

	result := r.database.First(&entity, id)

	return entity, result.Error
}

func (r *alertRepository) GetMany(filters map[string]string) ([]models.Alert, error) {
	var entities []models.Alert = []models.Alert{}

	query := r.database.Order("updated_at desc")

	if value, ok := filters["search"]; ok {
		if value != "" {
			query = query.Where("name ilike ?", "%"+value+"%")
		}
	}

	res := query.Find(&entities)

	return entities, res.Error
}

func (c *alertRepository) Insert(data models.Alert) (models.Alert, error) {
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()

	res := c.database.Create(&data)

	return data, res.Error
}

func (c *alertRepository) Save(data models.Alert) (models.Alert, error) {
	data.UpdatedAt = time.Now()

	res := c.database.Save(&data)

	return data, res.Error
}

func (c *alertRepository) Delete(id int) (models.Alert, error) {
	var model models.Alert

	result := c.database.First(&model, id)
	if result.Error != nil {
		return model, result.Error
	}

	res := c.database.Delete(&model)

	return model, res.Error
}

func NewAlertRepository(database database.Database) AlertRepository {
	return &alertRepository{
		database: database,
	}
}
