package models

import (
	"github.com/lib/pq"
)

type Rating struct {
	Average float64 `json:"average"`
	Count   int     `json:"count"`
}

type Product struct {
	ID               uint     `json:"id" gorm:"primaryKey"`
	Name             string   `json:"name"`
	Price            float64  `json:"price"`
	Quantity         int      `json:"quantity"`
	Description      string   `json:"description"`
	Category         string   `json:"category"`
	Image            string   `json:"image"`
	Title            string   `json:"title,omitempty"`
	Discount         bool     `json:"discount,omitempty"`
	DiscountedPrice  float64  `json:"discountedPrice,omitempty"`
	AdditionalImages pq.StringArray `json:"additionalImages,omitempty" gorm:"type:text[]; serializer:json"`
	Rating           Rating   `json:"rating,omitempty" gorm:"embedded"`
}
