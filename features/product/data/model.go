package data

import (
	c "ecommerce/features/cart/data"
	"ecommerce/features/product"

	"gorm.io/gorm"
)

type Products struct {
	gorm.Model
	Title       string
	Price       uint
	Description string
	Image       string
	UserID      uint
	// User        User
	Cartss []c.Carts `gorm:"foreignKey:ProductID"`
}
type User struct {
	gorm.Model
	Avatar    string
	Name      string
	Email     string
	Address   string
	Password  string
	Productss []Products
	// UserID    uint
}

func ToCores(data Products) product.CoreProduct {
	return product.CoreProduct{
		ID:          data.ID,
		Title:       data.Title,
		Price:       data.Price,
		Description: data.Description,
		Image:       data.Image,
		// UserID:      data.User.ID,
		// Seller:      data.User.Name,
	}
}

func CoreToData(data product.CoreProduct) Products {
	return Products{
		Model:       gorm.Model{ID: data.ID},
		Title:       data.Title,
		Price:       data.Price,
		Description: data.Description,
		Image:       data.Image,
		UserID:      data.UserID,
		// Seller:      data.User.Name,
	}
}
func ToCoresArr(data []Products) []product.CoreProduct {
	arrRes := []product.CoreProduct{}
	for _, v := range data {
		tmp := ToCores(v)
		arrRes = append(arrRes, tmp)
	}
	return arrRes
}
