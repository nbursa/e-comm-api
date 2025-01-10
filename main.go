package main

import (
	"e-comm-backend/controllers"
	"e-comm-backend/database"
	"e-comm-backend/models"
	"e-comm-backend/routes"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth_gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
    db, _ := gorm.Open(sqlite.Open("shop.db"), &gorm.Config{})
    db.AutoMigrate(&models.Product{})

		// Seed the database
    database.SeedDatabase(db)

    r := gin.Default()
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        AllowCredentials: false,
    }))

		r.Static("/static", "./static")

		// Rate limiting middleware
    limiter := tollbooth.NewLimiter(1, nil) // 1 request per second
    r.Use(tollbooth_gin.LimitHandler(limiter))

    productController := &controllers.ProductController{DB: db}
    routes.RegisterRoutes(r, productController)

    r.Run()
}