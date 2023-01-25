package data

import (
	"ecommerce/features/cart"

	"gorm.io/gorm"
)

type Carts struct {
	gorm.Model
	Qty         uint
	Total_price uint
	UserID      uint
	// User        User
	ProductID uint
	Product   Product
}

// type X struct {
// 	Title string
// 	Price uint
// 	Image string
// }

type Product struct {
	gorm.Model
	Title       string
	Price       uint
	Description string
	Image       string
}

// type User struct {
// 	gorm.Model
// 	Avatar   string
// 	Name     string
// 	Email    string
// 	Address  string
// 	Password string
// }

func ToCores(data Carts) cart.CoreCart {
	return cart.CoreCart{
		ID:          data.ID,
		Title:       data.Product.Title,
		Qty:         data.Qty,
		Price:       data.Product.Price,
		Image:       data.Product.Image,
		Total_Price: data.Total_price,
		UserID:      data.UserID,
		ProductID:   data.Product.ID,
	}
}
func ToCoresArr(data []Carts) []cart.CoreCart {
	arrRes := []cart.CoreCart{}
	for _, v := range data {
		tmp := ToCores(v)
		arrRes = append(arrRes, tmp)
	}
	return arrRes
}

func CoreToData(data cart.CoreCart) Carts {
	return Carts{
		Model:       gorm.Model{ID: data.ID},
		Qty:         data.Qty,
		Total_price: data.Total_Price,
		UserID:      data.UserID,
		ProductID:   data.ProductID,
	}
}
