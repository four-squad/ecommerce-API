package handler

import "ecommerce/features/product"

type ProductResponse struct {
	Title       string `validate:"required" json:"title" form:"title"`
	Price       uint   `validate:"required" json:"price" form:"price"`
	Description string `json:"description" form:"description"`
	Image       string `json:"image" form:"image"`
}

func ToResponse(data product.CoreProduct) ProductResponse {
	return ProductResponse{
		Title:       data.Title,
		Price:       data.Price,
		Description: data.Description,
		Image:       data.Image,
	}
}

func ToResponseArr(data []product.CoreProduct) []ProductResponse {
	res := []ProductResponse{}
	for _, v := range data {
		tmp := ToResponse(v)
		res = append(res, tmp)
	}
	return res
}
