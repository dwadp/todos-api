package repositories

import (
	"time"

	"github.com/dwadp/todos-api/internal/todo"
	"github.com/dwadp/todos-api/internal/todo/repository"
)

func (m *MysqlRepositoryTestSuite) TestMysqlTodoRepository_CreateTodo() {
	now := time.Now()
	newTodo := &todo.Todo{
		Title:       "sample title",
		Description: "sample description",
		IsDone:      false,
		DueDate:     now,
	}

	repo := repository.NewTodoMysqlRepository(m.gormDB)

	m.Assert().NoError(repo.Create(newTodo))

	var result todo.Todo
	m.Assert().NoError(m.gormDB.First(&result, todo.Todo{ID: newTodo.ID}).Error)
	m.Assert().Equal(newTodo.Title, result.Title)
	m.Assert().Equal(newTodo.Description, result.Description)
	m.Assert().Equal(newTodo.IsDone, result.IsDone)
	m.Assert().Equal(newTodo.DueDate.Format("2006-02-01 15:04:05"), result.DueDate.Format("2006-02-01 15:04:05"))
}

func (m *MysqlRepositoryTestSuite) TestMysqlTodoRepository_GetByID() {
	newTodo := &todo.Todo{
		Title:       "first task",
		Description: "first task description",
		IsDone:      true,
		DueDate:     time.Now(),
	}

	repo := repository.NewTodoMysqlRepository(m.gormDB)

	m.Assert().NoError(repo.Create(newTodo))

	result, err := repo.GetByID(newTodo.ID)

	m.Assert().NoError(err)
	m.Assert().NotNil(result)
	m.Assert().Equal(newTodo.Title, result.Title)
	m.Assert().Equal(newTodo.Description, result.Description)
	m.Assert().Equal(newTodo.IsDone, result.IsDone)
	m.Assert().Equal(newTodo.DueDate.Format("2006-02-01 15:04:05"), result.DueDate.Format("2006-02-01 15:04:05"))
}

func (m *MysqlRepositoryTestSuite) TestMysqlTodoRepository_GetAll() {
	firstTask := &todo.Todo{
		Title:       "first task",
		Description: "first task description",
		IsDone:      true,
		DueDate:     time.Now(),
	}

	secondTask := &todo.Todo{
		Title:       "second task",
		Description: "second task description",
		IsDone:      false,
		DueDate:     time.Now().Add(24 * time.Hour),
	}

	repo := repository.NewTodoMysqlRepository(m.gormDB)

	m.Assert().NoError(repo.Create(firstTask))
	m.Assert().NoError(repo.Create(secondTask))

	result, err := repo.GetAll()
	m.Assert().NoError(err)
	m.Assert().Equal(len(result), 2)
}
