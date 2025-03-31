package main

import (
	"api-project/controllers"
	"api-project/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const idRequest = "/:id"

const productRequest = "/products"
const productIdRequest = productRequest + idRequest

const categoryRequest = "/categories"
const categoryIdRequest = categoryRequest + idRequest

const cartRequest = "/cart"
const cartIdRequest = cartRequest + idRequest

func main() {
	e := echo.New()

	database.InitDB()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173", "http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
	}))

	e.GET(productRequest, controllers.GetAllProducts)
	e.GET(productIdRequest, controllers.GetProductById)
	e.POST(productRequest, controllers.CreateProduct)
	e.PUT(productIdRequest, controllers.UpdateProductById)
	e.DELETE(productIdRequest, controllers.DeleteProduct)

	e.GET(categoryRequest, controllers.GetAllCategories)
	e.GET(categoryIdRequest, controllers.GetCategoryById)
	e.POST(categoryRequest, controllers.CreateCategory)
	e.PUT(categoryIdRequest, controllers.UpdateCategoryById)
	e.DELETE(categoryIdRequest, controllers.DeleteCategory)

	e.GET(cartRequest, controllers.GetAllCategories)
	e.GET(cartIdRequest, controllers.GetCategoryById)
	e.POST(cartRequest, controllers.CreateCategory)
	e.PUT(cartIdRequest, controllers.UpdateCategoryById)
	e.DELETE(cartIdRequest, controllers.DeleteCategory)

	e.POST("/checkout", controllers.Checkout)

	e.Logger.Fatal(e.Start(":8080"))
}
