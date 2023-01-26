package helper

import (
	"ecommerce/config"
	"fmt"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

func MidtransSnapClient() snap.Client {
	s := snap.Client{}
	s.New(config.ServerKey, midtrans.Sandbox)
	fmt.Println(config.ServerKey)
	return s
}
