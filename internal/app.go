package internal

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/dwadp/todos-api/db"
	"github.com/dwadp/todos-api/internal/todo"
	"github.com/gofiber/fiber/v2"
)

func NewApp() *fiber.App {
	app := fiber.New()

	dbConn, err := db.Connect("root:@tcp(127.0.0.1:3306)/todos?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}

	dbConn.AutoMigrate(&todo.Todo{})

	// todoRepository := todoRepo.NewTodoMysqlRepository(dbConn)

	// routes.NewTodos(app, todoRepository).Register()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		fmt.Println("os interupt: gracefully shutting down the server")

		db, _ := dbConn.DB()

		if err := db.Close(); err != nil {
			log.Fatalf("error closing database connection: [%v]", err)
		}

		if err := app.Shutdown(); err != nil {
			log.Fatalf("error shutting down application: [%v]", err)
		}
	}()

	return app
}
