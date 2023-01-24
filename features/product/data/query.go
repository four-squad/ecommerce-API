package data

import (
	"ecommerce/features/product"
	"errors"
	"fmt"
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
	cnv.User.ID = id
	fmt.Println("======data1=====")
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
	newProduct.UserID = cnv.ID
	fmt.Println("======data2=====")
	return ToCores(cnv), nil
}

func (pq *productQuery) GetById(idProduct uint) ([]product.CoreProduct, error) {
	var sementara []Products

	if err := pq.db.Preload("User").Where("id = ?", idProduct).Find(&sementara).Error; err != nil {
		log.Println("Get By ID query error", err.Error())
		return ToCoresArr(sementara), err
	}
	X := ToCoresArr(sementara)
	return X, nil
}

func (pq *productQuery) Update(userId uint, productId uint, updatedData product.CoreProduct) (product.CoreProduct, error) {
	cnv := CoreToData(updatedData)
	qry := pq.db.Where("user_id = ? AND id = ?", userId, productId).Updates(&cnv)
	if qry.RowsAffected <= 0 {
		log.Println("update product query error : data not found")
		return product.CoreProduct{}, errors.New("not found")
	}

	if err := qry.Error; err != nil {
		log.Println("update book query error :", err.Error())
		return product.CoreProduct{}, err
	}
	// cnv.ID = id
	Y := ToCores(cnv)
	return Y, nil
}

func (pq *productQuery) Delete(userId uint, productId uint) error {
	var sementara Products
	if err := pq.db.Where("user_id = ? AND id = ?", userId, productId).Delete(&sementara).Error; err != nil {
		log.Println("Get By ID query error", err.Error())
		return err
	}
	return nil
}
