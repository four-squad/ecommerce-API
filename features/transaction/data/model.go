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
