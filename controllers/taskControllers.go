package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kweku-xvi/todolist-api/initializers"
	"github.com/kweku-xvi/todolist-api/models"
)

func CreateTask(c *gin.Context) {
	var body struct {
		Title       string
		Description string
		Priority    string
	}
	c.Bind(&body)

	task := models.Task{
		Title:       body.Title,
		Description: body.Description,
		Priority:    body.Priority,
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
	result := initializers.DB.First(&task, id)

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
		Title       string
		Description string
		Priority    string
	}
	c.Bind(&body)

	var task models.Task
	result := initializers.DB.First(&task, id)

	if result.Error != nil {
		c.JSON(404, gin.H{
			"message": "task not found",
		})
	} else {
		initializers.DB.Model(&task).Updates(models.Task{Title: body.Title,
			Description: body.Description,
			Priority:    body.Priority,
		})
		c.JSON(200, gin.H{
			"task": task,
		})
	}
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")

	var task models.Task
	result := initializers.DB.First(&task, id)

	if result.Error != nil {
		c.JSON(404, gin.H{
			"message": "task not found",
		})
	} else {
		initializers.DB.Delete(&models.Task{}, id)
		c.JSON(200, gin.H{
			"message": "task successfully deleted",
		})
	}

}
