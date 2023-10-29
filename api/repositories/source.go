package repositories

import (
	"api/database"
	"api/models"
	"time"
)

type sourceRepository struct {
	database database.Database
}

type SourceRepository interface {
	GetByCode(string) (models.Source, error)
	Save(models.Source) (models.Source, error)
}

func (r *sourceRepository) GetByCode(code string) (models.Source, error) {
	var entity models.Source

	result := r.database.Where("code = ?", code).First(&entity)

	return entity, result.Error
}

func (c *sourceRepository) Save(data models.Source) (models.Source, error) {
	data.UpdatedAt = time.Now()

	res := c.database.Save(&data)

	return data, res.Error
}

func NewSourceRepository(database database.Database) SourceRepository {
	return &sourceRepository{
		database: database,
	}
}
