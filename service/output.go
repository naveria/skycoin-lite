package service

type Output struct {
	Hash             *string `protobuf:"bytes,10,opt,name=hash" json:"hash,omitempty"`
	SrcTx            *string `protobuf:"bytes,11,opt,name=src_tx" json:"src_tx,omitempty"`
	Address          *string `protobuf:"bytes,12,opt,name=address" json:"address,omitempty"`
	Coins            *uint64 `protobuf:"varint,13,opt,name=coins" json:"coins,omitempty"`
	Hours            *uint64 `protobuf:"varint,14,opt,name=hours" json:"hours,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}
