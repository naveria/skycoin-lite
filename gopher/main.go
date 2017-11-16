package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/skycoin/skycoin-lite-gopher/mobile"
)

func main() {
	js.Global.Set("Cipher", map[string]interface{}{
		"GenerateAddresses": mobile.GetAddresses("nest*", 3),
	})
}
