package model

type Product struct {
	ProductId   string   `json:"productId" binding:"required"`
	ProductName string   `json:"productName" binding:"required"`
	Category    Category `json:"category" binding:"required"`
	IsActive    bool     `json:"is_active" binding:"required"`
}
