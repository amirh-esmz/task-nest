package category

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CategoryController struct {
	CategoryService *CategoryService
}

func NewCategoryController(categoryService *CategoryService) *CategoryController {
	return &CategoryController{CategoryService: categoryService}
}

// @Summary Create a new category
// @Description Adds a new category to the database
// @Tags categories
// @Accept json
// @Produce json
// @Param category body Category true "category Data"
// @Success 201 {object} Category
// @Failure 400 {object} map[string]string
// @Router /categories [post]
func (tc *CategoryController) Creat(c echo.Context) error {
	var category Category
	if err := c.Bind(&category); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}
	result, err := tc.CategoryService.Create(context.Background(), category)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusCreated, result)
}

// @Summary Get a category by ID
// @Description Retrieves a category by its ID
// @Tags categories
// @Produce json
// @Param id path string true "category ID"
// @Success 200 {object} Category
// @Failure 404 {object} map[string]string
// @Router /categories/{id} [get]
func (tc *CategoryController) GetById(c echo.Context) error {
	id := c.Param("id")
	task, err := tc.CategoryService.GetById(context.Background(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Task not found"})
	}
	return c.JSON(http.StatusOK, task)
}
