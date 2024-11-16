package task

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TaskController struct {
	TaskService *TaskService
}

func NewTaskController(taskService *TaskService) *TaskController {
	return &TaskController{TaskService: taskService}
}

// @Summary Create a new task
// @Description Adds a new task to the database
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body Task true "Task Data"
// @Success 201 {object} Task
// @Failure 400 {object} map[string]string
// @Router /tasks [post]
func (tc *TaskController) CreateTask(c echo.Context) error {
	var task Task
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}
	result, err := tc.TaskService.CreateTask(context.Background(), task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not create task"})
	}
	return c.JSON(http.StatusCreated, result)
}

// @Summary Get a task by ID
// @Description Retrieves a task by its ID
// @Tags tasks
// @Produce json
// @Param id path string true "Task ID"
// @Success 200 {object} Task
// @Failure 404 {object} map[string]string
// @Router /tasks/{id} [get]
func (tc *TaskController) GetTaskByID(c echo.Context) error {
	id := c.Param("id")
	task, err := tc.TaskService.GetTaskByID(context.Background(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Task not found"})
	}
	return c.JSON(http.StatusOK, task)
}

// Update a task
func (tc *TaskController) UpdateTask(c echo.Context) error {
	id := c.Param("id")
	var updatedTask Task
	if err := c.Bind(&updatedTask); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}
	result, err := tc.TaskService.UpdateTask(context.Background(), id, updatedTask)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not update task"})
	}
	return c.JSON(http.StatusOK, result)
}

// Delete a task
func (tc *TaskController) DeleteTask(c echo.Context) error {
	id := c.Param("id")
	result, err := tc.TaskService.DeleteTask(context.Background(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not delete task"})
	}
	return c.JSON(http.StatusOK, result)
}
