package model

type Product struct {
	ProductId   string `json:"productId" binding:"required"`
	ProductName string `json:"productName" binding:"required"`
	IsStatus    bool   `json:"isStatus" gorm:"default:true"`
	UrlPath     string
	ImgPath     string `json:"imgPath,omitempty"` // will be used for miltiport form type
}
