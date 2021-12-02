// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: sampleclient/busi/svc.proto

package busi

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BusiReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount int64 `protobuf:"varint,1,opt,name=Amount,proto3" json:"Amount,omitempty"`
	UserId int64 `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
}

func (x *BusiReq) Reset() {
	*x = BusiReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sampleclient_busi_svc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BusiReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BusiReq) ProtoMessage() {}

func (x *BusiReq) ProtoReflect() protoreflect.Message {
	mi := &file_sampleclient_busi_svc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BusiReq.ProtoReflect.Descriptor instead.
func (*BusiReq) Descriptor() ([]byte, []int) {
	return file_sampleclient_busi_svc_proto_rawDescGZIP(), []int{0}
}

func (x *BusiReq) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *BusiReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

var File_sampleclient_busi_svc_proto protoreflect.FileDescriptor

var file_sampleclient_busi_svc_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x62,
	0x75, 0x73, 0x69, 0x2f, 0x73, 0x76, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x62,
	0x75, 0x73, 0x69, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x39, 0x0a, 0x07, 0x42, 0x75, 0x73, 0x69, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x32, 0x6f, 0x0a, 0x04, 0x42,
	0x75, 0x73, 0x69, 0x12, 0x32, 0x0a, 0x07, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x49, 0x6e, 0x12, 0x0d,
	0x2e, 0x62, 0x75, 0x73, 0x69, 0x2e, 0x42, 0x75, 0x73, 0x69, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x08, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x4f, 0x75, 0x74, 0x12, 0x0d, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x2e, 0x42, 0x75, 0x73, 0x69, 0x52,
	0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x08, 0x5a, 0x06,
	0x2e, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sampleclient_busi_svc_proto_rawDescOnce sync.Once
	file_sampleclient_busi_svc_proto_rawDescData = file_sampleclient_busi_svc_proto_rawDesc
)

func file_sampleclient_busi_svc_proto_rawDescGZIP() []byte {
	file_sampleclient_busi_svc_proto_rawDescOnce.Do(func() {
		file_sampleclient_busi_svc_proto_rawDescData = protoimpl.X.CompressGZIP(file_sampleclient_busi_svc_proto_rawDescData)
	})
	return file_sampleclient_busi_svc_proto_rawDescData
}

var file_sampleclient_busi_svc_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_sampleclient_busi_svc_proto_goTypes = []interface{}{
	(*BusiReq)(nil),       // 0: busi.BusiReq
	(*emptypb.Empty)(nil), // 1: google.protobuf.Empty
}
var file_sampleclient_busi_svc_proto_depIdxs = []int32{
	0, // 0: busi.Busi.TransIn:input_type -> busi.BusiReq
	0, // 1: busi.Busi.TransOut:input_type -> busi.BusiReq
	1, // 2: busi.Busi.TransIn:output_type -> google.protobuf.Empty
	1, // 3: busi.Busi.TransOut:output_type -> google.protobuf.Empty
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sampleclient_busi_svc_proto_init() }
func file_sampleclient_busi_svc_proto_init() {
	if File_sampleclient_busi_svc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sampleclient_busi_svc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BusiReq); i {
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
			RawDescriptor: file_sampleclient_busi_svc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sampleclient_busi_svc_proto_goTypes,
		DependencyIndexes: file_sampleclient_busi_svc_proto_depIdxs,
		MessageInfos:      file_sampleclient_busi_svc_proto_msgTypes,
	}.Build()
	File_sampleclient_busi_svc_proto = out.File
	file_sampleclient_busi_svc_proto_rawDesc = nil
	file_sampleclient_busi_svc_proto_goTypes = nil
	file_sampleclient_busi_svc_proto_depIdxs = nil
}
