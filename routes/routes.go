package routes

import (
	"e-comm-api/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, productController *controllers.ProductController) {

	api := r.Group("/api")
	{
		api.GET("/products/categories", productController.GetCategories)
		api.GET("/products/category/:category", productController.GetProductsByCategory)
		api.GET("/products", productController.GetProducts)
		api.GET("/products/:id", productController.GetProduct)
		api.POST("/products", productController.CreateProduct)
		api.PUT("/products/:id", productController.UpdateProduct)
		api.DELETE("/products/:id", productController.DeleteProduct)
	}

	r.NoRoute(func(c *gin.Context) {
		c.File("./static/index.html")
	})
}
