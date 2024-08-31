// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.12.4
// source: exchange.proto

package exchange_v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Market int32

const (
	Market_usdtrub Market = 0
	Market_usdtusd Market = 1
	Market_usdteur Market = 2
)

// Enum value maps for Market.
var (
	Market_name = map[int32]string{
		0: "usdtrub",
		1: "usdtusd",
		2: "usdteur",
	}
	Market_value = map[string]int32{
		"usdtrub": 0,
		"usdtusd": 1,
		"usdteur": 2,
	}
)

func (x Market) Enum() *Market {
	p := new(Market)
	*p = x
	return p
}

func (x Market) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Market) Descriptor() protoreflect.EnumDescriptor {
	return file_exchange_proto_enumTypes[0].Descriptor()
}

func (Market) Type() protoreflect.EnumType {
	return &file_exchange_proto_enumTypes[0]
}

func (x Market) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Market.Descriptor instead.
func (Market) EnumDescriptor() ([]byte, []int) {
	return file_exchange_proto_rawDescGZIP(), []int{0}
}

type GetRatesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Market Market `protobuf:"varint,1,opt,name=market,proto3,enum=exchange_v1.Market" json:"market,omitempty"`
}

func (x *GetRatesRequest) Reset() {
	*x = GetRatesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exchange_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRatesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRatesRequest) ProtoMessage() {}

func (x *GetRatesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_exchange_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRatesRequest.ProtoReflect.Descriptor instead.
func (*GetRatesRequest) Descriptor() ([]byte, []int) {
	return file_exchange_proto_rawDescGZIP(), []int{0}
}

func (x *GetRatesRequest) GetMarket() Market {
	if x != nil {
		return x.Market
	}
	return Market_usdtrub
}

type GetRatesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp int64   `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Market    Market  `protobuf:"varint,2,opt,name=market,proto3,enum=exchange_v1.Market" json:"market,omitempty"`
	Ask       float64 `protobuf:"fixed64,3,opt,name=ask,proto3" json:"ask,omitempty"`
	Bid       float64 `protobuf:"fixed64,4,opt,name=bid,proto3" json:"bid,omitempty"`
}

func (x *GetRatesResponse) Reset() {
	*x = GetRatesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exchange_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRatesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRatesResponse) ProtoMessage() {}

func (x *GetRatesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_exchange_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRatesResponse.ProtoReflect.Descriptor instead.
func (*GetRatesResponse) Descriptor() ([]byte, []int) {
	return file_exchange_proto_rawDescGZIP(), []int{1}
}

func (x *GetRatesResponse) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *GetRatesResponse) GetMarket() Market {
	if x != nil {
		return x.Market
	}
	return Market_usdtrub
}

func (x *GetRatesResponse) GetAsk() float64 {
	if x != nil {
		return x.Ask
	}
	return 0
}

func (x *GetRatesResponse) GetBid() float64 {
	if x != nil {
		return x.Bid
	}
	return 0
}

var File_exchange_proto protoreflect.FileDescriptor

var file_exchange_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0b, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x76, 0x31, 0x22, 0x3e, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x52, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2b, 0x0a, 0x06, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x13, 0x2e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x76, 0x31, 0x2e, 0x4d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x52, 0x06, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x22, 0x81, 0x01,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x52, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x12, 0x2b, 0x0a, 0x06, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x13, 0x2e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x76, 0x31, 0x2e, 0x4d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x52, 0x06, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x61, 0x73, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x61, 0x73, 0x6b, 0x12,
	0x10, 0x0a, 0x03, 0x62, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x62, 0x69,
	0x64, 0x2a, 0x2f, 0x0a, 0x06, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x12, 0x0b, 0x0a, 0x07, 0x75,
	0x73, 0x64, 0x74, 0x72, 0x75, 0x62, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x75, 0x73, 0x64, 0x74,
	0x75, 0x73, 0x64, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x75, 0x73, 0x64, 0x74, 0x65, 0x75, 0x72,
	0x10, 0x02, 0x32, 0x57, 0x0a, 0x0c, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x47, 0x52,
	0x50, 0x43, 0x12, 0x47, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x52, 0x61, 0x74, 0x65, 0x73, 0x12, 0x1c,
	0x2e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x65,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x61,
	0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2d, 0x5a, 0x2b, 0x47,
	0x61, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x76, 0x31, 0x3b, 0x65,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_exchange_proto_rawDescOnce sync.Once
	file_exchange_proto_rawDescData = file_exchange_proto_rawDesc
)

func file_exchange_proto_rawDescGZIP() []byte {
	file_exchange_proto_rawDescOnce.Do(func() {
		file_exchange_proto_rawDescData = protoimpl.X.CompressGZIP(file_exchange_proto_rawDescData)
	})
	return file_exchange_proto_rawDescData
}

var file_exchange_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_exchange_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_exchange_proto_goTypes = []interface{}{
	(Market)(0),              // 0: exchange_v1.Market
	(*GetRatesRequest)(nil),  // 1: exchange_v1.GetRatesRequest
	(*GetRatesResponse)(nil), // 2: exchange_v1.GetRatesResponse
}
var file_exchange_proto_depIdxs = []int32{
	0, // 0: exchange_v1.GetRatesRequest.market:type_name -> exchange_v1.Market
	0, // 1: exchange_v1.GetRatesResponse.market:type_name -> exchange_v1.Market
	1, // 2: exchange_v1.ExchangeGRPC.GetRates:input_type -> exchange_v1.GetRatesRequest
	2, // 3: exchange_v1.ExchangeGRPC.GetRates:output_type -> exchange_v1.GetRatesResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_exchange_proto_init() }
func file_exchange_proto_init() {
	if File_exchange_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_exchange_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRatesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_exchange_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRatesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_exchange_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_exchange_proto_goTypes,
		DependencyIndexes: file_exchange_proto_depIdxs,
		EnumInfos:         file_exchange_proto_enumTypes,
		MessageInfos:      file_exchange_proto_msgTypes,
	}.Build()
	File_exchange_proto = out.File
	file_exchange_proto_rawDesc = nil
	file_exchange_proto_goTypes = nil
	file_exchange_proto_depIdxs = nil
}
