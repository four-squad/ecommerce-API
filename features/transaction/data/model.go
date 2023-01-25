package data

import (
	"ecommerce/features/transaction"

	"gorm.io/gorm"
)

type Transactions struct {
	gorm.Model
	UserID     uint
	TotalPrice uint
	Address    string
	PaymentUrl string
	Status     string
}
type TransactionDetail struct {
	TransactionID uint `gorm:"primary_key"`
	ProductID     uint `gorm:"primary_key"`
	Title         string
	Price         uint
	Qty           uint
	TotalPrice    uint
}

func ToCore(data Transactions) transaction.Core {
	return transaction.Core{
		ID:         data.ID,
		TotalPrice: data.TotalPrice,
		Address:    data.Address,
		PaymentUrl: data.PaymentUrl,
		Status:     data.Status,
	}
}

func CoreToData(data transaction.Core) Transactions {
	return Transactions{
		Model:      gorm.Model{ID: data.ID},
		TotalPrice: data.TotalPrice,
		Address:    data.Address,
		PaymentUrl: data.PaymentUrl,
		Status:     data.Status,
	}
}
