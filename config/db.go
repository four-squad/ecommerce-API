package config

import (
	cart "ecommerce/features/cart/data"
	product "ecommerce/features/product/data"

	trx "ecommerce/features/transaction/data"

	transaction_detail "ecommerce/features/transaction_detail/data"
	user "ecommerce/features/user/data"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(ac AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		ac.DBUser, ac.DBPass, ac.DBHost, ac.DBPort, ac.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Database connection error : ", err.Error())
		return nil
	}

	return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(user.Users{})
	db.AutoMigrate(product.Products{})
	db.AutoMigrate(cart.Carts{})
	db.AutoMigrate(trx.Transactions{})
	db.AutoMigrate(transaction_detail.Transaction_Detail{})
}
