package handler

import (
	"ecommerce/features/product"
	"ecommerce/helper"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type productControll struct {
	srv product.ProductService
}

func New(srv product.ProductService) product.ProductHandler {
	return &productControll{
		srv: srv,
	}
}

func (pc *productControll) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := pc.srv.GetAll()
		if err != nil {
			log.Println("trouble :  ", err.Error())
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		return c.JSON(helper.PrintSuccessReponse(http.StatusOK, "berhasil menampilkan semua product", ToResponseArr(res)))
	}
}

func (pc *productControll) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")
		input := AddReq{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "kesalahan input")
		}
		newProduct := ToCore(input)
		//-----------
		// Read file
		//-----------
		file, err := c.FormFile("image")
		if err != nil {
			file = nil
		}
		fmt.Println("====handler======")
		_, err = pc.srv.Add(*newProduct, token, file)
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		fmt.Println("====handler2======")
		return c.JSON(http.StatusOK, helper.SuccessResponse("posting product berhasil"))

	}
}

func (pc *productControll) GetById() echo.HandlerFunc {
	return func(c echo.Context) error {
		// id := c.QueryParam("idPost")
		// idPost, err := strconv.Atoi(id)
		// if err != nil {
		// 	log.Println("trouble convert param id post:  ", err.Error())
		// 	return c.JSON(helper.PrintErrorResponse(err.Error()))
		// }
		idProduct, errBind := strconv.Atoi(c.Param("id"))
		if errBind != nil {
			return c.JSON(helper.PrintErrorResponse("Data not found"))
		}

		res, err2 := pc.srv.GetById(uint(idProduct))
		if err2 != nil {
			return c.JSON(helper.PrintErrorResponse(err2.Error()))
		}
		return c.JSON(helper.PrintSuccessReponse(http.StatusOK, "berhasil menampilkan product", ToResponseArr(res)))
	}
}

func (pc *productControll) Update() echo.HandlerFunc {
	return func(c echo.Context) error {

		tes, errBind := strconv.Atoi(c.Param("id"))
		if errBind != nil {
			return c.JSON(helper.PrintErrorResponse("Error binding data"))
		}
		token := c.Get("user")
		tmp := AddReq{}
		if err := c.Bind(&tmp); err != nil {
			return c.JSON(http.StatusBadRequest, "kesalahan input")
		}
		//-----------
		// Read file
		//-----------
		file, err := c.FormFile("image")
		if err != nil {
			file = nil
		}
		_, err2 := pc.srv.Update(token, uint(tes), *ToCore(tmp), file)
		if err2 != nil {
			return c.JSON(helper.PrintErrorResponse(err2.Error()))
		}
		return c.JSON(helper.PrintSuccessReponse(http.StatusOK, "update berhasil"))
	}
}

func (pc *productControll) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		tes, errBind := strconv.Atoi(c.Param("id"))
		if errBind != nil {
			return c.JSON(helper.PrintErrorResponse("Kesalahan input"))
		}
		token := c.Get("user")
		err2 := pc.srv.Delete(token, uint(tes))
		if err2 != nil {
			return c.JSON(helper.PrintErrorResponse(err2.Error()))
		}
		return c.JSON(helper.PrintSuccessReponse(http.StatusOK, "delete product berhasil"))
	}
}
