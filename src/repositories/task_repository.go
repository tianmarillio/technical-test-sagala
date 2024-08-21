package repositories

import (
	"github.com/tianmarillio/technical-test-sagala/src/models"
	"gorm.io/gorm"
)

type TaskRepository interface {
	Create(*models.Task) (*models.Task, error)
	FindAll() ([]models.Task, error)
	FindByID(id uint) (*models.Task, error)
	Update(task *models.Task) error
	Delete(id uint) error
}

type GormTaskRepository struct {
	db *gorm.DB
}

func NewGormTaskRepository(db *gorm.DB) *GormTaskRepository {
	return &GormTaskRepository{
		db: db,
	}
}

func (r *GormTaskRepository) Create(task *models.Task) (*models.Task, error) {
	err := r.db.Create(task).Error

	if err != nil {
		return nil, err
	}

	return task, nil
}

func (r *GormTaskRepository) FindAll() ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.Find(&tasks).Error

	return tasks, err
}

func (r *GormTaskRepository) FindByID(id uint) (*models.Task, error) {
	var task models.Task
	err := r.db.First(&task, id).Error

	return &task, err
}

func (r *GormTaskRepository) Update(task *models.Task) error {
	return r.db.Save(task).Error
}

func (r *GormTaskRepository) Delete(id uint) error {
	// Soft delete
	return r.db.Delete(&models.Task{}, id).Error
}
