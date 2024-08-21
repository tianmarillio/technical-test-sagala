package models

import (
	"time"

	"gorm.io/gorm"
)

type TaskStatus string

const (
	TaskStatusWaitingList TaskStatus = "waiting_list"
	TaskStatusInProgress  TaskStatus = "in_progress"
	TaskStatusDone        TaskStatus = "done"
)

type Task struct {
	ID          uint `gorm:"primarykey"`
	Title       string
	Description string
	Status      TaskStatus `gorm:"type:task_status"`
	DueDate     time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
