// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: crud.proto

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

type RegisterReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *RegisterReq) Reset() {
	*x = RegisterReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crud_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterReq) ProtoMessage() {}

func (x *RegisterReq) ProtoReflect() protoreflect.Message {
	mi := &file_crud_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterReq.ProtoReflect.Descriptor instead.
func (*RegisterReq) Descriptor() ([]byte, []int) {
	return file_crud_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *RegisterReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type RegisterRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Ok   bool   `protobuf:"varint,3,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *RegisterRes) Reset() {
	*x = RegisterRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crud_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRes) ProtoMessage() {}

func (x *RegisterRes) ProtoReflect() protoreflect.Message {
	mi := &file_crud_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRes.ProtoReflect.Descriptor instead.
func (*RegisterRes) Descriptor() ([]byte, []int) {
	return file_crud_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterRes) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *RegisterRes) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *RegisterRes) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type LoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	LastIp   string `protobuf:"bytes,3,opt,name=last_ip,json=lastIp,proto3" json:"last_ip,omitempty"`
}

func (x *LoginReq) Reset() {
	*x = LoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crud_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReq) ProtoMessage() {}

func (x *LoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_crud_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReq.ProtoReflect.Descriptor instead.
func (*LoginReq) Descriptor() ([]byte, []int) {
	return file_crud_proto_rawDescGZIP(), []int{2}
}

func (x *LoginReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *LoginReq) GetLastIp() string {
	if x != nil {
		return x.LastIp
	}
	return ""
}

type LoginRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32          `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string         `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Ok   bool           `protobuf:"varint,3,opt,name=ok,proto3" json:"ok,omitempty"`
	Data *LoginRes_Data `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *LoginRes) Reset() {
	*x = LoginRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crud_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRes) ProtoMessage() {}

func (x *LoginRes) ProtoReflect() protoreflect.Message {
	mi := &file_crud_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRes.ProtoReflect.Descriptor instead.
func (*LoginRes) Descriptor() ([]byte, []int) {
	return file_crud_proto_rawDescGZIP(), []int{3}
}

func (x *LoginRes) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *LoginRes) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *LoginRes) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *LoginRes) GetData() *LoginRes_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type CreateCollectionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Action      int32 `protobuf:"varint,2,opt,name=action,proto3" json:"action,omitempty"`
	CollectType int32 `protobuf:"varint,3,opt,name=collect_type,json=collectType,proto3" json:"collect_type,omitempty"`
	ObjType     int32 `protobuf:"varint,4,opt,name=obj_type,json=objType,proto3" json:"obj_type,omitempty"`
	ObjId       int64 `protobuf:"varint,5,opt,name=obj_id,json=objId,proto3" json:"obj_id,omitempty"`
}

func (x *CreateCollectionReq) Reset() {
	*x = CreateCollectionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crud_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCollectionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCollectionReq) ProtoMessage() {}

func (x *CreateCollectionReq) ProtoReflect() protoreflect.Message {
	mi := &file_crud_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCollectionReq.ProtoReflect.Descriptor instead.
func (*CreateCollectionReq) Descriptor() ([]byte, []int) {
	return file_crud_proto_rawDescGZIP(), []int{4}
}

func (x *CreateCollectionReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateCollectionReq) GetAction() int32 {
	if x != nil {
		return x.Action
	}
	return 0
}

func (x *CreateCollectionReq) GetCollectType() int32 {
	if x != nil {
		return x.CollectType
	}
	return 0
}

func (x *CreateCollectionReq) GetObjType() int32 {
	if x != nil {
		return x.ObjType
	}
	return 0
}

func (x *CreateCollectionReq) GetObjId() int64 {
	if x != nil {
		return x.ObjId
	}
	return 0
}

type CreateCollectionRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Ok   bool   `protobuf:"varint,3,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *CreateCollectionRes) Reset() {
	*x = CreateCollectionRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crud_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCollectionRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCollectionRes) ProtoMessage() {}

func (x *CreateCollectionRes) ProtoReflect() protoreflect.Message {
	mi := &file_crud_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCollectionRes.ProtoReflect.Descriptor instead.
func (*CreateCollectionRes) Descriptor() ([]byte, []int) {
	return file_crud_proto_rawDescGZIP(), []int{5}
}

func (x *CreateCollectionRes) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CreateCollectionRes) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *CreateCollectionRes) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type DeleteCollectionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CollectionId int64 `protobuf:"varint,1,opt,name=collection_id,json=collectionId,proto3" json:"collection_id,omitempty"`
}

func (x *DeleteCollectionReq) Reset() {
	*x = DeleteCollectionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crud_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCollectionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCollectionReq) ProtoMessage() {}

func (x *DeleteCollectionReq) ProtoReflect() protoreflect.Message {
	mi := &file_crud_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCollectionReq.ProtoReflect.Descriptor instead.
func (*DeleteCollectionReq) Descriptor() ([]byte, []int) {
	return file_crud_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteCollectionReq) GetCollectionId() int64 {
	if x != nil {
		return x.CollectionId
	}
	return 0
}

type DeleteCollectionRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Ok   bool   `protobuf:"varint,3,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *DeleteCollectionRes) Reset() {
	*x = DeleteCollectionRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crud_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCollectionRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCollectionRes) ProtoMessage() {}

func (x *DeleteCollectionRes) ProtoReflect() protoreflect.Message {
	mi := &file_crud_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCollectionRes.ProtoReflect.Descriptor instead.
func (*DeleteCollectionRes) Descriptor() ([]byte, []int) {
	return file_crud_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteCollectionRes) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *DeleteCollectionRes) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *DeleteCollectionRes) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type ChangeNicknameReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Nickname string `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
}

func (x *ChangeNicknameReq) Reset() {
	*x = ChangeNicknameReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crud_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeNicknameReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeNicknameReq) ProtoMessage() {}

func (x *ChangeNicknameReq) ProtoReflect() protoreflect.Message {
	mi := &file_crud_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeNicknameReq.ProtoReflect.Descriptor instead.
func (*ChangeNicknameReq) Descriptor() ([]byte, []int) {
	return file_crud_proto_rawDescGZIP(), []int{8}
}

func (x *ChangeNicknameReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ChangeNicknameReq) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

type ChangeNicknameRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Ok   bool   `protobuf:"varint,3,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *ChangeNicknameRes) Reset() {
	*x = ChangeNicknameRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crud_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeNicknameRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeNicknameRes) ProtoMessage() {}

func (x *ChangeNicknameRes) ProtoReflect() protoreflect.Message {
	mi := &file_crud_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeNicknameRes.ProtoReflect.Descriptor instead.
func (*ChangeNicknameRes) Descriptor() ([]byte, []int) {
	return file_crud_proto_rawDescGZIP(), []int{9}
}

func (x *ChangeNicknameRes) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ChangeNicknameRes) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *ChangeNicknameRes) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type ChangeFollowerReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Action int32 `protobuf:"varint,2,opt,name=action,proto3" json:"action,omitempty"`
}

func (x *ChangeFollowerReq) Reset() {
	*x = ChangeFollowerReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crud_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeFollowerReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeFollowerReq) ProtoMessage() {}

func (x *ChangeFollowerReq) ProtoReflect() protoreflect.Message {
	mi := &file_crud_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeFollowerReq.ProtoReflect.Descriptor instead.
func (*ChangeFollowerReq) Descriptor() ([]byte, []int) {
	return file_crud_proto_rawDescGZIP(), []int{10}
}

func (x *ChangeFollowerReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ChangeFollowerReq) GetAction() int32 {
	if x != nil {
		return x.Action
	}
	return 0
}

type ChangeFollowerRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Ok   bool   `protobuf:"varint,3,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *ChangeFollowerRes) Reset() {
	*x = ChangeFollowerRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crud_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeFollowerRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeFollowerRes) ProtoMessage() {}

func (x *ChangeFollowerRes) ProtoReflect() protoreflect.Message {
	mi := &file_crud_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeFollowerRes.ProtoReflect.Descriptor instead.
func (*ChangeFollowerRes) Descriptor() ([]byte, []int) {
	return file_crud_proto_rawDescGZIP(), []int{11}
}

func (x *ChangeFollowerRes) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ChangeFollowerRes) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *ChangeFollowerRes) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type LoginRes_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthToken string `protobuf:"bytes,1,opt,name=auth_token,json=authToken,proto3" json:"auth_token,omitempty"`
}

func (x *LoginRes_Data) Reset() {
	*x = LoginRes_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crud_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRes_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRes_Data) ProtoMessage() {}

func (x *LoginRes_Data) ProtoReflect() protoreflect.Message {
	mi := &file_crud_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRes_Data.ProtoReflect.Descriptor instead.
func (*LoginRes_Data) Descriptor() ([]byte, []int) {
	return file_crud_proto_rawDescGZIP(), []int{3, 0}
}

func (x *LoginRes_Data) GetAuthToken() string {
	if x != nil {
		return x.AuthToken
	}
	return ""
}

var File_crud_proto protoreflect.FileDescriptor

var file_crud_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x72,
	0x75, 0x64, 0x22, 0x45, 0x0a, 0x0b, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x43, 0x0a, 0x0b, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x0e,
	0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x5b,
	0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x69, 0x70, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x61, 0x73, 0x74, 0x49, 0x70, 0x22, 0x90, 0x01, 0x0a, 0x08,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x0e,
	0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x12, 0x27,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63,
	0x72, 0x75, 0x64, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x25, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x1d, 0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x9b,
	0x01, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x62,
	0x6a, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6f, 0x62,
	0x6a, 0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x5f, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6f, 0x62, 0x6a, 0x49, 0x64, 0x22, 0x4b, 0x0a, 0x13,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x3a, 0x0a, 0x13, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x4b, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d,
	0x73, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02,
	0x6f, 0x6b, 0x22, 0x3f, 0x0a, 0x11, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4e, 0x69, 0x63, 0x6b,
	0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x49, 0x0a, 0x11, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4e, 0x69, 0x63,
	0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x0e,
	0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x44,
	0x0a, 0x11, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x49, 0x0a, 0x11, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x52, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12,
	0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x32,
	0xfd, 0x02, 0x0a, 0x04, 0x43, 0x72, 0x75, 0x64, 0x12, 0x27, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x12, 0x0e, 0x2e, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x1a, 0x0e, 0x2e, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x73, 0x12, 0x30, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x11, 0x2e,
	0x63, 0x72, 0x75, 0x64, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x1a, 0x11, 0x2e, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x12, 0x48, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x2e, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x1a, 0x19, 0x2e, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x48, 0x0a,
	0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x19, 0x2e, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x63,
	0x72, 0x75, 0x64, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x42, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x4e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x2e, 0x63, 0x72, 0x75, 0x64,
	0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x52,
	0x65, 0x71, 0x1a, 0x17, 0x2e, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x4e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x12, 0x42, 0x0a, 0x0e, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x17, 0x2e,
	0x63, 0x72, 0x75, 0x64, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x52, 0x65, 0x73, 0x42,
	0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_crud_proto_rawDescOnce sync.Once
	file_crud_proto_rawDescData = file_crud_proto_rawDesc
)

func file_crud_proto_rawDescGZIP() []byte {
	file_crud_proto_rawDescOnce.Do(func() {
		file_crud_proto_rawDescData = protoimpl.X.CompressGZIP(file_crud_proto_rawDescData)
	})
	return file_crud_proto_rawDescData
}

var file_crud_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_crud_proto_goTypes = []interface{}{
	(*RegisterReq)(nil),         // 0: crud.RegisterReq
	(*RegisterRes)(nil),         // 1: crud.RegisterRes
	(*LoginReq)(nil),            // 2: crud.LoginReq
	(*LoginRes)(nil),            // 3: crud.LoginRes
	(*CreateCollectionReq)(nil), // 4: crud.CreateCollectionReq
	(*CreateCollectionRes)(nil), // 5: crud.CreateCollectionRes
	(*DeleteCollectionReq)(nil), // 6: crud.DeleteCollectionReq
	(*DeleteCollectionRes)(nil), // 7: crud.DeleteCollectionRes
	(*ChangeNicknameReq)(nil),   // 8: crud.ChangeNicknameReq
	(*ChangeNicknameRes)(nil),   // 9: crud.ChangeNicknameRes
	(*ChangeFollowerReq)(nil),   // 10: crud.ChangeFollowerReq
	(*ChangeFollowerRes)(nil),   // 11: crud.ChangeFollowerRes
	(*LoginRes_Data)(nil),       // 12: crud.LoginRes.Data
}
var file_crud_proto_depIdxs = []int32{
	12, // 0: crud.LoginRes.data:type_name -> crud.LoginRes.Data
	2,  // 1: crud.Crud.Login:input_type -> crud.LoginReq
	0,  // 2: crud.Crud.Register:input_type -> crud.RegisterReq
	4,  // 3: crud.Crud.CreateCollection:input_type -> crud.CreateCollectionReq
	6,  // 4: crud.Crud.DeleteCollection:input_type -> crud.DeleteCollectionReq
	8,  // 5: crud.Crud.ChangeNickName:input_type -> crud.ChangeNicknameReq
	10, // 6: crud.Crud.ChangeFollower:input_type -> crud.ChangeFollowerReq
	3,  // 7: crud.Crud.Login:output_type -> crud.LoginRes
	1,  // 8: crud.Crud.Register:output_type -> crud.RegisterRes
	5,  // 9: crud.Crud.CreateCollection:output_type -> crud.CreateCollectionRes
	7,  // 10: crud.Crud.DeleteCollection:output_type -> crud.DeleteCollectionRes
	9,  // 11: crud.Crud.ChangeNickName:output_type -> crud.ChangeNicknameRes
	11, // 12: crud.Crud.ChangeFollower:output_type -> crud.ChangeFollowerRes
	7,  // [7:13] is the sub-list for method output_type
	1,  // [1:7] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_crud_proto_init() }
func file_crud_proto_init() {
	if File_crud_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_crud_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterReq); i {
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
		file_crud_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRes); i {
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
		file_crud_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReq); i {
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
		file_crud_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRes); i {
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
		file_crud_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCollectionReq); i {
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
		file_crud_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCollectionRes); i {
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
		file_crud_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCollectionReq); i {
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
		file_crud_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCollectionRes); i {
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
		file_crud_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeNicknameReq); i {
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
		file_crud_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeNicknameRes); i {
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
		file_crud_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeFollowerReq); i {
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
		file_crud_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeFollowerRes); i {
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
		file_crud_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRes_Data); i {
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
			RawDescriptor: file_crud_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_crud_proto_goTypes,
		DependencyIndexes: file_crud_proto_depIdxs,
		MessageInfos:      file_crud_proto_msgTypes,
	}.Build()
	File_crud_proto = out.File
	file_crud_proto_rawDesc = nil
	file_crud_proto_goTypes = nil
	file_crud_proto_depIdxs = nil
}
