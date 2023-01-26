package helper

import (
	"ecommerce/config"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

func MidtransSnapClient() snap.Client {
	s := snap.Client{}
	s.New(config.ServerKey, midtrans.Sandbox)
	return s
}
