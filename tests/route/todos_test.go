package routes

import (
	"encoding/json"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/dwadp/todos-api/internal"
	"github.com/dwadp/todos-api/internal/todo"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestTodosRoute(t *testing.T) {
	app := internal.NewApp()

	t.Run("get all todos should return 200 OK and the expected body", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/v1/todos", nil)

		resp, err := app.Test(req)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		expectedBody := []todo.Todo{
			{
				ID:          1,
				Title:       "Clean the room",
				Description: "Clean the fucking room",
				IsDone:      true,
				DueDate:     time.Now(),
			},
			{
				ID:          2,
				Title:       "Buy a macbook pro",
				Description: "Buy a fucking macbook pro",
				IsDone:      false,
				DueDate:     time.Now().Add(time.Hour * 24),
			},
		}

		resultBody := make([]todo.Todo, 2)
		if err := json.NewDecoder(resp.Body).Decode(&resultBody); err != nil {
			t.Error(err)
		}

		for k, v := range resultBody {
			assert.Equal(t, expectedBody[k].ID, v.ID)
			assert.Equal(t, expectedBody[k].Title, v.Title)
			assert.Equal(t, expectedBody[k].Description, v.Description)
			assert.Equal(t, expectedBody[k].IsDone, v.IsDone)

			expectedDueDate := expectedBody[k].DueDate
			resultDueDate := v.DueDate

			assert.Equal(t, expectedDueDate.Year(), resultDueDate.Year())
			assert.Equal(t, expectedDueDate.Month(), resultDueDate.Month())
			assert.Equal(t, expectedDueDate.Day(), resultDueDate.Day())
		}

		assert.Equal(t, len(resultBody), 2)
	})

}
