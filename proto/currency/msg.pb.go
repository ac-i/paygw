//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: paygw/proto/currency/msg.proto

package currency

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

//
type Currency struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//
	Code3 string `protobuf:"bytes,1,opt,name=code3,proto3" json:"code3,omitempty"`
	//
	Amount100 int64 `protobuf:"varint,3,opt,name=amount100,proto3" json:"amount100,omitempty"`
}

func (x *Currency) Reset() {
	*x = Currency{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paygw_proto_currency_msg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Currency) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Currency) ProtoMessage() {}

func (x *Currency) ProtoReflect() protoreflect.Message {
	mi := &file_paygw_proto_currency_msg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Currency.ProtoReflect.Descriptor instead.
func (*Currency) Descriptor() ([]byte, []int) {
	return file_paygw_proto_currency_msg_proto_rawDescGZIP(), []int{0}
}

func (x *Currency) GetCode3() string {
	if x != nil {
		return x.Code3
	}
	return ""
}

func (x *Currency) GetAmount100() int64 {
	if x != nil {
		return x.Amount100
	}
	return 0
}

//
type Log100 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// sum amount100
	Sum int64 `protobuf:"varint,1,opt,name=sum,proto3" json:"sum,omitempty"`
	//
	Count int32 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	// amount100
	Amounts []int64 `protobuf:"varint,3,rep,packed,name=amounts,proto3" json:"amounts,omitempty"`
	// received at ns
	Ns []int64 `protobuf:"varint,4,rep,packed,name=ns,proto3" json:"ns,omitempty"`
}

func (x *Log100) Reset() {
	*x = Log100{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paygw_proto_currency_msg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Log100) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log100) ProtoMessage() {}

func (x *Log100) ProtoReflect() protoreflect.Message {
	mi := &file_paygw_proto_currency_msg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log100.ProtoReflect.Descriptor instead.
func (*Log100) Descriptor() ([]byte, []int) {
	return file_paygw_proto_currency_msg_proto_rawDescGZIP(), []int{1}
}

func (x *Log100) GetSum() int64 {
	if x != nil {
		return x.Sum
	}
	return 0
}

func (x *Log100) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *Log100) GetAmounts() []int64 {
	if x != nil {
		return x.Amounts
	}
	return nil
}

func (x *Log100) GetNs() []int64 {
	if x != nil {
		return x.Ns
	}
	return nil
}

var File_paygw_proto_currency_msg_proto protoreflect.FileDescriptor

var file_paygw_proto_currency_msg_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x61, 0x79, 0x67, 0x77, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2f, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x22, 0x3e, 0x0a, 0x08, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x64, 0x65, 0x33, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x64,
	0x65, 0x33, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x31, 0x30, 0x30, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x31, 0x30, 0x30,
	0x22, 0x5a, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x31, 0x30, 0x30, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x75,
	0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x73, 0x75, 0x6d, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x03, 0x52, 0x07, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x0e, 0x0a, 0x02,
	0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x03, 0x52, 0x02, 0x6e, 0x73, 0x42, 0x16, 0x5a, 0x14,
	0x70, 0x61, 0x79, 0x67, 0x77, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_paygw_proto_currency_msg_proto_rawDescOnce sync.Once
	file_paygw_proto_currency_msg_proto_rawDescData = file_paygw_proto_currency_msg_proto_rawDesc
)

func file_paygw_proto_currency_msg_proto_rawDescGZIP() []byte {
	file_paygw_proto_currency_msg_proto_rawDescOnce.Do(func() {
		file_paygw_proto_currency_msg_proto_rawDescData = protoimpl.X.CompressGZIP(file_paygw_proto_currency_msg_proto_rawDescData)
	})
	return file_paygw_proto_currency_msg_proto_rawDescData
}

var file_paygw_proto_currency_msg_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_paygw_proto_currency_msg_proto_goTypes = []interface{}{
	(*Currency)(nil), // 0: proto.currency.Currency
	(*Log100)(nil),   // 1: proto.currency.Log100
}
var file_paygw_proto_currency_msg_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_paygw_proto_currency_msg_proto_init() }
func file_paygw_proto_currency_msg_proto_init() {
	if File_paygw_proto_currency_msg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_paygw_proto_currency_msg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Currency); i {
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
		file_paygw_proto_currency_msg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Log100); i {
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
			RawDescriptor: file_paygw_proto_currency_msg_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_paygw_proto_currency_msg_proto_goTypes,
		DependencyIndexes: file_paygw_proto_currency_msg_proto_depIdxs,
		MessageInfos:      file_paygw_proto_currency_msg_proto_msgTypes,
	}.Build()
	File_paygw_proto_currency_msg_proto = out.File
	file_paygw_proto_currency_msg_proto_rawDesc = nil
	file_paygw_proto_currency_msg_proto_goTypes = nil
	file_paygw_proto_currency_msg_proto_depIdxs = nil
}
