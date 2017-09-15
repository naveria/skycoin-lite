package liteclient

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
