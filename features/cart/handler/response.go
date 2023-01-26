package handler

import "ecommerce/features/cart"

type CartResponse struct {
	ID          uint   `json:"id"`
	Title       string `validate:"required" json:"title" form:"title"`
	Qty         uint   `json:"qty" form:"qty"`
	Price       uint   `validate:"required" json:"price" form:"price"`
	Image       string `json:"image" form:"image"`
	Total_Price uint   `json:"total_price" form:"total_price"`
}

// type GetAllRespon struct {
// 	ID    uint   `json:"id" form:"id"`
// 	Title string `validate:"required" json:"title" form:"title"`
// 	Price uint   `validate:"required" json:"price" form:"price"`
// 	Image string `json:"image" form:"image"`
// }

func ToResponse(data cart.CoreCart) CartResponse {
	return CartResponse{
		ID:          data.ID,
		Title:       data.Title,
		Qty:         data.Qty,
		Price:       data.Price,
		Image:       data.Image,
		Total_Price: data.Total_Price,
	}
}

func ToResponseArr(data []cart.CoreCart) []CartResponse {
	res := []CartResponse{}
	for _, v := range data {
		tmp := ToResponse(v)
		res = append(res, tmp)
	}
	return res
}
