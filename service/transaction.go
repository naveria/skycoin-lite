package service

import (
	"fmt"
	"github.com/skycoin/skycoin-exchange/src/pp"
	"github.com/skycoin/mobile/sknet"
)

type InjectTxnReq struct {
	CoinType         *string `protobuf:"bytes,10,opt,name=coin_type" json:"coin_type,omitempty"`
	Tx               *string `protobuf:"bytes,20,opt,name=tx" json:"tx,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

type InjectTxnRes struct {
	Result           *Result `protobuf:"bytes,1,req,name=result" json:"result,omitempty"`
	Txid             *string `protobuf:"bytes,10,opt,name=txid" json:"txid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *InjectTxnRes) GetTxid() string {
	if m != nil && m.Txid != nil {
		return *m.Txid
	}
	return ""
}

func InjectTransaction(rawtx string) (string, error) {
	req := InjectTxnReq{
		CoinType: pp.PtrString("skycoin"),
		Tx:       pp.PtrString(rawtx),
	}
	res := InjectTxnRes{}
	if err := sknet.EncryGet(NodeAddress, "/inject/tx", req, &res); err != nil {
		return "", err
	}

	if !res.Result.GetSuccess() {
		return "", fmt.Errorf("broadcast tx failed: %v", res.Result.GetReason())
	}

	return res.GetTxid(), nil
}
