package routes

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/dwadp/todos-api/internal"
	"github.com/dwadp/todos-api/internal/routes"
	"github.com/dwadp/todos-api/internal/todo"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTodoRepo struct {
	mock.Mock
}

func (m MockTodoRepo) GetAll() ([]todo.Todo, error) {
	args := m.Called()
	return args.Get(0).([]todo.Todo), args.Error(1)
}

func (m MockTodoRepo) Create(t *todo.Todo) error {
	return nil
}

func (m MockTodoRepo) GetByID(id int) (*todo.Todo, error) {
	return nil, nil
}

func TestTodosRoute(t *testing.T) {

	t.Run("it returns 2xx if there was no error", func(t *testing.T) {
		app := internal.NewApp()
		todoRepoMock := new(MockTodoRepo)
		givenTodos := []todo.Todo{
			{
				ID:          1,
				Title:       "example todo 1",
				Description: "desc of example todo 1",
				IsDone:      false,
				DueDate:     time.Now(),
			},
			{
				ID:          2,
				Title:       "example todo 2",
				Description: "desc of example todo 2",
				IsDone:      false,
				DueDate:     time.Now(),
			},
		}
		todoRepoMock.On("GetAll").Return(givenTodos, nil)

		routes.NewTodos(app, todoRepoMock).Register()
		req := httptest.NewRequest("GET", "/v1/todos", nil)

		resp, err := app.Test(req)
		if err != nil {
			t.Error(err)
		}

		var actualTodos []todo.Todo
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		assert.NoError(t, json.NewDecoder(resp.Body).Decode(&actualTodos))
		assert.Equal(t, len(givenTodos), len(actualTodos))
	})

	t.Run("it returns 5xx if there was an error in the repository", func(t *testing.T) {
		app := internal.NewApp()
		todoRepoMock := new(MockTodoRepo)
		givenTodos := []todo.Todo{}

		todoRepoMock.On("GetAll").
			Return(givenTodos, fmt.Errorf("error querying to the database"))

		routes.NewTodos(app, todoRepoMock).Register()
		req := httptest.NewRequest("GET", "/v1/todos", nil)

		resp, err := app.Test(req)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)
	})
}
