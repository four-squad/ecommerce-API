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
	// Profile() echo.HandlerFunc
	// Update() echo.HandlerFunc
	// Deactivate() echo.HandlerFunc
}

type CartService interface {
	Add(token interface{}, idProduct uint) error
	GetByIdC(token interface{}, idCart uint) (CoreCart, error)
	// Profile(token interface{}) (interface{}, error)
	// Update(formHeader multipart.FileHeader, token interface{}, updatedProfile Core) (Core, error)
	// Deactivate(token interface{}) error
}

type CartData interface {
	Add(idUser uint, idProduct uint) error
	GetByIdC(idUser uint, idCart uint) (CoreCart, error)
	// Profile(id uint) (interface{}, error)
	// Update(id uint, updatedProfile Core) (Core, error)
	// Deactivate(id uint) error
}
