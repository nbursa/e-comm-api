package main

import (
	"e-comm-backend/controllers"
	"e-comm-backend/database"
	"e-comm-backend/models"
	"e-comm-backend/routes"
	"os"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth_gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
		if os.Getenv("GO_ENV") != "production" {
			err := godotenv.Load()
			if err != nil {
					panic("Error loading .env file")
			}
		}

    prodOrigin := os.Getenv("CORS_ORIGIN")
    devOrigin := os.Getenv("DEV_ORIGIN")

		var allowedOrigins []string
    if os.Getenv("GO_ENV") == "production" {
        allowedOrigins = []string{prodOrigin}
    } else {
        allowedOrigins = []string{devOrigin}
    }

    db, err := gorm.Open(sqlite.Open("shop.db"), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
    db.AutoMigrate(&models.Product{})

		// Seed the database
    database.SeedDatabase(db)

    r := gin.Default()
    r.Use(cors.New(cors.Config{
        AllowOrigins:     allowedOrigins,
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


		port := os.Getenv("PORT")
    if port == "" {
        port = "8080" 
    }

    r.Run(":" + port)
}