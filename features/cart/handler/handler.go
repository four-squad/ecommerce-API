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
		idProduct, errBind := strconv.Atoi(c.Param("idProduct"))
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
func (cc *cartControll) GetByIdC() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")
		idCart, errBind := strconv.Atoi(c.Param("idCart"))
		if errBind != nil {
			return c.JSON(helper.PrintErrorResponse("Data not found"))
		}
		res, err2 := cc.srv.GetByIdC(token, uint(idCart))
		if err2 != nil {
			return c.JSON(helper.PrintErrorResponse(err2.Error()))
		}
		return c.JSON(helper.PrintSuccessReponse(http.StatusOK, "berhasil menampilkan cart", ToResponse(res)))
	}
}
func (cc *cartControll) GetByIdU() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")

		res, err2 := cc.srv.GetByIdU(token)
		if err2 != nil {
			return c.JSON(helper.PrintErrorResponse(err2.Error()))
		}
		return c.JSON(helper.PrintSuccessReponse(http.StatusOK, "berhasil menampilkan cart", ToResponseArr(res)))
	}
}
func (cc *cartControll) Update() echo.HandlerFunc {
	return func(c echo.Context) error {

		idCart, errBind := strconv.Atoi(c.Param("idCart"))
		if errBind != nil {
			return c.JSON(helper.PrintErrorResponse("Error binding data"))
		}
		token := c.Get("user")
		updatedData := UptReq{}
		if err := c.Bind(&updatedData); err != nil {
			return c.JSON(http.StatusBadRequest, "Invalid user input")
		}

		_, err2 := cc.srv.Update(token, uint(idCart), *ToCore(updatedData))
		if err2 != nil {
			return c.JSON(helper.PrintErrorResponse(err2.Error()))
		}
		return c.JSON(http.StatusOK, helper.SuccessResponse("Checkout berhasil"))
	}
}
func (cc *cartControll) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		idCart, errBind := strconv.Atoi(c.Param("idCart"))
		if errBind != nil {
			return c.JSON(helper.PrintErrorResponse("Invalid user input"))
		}
		token := c.Get("user")
		err2 := cc.srv.Delete(token, uint(idCart))
		if err2 != nil {
			return c.JSON(helper.PrintErrorResponse(err2.Error()))
		}
		return c.JSON(http.StatusOK, helper.SuccessResponse("Delete cart berhasil"))
	}
}
