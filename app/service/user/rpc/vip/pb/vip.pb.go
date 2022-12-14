// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: vip.proto

package pb

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

type UpgradeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UpgradeReq) Reset() {
	*x = UpgradeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vip_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpgradeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpgradeReq) ProtoMessage() {}

func (x *UpgradeReq) ProtoReflect() protoreflect.Message {
	mi := &file_vip_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpgradeReq.ProtoReflect.Descriptor instead.
func (*UpgradeReq) Descriptor() ([]byte, []int) {
	return file_vip_proto_rawDescGZIP(), []int{0}
}

func (x *UpgradeReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UpgradeRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32            `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string           `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Ok   bool             `protobuf:"varint,3,opt,name=ok,proto3" json:"ok,omitempty"`
	Data *UpgradeRes_Data `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *UpgradeRes) Reset() {
	*x = UpgradeRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vip_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpgradeRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpgradeRes) ProtoMessage() {}

func (x *UpgradeRes) ProtoReflect() protoreflect.Message {
	mi := &file_vip_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpgradeRes.ProtoReflect.Descriptor instead.
func (*UpgradeRes) Descriptor() ([]byte, []int) {
	return file_vip_proto_rawDescGZIP(), []int{1}
}

func (x *UpgradeRes) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *UpgradeRes) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *UpgradeRes) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *UpgradeRes) GetData() *UpgradeRes_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type ResetReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ResetReq) Reset() {
	*x = ResetReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vip_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetReq) ProtoMessage() {}

func (x *ResetReq) ProtoReflect() protoreflect.Message {
	mi := &file_vip_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetReq.ProtoReflect.Descriptor instead.
func (*ResetReq) Descriptor() ([]byte, []int) {
	return file_vip_proto_rawDescGZIP(), []int{2}
}

func (x *ResetReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ResetRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Ok   bool   `protobuf:"varint,3,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *ResetRes) Reset() {
	*x = ResetRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vip_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetRes) ProtoMessage() {}

func (x *ResetRes) ProtoReflect() protoreflect.Message {
	mi := &file_vip_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetRes.ProtoReflect.Descriptor instead.
func (*ResetRes) Descriptor() ([]byte, []int) {
	return file_vip_proto_rawDescGZIP(), []int{3}
}

func (x *ResetRes) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ResetRes) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *ResetRes) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type UpgradeRes_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VipLevel int32 `protobuf:"varint,1,opt,name=vip_level,json=vipLevel,proto3" json:"vip_level,omitempty"`
}

func (x *UpgradeRes_Data) Reset() {
	*x = UpgradeRes_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vip_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpgradeRes_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpgradeRes_Data) ProtoMessage() {}

func (x *UpgradeRes_Data) ProtoReflect() protoreflect.Message {
	mi := &file_vip_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpgradeRes_Data.ProtoReflect.Descriptor instead.
func (*UpgradeRes_Data) Descriptor() ([]byte, []int) {
	return file_vip_proto_rawDescGZIP(), []int{1, 0}
}

func (x *UpgradeRes_Data) GetVipLevel() int32 {
	if x != nil {
		return x.VipLevel
	}
	return 0
}

var File_vip_proto protoreflect.FileDescriptor

var file_vip_proto_rawDesc = []byte{
	0x0a, 0x09, 0x76, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x76, 0x69, 0x70,
	0x22, 0x1c, 0x0a, 0x0a, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x91,
	0x01, 0x0a, 0x0a, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x02, 0x6f, 0x6b, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x76, 0x69, 0x70, 0x2e, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x52,
	0x65, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x23, 0x0a,
	0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x76, 0x69, 0x70, 0x5f, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x76, 0x69, 0x70, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x22, 0x1a, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x40,
	0x0a, 0x08, 0x52, 0x65, 0x73, 0x65, 0x74, 0x52, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67,
	0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b,
	0x32, 0x59, 0x0a, 0x03, 0x56, 0x69, 0x70, 0x12, 0x2b, 0x0a, 0x07, 0x55, 0x70, 0x67, 0x72, 0x61,
	0x64, 0x65, 0x12, 0x0f, 0x2e, 0x76, 0x69, 0x70, 0x2e, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65,
	0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x76, 0x69, 0x70, 0x2e, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64,
	0x65, 0x52, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x05, 0x52, 0x65, 0x73, 0x65, 0x74, 0x12, 0x0d, 0x2e,
	0x76, 0x69, 0x70, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x76,
	0x69, 0x70, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x52, 0x65, 0x73, 0x42, 0x06, 0x5a, 0x04, 0x2e,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vip_proto_rawDescOnce sync.Once
	file_vip_proto_rawDescData = file_vip_proto_rawDesc
)

func file_vip_proto_rawDescGZIP() []byte {
	file_vip_proto_rawDescOnce.Do(func() {
		file_vip_proto_rawDescData = protoimpl.X.CompressGZIP(file_vip_proto_rawDescData)
	})
	return file_vip_proto_rawDescData
}

var file_vip_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_vip_proto_goTypes = []interface{}{
	(*UpgradeReq)(nil),      // 0: vip.UpgradeReq
	(*UpgradeRes)(nil),      // 1: vip.UpgradeRes
	(*ResetReq)(nil),        // 2: vip.ResetReq
	(*ResetRes)(nil),        // 3: vip.ResetRes
	(*UpgradeRes_Data)(nil), // 4: vip.UpgradeRes.Data
}
var file_vip_proto_depIdxs = []int32{
	4, // 0: vip.UpgradeRes.data:type_name -> vip.UpgradeRes.Data
	0, // 1: vip.Vip.Upgrade:input_type -> vip.UpgradeReq
	2, // 2: vip.Vip.Reset:input_type -> vip.ResetReq
	1, // 3: vip.Vip.Upgrade:output_type -> vip.UpgradeRes
	3, // 4: vip.Vip.Reset:output_type -> vip.ResetRes
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_vip_proto_init() }
func file_vip_proto_init() {
	if File_vip_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vip_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpgradeReq); i {
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
		file_vip_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpgradeRes); i {
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
		file_vip_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetReq); i {
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
		file_vip_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetRes); i {
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
		file_vip_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpgradeRes_Data); i {
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
			RawDescriptor: file_vip_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_vip_proto_goTypes,
		DependencyIndexes: file_vip_proto_depIdxs,
		MessageInfos:      file_vip_proto_msgTypes,
	}.Build()
	File_vip_proto = out.File
	file_vip_proto_rawDesc = nil
	file_vip_proto_goTypes = nil
	file_vip_proto_depIdxs = nil
}
