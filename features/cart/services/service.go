package services

import (
	"ecommerce/features/cart"
	"ecommerce/helper"
	"errors"
	"fmt"
	"strings"
)

type cartUseCase struct {
	qry cart.CartData
}

func New(cd cart.CartData) cart.CartService {
	return &cartUseCase{
		qry: cd,
	}
}

func (cuu *cartUseCase) Add(token interface{}, idProduct uint) error {
	fmt.Println("======token=====")
	idUser := helper.ExtractToken(token)
	fmt.Println("======srv1=====")
	err := cuu.qry.Add(uint(idUser), idProduct)
	fmt.Println("======srv2=====")
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "Please input correctly"
		} else {
			msg = "There is a problem with the server"
		}
		return errors.New(msg)
	}

	return nil

}
func (cuu *cartUseCase) GetByIdC(token interface{}, idCart uint) (cart.CoreCart, error) {
	fmt.Println("======srv1=====")
	idUser := helper.ExtractToken(token)
	res, err := cuu.qry.GetByIdC(uint(idUser), idCart)
	fmt.Println("======srv2=====")
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data not found"
		} else {
			msg = "There is a problem with the server"
		}
		return cart.CoreCart{}, errors.New(msg)
	}

	return res, nil
}
