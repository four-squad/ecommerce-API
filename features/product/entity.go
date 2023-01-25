package product

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

type CoreProduct struct {
	ID          uint
	Title       string
	Price       uint
	Description string
	Image       string
	UserID      uint
	// Seller      string
	User CoreUser
}

type CoreUser struct {
	ID   uint
	Name string
}

type ProductHandler interface {
	Add() echo.HandlerFunc
	GetAll() echo.HandlerFunc
	GetById() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type ProductService interface {
	Add(newContent CoreProduct, token interface{}, image *multipart.FileHeader) (CoreProduct, error)
	GetAll() ([]CoreProduct, error)
	GetById(tes uint) ([]CoreProduct, error)
	Update(token interface{}, id uint, updatedData CoreProduct, file *multipart.FileHeader) (CoreProduct, error)
	Delete(token interface{}, contentId uint) error
}

type ProductData interface {
	Add(newContent CoreProduct, id uint) (CoreProduct, error)
	GetAll() ([]CoreProduct, error)
	GetById(tes uint) ([]CoreProduct, error)
	Update(userId uint, contentId uint, updatedData CoreProduct) (CoreProduct, error)
	Delete(userId uint, contentId uint) error
}
