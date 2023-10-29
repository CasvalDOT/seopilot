package repositories

import (
	"api/database"
	"api/models"
)

type roleRepository struct {
	database database.Database
}

type RoleRepository interface {
	GetOne(id int) (models.Role, error)
	GetMany(map[string]interface{}) ([]models.Role, error)
}

func (r *roleRepository) GetOne(id int) (models.Role, error) {
	var entity models.Role

	result := r.database.First(&entity, id)

	return entity, result.Error
}

func (r *roleRepository) GetMany(filters map[string]interface{}) ([]models.Role, error) {
	var entities []models.Role = []models.Role{}

	query := r.database.Model(&models.Role{})

	if value, ok := filters["ids"]; ok {
		query = query.Where("id in (?)", value)
	}

	res := query.Find(&entities)

	return entities, res.Error
}

func NewRoleRepository(database database.Database) RoleRepository {
	return &roleRepository{
		database: database,
	}
}
