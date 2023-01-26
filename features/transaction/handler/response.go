package handler

import "ecommerce/features/transaction"

type AddResponse struct {
	PaymentUrl string `json:"payment_url"`
}

func ToResponse(data transaction.Core) AddResponse {
	return AddResponse{
		PaymentUrl: data.PaymentUrl,
	}
}
