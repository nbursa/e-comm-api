package models

type Product struct {
    ID          uint   `json:"id" gorm:"primaryKey"`
    Name        string `json:"name"`
    Price       int    `json:"price"`
    Quantity    int    `json:"quantity"`
    Description string `json:"description"`
    Category    string `json:"category"`
    Image       string `json:"image"`
}