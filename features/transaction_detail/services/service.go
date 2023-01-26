package services

import (
	"ecommerce/features/transaction_detail"

	"github.com/go-playground/validator"
)

type transaction_detailUseCase struct {
	qry transaction_detail.Transaction_DetailData
	vld *validator.Validate
}

func New(td transaction_detail.Transaction_DetailData) transaction_detail.Transaction_DetailService {
	return &transaction_detailUseCase{
		qry: td,
		vld: validator.New(),
	}
}
