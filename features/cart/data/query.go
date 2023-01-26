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
	fmt.Println("======data2=====")
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
func (cq *cartQuery) GetByIdC(idUser uint, idCart uint) (cart.CoreCart, error) {
	var sementara Carts
	sementara.UserID = idUser
	sementara.ID = idCart
	fmt.Println("======dt1=====")
	if err := cq.db.Preload("Product").Find(&sementara).Error; err != nil {
		log.Println("Get By ID query error", err.Error())
		return ToCores(sementara), err
	}
	X := ToCores(sementara)
	fmt.Println("======dt2=====")
	return X, nil
}
func (cq *cartQuery) GetByIdU(idUser uint) ([]cart.CoreCart, error) {
	var sementara []Carts
	// sementara.UserID = idUser
	fmt.Println("======dt1=====")
	if err := cq.db.Preload("Product").Where("user_id = ?", idUser).Find(&sementara).Error; err != nil {
		log.Println("Get By ID query error", err.Error())
		return ToCoresArr(sementara), err
	}
	X := ToCoresArr(sementara)
	fmt.Println("======dt2=====")
	return X, nil
}
func (cq *cartQuery) Update(idUser uint, idCart uint, updatedData cart.CoreCart) (cart.CoreCart, error) {
	cnv := CoreToData(updatedData)
	qry := cq.db.Where("user_id = ? AND id = ?", idUser, idCart).Updates(&cnv)
	if qry.RowsAffected <= 0 {
		log.Println("update product query error : data not found")
		return cart.CoreCart{}, errors.New("not found")
	}

	if err := qry.Error; err != nil {
		log.Println("update cart query error :", err.Error())
		return cart.CoreCart{}, err
	}
	// cnv.ID = id
	Y := ToCores(cnv)
	return Y, nil
}
func (cq *cartQuery) Delete(idUser uint, idCart uint) error {
	var sementara Carts
	if err := cq.db.Where("user_id = ? AND id = ?", idUser, idCart).Delete(&sementara).Error; err != nil {
		log.Println("Get By ID query error", err.Error())
		return err
	}
	return nil
}
