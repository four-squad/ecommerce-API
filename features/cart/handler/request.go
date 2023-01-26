package handler

import "ecommerce/features/cart"

type UptReq struct {
	Qty         uint `validate:"required" json:"qty" form:"qty"`
	Total_Price uint `json:"total_price" form:"total_price"`
}

func ToCore(data interface{}) *cart.CoreCart {
	res := cart.CoreCart{}

	switch data.(type) {
	case UptReq:
		cnv := data.(UptReq)
		res.Qty = cnv.Qty
		res.Total_Price = cnv.Total_Price
	default:
		return nil
	}

	return &res
}
