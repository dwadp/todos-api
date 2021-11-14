package internal

import (
	"log"

	"github.com/dwadp/todos-api/db"
	"github.com/dwadp/todos-api/internal/routes"
	"github.com/dwadp/todos-api/internal/todo"
	todoRepo "github.com/dwadp/todos-api/internal/todo/repository"
	"github.com/gofiber/fiber/v2"
)

func NewApp() *fiber.App {
	app := fiber.New()

	dbConn, err := db.Connect("root:@tcp(127.0.0.1:3306)/todos?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}

	dbConn.AutoMigrate(&todo.Todo{})

	todoRepository := todoRepo.NewTodoMysqlRepository(dbConn)

	routes.NewTodos(app, todoRepository).Register()

	return app
}
