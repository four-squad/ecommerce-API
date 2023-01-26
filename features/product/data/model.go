package data

import (
	c "ecommerce/features/cart/data"
	"ecommerce/features/product"

	x "ecommerce/features/transaction_detail/data"

	"gorm.io/gorm"
)

type Products struct {
	gorm.Model
	Title             string
	Price             uint
	Description       string
	Image             string
	UserID            uint
	User              Users
	Cartss            []c.Carts              `gorm:"foreignKey:ProductID"`
	TransactionDetail []x.Transaction_Detail `gorm:"foreignKey:ProductID"`
}

type Users struct {
	gorm.Model
	Avatar   string
	Name     string
	Email    string
	Address  string
	Password string
	Cartss   []c.Carts  `gorm:"foreignKey:UserID"`
	Product  []Products `gorm:"foreignKey:UserID"`
	// Transaction []Transactions `gorm:"foreignKey:UserID"`
}

func ToCores(data Products) product.CoreProduct {
	return product.CoreProduct{
		ID:          data.ID,
		Title:       data.Title,
		Price:       data.Price,
		Description: data.Description,
		Image:       data.Image,
		User: product.CoreUser{
			ID:   data.User.ID,
			Name: data.User.Name,
		},
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
