package controllers

import (
	"e-comm-backend/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductController struct {
	DB *gorm.DB
}

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Code    string `json:"code"`
	Detail  string `json:"detail,omitempty"`
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
	var total int64

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	if err := pc.DB.Model(&models.Product{}).Where("category = ?", category).Count(&total).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error": ErrorResponse{
				Message: "Failed to count products in category",
				Code:    "CATEGORY_PRODUCTS_COUNT_ERROR",
				Detail:  err.Error(),
			},
		})
		return
	}

	if err := pc.DB.Where("category = ?", category).Limit(limit).Offset(offset).Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error": ErrorResponse{
				Message: "Failed to fetch products in category",
				Code:    "CATEGORY_PRODUCTS_FETCH_ERROR",
				Detail:  err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   products,
		"meta": gin.H{
			"total": total,
			"page":  page,
			"limit": limit,
		},
	})
}

func (pc *ProductController) GetProducts(c *gin.Context) {
	var products []models.Product
	var total int64

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	if err := pc.DB.Model(&models.Product{}).Count(&total).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error": ErrorResponse{
				Message: "Failed to count products",
				Code:    "PRODUCTS_COUNT_ERROR",
				Detail:  err.Error(),
			},
		})
		return
	}

	if err := pc.DB.Limit(limit).Offset(offset).Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error": ErrorResponse{
				Message: "Failed to fetch products",
				Code:    "PRODUCTS_FETCH_ERROR",
				Detail:  err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   products,
		"meta": gin.H{
			"total": total,
			"page":  page,
			"limit": limit,
		},
	})
}

func (pc *ProductController) GetProduct(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if _, err := strconv.Atoi(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error": ErrorResponse{
				Message: "Invalid product ID format",
				Code:    "INVALID_PRODUCT_ID_FORMAT",
			},
		})
		return
	}

	if err := pc.DB.First(&product, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"status": "error",
				"error": ErrorResponse{
					Message: "Product not found",
					Code:    "PRODUCT_NOT_FOUND",
					Detail:  fmt.Sprintf("Product with ID %s does not exist", id),
				},
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error": ErrorResponse{
				Message: "Failed to fetch product",
				Code:    "PRODUCT_FETCH_ERROR",
				Detail:  err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   product,
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
