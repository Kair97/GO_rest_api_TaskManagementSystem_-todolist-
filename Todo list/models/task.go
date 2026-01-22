package models

type Task struct {
	ID          int     `json:"id,omitempty"`
	Title       string  `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Completed   bool    `json:"completed"`
	Priority    int     `json:"priority,omitempty"`
}

type CreateTask struct {
	Title       string  `json:"title" binding:"required,min=3"`
	Description *string `json:"description,omitempty" binding:"omitempty,min=3"`
	Priority    int     `json:"priority" binding:"required"`
}

type UpdateTask struct {
	Title       string  `json:"title,,omitempty" binding:"required"`
	Description *string `json:"description,omitempty" binding:"omitempty,min=3"`
	Completed   bool    `json:"completed"`
	Priority    int     `json:"priority,omitempty" binding:"required"`
}

type PatchTask struct {
	Title       *string `json:"title" binding:"omitempty,min=3"`
	Description *string `json:"description,omitempty" binding:"omitempty,min=3"`
	Completed   *bool   `json:"completed"`
	Priority    *int    `json:"priority" binding:"omitempty,min=1,max=5"`
}
