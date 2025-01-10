package database

import (
	"e-comm-backend/models"
	"encoding/json"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func SeedDatabase(db *gorm.DB) {
	err := godotenv.Load()
    if err != nil {
        panic("Error loading .env file")
    }
		
	// Check if the database is empty and seed it with mock data
	var count int64
	baseURL := os.Getenv("BASE_URL")
	db.Model(&models.Product{}).Count(&count)
	if count == 0 {
		// Add mock products
		products := []models.Product{
			{Name: "Smart TV", Price: 1000, Quantity: 5, Description: "4K smart TV", Category: "Electronics", Image: baseURL + "/static/images/smarttv.webp"},
		{Name: "Gaming Laptop", Price: 2000, Quantity: 3, Description: "High-end gaming laptop", Category: "Electronics", Image: baseURL + "/static/images/gaminglaptop.webp"},
		{Name: "Running Shoes", Price: 150, Quantity: 10, Description: "Lightweight running shoes", Category: "Clothing", Image: baseURL + "/static/images/runningshoes.webp"},
		{Name: "Sofa", Price: 800, Quantity: 2, Description: "Leather sofa", Category: "Furniture", Image: baseURL + "/static/images/sofa.webp"},
		{Name: "Refrigerator", Price: 1200, Quantity: 5, Description: "Large refrigerator", Category: "Home Appliances", Image: baseURL + "/static/images/refrigerator.webp"},
		{Name: "Dining Table", Price: 500, Quantity: 3, Description: "Wooden dining table", Category: "Furniture", Image: baseURL + "/static/images/diningtable.webp"},
		{Name: "Laptop", Price: 1200, Quantity: 10, Description: "High-performance laptop", Category: "Electronics", Image: baseURL + "/static/images/laptop.webp"},
		{Name: "Smartphone", Price: 800, Quantity: 20, Description: "Latest smartphone", Category: "Electronics", Image: baseURL + "/static/images/smartphone.webp"},
		{Name: "T-Shirt", Price: 20, Quantity: 50, Description: "Cotton t-shirt", Category: "Clothing", Image: baseURL + "/static/images/tshirt.webp"},
		{Name: "Sneakers", Price: 100, Quantity: 15, Description: "Comfortable running shoes", Category: "Clothing", Image: baseURL + "/static/images/sneakers.webp"},
		{Name: "Blender", Price: 50, Quantity: 10, Description: "Kitchen blender", Category: "Home Appliances", Image: baseURL + "/static/images/blender.webp"},
		{Name: "Chair", Price: 75, Quantity: 5, Description: "Office chair", Category: "Furniture", Image: baseURL + "/static/images/chair.webp"},
		{Name: "Headphones", Price: 150, Quantity: 25, Description: "Noise-cancelling headphones", Category: "Electronics", Image: baseURL + "/static/images/headphones.webp"},
		{Name: "Keyboard", Price: 40, Quantity: 30, Description: "Mechanical keyboard", Category: "Electronics", Image: baseURL + "/static/images/keyboard.webp"},
		{Name: "Desk Lamp", Price: 35, Quantity: 10, Description: "LED desk lamp", Category: "Home Appliances", Image: baseURL + "/static/images/desklamp.webp"},
		{Name: "Backpack", Price: 60, Quantity: 20, Description: "Durable travel backpack", Category: "Accessories", Image: baseURL + "/static/images/backpack.jpg"},
		{Name: "Wristwatch", Price: 200, Quantity: 15, Description: "Luxury wristwatch", Category: "Accessories", Image: baseURL + "/static/images/wristwatch.webp"},
		{Name: "Coffee Maker", Price: 100, Quantity: 8, Description: "Programmable coffee maker", Category: "Home Appliances", Image: baseURL + "/static/images/coffeemaker.webp"},
		{Name: "Microwave", Price: 150, Quantity: 5, Description: "Compact microwave oven", Category: "Home Appliances", Image: baseURL + "/static/images/microwave.webp"},
		{Name: "Running Shorts", Price: 25, Quantity: 40, Description: "Breathable running shorts", Category: "Clothing", Image: baseURL + "/static/images/runningshorts.webp"},
		{Name: "Gaming Mouse", Price: 50, Quantity: 30, Description: "High-precision gaming mouse", Category: "Electronics", Image: baseURL + "/static/images/gamingmouse.webp"},
		{Name: "Tablet", Price: 500, Quantity: 10, Description: "High-resolution tablet", Category: "Electronics", Image: baseURL + "/static/images/tablet.webp"},
		{Name: "Water Bottle", Price: 20, Quantity: 50, Description: "Insulated stainless steel water bottle", Category: "Accessories", Image: baseURL + "/static/images/waterbottle.webp"},
		{Name: "Yoga Mat", Price: 30, Quantity: 25, Description: "Non-slip yoga mat", Category: "Fitness", Image: baseURL + "/static/images/yogamat.webp"},
		{Name: "Dumbbells", Price: 100, Quantity: 15, Description: "Adjustable dumbbells", Category: "Fitness", Image: baseURL + "/static/images/dumbbells.webp"},
		{Name: "Sunglasses", Price: 120, Quantity: 20, Description: "Polarized sunglasses", Category: "Accessories", Image: baseURL + "/static/images/sunglasses.webp"},
		{Name: "Tennis Racket", Price: 150, Quantity: 10, Description: "Lightweight tennis racket", Category: "Sports", Image: baseURL + "/static/images/tennisracket.webp"},
		{Name: "Soccer Ball", Price: 30, Quantity: 20, Description: "Official size soccer ball", Category: "Sports", Image: baseURL + "/static/images/soccerball.webp"},
		{Name: "Basketball", Price: 40, Quantity: 15, Description: "Indoor/outdoor basketball", Category: "Sports", Image: baseURL + "/static/images/basketball.webp"},
		{Name: "Running Jacket", Price: 75, Quantity: 10, Description: "Waterproof running jacket", Category: "Clothing", Image: baseURL + "/static/images/runningjacket.webp"},
		}

		// Save products to the database
		db.Create(&products)

		// Create JSON file
		file, err := os.Create("products.json")
		if err != nil {
				panic(err)
		}
		defer file.Close()

		// Marshal products to JSON
		jsonData, err := json.MarshalIndent(products, "", "  ")
		if err != nil {
				panic(err)
		}

		// Write JSON data to file
		_, err = file.Write(jsonData)
		if err != nil {
				panic(err)
		}
	}
}