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
	ID          uint        `json:"id" gorm:"primarykey"`
	Title       string      `json:"title"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status" gorm:"type:task_status"`
	DueDate     time.Time  `json:"due_date"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
