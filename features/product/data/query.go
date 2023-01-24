package data

import (
	"ecommerce/features/product"
	"errors"
	"log"
	"strings"

	"gorm.io/gorm"
)

type productQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) product.ProductData {
	return &productQuery{
		db: db,
	}
}

func (pq *productQuery) GetAll() ([]product.CoreProduct, error) {
	var sementara []Products

	if err := pq.db.Find(&sementara).Error; err != nil {
		log.Println("Get By ID query error", err.Error())
		return ToCoresArr(sementara), err
	}
	Y := ToCoresArr(sementara)
	return Y, nil
}

func (pq *productQuery) Add(newProduct product.CoreProduct, id uint) (product.CoreProduct, error) {
	cnv := CoreToData(newProduct)
	cnv.ID = id

	err := pq.db.Create(&cnv).Error

	if err != nil {
		log.Println("register query error", err.Error())
		msg := ""
		if strings.Contains(err.Error(), "Duplicate") {
			msg = "data is duplicated"
		} else {
			msg = "data tidak bisa diolah"
		}
		return product.CoreProduct{}, errors.New(msg)
	}
	// newBook.Users.ID = cnv.ID
	// newBook.Users.Name = cnv.User.Name

	return newProduct, nil
}

func (pq *productQuery) GetById(idUser uint, idProduct uint) ([]product.CoreProduct, error) {
	var sementara []Products

	if err := pq.db.Where("product_id = ?", idProduct).Find(&sementara).Error; err != nil {
		log.Println("Get By ID query error", err.Error())
		return ToCoresArr(sementara), err
	}
	X := ToCoresArr(sementara)
	return X, nil
}
