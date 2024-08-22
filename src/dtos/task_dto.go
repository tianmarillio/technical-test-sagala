package dtos

type CreateTaskDTO struct {
	Title       string  `json:"title" binding:"required"`
	Description *string `json:"description"`
	Status      *string `json:"status"`
	DueDate     *string `json:"due_date"`
}

type UpdateTaskDTO struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Status      *string `json:"status"`
	DueDate     *string `json:"due_date"`
}

type TaskQueryParams struct {
	// Sort format example: "title asc", "status desc"
	Sort string `json:"sort"`
}
