package data

import (
	"ecommerce/features/transaction_detail"

	"gorm.io/gorm"
)

type transaction_detailQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) transaction_detail.Transaction_DetailData {
	return &transaction_detailQuery{
		db: db,
	}
}
