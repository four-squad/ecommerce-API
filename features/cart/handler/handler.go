package handler

import (
	"ecommerce/features/cart"
	"ecommerce/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type cartControll struct {
	srv cart.CartService
}

func New(srv cart.CartService) cart.CartHandler {
	return &cartControll{
		srv: srv,
	}
}

func (cc *cartControll) Add() echo.HandlerFunc {
	return func(c echo.Context) error {

		token := c.Get("user")
		idProduct, errBind := strconv.Atoi(c.Param("id"))
		if errBind != nil {
			return c.JSON(helper.PrintErrorResponse("Data not found"))
		}
		err := cc.srv.Add(token, uint(idProduct))
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}

		return c.JSON(http.StatusOK, helper.SuccessResponse("Add cart berhasil"))
	}
}
