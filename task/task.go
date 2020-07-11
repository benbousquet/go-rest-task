package task

import (
	"github.com/benbousquet/go-rest-task/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

type Task struct {
	gorm.Model
	Title     string `json:"title"`
	Creator   string `json:"creator"`
	Completed bool   `json:"completed"`
}

func GetTasks(c *fiber.Ctx) {

	var tasks []Task
	database.DBConn.Find(&tasks)
	c.JSON(tasks)
}

func GetTask(c *fiber.Ctx) {

	c.Send("single task with id")
}

func NewTask(c *fiber.Ctx) {

	c.Send("add a task")
}

func DeleteTask(c *fiber.Ctx) {

	c.Send("delete a task")
}
