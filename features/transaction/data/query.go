package data

import (
	"ecommerce/features/transaction"
	"log"

	"gorm.io/gorm"
)

type trxQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) transaction.TrxData {
	return &trxQuery{
		db: db,
	}
}

func (tq *trxQuery) Add(cartID uint, userID uint, newTrx transaction.Core) error {
	cnv := CoreToData(newTrx)

	cnv.UserID = userID

	err := tq.db.Create(&cnv).Error
	if err != nil {
		log.Println("Create query error : ", err.Error())
		return err
	}

	newTrx.ID = cnv.ID

	return nil
}
