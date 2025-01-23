package routes

import (
	"e-comm-api/controllers"
	"e-comm-api/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, productController *controllers.ProductController, userController *controllers.UserController, db *gorm.DB) {

	api := r.Group("/api")
	{
		api.POST("/auth/register", userController.RegisterUser)
		api.POST("/auth/login", userController.LoginUser)
		api.POST("/auth/reset-password", userController.ResetPassword)
		api.POST("/auth/complete-reset-password", userController.CompleteResetPassword) 
		api.GET("/products/categories", productController.GetCategories)
		api.GET("/products/category/:category", productController.GetProductsByCategory)
		api.GET("/products", productController.GetProducts)
		api.GET("/products/:id", productController.GetProduct)
	}

	api.Use(middlewares.AuthMiddleware(middlewares.ValidateToken(db)))
	{
		api.POST("/products", productController.CreateProduct)
		api.PUT("/products/:id", productController.UpdateProduct)
		api.DELETE("/products/:id", productController.DeleteProduct)
		api.GET("/auth/user", userController.GetUser)
		api.PUT("/auth/user", userController.UpdateUser)
		api.PUT("/auth/change-password", userController.ChangePassword)
	}

	r.NoRoute(func(c *gin.Context) {
		c.File("./static/index.html")
	})
}
