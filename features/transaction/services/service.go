package services

import (
	"ecommerce/features/transaction"
	"ecommerce/helper"
	"errors"
	"strings"
)

type trxSrv struct {
	data transaction.TrxData
}

func New(c transaction.TrxData) transaction.TrxService {
	return &trxSrv{
		data: c,
	}
}

func (ts *trxSrv) Add(token interface{}, newTrx transaction.Core) (transaction.Core, error) {
	userID := helper.ExtractToken(token)

	res, err := ts.data.Add(uint(userID), newTrx)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "user not found"
		} else {
			msg = "unable to process the data"
		}
		return transaction.Core{}, errors.New(msg)
	}

	return res, nil
}
