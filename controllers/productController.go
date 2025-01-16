package controllers

import (
	"e-comm-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductController struct {
    DB *gorm.DB
}

func (pc *ProductController) GetCategories(c *gin.Context) {
	var categories []string
	if err := pc.DB.Model(&models.Product{}).Distinct().Pluck("Category", &categories).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
	}
	c.JSON(http.StatusOK, categories)
}

func (pc *ProductController) GetProductsByCategory(c *gin.Context) {
	category := c.Param("category")

	var products []models.Product
	if err := pc.DB.Where("category = ?", category).Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}

	c.JSON(http.StatusOK, products)
}

func (pc *ProductController) GetProducts(c *gin.Context) {
    var products []models.Product
    if err := pc.DB.Find(&products).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, products)
}

// func (pc *ProductController) GetProduct(c *gin.Context) {
//     var product models.Product
//     id := c.Param("id")
//     if err := pc.DB.First(&product, id).Error; err != nil {
//         c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
//         return
//     }
//     c.JSON(http.StatusOK, product)
// }
func (pc *ProductController) GetProduct(c *gin.Context) {
	// Get ID from URL parameter
	id := c.Param("id")

	// Validate ID
	if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
			return
	}

	// Query database for product
	var product models.Product
	if err := pc.DB.First(&product, id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
					c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
					return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
			return
	}

	// Return success response
	c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"data": product,
	})
}

func (pc *ProductController) CreateProduct(c *gin.Context) {
    var product models.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    pc.DB.Create(&product)
    c.JSON(http.StatusCreated, product)
}

func (pc *ProductController) UpdateProduct(c *gin.Context) {
    var product models.Product
    id := c.Param("id")
    if err := pc.DB.First(&product, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    pc.DB.Save(&product)
    c.JSON(http.StatusOK, product)
}

func (pc *ProductController) DeleteProduct(c *gin.Context) {
    var product models.Product
    id := c.Param("id")
    if err := pc.DB.First(&product, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }
    pc.DB.Delete(&product)
    c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}