package dtos

type CreateTaskDTO struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Status      string `json:"status"`
	DueDate     string `json:"due_date" binding:"required"`
}

// FIXME: partial
type UpdateTaskDTO struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Status      *string `json:"status"`
	DueDate     *string `json:"due_date"`
}
