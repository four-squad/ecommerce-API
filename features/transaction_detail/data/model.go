package data

import (
	c "ecommerce/features/cart/data"
	"ecommerce/features/transaction_detail"

	"gorm.io/gorm"
)

type Transaction_Detail struct {
	TransactionID uint `gorm:"primaryKey"`
	ProductID     uint `gorm:"primaryKey"`
	Title         string
	Price         uint
	Qty           uint
	TotalPrice    uint
	Transaction   Transactions
	Product       Products
}
type Transactions struct {
	gorm.Model
	TotalPrice uint
	Address    string
	PaymentUrl string
	Status     string `gorm:"default:pending"`
	UserID     uint
	// User              Users
	TransactionDetail []Transaction_Detail `gorm:"foreignKey:TransactionID"`
}
type Products struct {
	gorm.Model
	Title       string
	Price       uint
	Description string
	Image       string
	UserID      uint
	// User              Users
	Cartss            []c.Carts            `gorm:"foreignKey:ProductID"`
	TransactionDetail []Transaction_Detail `gorm:"foreignKey:ProductID"`
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
}

func ToCores(data Transaction_Detail) transaction_detail.CoreTransaction_Detail {
	return transaction_detail.CoreTransaction_Detail{
		ID: data.TransactionID,
		// NameBuyer: data.Product.User.Name,
		// NameSeller:   "",
		ProductName:  data.Product.Title,
		ProductPrice: data.Price,
		Qty:          data.Qty,
		TotalPrice:   data.TotalPrice,
		Image:        data.Product.Image,
		UpdatedAt:    data.Product.UpdatedAt,
	}
}
func ToCores2(data Transaction_Detail) transaction_detail.CoreTransaction_Detail {
	return transaction_detail.CoreTransaction_Detail{
		ID: data.TransactionID,
		// NameBuyer: data.Product.User.Name,
		// NameSeller:   data.Product.User.Name,
		ProductName:  data.Product.Title,
		ProductPrice: data.Price,
		Qty:          data.Qty,
		TotalPrice:   data.TotalPrice,
		Image:        data.Product.Image,
		UpdatedAt:    data.Product.UpdatedAt,
	}
}

func CoreToData(data transaction_detail.CoreTransaction_Detail) Transaction_Detail {
	return Transaction_Detail{
		TransactionID: data.ID,
		ProductID:     data.ProductID,
		Title:         data.ProductName,
		Price:         data.ProductPrice,
		Qty:           data.Qty,
		TotalPrice:    data.TotalPrice,
		Transaction: Transactions{
			Model:      gorm.Model{},
			TotalPrice: 0,
			Address:    "",
			PaymentUrl: "",
			Status:     "",
			UserID:     0,
		},
		Product: Products{
			Model:       gorm.Model{UpdatedAt: data.UpdatedAt},
			Title:       data.ProductName,
			Description: "",
			Image:       data.Image,
			UserID:      0,
			// User: Users{
			// 	Model:    gorm.Model{},
			// 	Avatar:   "",
			// 	Name:     data.NameBuyer,
			// 	Email:    "",
			// 	Address:  "",
			// 	Password: "",
			// },
		},
	}
}

func ToCoresArr(data []Transaction_Detail) []transaction_detail.CoreTransaction_Detail {
	arrRes := []transaction_detail.CoreTransaction_Detail{}
	for _, v := range data {
		tmp := ToCores(v)
		arrRes = append(arrRes, tmp)
	}
	return arrRes
}
