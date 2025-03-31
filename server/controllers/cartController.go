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

func GetAllCarts(c echo.Context) error {
	var carts []models.Cart
	database.DB.Preload("Products").Find(&carts)
	return c.JSON(http.StatusOK, carts)
}

func GetCartById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": commons.InvalidIdError})
	}

	var cart models.Cart
	if err := database.DB.Preload("Products").First(&cart, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"error": commons.CartNotFoundError})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": commons.CartNotFoundError})
	}

	return c.JSON(http.StatusOK, cart)
}

func CreateCart(c echo.Context) error {
	cart := new(models.Cart)
	if err := c.Bind(cart); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": commons.InvalidRequestError})
	}
	database.DB.Create(cart)
	return c.JSON(http.StatusCreated, cart)
}

func DeleteCart(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": commons.InvalidIdError})
	}

	cart, status, errResponse := findCartByID(id)
	if cart == nil {
		return c.JSON(status, errResponse)
	}

	database.DB.Delete(&cart)
	return c.JSON(http.StatusOK, map[string]string{"message": "Cart deleted successfully"})
}

func UpdateCartById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": commons.InvalidIdError})
	}

	cart, status, errResponse := findCartByID(id)
	if cart == nil {
		return c.JSON(status, errResponse)
	}

	updatedCart := new(models.Cart)
	if err := c.Bind(updatedCart); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": commons.InvalidRequestError})
	}
	database.DB.Model(&cart).Updates(updatedCart)
	return c.JSON(http.StatusOK, cart)
}

func findCartByID(id int) (*models.Cart, int, map[string]string) {
	var cart models.Cart
	if err := database.DB.First(&cart, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, http.StatusNotFound, map[string]string{"error": commons.CartNotFoundError}
		}
		return nil, http.StatusInternalServerError, map[string]string{"error": commons.DatabaseError}
	}
	return &cart, http.StatusOK, nil
}
