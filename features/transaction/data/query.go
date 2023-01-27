package data

import (
	"ecommerce/features/transaction"
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

func (tq *trxQuery) Add(userID uint, newTrx transaction.Core) (transaction.Core, error) {
	cnv := CoreToData(newTrx)

	getCart := []map[string]interface{}{}
	if err := tq.db.Raw("SELECT products.id, products.title, products.price, carts.qty FROM carts JOIN products ON carts.product_id = products.id WHERE carts.user_id = ?", userID).Find(&getCart).Error; err != nil {
		log.Println("Get cart query error : ", err.Error())
		return transaction.Core{}, err
	}

	cnv.Address = newTrx.Address
	cnv.UserID = userID

	err := tq.db.Create(&cnv).Error
	if err != nil {
		log.Println("Create query error : ", err.Error())
		return transaction.Core{}, err
	}

	trxID := cnv.ID

	var productID uint64
	var title string
	var price int64
	var qty int64
	var total int64

	for i := 0; i < len(getCart); i++ {
		for k, v := range getCart[i] {
			switch k {
			case "id":
				productID = v.(uint64)
			case "title":
				title = v.(string)
			case "price":
				price = v.(int64)
			case "qty":
				qty = v.(int64)
			}
		}

		total = total + (price * qty)

		if err := tq.db.Exec("INSERT into transaction_details(transaction_id, product_id, title, price, qty, total_price) VALUE(?,?,?,?,?,?)", trxID, productID, title, price, qty, total).Error; err != nil {
			tq.db.Rollback()
			log.Println("Insert new transaction detail query error : ", err.Error())
			return transaction.Core{}, err
		}
		fmt.Println(trxID, productID, title, price, qty, total)
	}

	if err := tq.db.Exec("UPDATE transactions SET total_price = ? WHERE id = ?", total, trxID).Error; err != nil {
		log.Println("Get cart query error : ", err.Error())
		return transaction.Core{}, err
	}

	cnv.TotalPrice = uint(total)

	return ToCore(cnv), nil
}
