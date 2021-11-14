package todo

import "time"

type Todo struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	Title       string    `json:"string"`
	Description string    `json:"description"`
	IsDone      bool      `json:"is_done"`
	DueDate     time.Time `json:"due_date"`
}

func (t Todo) TableName() string {
	return "todos"
}
