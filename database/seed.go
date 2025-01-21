package database

import (
	"e-comm-api/models"
	"encoding/json"
	"fmt"
	"os"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

func SeedDatabase(db *gorm.DB) {
	// Check if the database is empty and seed it with mock data
	var count int64
	db.Model(&models.Product{}).Count(&count)
	fmt.Printf("Product count before seeding: %d\n", count)

	if count == 0 {
		// Read products from JSON file
		file, err := os.ReadFile("products.json")
		if err != nil {
			fmt.Printf("Error reading products.json: %v\n", err)
			return
		}

		var products []models.Product
		if err := json.Unmarshal(file, &products); err != nil {
			fmt.Printf("Error unmarshalling products.json: %v\n", err)
			return
		}

		// Convert AdditionalImages to datatypes.JSON
		for i := range products {
			images, err := json.Marshal(products[i].AdditionalImages)
			if err != nil {
				fmt.Printf("Error marshalling additionalImages: %v\n", err)
				return
			}
			var additionalImages pq.StringArray
			if err := json.Unmarshal(images, &additionalImages); err != nil {
				fmt.Printf("Error unmarshalling additionalImages: %v\n", err)
				return
			}
			products[i].AdditionalImages = additionalImages
		}

		// Insert products one by one to identify the problematic entry
		for _, product := range products {
			fmt.Printf("Inserting product: %+v\n", product)
			result := db.Create(&product)
			if result.Error != nil {
				fmt.Printf("Error seeding product: %v\n", result.Error)
				fmt.Printf("Product: %+v\n", product)
				return
			}
		}

		fmt.Printf("Seeded %d products\n", len(products))
	}

	// Verify the product count after seeding
	db.Model(&models.Product{}).Count(&count)
	fmt.Printf("Product count after seeding: %d\n", count)
}
