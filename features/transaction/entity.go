package transaction

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID           uint
	SellerName   string
	BuyerName    string
	Address      string
	ProductName  string
	ProductPrice uint
	Qty          uint
	TotalPrice   uint
	PaymentUrl   string
	Status       string
	UpdatedAt    time.Time
	UserID       uint
}

type TrxHandler interface {
	Add() echo.HandlerFunc
	// BuyHistory() echo.HandlerFunc
	// SellHistory() echo.HandlerFunc
}

type TrxService interface {
	Add(cartID uint, token interface{}, newTrx Core) error
	// BuyHistory(token interface{}) ([]Core, error)
	// SellHistory(token interface{}) ([]Core, error)
}

type TrxData interface {
	Add(cartID uint, userID uint, newTrx Core) error
	// BuyHistory(userID uint) ([]Core, error)
	// SellHistory(userID uint) ([]Core, error)
}
