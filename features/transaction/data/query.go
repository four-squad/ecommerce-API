package data

import (
	"ecommerce/features/transaction"
	"errors"
	"fmt"
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
	fmt.Println("qry")
	cnv := CoreToData(newTrx)

	getCart := map[string]interface{}{}
	if err := tq.db.Raw("SELECT products.id, products.title, products.price, carts.qty FROM carts JOIN products ON carts.product_id = products.id WHERE carts.id = ?", cartID).Find(&getCart).Error; err != nil {
		log.Println("Get cart query error : ", err.Error())
		return err
	}

	if getCart["price"] == nil {
		return errors.New("Cart not found")
	}

	fmt.Println(getCart)

	price := getCart["price"].(int64)
	qty := getCart["qty"].(int64)

	// priceInt, _ := strconv.Atoi(price)
	// qtyInt, _ := strconv.Atoi(qty)

	totalPrice := price * qty

	cnv.UserID = userID
	cnv.TotalPrice = uint(totalPrice)

	err := tq.db.Create(&cnv).Error
	if err != nil {
		log.Println("Create query error : ", err.Error())
		return err
	}

	trxID := cnv.ID
	fmt.Println("trx id", trxID)

	if err := tq.db.Raw("INSERT into transaction_details(transaction_id, product_id, title, price, qty, total_price) VALUE(?,?,?,?,?,?)", trxID, getCart["id"], getCart["title"], uint(price), uint(qty), uint(totalPrice)).Error; err != nil {
		log.Println("Insert new transaction detail query error : ", err.Error())
		return err
	}

	fmt.Println("query good")

	return nil
}
