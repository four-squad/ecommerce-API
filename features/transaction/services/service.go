package services

import (
	"ecommerce/features/transaction"
	"ecommerce/helper"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
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

	fmt.Println("servriecds", res)

	id := strconv.Itoa(int(res.ID))

	s := helper.MidtransSnapClient()
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  id,
			GrossAmt: int64(res.TotalPrice),
		},
	}
	snapResp, _ := s.CreateTransaction(req)

	res.PaymentUrl = snapResp.RedirectURL

	return res, nil
}
