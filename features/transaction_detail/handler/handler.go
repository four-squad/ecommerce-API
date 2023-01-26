package handler

import (
	"ecommerce/features/transaction_detail"
)

type transaction_detailControll struct {
	srv transaction_detail.Transaction_DetailService
}

func New(srv transaction_detail.Transaction_DetailService) transaction_detail.Transaction_DetailHandler {
	return &transaction_detailControll{
		srv: srv,
	}
}
