package routes

import (
	"github.com/dwadp/todos-api/internal/todo"
	"github.com/gofiber/fiber/v2"
)

type todosRoute struct {
	app      *fiber.App
	todoRepo todo.Repository
}

func NewTodos(app *fiber.App, todoRepo todo.Repository) *todosRoute {
	return &todosRoute{app, todoRepo}
}

func (t *todosRoute) Register() {
	routes := t.app.Group("/v1/todos")

	routes.Get("/", t.getAll)
}

func (t *todosRoute) getAll(ctx *fiber.Ctx) error {
	result, err := t.todoRepo.GetAll()

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]string{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(result)
}
