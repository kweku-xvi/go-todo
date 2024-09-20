package controllers

import (
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kweku-xvi/todolist-api/initializers"
	"github.com/kweku-xvi/todolist-api/models"
)

func GenerateUID(n int) string {
	const lettersAndNumbers = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var uid string

	for i := 0; i < n; i++ {
		uid += string(lettersAndNumbers[rand.Intn(len(lettersAndNumbers))])
	}
	return uid
}

func CreateTask(c *gin.Context) {
	var body struct {
		Title       string    `json:"title" form:"title"`
		Description string    `json:"description" form:"description"`
		Priority    string    `json:"priority" form:"priority"`
		Deadline    time.Time `json:"deadline" form:"deadline"`
		Status      string    `json:"status" form:"status"`
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	task := models.Task{
		ID:          GenerateUID(13),
		Title:       body.Title,
		Description: body.Description,
		Priority:    body.Priority,
		Deadline:    body.Deadline,
		Status:      body.Status,
		CreatedAt:   time.Now(),
	}

	result := initializers.DB.Create(&task)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"message": task,
	})
}

func GetTasks(c *gin.Context) {
	var tasks []models.Task

	initializers.DB.Find(&tasks)

	c.JSON(200, gin.H{
		"tasks": tasks,
	})
}

func GetSpecificTask(c *gin.Context) {
	id := c.Param("id")

	var task models.Task
	result := initializers.DB.First(&task, "id = ?",id)

	if result.Error != nil {
		c.JSON(404, gin.H{
			"message": "task not found",
		})
	} else {
		c.JSON(200, gin.H{
			"task": task,
		})
	}
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Title       string    `json:"title" form:"title"`
		Description string    `json:"description" form:"description"`
		Priority    string    `json:"priority" form:"priority"`
		Deadline    time.Time `json:"deadline" form:"deadline"`
		Status      string    `json:"status" form:"status"`
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	var task models.Task
	result := initializers.DB.First(&task, "id = ?",id)

	if result.Error != nil {
		c.JSON(404, gin.H{
			"message": "task not found",
		})
	} else {
		initializers.DB.Model(&task).Updates(models.Task{Title: body.Title,
			Description: body.Description,
			Priority:    body.Priority,
			Deadline:    body.Deadline,
			Status:      body.Status,
			UpdatedAt:   time.Now(),
		})
		c.JSON(200, gin.H{
			"task": task,
		})
	}
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")

	var task models.Task
	result := initializers.DB.First(&task, "id = ?",id)

	if result.Error != nil {
		c.JSON(404, gin.H{
			"message": "task not found",
		})
	} else {
		initializers.DB.Delete(&task, "id = ?",id)
		c.JSON(200, gin.H{
			"message": "task successfully deleted",
		})
	}

}
