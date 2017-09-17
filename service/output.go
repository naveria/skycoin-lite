package service

import (
	"github.com/skycoin/mobile/sknet"
)

type Output struct {
	Hash             *string `protobuf:"bytes,10,opt,name=hash" json:"hash,omitempty"`
	SrcTx            *string `protobuf:"bytes,11,opt,name=src_tx" json:"src_tx,omitempty"`
	Address          *string `protobuf:"bytes,12,opt,name=address" json:"address,omitempty"`
	Coins            *uint64 `protobuf:"varint,13,opt,name=coins" json:"coins,omitempty"`
	Hours            *uint64 `protobuf:"varint,14,opt,name=hours" json:"hours,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

type OutputRequest struct {
	CoinType  	*string  `protobuf:"bytes,1,opt,name=coin_type" json:"coin_type,omitempty"`
	Addresses	[]string `protobuf:"bytes,10,rep,name=addresses" json:"addresses,omitempty"`
}

type OutputResponse struct {
	Result           *Result    `protobuf:"bytes,1,opt,name=result" json:"result,omitempty"`
	Outputs         []*Output `protobuf:"bytes,12,rep,name=sky_utxos" json:"sky_utxos,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *Output) GetAddress() string {
	if m != nil && m.Address != nil {
		return *m.Address
	}
	return ""
}

func (m *Output) GetCoins() uint64 {
	if m != nil && m.Coins != nil {
		return *m.Coins
	}
	return 0
}

func (m *Output) GetHours() uint64 {
	if m != nil && m.Hours != nil {
		return *m.Hours
	}
	return 0
}

func GetOutputs(addrs []string) ([]*Output, error) {
	name := "skycoin"
	req := OutputRequest{&name, addrs}
	res := OutputResponse{}


	if len(addrs) == 0 {
		return []*Output{}, nil
	}

	if err := sknet.EncryGet(NodeAddress, "/get/utxos", req, &res); err != nil {
		return nil, err
	}

	return res.Outputs, nil
}