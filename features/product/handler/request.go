package handler

import "ecommerce/features/product"

type AddReq struct {
	Title       string `validate:"required" json:"title" form:"title"`
	Price       uint   `validate:"required" json:"price" form:"price"`
	Description string `json:"description" form:"description"`
	Image       string `json:"image" form:"image"`
}

func ToCore(data interface{}) *product.CoreProduct {
	res := product.CoreProduct{}

	switch data.(type) {
	case AddReq:
		cnv := data.(AddReq)
		res.Title = cnv.Title
		res.Price = cnv.Price
		res.Description = cnv.Description
		res.Image = cnv.Image
	default:
		return nil
	}

	return &res
}
