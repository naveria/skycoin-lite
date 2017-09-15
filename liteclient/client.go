package liteclient

import (
	"github.com/skycoin/mobile/service"
	"github.com/skycoin/skycoin/src/cipher"
	"errors"
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

func PrepareTx(wlt Wallet, toAddr string, amount uint64) {

	// TODO: Add address and wallet validation
	addresses, _ := Addresses(wlt.Seed, wlt.Addresses);
	stringifiedAddresses := make([]string, len(addresses))
	for i, address := range addresses {
		stringifiedAddresses[i] = address.Address
	}

	totalUtxos, err := service.GetOutputs(stringifiedAddresses)

	utxos, err := getSufficientOutputs(totalUtxos, amount)
	// TODO: add catch for err
}

func getSufficientOutputs(utxos []*service.Output, amt uint64) ([]*service.Output, error) {
	outMap := make(map[string][]*service.Output)
	for _, u := range utxos {
		outMap[u.GetAddress()] = append(outMap[u.GetAddress()], u)
	}

	allUtxos := []*service.Output{}
	var allBal uint64
	for _, utxos := range outMap {
		allBal += func(utxos []*service.Output) uint64 {
			var bal uint64
			for _, u := range utxos {
				if u.GetCoins() == 0 {
					continue
				}
				bal += u.GetCoins()
			}
			return bal
		}(utxos)

		allUtxos = append(allUtxos, utxos...)
		if allBal >= amt {
			return allUtxos, nil
		}
	}

	return nil, errors.New("insufficient balance")
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

	return addresses, nil
}

func AddressesWithBalance(addresses []Address) ([]Address, error) {
	stringifiedAddresses := make([]string, len(addresses))
	for i, address := range addresses {
		stringifiedAddresses[i] = address.Address
	}

	outputs, _ := service.GetOutputs(stringifiedAddresses)

	for _, output := range outputs {
		for i, address := range addresses {
			if address.Address == output.GetAddress() {
				addresses[i].Coins += output.GetCoins()
				addresses[i].Hours += output.GetHours()
			}
		}
	}

	return addresses, nil
}

