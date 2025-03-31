package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"api-project/commons"
	"api-project/database"
	"api-project/models"
)

func GetAllCategories(c echo.Context) error {
	var categories []models.Category
	database.DB.Preload("Products").Find(&categories)
	return c.JSON(http.StatusOK, categories)
}

func GetCategoryById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": commons.InvalidIdError})
	}

	var category models.Category
	if err := database.DB.Preload("Products").First(&category, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"error": commons.CategoryNotFoundError})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": commons.DatabaseError})
	}

	return c.JSON(http.StatusOK, category)
}

func CreateCategory(c echo.Context) error {
	category := new(models.Category)
	if err := c.Bind(category); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": commons.InvalidRequestError})
	}
	database.DB.Create(category)
	return c.JSON(http.StatusCreated, category)
}

func DeleteCategory(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": commons.InvalidIdError})
	}

	category, status, errResponse := findCategoryByID(id)
	if category == nil {
		return c.JSON(status, errResponse)
	}

	database.DB.Delete(&category)
	return c.JSON(http.StatusOK, map[string]string{"message": "Category deleted successfully"})
}

func UpdateCategoryById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": commons.InvalidIdError})
	}

	category, status, errResponse := findCategoryByID(id)
	if category == nil {
		return c.JSON(status, errResponse)
	}

	updatedCategory := new(models.Category)
	if err := c.Bind(updatedCategory); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": commons.InvalidRequestError})
	}
	database.DB.Model(&category).Updates(updatedCategory)
	return c.JSON(http.StatusOK, category)
}

func findCategoryByID(id int) (*models.Category, int, map[string]string) {
	var category models.Category
	if err := database.DB.First(&category, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, http.StatusNotFound, map[string]string{"error": commons.CategoryNotFoundError}
		}
		return nil, http.StatusInternalServerError, map[string]string{"error": commons.DatabaseError}
	}
	return &category, http.StatusOK, nil
}
