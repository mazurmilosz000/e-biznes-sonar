package controllers

import (
	"api-project/models"
	"net/http"

	"api-project/database"

	"github.com/labstack/echo/v4"
)

type CheckoutRequest struct {
	Items []struct {
		ProductID uint    `json:"product_id"`
		Name      string  `json:"name"`
		Price     float64 `json:"price"`
		Quantity  int     `json:"quantity"`
	} `json:"items"`
}

func Checkout(c echo.Context) error {

	var req CheckoutRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	var totalPrice float64
	for _, item := range req.Items {
		totalPrice += item.Price * float64(item.Quantity)
	}

	order := models.Order{
		TotalPrice: totalPrice,
	}
	database.DB.Create(&order)

	var orderItems []models.OrderItem
	for _, item := range req.Items {
		orderItems = append(orderItems, models.OrderItem{
			OrderID:   order.ID,
			ProductID: item.ProductID,
			Name:      item.Name,
			Price:     item.Price,
			Quantity:  item.Quantity,
		})
	}
	database.DB.Create(&orderItems)

	return c.JSON(http.StatusOK, map[string]string{"message": "Order placed successfully"})
}
