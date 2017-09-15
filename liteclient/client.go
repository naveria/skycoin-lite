package liteclient

import (
	"github.com/skycoin/skycoin/src/cipher"
)

type Wallet struct {
	Seed string
	Addresses int
}

type Address struct {
	Address string
	Secret string
	Coins uint64
	Hours uint64
}

func Addresses(seed string, amount int) ([]Address, error) {
	_, secretKeys := cipher.GenerateDeterministicKeyPairsSeed([]byte(seed), amount)
	addresses := make([]Address, amount)
	for i, sec := range secretKeys {
		pub := cipher.PubKeyFromSecKey(sec)
		address := Address{
			cipher.AddressFromPubKey(pub).String(),
			sec.Hex(),
			0,
			0,
		}
		addresses[i] = address
	}
