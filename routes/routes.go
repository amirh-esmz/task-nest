package routes

import (
	"github.com/amirh-esmz/task-nest/internal/category"
	"github.com/amirh-esmz/task-nest/internal/task"
	"github.com/labstack/echo/v4"
)

func RegisterTaskRoutes(e *echo.Echo, taskController *task.TaskController) {
	e.POST("/tasks", taskController.CreateTask)
	e.GET("/tasks/:id", taskController.GetTaskByID)
	e.PUT("/tasks/:id", taskController.UpdateTask)
	e.DELETE("/tasks/:id", taskController.DeleteTask)
}

func RegisterCategoryRoutes(e *echo.Echo, categoryController *category.CategoryController) {
	e.POST("/categories", categoryController.Creat)
	e.GET("/tasks/:id", categoryController.GetById)

}
