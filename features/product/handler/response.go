package handler

import "ecommerce/features/product"

type ProductResponse struct {
	Title       string `validate:"required" json:"title" form:"title"`
	Price       uint   `validate:"required" json:"price" form:"price"`
	Description string `json:"description" form:"description"`
	Image       string `json:"image" form:"image"`
	// Seller      string       `json:"seller_name" form:"seller_name"`
	// User UserResponse `json:"user" form:"user"`
	Name string `json:"seller" form:"seller"`
}

type GetAllRespon struct {
	ID    uint   `json:"id" form:"id"`
	Title string `validate:"required" json:"title" form:"title"`
	Price uint   `validate:"required" json:"price" form:"price"`
	Image string `json:"image" form:"image"`
}

// type UserResponse struct {
// 	// ID   uint   `json:"id" form:"id"`
// 	Name string `json:"seller" form:"seller"`
// }

func ToResponse(data product.CoreProduct) GetAllRespon {
	return GetAllRespon{
		ID:    data.ID,
		Title: data.Title,
		Price: data.Price,
		// Description: data.Description,
		Image: data.Image,
	}
}

func ToResponseArr(data []product.CoreProduct) []GetAllRespon {
	res := []GetAllRespon{}
	for _, v := range data {
		tmp := ToResponse(v)
		res = append(res, tmp)
	}
	return res
}

func ToResponse2(data product.CoreProduct) ProductResponse {
	return ProductResponse{
		Title:       data.Title,
		Price:       data.Price,
		Description: data.Description,
		Image:       data.Image,
		// User: UserResponse{
		// 	// ID:   data.User.ID,
		// 	Name: data.User.Name,
		// },
		Name: data.User.Name,
	}
}

func ToResponseArr2(data []product.CoreProduct) []ProductResponse {
	res := []ProductResponse{}
	for _, v := range data {
		tmp := ToResponse2(v)
		res = append(res, tmp)
	}
	return res
}
