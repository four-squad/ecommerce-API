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

	getCart := map[string]interface{}{}
	if err := tq.db.Raw("SELECT products.title, products.price, carts.qty FROM carts JOIN poducts ON carts.product_id = products.id WHERE carts.id = ?", cartID).Find(&getCart).Error; err != nil {
		log.Println("Get User Content by username query error : ", err.Error())
		return err
	}

	// if err := tq.db.Raw("INSERT into transaction_details(transaction_id, product_id, title, price, qty, total_price) VALUE()", cartID).Error; err != nil {
	// 	log.Println("Get User Content by username query error : ", err.Error())
	// 	return err
	// }

	newTrx.ID = cnv.ID

	return nil
}
