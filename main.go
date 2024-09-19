package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kweku-xvi/todolist-api/controllers"
	"github.com/kweku-xvi/todolist-api/initializers"
)

func init() {
	initializers.LoadDotEnv()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/tasks", controllers.CreateTask)
	r.GET("/tasks", controllers.GetTasks)
	r.GET("/tasks/:id", controllers.GetSpecificTask)
	r.PUT("/tasks/:id", controllers.UpdateTask)
	r.DELETE("/tasks/:id", controllers.DeleteTask)

	r.Run()
}
