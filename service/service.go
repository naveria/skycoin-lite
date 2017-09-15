package service

type Result struct {
	Success          *bool   `protobuf:"varint,1,req,name=success" json:"success,omitempty"`
	Errcode          *int32  `protobuf:"varint,2,opt,name=errcode" json:"errcode,omitempty"`
	Reason           *string `protobuf:"bytes,3,opt,name=reason" json:"reason,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}
