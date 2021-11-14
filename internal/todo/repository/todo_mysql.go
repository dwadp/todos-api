package repository

import (
	"github.com/dwadp/todos-api/internal/todo"
	"gorm.io/gorm"
)

type todoMysqlRepository struct {
	conn *gorm.DB
}

func NewTodoMysqlRepository(conn *gorm.DB) *todoMysqlRepository {
	return &todoMysqlRepository{conn}
}

func (t *todoMysqlRepository) Create(td *todo.Todo) error {
	return t.conn.Create(td).Error
}

func (t *todoMysqlRepository) GetAll() ([]todo.Todo, error) {
	var result []todo.Todo
	if err := t.conn.Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (t *todoMysqlRepository) GetByID(id int) (*todo.Todo, error) {
	result := &todo.Todo{}
	tx := t.conn.Where("id = ?", id).First(result)
	return result, tx.Error
}
