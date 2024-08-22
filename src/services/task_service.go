package services

import (
	"errors"

	"github.com/jinzhu/now"
	"github.com/tianmarillio/technical-test-sagala/src/dtos"
	"github.com/tianmarillio/technical-test-sagala/src/models"
	"github.com/tianmarillio/technical-test-sagala/src/repositories"
)

type TaskService struct {
	repository repositories.TaskRepository
}

func NewTaskService(r repositories.TaskRepository) *TaskService {
	return &TaskService{repository: r}
}

func (s *TaskService) CreateTask(dto dtos.CreateTaskDTO) (*models.Task, error) {
	task := &models.Task{
		Title: dto.Title,
	}

	if dto.Description != nil {
		task.Description = *dto.Description
	}

	// Validate dto.Status enum
	if dto.Status != nil {
		status := models.TaskStatus(*dto.Status)
		task.Status = status
	} else {
		defaultStatus := models.TaskStatusWaitingList
		task.Status = defaultStatus
	}

	// Parse dueDate from string
	if dto.DueDate != nil {
		dueDate, err := now.Parse(*dto.DueDate)
		if err != nil {
			return nil, err
		}

		task.DueDate = dueDate
	}

	task, err := s.repository.Create(task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (s *TaskService) GetTasks(queryParams dtos.TaskQueryParams) ([]models.Task, error) {
	return s.repository.FindAll(queryParams)
}

func (s *TaskService) GetTask(id uint) (*models.Task, error) {
	return s.repository.FindByID(id)
}

func (s *TaskService) UpdateTask(id uint, dto dtos.UpdateTaskDTO) error {
	task, err := s.repository.FindByID(id)
	if err != nil {
		return err
	}

	if task == nil {
		return errors.New("task not found")
	}

	if dto.Title != nil {
		task.Title = *dto.Title
	}

	if dto.Description != nil {
		task.Description = *dto.Description
	}

	if dto.Status != nil {
		status := models.TaskStatus(*dto.Status)
		task.Status = status
	}

	if dto.DueDate != nil {
		dueDate, err := now.Parse(*dto.DueDate)
		if err != nil {
			return err
		}

		task.DueDate = dueDate
	}

	return s.repository.Update(task)
}

func (s *TaskService) DeleteTask(id uint) error {
	return s.repository.Delete(id)
}
