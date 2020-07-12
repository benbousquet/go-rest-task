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
	var task Task
	id := c.Params("id")
	database.DBConn.First(&task, id)
	if task.Title == "" {
		c.Status(500).Send("Cannot find a task with ID of ", id)
		return
	}
	c.JSON(task)
}

func NewTask(c *fiber.Ctx) {
	// create new task
	task := new(Task)
	// put post req body onto task struct
	if err := c.BodyParser(task); err != nil {
		c.Status(503).Send(err)
		return
	}
	// put into db
	database.DBConn.Create(&task)
	c.JSON(task)
}

func DeleteTask(c *fiber.Ctx) {
	var task Task
	// get param and check if its in db
	id := c.Params("id")
	database.DBConn.First(&task, id)
	if task.Title == "" {
		c.Status(500).Send("No task could be found with that ID")
		return
	}
	// delete task
	database.DBConn.Delete(task)
	c.Send("Successfully Deleted Task with ID of ", id)
}
