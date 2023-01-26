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
func (cuu *cartUseCase) GetByIdC(token interface{}, idCart uint) (cart.CoreCart, error) {
	idUser := helper.ExtractToken(token)
	res, err := cuu.qry.GetByIdC(uint(idUser), idCart)
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
func (cuu *cartUseCase) GetByIdU(token interface{}) ([]cart.CoreCart, error) {
	idUser := helper.ExtractToken(token)
	res, err := cuu.qry.GetByIdU(uint(idUser))
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data not found"
		} else {
			msg = "There is a problem with the server"
		}
		return []cart.CoreCart{}, errors.New(msg)
	}

	return res, nil
}
func (cuu *cartUseCase) Update(token interface{}, idCart uint, updatedData cart.CoreCart) (cart.CoreCart, error) {
	idUser := helper.ExtractToken(token)

	res, err := cuu.qry.Update(uint(idUser), idCart, updatedData)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "Please input correctly"
		} else {
			msg = "There is a problem with the server"
		}
		return cart.CoreCart{}, errors.New(msg)
	}

	return res, nil
}
func (cuu *cartUseCase) Delete(token interface{}, idCart uint) error {
	idUser := helper.ExtractToken(token)
	err := cuu.qry.Delete(uint(idUser), idCart)
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
