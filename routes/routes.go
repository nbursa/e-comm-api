package routes

import (
	"e-comm-backend/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, productController *controllers.ProductController) {
    r.OPTIONS("/*path", func(c *gin.Context) {
        c.Header("Access-Control-Allow-Origin", c.Request.Header.Get("Origin"))
        c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
        c.Status(http.StatusNoContent)
    })

    r.GET("/products/categories", productController.GetCategories)
		r.GET("/products", productController.GetProducts)
    r.GET("/products/:id", productController.GetProduct)
		r.POST("/products", productController.CreateProduct)
    r.PUT("/products/:id", productController.UpdateProduct)
    r.DELETE("/products/:id", productController.DeleteProduct)
}