package main

import (
	"github.com/kweku-xvi/todolist-api/initializers"
	"github.com/kweku-xvi/todolist-api/models"
)

func init() {
	initializers.LoadDotEnv()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Task{})
}
