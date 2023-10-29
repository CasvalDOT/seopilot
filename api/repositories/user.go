package repositories

import (
	"api/database"
	"api/models"
	"time"
)

type userRepository struct {
	database database.Database
}

type UserRepository interface {
	GetOne(id int) (models.User, error)
	GetMany(map[string]string) ([]models.User, error)
	GetByEmail(email string) (models.User, error)
	Insert(models.User) (models.User, error)
	Save(models.User) (models.User, error)
	Delete(int) (models.User, error)
}

func (r *userRepository) GetOne(id int) (models.User, error) {
	var entity models.User

	result := r.database.Preload("Roles").First(&entity, id)

	return entity, result.Error
}

func (r *userRepository) GetByEmail(email string) (models.User, error) {
	var entity models.User

	result := r.database.Where("email = ?", email).First(&entity)

	return entity, result.Error
}

func (r *userRepository) GetMany(filters map[string]string) ([]models.User, error) {
	var entities []models.User = []models.User{}

	query := r.database.Preload("Roles")

	if value, ok := filters["search"]; ok {
		if value != "" {
			query = query.Where("name ilike ?", "%"+value+"%")
		}
	}

	res := query.Find(&entities)

	return entities, res.Error
}

func (c *userRepository) Insert(data models.User) (models.User, error) {
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()

	res := c.database.Create(&data)

	c.database.Model(&data).Association("Roles").Replace(data.Roles)

	return data, res.Error
}

func (c *userRepository) Save(data models.User) (models.User, error) {
	data.UpdatedAt = time.Now()

	c.database.Model(&data).Association("Roles").Replace(data.Roles)
	res := c.database.Save(&data)

	return data, res.Error
}

func (c *userRepository) Delete(id int) (models.User, error) {
	var model models.User

	result := c.database.First(&model, id)
	if result.Error != nil {
		return model, result.Error
	}

	res := c.database.Delete(&model)

	return model, res.Error
}

func NewUserRepository(database database.Database) UserRepository {
	return &userRepository{
		database: database,
	}
}
