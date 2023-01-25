package handler

import "ecommerce/features/transaction"

type AddTrxRequest struct {
	Address string `json:"address" form:"address"`
}

func ToCore(data interface{}) *transaction.Core {
	res := transaction.Core{}

	switch data.(type) {
	case AddTrxRequest:
		cnv := data.(AddTrxRequest)
		res.Address = cnv.Address
	default:
		return nil
	}

	return &res
}
