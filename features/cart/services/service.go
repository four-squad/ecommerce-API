package services

import (
	"ecommerce/features/cart"
	"ecommerce/helper"
	"errors"
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
	idUser := helper.ExtractToken(token)

	err := cuu.qry.Add(uint(idUser), idProduct)

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
