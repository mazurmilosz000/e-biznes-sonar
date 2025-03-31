package controllers

import (
	"api-project/database"
	"api-project/models"
	"net/http"
	"strconv"

	"api-project/commons"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetAllProducts(c echo.Context) error {
	var products []models.Product
	database.DB.Preload("Category").Find(&products)
	return c.JSON(http.StatusOK, products)
}

func GetProductById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": commons.InvalidIdError})
	}

	var product models.Product
	if err := database.DB.Preload("Category").First(&product, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"error": commons.ProductNotFoundError})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": commons.DatabaseError})
	}

	return c.JSON(http.StatusOK, product)
}

func CreateProduct(c echo.Context) error {
	product := new(models.Product)
	if err := c.Bind(product); err != nil {
		return err
	}

	database.DB.Create(product)
	return c.JSON(http.StatusCreated, product)
}

func DeleteProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": commons.InvalidIdError})
	}

	product, status, errResponse := findProductByID(id)
	if product == nil {
		return c.JSON(status, errResponse)
	}

	database.DB.Delete(&product)
	return c.JSON(http.StatusOK, map[string]string{"message": "Product deleted successfully"})
}

func UpdateProductById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": commons.InvalidIdError})
	}

	product, status, errResponse := findProductByID(id)
	if product == nil {
		return c.JSON(status, errResponse)
	}

	updatedProduct := new(models.Product)
	if err := c.Bind(updatedProduct); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": commons.InvalidRequestError})
	}
	database.DB.Model(&product).Updates(updatedProduct)
	return c.JSON(http.StatusOK, product)
}

func findProductByID(id int) (*models.Product, int, map[string]string) {
	var product models.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, http.StatusNotFound, map[string]string{"error": commons.ProductNotFoundError}
		}
		return nil, http.StatusInternalServerError, map[string]string{"error": commons.DatabaseError}
	}
	return &product, http.StatusOK, nil
}
