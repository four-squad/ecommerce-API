package services

import (
	"ecommerce/features/transaction"
	"ecommerce/helper"
	"errors"
	"fmt"
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

func (ts *trxSrv) Add(cartID uint, token interface{}, newTrx transaction.Core) error {
	userID := helper.ExtractToken(token)

	fmt.Println("srv")

	err := ts.data.Add(cartID, uint(userID), newTrx)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "user not found"
		} else {
			msg = "unable to process the data"
		}
		return errors.New(msg)
	}

	return nil
}
