// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: info.proto

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

type GetObjInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ObjType int32 `protobuf:"varint,2,opt,name=obj_type,json=objType,proto3" json:"obj_type,omitempty"` // 0:问题 1:回答 2:文章
}

func (x *GetObjInfoReq) Reset() {
	*x = GetObjInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetObjInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetObjInfoReq) ProtoMessage() {}

func (x *GetObjInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetObjInfoReq.ProtoReflect.Descriptor instead.
func (*GetObjInfoReq) Descriptor() ([]byte, []int) {
	return file_info_proto_rawDescGZIP(), []int{0}
}

func (x *GetObjInfoReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetObjInfoReq) GetObjType() int32 {
	if x != nil {
		return x.ObjType
	}
	return 0
}

type GetObjInfoRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32               `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string              `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Ok   bool                `protobuf:"varint,3,opt,name=ok,proto3" json:"ok,omitempty"`
	Data *GetObjInfoRes_Data `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetObjInfoRes) Reset() {
	*x = GetObjInfoRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_info_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetObjInfoRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetObjInfoRes) ProtoMessage() {}

func (x *GetObjInfoRes) ProtoReflect() protoreflect.Message {
	mi := &file_info_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetObjInfoRes.ProtoReflect.Descriptor instead.
func (*GetObjInfoRes) Descriptor() ([]byte, []int) {
	return file_info_proto_rawDescGZIP(), []int{1}
}

func (x *GetObjInfoRes) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetObjInfoRes) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GetObjInfoRes) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *GetObjInfoRes) GetData() *GetObjInfoRes_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetPersonalInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetPersonalInfoReq) Reset() {
	*x = GetPersonalInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_info_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPersonalInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPersonalInfoReq) ProtoMessage() {}

func (x *GetPersonalInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_info_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPersonalInfoReq.ProtoReflect.Descriptor instead.
func (*GetPersonalInfoReq) Descriptor() ([]byte, []int) {
	return file_info_proto_rawDescGZIP(), []int{2}
}

func (x *GetPersonalInfoReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetPersonalInfoRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32                    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string                   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Ok   bool                     `protobuf:"varint,3,opt,name=ok,proto3" json:"ok,omitempty"`
	Data *GetPersonalInfoRes_Data `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetPersonalInfoRes) Reset() {
	*x = GetPersonalInfoRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_info_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPersonalInfoRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPersonalInfoRes) ProtoMessage() {}

func (x *GetPersonalInfoRes) ProtoReflect() protoreflect.Message {
	mi := &file_info_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPersonalInfoRes.ProtoReflect.Descriptor instead.
func (*GetPersonalInfoRes) Descriptor() ([]byte, []int) {
	return file_info_proto_rawDescGZIP(), []int{3}
}

func (x *GetPersonalInfoRes) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetPersonalInfoRes) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GetPersonalInfoRes) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *GetPersonalInfoRes) GetData() *GetPersonalInfoRes_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetCollectionInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId         int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CollectionType int32 `protobuf:"varint,2,opt,name=collection_type,json=collectionType,proto3" json:"collection_type,omitempty"`
}

func (x *GetCollectionInfoReq) Reset() {
	*x = GetCollectionInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_info_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCollectionInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCollectionInfoReq) ProtoMessage() {}

func (x *GetCollectionInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_info_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCollectionInfoReq.ProtoReflect.Descriptor instead.
func (*GetCollectionInfoReq) Descriptor() ([]byte, []int) {
	return file_info_proto_rawDescGZIP(), []int{4}
}

func (x *GetCollectionInfoReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetCollectionInfoReq) GetCollectionType() int32 {
	if x != nil {
		return x.CollectionType
	}
	return 0
}

type GetCollectionInfoRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32                      `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string                     `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Ok   bool                       `protobuf:"varint,3,opt,name=ok,proto3" json:"ok,omitempty"`
	Data *GetCollectionInfoRes_Data `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetCollectionInfoRes) Reset() {
	*x = GetCollectionInfoRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_info_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCollectionInfoRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCollectionInfoRes) ProtoMessage() {}

func (x *GetCollectionInfoRes) ProtoReflect() protoreflect.Message {
	mi := &file_info_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCollectionInfoRes.ProtoReflect.Descriptor instead.
func (*GetCollectionInfoRes) Descriptor() ([]byte, []int) {
	return file_info_proto_rawDescGZIP(), []int{5}
}

func (x *GetCollectionInfoRes) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetCollectionInfoRes) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GetCollectionInfoRes) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *GetCollectionInfoRes) GetData() *GetCollectionInfoRes_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetNotificationInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	MessageType int32 `protobuf:"varint,2,opt,name=message_type,json=messageType,proto3" json:"message_type,omitempty"`
}

func (x *GetNotificationInfoReq) Reset() {
	*x = GetNotificationInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_info_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNotificationInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNotificationInfoReq) ProtoMessage() {}

func (x *GetNotificationInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_info_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNotificationInfoReq.ProtoReflect.Descriptor instead.
func (*GetNotificationInfoReq) Descriptor() ([]byte, []int) {
	return file_info_proto_rawDescGZIP(), []int{6}
}

func (x *GetNotificationInfoReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetNotificationInfoReq) GetMessageType() int32 {
	if x != nil {
		return x.MessageType
	}
	return 0
}

type GetNotificationInfoRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32                        `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string                       `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Ok   bool                         `protobuf:"varint,3,opt,name=ok,proto3" json:"ok,omitempty"`
	Data *GetNotificationInfoRes_Data `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetNotificationInfoRes) Reset() {
	*x = GetNotificationInfoRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_info_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNotificationInfoRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNotificationInfoRes) ProtoMessage() {}

func (x *GetNotificationInfoRes) ProtoReflect() protoreflect.Message {
	mi := &file_info_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNotificationInfoRes.ProtoReflect.Descriptor instead.
func (*GetNotificationInfoRes) Descriptor() ([]byte, []int) {
	return file_info_proto_rawDescGZIP(), []int{7}
}

func (x *GetNotificationInfoRes) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetNotificationInfoRes) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GetNotificationInfoRes) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *GetNotificationInfoRes) GetData() *GetNotificationInfoRes_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetObjInfoRes_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []int64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *GetObjInfoRes_Data) Reset() {
	*x = GetObjInfoRes_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_info_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetObjInfoRes_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetObjInfoRes_Data) ProtoMessage() {}

func (x *GetObjInfoRes_Data) ProtoReflect() protoreflect.Message {
	mi := &file_info_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetObjInfoRes_Data.ProtoReflect.Descriptor instead.
func (*GetObjInfoRes_Data) Descriptor() ([]byte, []int) {
	return file_info_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GetObjInfoRes_Data) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type GetPersonalInfoRes_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Nickname string `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Email    string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phone    string `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Vip      int32  `protobuf:"varint,5,opt,name=vip,proto3" json:"vip,omitempty"`
}

func (x *GetPersonalInfoRes_Data) Reset() {
	*x = GetPersonalInfoRes_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_info_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPersonalInfoRes_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPersonalInfoRes_Data) ProtoMessage() {}

func (x *GetPersonalInfoRes_Data) ProtoReflect() protoreflect.Message {
	mi := &file_info_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPersonalInfoRes_Data.ProtoReflect.Descriptor instead.
func (*GetPersonalInfoRes_Data) Descriptor() ([]byte, []int) {
	return file_info_proto_rawDescGZIP(), []int{3, 0}
}

func (x *GetPersonalInfoRes_Data) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *GetPersonalInfoRes_Data) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *GetPersonalInfoRes_Data) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GetPersonalInfoRes_Data) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *GetPersonalInfoRes_Data) GetVip() int32 {
	if x != nil {
		return x.Vip
	}
	return 0
}

type GetCollectionInfoRes_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjType []int32 `protobuf:"varint,1,rep,packed,name=obj_type,json=objType,proto3" json:"obj_type,omitempty"` // 0:回答 1:文章
	ObjId   []int64 `protobuf:"varint,2,rep,packed,name=obj_id,json=objId,proto3" json:"obj_id,omitempty"`
}

func (x *GetCollectionInfoRes_Data) Reset() {
	*x = GetCollectionInfoRes_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_info_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCollectionInfoRes_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCollectionInfoRes_Data) ProtoMessage() {}

func (x *GetCollectionInfoRes_Data) ProtoReflect() protoreflect.Message {
	mi := &file_info_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCollectionInfoRes_Data.ProtoReflect.Descriptor instead.
func (*GetCollectionInfoRes_Data) Descriptor() ([]byte, []int) {
	return file_info_proto_rawDescGZIP(), []int{5, 0}
}

func (x *GetCollectionInfoRes_Data) GetObjType() []int32 {
	if x != nil {
		return x.ObjType
	}
	return nil
}

func (x *GetCollectionInfoRes_Data) GetObjId() []int64 {
	if x != nil {
		return x.ObjId
	}
	return nil
}

type GetNotificationInfoRes_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageId []int64 `protobuf:"varint,1,rep,packed,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
}

func (x *GetNotificationInfoRes_Data) Reset() {
	*x = GetNotificationInfoRes_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_info_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNotificationInfoRes_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNotificationInfoRes_Data) ProtoMessage() {}

func (x *GetNotificationInfoRes_Data) ProtoReflect() protoreflect.Message {
	mi := &file_info_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNotificationInfoRes_Data.ProtoReflect.Descriptor instead.
func (*GetNotificationInfoRes_Data) Descriptor() ([]byte, []int) {
	return file_info_proto_rawDescGZIP(), []int{7, 0}
}

func (x *GetNotificationInfoRes_Data) GetMessageId() []int64 {
	if x != nil {
		return x.MessageId
	}
	return nil
}

var File_info_proto protoreflect.FileDescriptor

var file_info_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x22, 0x43, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08,
	0x6f, 0x62, 0x6a, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x6f, 0x62, 0x6a, 0x54, 0x79, 0x70, 0x65, 0x22, 0x8d, 0x01, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4f,
	0x62, 0x6a, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12,
	0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x12,
	0x2c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x18, 0x0a,
	0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x2d, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0xfb, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x02, 0x6f, 0x6b, 0x12, 0x31, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x7c, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69,
	0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69,
	0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x69, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x76, 0x69, 0x70, 0x22, 0x58, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e,
	0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0xbb,
	0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x0e, 0x0a,
	0x02, 0x6f, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x12, 0x33, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x69, 0x6e,
	0x66, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x1a, 0x38, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x62,
	0x6a, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x07, 0x6f, 0x62,
	0x6a, 0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x03, 0x52, 0x05, 0x6f, 0x62, 0x6a, 0x49, 0x64, 0x22, 0x54, 0x0a, 0x16,
	0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x22, 0xac, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x02, 0x6f, 0x6b, 0x12, 0x35, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x25, 0x0a, 0x04, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49,
	0x64, 0x32, 0xa5, 0x02, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x36, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x4f, 0x62, 0x6a, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x13, 0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x2e,
	0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e,
	0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x73, 0x12, 0x45, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61,
	0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x47, 0x65, 0x74,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a,
	0x18, 0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x12, 0x4b, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a,
	0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x69, 0x6e, 0x66,
	0x6f, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x12, 0x51, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x2e,
	0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x69, 0x6e,
	0x66, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_info_proto_rawDescOnce sync.Once
	file_info_proto_rawDescData = file_info_proto_rawDesc
)

func file_info_proto_rawDescGZIP() []byte {
	file_info_proto_rawDescOnce.Do(func() {
		file_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_info_proto_rawDescData)
	})
	return file_info_proto_rawDescData
}

var file_info_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_info_proto_goTypes = []interface{}{
	(*GetObjInfoReq)(nil),               // 0: info.GetObjInfoReq
	(*GetObjInfoRes)(nil),               // 1: info.GetObjInfoRes
	(*GetPersonalInfoReq)(nil),          // 2: info.GetPersonalInfoReq
	(*GetPersonalInfoRes)(nil),          // 3: info.GetPersonalInfoRes
	(*GetCollectionInfoReq)(nil),        // 4: info.GetCollectionInfoReq
	(*GetCollectionInfoRes)(nil),        // 5: info.GetCollectionInfoRes
	(*GetNotificationInfoReq)(nil),      // 6: info.GetNotificationInfoReq
	(*GetNotificationInfoRes)(nil),      // 7: info.GetNotificationInfoRes
	(*GetObjInfoRes_Data)(nil),          // 8: info.GetObjInfoRes.Data
	(*GetPersonalInfoRes_Data)(nil),     // 9: info.GetPersonalInfoRes.Data
	(*GetCollectionInfoRes_Data)(nil),   // 10: info.GetCollectionInfoRes.Data
	(*GetNotificationInfoRes_Data)(nil), // 11: info.GetNotificationInfoRes.Data
}
var file_info_proto_depIdxs = []int32{
	8,  // 0: info.GetObjInfoRes.data:type_name -> info.GetObjInfoRes.Data
	9,  // 1: info.GetPersonalInfoRes.data:type_name -> info.GetPersonalInfoRes.Data
	10, // 2: info.GetCollectionInfoRes.data:type_name -> info.GetCollectionInfoRes.Data
	11, // 3: info.GetNotificationInfoRes.data:type_name -> info.GetNotificationInfoRes.Data
	0,  // 4: info.Info.GetObjInfo:input_type -> info.GetObjInfoReq
	2,  // 5: info.Info.GetPersonalInfo:input_type -> info.GetPersonalInfoReq
	4,  // 6: info.Info.GetCollectionInfo:input_type -> info.GetCollectionInfoReq
	6,  // 7: info.Info.GetNotificationInfo:input_type -> info.GetNotificationInfoReq
	1,  // 8: info.Info.GetObjInfo:output_type -> info.GetObjInfoRes
	3,  // 9: info.Info.GetPersonalInfo:output_type -> info.GetPersonalInfoRes
	5,  // 10: info.Info.GetCollectionInfo:output_type -> info.GetCollectionInfoRes
	7,  // 11: info.Info.GetNotificationInfo:output_type -> info.GetNotificationInfoRes
	8,  // [8:12] is the sub-list for method output_type
	4,  // [4:8] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_info_proto_init() }
func file_info_proto_init() {
	if File_info_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_info_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetObjInfoReq); i {
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
		file_info_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetObjInfoRes); i {
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
		file_info_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPersonalInfoReq); i {
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
		file_info_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPersonalInfoRes); i {
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
		file_info_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCollectionInfoReq); i {
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
		file_info_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCollectionInfoRes); i {
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
		file_info_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNotificationInfoReq); i {
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
		file_info_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNotificationInfoRes); i {
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
		file_info_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetObjInfoRes_Data); i {
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
		file_info_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPersonalInfoRes_Data); i {
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
		file_info_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCollectionInfoRes_Data); i {
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
		file_info_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNotificationInfoRes_Data); i {
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
			RawDescriptor: file_info_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_info_proto_goTypes,
		DependencyIndexes: file_info_proto_depIdxs,
		MessageInfos:      file_info_proto_msgTypes,
	}.Build()
	File_info_proto = out.File
	file_info_proto_rawDesc = nil
	file_info_proto_goTypes = nil
	file_info_proto_depIdxs = nil
}
