package main

import (
	"context"
	"log"
	"net/http"

	"github.com/amirh-esmz/task-nest/internal/category"
	"github.com/amirh-esmz/task-nest/internal/task"
	"github.com/amirh-esmz/task-nest/routes"

	_ "github.com/amirh-esmz/task-nest/docs"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017") // todo : read from env
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	taskCollection := client.Database("task-nest").Collection("tasks") // todo : improve

	taskService := task.NewTaskService(taskCollection)
	taskController := task.NewTaskController(taskService)

	categoryCollection := client.Database("task-nest").Collection("categories")

	categoryService := category.NewCategoryService(categoryCollection)
	categoryController := category.NewCategoryController(categoryService)

	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	routes.RegisterTaskRoutes(e, taskController)
	routes.RegisterCategoryRoutes(e, categoryController)

	e.GET("/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Start("127.0.0.1:8080")
}
