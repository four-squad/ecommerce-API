package data

import (
	"ecommerce/features/cart"
	"errors"
	"fmt"
	"log"
	"strings"

	"gorm.io/gorm"
)

type cartQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) cart.CartData {
	return &cartQuery{
		db: db,
	}
}

func (cq *cartQuery) Add(idUser uint, idProduct uint) error {
	cnv := Carts{}
	cnv.UserID = idUser
	cnv.ProductID = idProduct
	fmt.Println("======data1=====")
	err := cq.db.Create(&cnv).Error
	// err2 := cq.db.Model(&X{}).Select("name").Find(&cnv).Error

	if err != nil {
		log.Println("register query error", err.Error())
		msg := ""
		if strings.Contains(err.Error(), "Duplicate") {
			msg = "data is duplicated"
		} else {
			msg = "data tidak bisa diolah"
		}
		return errors.New(msg)
	}

	return nil
}
