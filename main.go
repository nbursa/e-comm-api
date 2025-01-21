package main

import (
	"e-comm-api/controllers"
	"e-comm-api/database"
	"e-comm-api/models"
	"e-comm-api/routes"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

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

	// Set CORS origins based on environment
	origin := os.Getenv("CORS_ORIGIN")
	var allowedOrigins []string = strings.Split(origin, ",")

	fmt.Printf("[INFO] Allowed Origins: %v\n", allowedOrigins)

	// Database setup
	db, err := gorm.Open(sqlite.Open("shop.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Product{})
	database.SeedDatabase(db)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
		AllowWildcard:    true,
	}))

	limiter := tollbooth.NewLimiter(10, nil)
	r.Use(tollbooth_gin.LimitHandler(limiter))

	// Serve static files with a fallback to placeholder.webp
	r.Use(func(c *gin.Context) {
		c.Next()
		if c.Writer.Status() == http.StatusNotFound && strings.HasPrefix(c.Request.URL.Path, "/api/static/images/") {
			c.File("./static/placeholder.webp")
		}
	})

	r.Static("/api/static/images", "./static/images")

	productController := &controllers.ProductController{DB: db}
	routes.RegisterRoutes(r, productController)

	// SPA fallback
	r.NoRoute(func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/api") {
			c.JSON(404, gin.H{"error": "API route not found"})
			return
		}
		c.File("./static/index.html")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
