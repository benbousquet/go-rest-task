package main

import (
	"fmt"

	"github.com/benbousquet/go-rest-task/database"
	"github.com/benbousquet/go-rest-task/task"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=tasks password=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connected")

	database.DBConn.AutoMigrate(&task.Task{})
}

func setup(app *fiber.App) {

	api := app.Group("/api/v1")

	api.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, World 👋!")
	})

	// GET /tasks returns all tasks
	api.Get("/tasks", task.GetTasks)

	// GET /task/:id returns single task with id
	api.Get("/task/:id", task.GetTask)

	// POST /task add a task
	api.Post("/task", task.NewTask)

	// DELETE /task/:id returns all tasks
	api.Delete("/task/:id", task.DeleteTask)
}

func main() {
	app := fiber.New()

	initDatabase()
	defer database.DBConn.Close()
	// setup routes
	setup(app)

	app.Listen(3000)
}
