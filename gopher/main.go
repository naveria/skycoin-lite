package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/skycoin/skycoin-lite-gopher/liteclient"
)

func main() {
	js.Global.Set("Cipher", map[string]interface{}{
		"GenerateAddresses": liteclient.GenerateAddress,
		"PrepareTransaction": liteclient.PrepareTransaction,
	})
}
