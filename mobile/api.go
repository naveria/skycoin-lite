package mobile

import (
	"github.com/skycoin/skycoin/src/cipher/go-bip39"
	"github.com/skycoin/skycoin-lite/liteclient"
	"encoding/json"
	"github.com/skycoin/skycoin-lite/service"
	"os"
	//"strconv"
)

// Returns a series of addresses owned/generated by the seed
func GetAddresses(seed string, amount int) (string, error) {
	addresses, err := liteclient.Addresses(seed, amount)
	response, err := json.Marshal(addresses)
	return string(response), err
}

// Returns addresses with balances
func GetBalances(seed string, amount int) (string, error) {
	addresses, err := liteclient.Addresses(seed, amount)
	completeAddresses, err := liteclient.AddressesWithBalance(addresses)
	if (err != nil) {
		os.Stderr.WriteString("got balances error "+err.Error()+"\n")
		return "error", err
	}
	response, err := json.Marshal(completeAddresses)
	return string(response), err
}

// Returns outputs
func GetOutputs(seed string, amount int) (string, error) {
	addresses, err := liteclient.Addresses(seed, amount)
	outputs, err := liteclient.Outputs(addresses)
	response, err := json.Marshal(outputs)
	return string(response), err
}

// Returns a transaction ID
func PostTransaction(seed string, addresses int, destinationAddress string, amount int) (string, error) {
	wallet := liteclient.Wallet{seed, addresses}
	//os.Stderr.WriteString("sending "+strconv.Itoa(amount)+"\n")
	return liteclient.Send(wallet, destinationAddress, uint64(amount))
}

func SetBackendNodeAddress(nodeAddr string) {
	service.NodeAddress = nodeAddr;
}

// Returns a nmemonic string
func GetSeed() (string) {
	entropy, err := bip39.NewEntropy(128)
	if err != nil {
		panic(err)
	}

	sd, err := bip39.NewMnemonic(entropy)
	if err != nil {
		panic(err)
	}
	return sd
}
