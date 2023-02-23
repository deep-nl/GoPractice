package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"go-base/logging"
)

func main() {
	fromKey := "92b5aef7cd54a40d541c66bf28a503122083ed60f3a758bb28d427c96484dbb4"
	privateKey, err := crypto.HexToECDSA(fromKey)
	if err != nil {
		logging.Logger.Errorln(err)
	}
	fmt.Println(privateKey)
	// how to convert privatekey to a modnscalar
}
