package cart

import "github.com/labstack/echo/v4"

type CoreCart struct {
	ID          uint
	Title       string
	Qty         uint
	Price       uint
	Image       string
	Total_Price uint
	UserID      uint
	ProductID   uint
}

type CartHandler interface {
	Add() echo.HandlerFunc
	GetByIdC() echo.HandlerFunc
	GetByIdU() echo.HandlerFunc
	// Profile() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type CartService interface {
	Add(token interface{}, idProduct uint) error
	GetByIdC(token interface{}, idCart uint) (CoreCart, error)
	GetByIdU(token interface{}) ([]CoreCart, error)
	// Profile(token interface{}) (interface{}, error)
	Update(token interface{}, idCart uint, updatedData CoreCart) (CoreCart, error)
	Delete(token interface{}, idCart uint) error
}

type CartData interface {
	Add(idUser uint, idProduct uint) error
	GetByIdC(idUser uint, idCart uint) (CoreCart, error)
	GetByIdU(idUser uint) ([]CoreCart, error)
	// Profile(id uint) (interface{}, error)
	Update(idUser uint, idCart uint, updatedData CoreCart) (CoreCart, error)
	Delete(idUser uint, idCart uint) error
}
