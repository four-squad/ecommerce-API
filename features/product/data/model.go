package data

import (
	"ecommerce/features/product"

	"gorm.io/gorm"
)

type Products struct {
	gorm.Model
	Title       string
	Price       uint
	Description string
	Image       string
}

func ToCores(data Products) product.CoreProduct {
	return product.CoreProduct{
		ID:          data.ID,
		Price:       data.Price,
		Description: data.Description,
		Image:       data.Image,
	}
}

func ToCoresArr(data []Products) []product.CoreProduct {
	arrRes := []product.CoreProduct{}
	for _, v := range data {
		tmp := ToCores(v)
		arrRes = append(arrRes, tmp)
	}
	return arrRes
}

func CoreToData(data product.CoreProduct) Products {
	return Products{
		Model:       gorm.Model{ID: data.ID},
		Title:       data.Title,
		Price:       data.Price,
		Description: data.Description,
		Image:       data.Image,
	}
}
