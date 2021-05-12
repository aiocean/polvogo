// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.2
// source: aiocean/user/v1/user_service.proto

package aiocean_user_v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UnlinkExternalAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orn string `protobuf:"bytes,1,opt,name=orn,proto3" json:"orn,omitempty"`
}

func (x *UnlinkExternalAccountRequest) Reset() {
	*x = UnlinkExternalAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aiocean_user_v1_user_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnlinkExternalAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnlinkExternalAccountRequest) ProtoMessage() {}

func (x *UnlinkExternalAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aiocean_user_v1_user_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnlinkExternalAccountRequest.ProtoReflect.Descriptor instead.
func (*UnlinkExternalAccountRequest) Descriptor() ([]byte, []int) {
	return file_aiocean_user_v1_user_service_proto_rawDescGZIP(), []int{0}
}

func (x *UnlinkExternalAccountRequest) GetOrn() string {
	if x != nil {
		return x.Orn
	}
	return ""
}

type UnlinkExternalAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UnlinkExternalAccountResponse) Reset() {
	*x = UnlinkExternalAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aiocean_user_v1_user_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnlinkExternalAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnlinkExternalAccountResponse) ProtoMessage() {}

func (x *UnlinkExternalAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aiocean_user_v1_user_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnlinkExternalAccountResponse.ProtoReflect.Descriptor instead.
func (*UnlinkExternalAccountResponse) Descriptor() ([]byte, []int) {
	return file_aiocean_user_v1_user_service_proto_rawDescGZIP(), []int{1}
}

type LinkExternalAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orn string `protobuf:"bytes,1,opt,name=orn,proto3" json:"orn,omitempty"`
}

func (x *LinkExternalAccountRequest) Reset() {
	*x = LinkExternalAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aiocean_user_v1_user_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LinkExternalAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LinkExternalAccountRequest) ProtoMessage() {}

func (x *LinkExternalAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aiocean_user_v1_user_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LinkExternalAccountRequest.ProtoReflect.Descriptor instead.
func (*LinkExternalAccountRequest) Descriptor() ([]byte, []int) {
	return file_aiocean_user_v1_user_service_proto_rawDescGZIP(), []int{2}
}

func (x *LinkExternalAccountRequest) GetOrn() string {
	if x != nil {
		return x.Orn
	}
	return ""
}

type LinkExternalAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LinkExternalAccountResponse) Reset() {
	*x = LinkExternalAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aiocean_user_v1_user_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LinkExternalAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LinkExternalAccountResponse) ProtoMessage() {}

func (x *LinkExternalAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aiocean_user_v1_user_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LinkExternalAccountResponse.ProtoReflect.Descriptor instead.
func (*LinkExternalAccountResponse) Descriptor() ([]byte, []int) {
	return file_aiocean_user_v1_user_service_proto_rawDescGZIP(), []int{3}
}

type DeleteUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orn string `protobuf:"bytes,1,opt,name=orn,proto3" json:"orn,omitempty"`
}

func (x *DeleteUserRequest) Reset() {
	*x = DeleteUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aiocean_user_v1_user_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserRequest) ProtoMessage() {}

func (x *DeleteUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aiocean_user_v1_user_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserRequest.ProtoReflect.Descriptor instead.
func (*DeleteUserRequest) Descriptor() ([]byte, []int) {
	return file_aiocean_user_v1_user_service_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteUserRequest) GetOrn() string {
	if x != nil {
		return x.Orn
	}
	return ""
}

type CreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orn  string `protobuf:"bytes,1,opt,name=orn,proto3" json:"orn,omitempty"`
	User *User  `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aiocean_user_v1_user_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aiocean_user_v1_user_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_aiocean_user_v1_user_service_proto_rawDescGZIP(), []int{5}
}

func (x *CreateUserRequest) GetOrn() string {
	if x != nil {
		return x.Orn
	}
	return ""
}

func (x *CreateUserRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type UpdateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orn        string                 `protobuf:"bytes,1,opt,name=orn,proto3" json:"orn,omitempty"`
	User       *User                  `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateUserRequest) Reset() {
	*x = UpdateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aiocean_user_v1_user_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserRequest) ProtoMessage() {}

func (x *UpdateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aiocean_user_v1_user_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return file_aiocean_user_v1_user_service_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateUserRequest) GetOrn() string {
	if x != nil {
		return x.Orn
	}
	return ""
}

func (x *UpdateUserRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *UpdateUserRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

type UpdateUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *UpdateUserResponse) Reset() {
	*x = UpdateUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aiocean_user_v1_user_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserResponse) ProtoMessage() {}

func (x *UpdateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aiocean_user_v1_user_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserResponse.ProtoReflect.Descriptor instead.
func (*UpdateUserResponse) Descriptor() ([]byte, []int) {
	return file_aiocean_user_v1_user_service_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateUserResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type CreateUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *CreateUserResponse) Reset() {
	*x = CreateUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aiocean_user_v1_user_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserResponse) ProtoMessage() {}

func (x *CreateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aiocean_user_v1_user_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserResponse.ProtoReflect.Descriptor instead.
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return file_aiocean_user_v1_user_service_proto_rawDescGZIP(), []int{8}
}

func (x *CreateUserResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type GetUserByOrnRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orn string `protobuf:"bytes,1,opt,name=orn,proto3" json:"orn,omitempty"`
}

func (x *GetUserByOrnRequest) Reset() {
	*x = GetUserByOrnRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aiocean_user_v1_user_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserByOrnRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByOrnRequest) ProtoMessage() {}

func (x *GetUserByOrnRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aiocean_user_v1_user_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByOrnRequest.ProtoReflect.Descriptor instead.
func (*GetUserByOrnRequest) Descriptor() ([]byte, []int) {
	return file_aiocean_user_v1_user_service_proto_rawDescGZIP(), []int{9}
}

func (x *GetUserByOrnRequest) GetOrn() string {
	if x != nil {
		return x.Orn
	}
	return ""
}

type GetUserByOrnResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetUserByOrnResponse) Reset() {
	*x = GetUserByOrnResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aiocean_user_v1_user_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserByOrnResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByOrnResponse) ProtoMessage() {}

func (x *GetUserByOrnResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aiocean_user_v1_user_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByOrnResponse.ProtoReflect.Descriptor instead.
func (*GetUserByOrnResponse) Descriptor() ([]byte, []int) {
	return file_aiocean_user_v1_user_service_proto_rawDescGZIP(), []int{10}
}

func (x *GetUserByOrnResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

var File_aiocean_user_v1_user_service_proto protoreflect.FileDescriptor

var file_aiocean_user_v1_user_service_proto_rawDesc = []byte{
	0x0a, 0x22, 0x61, 0x69, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x61, 0x69, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1a, 0x61, 0x69, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x30, 0x0a, 0x1c, 0x55, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x6f, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6f, 0x72,
	0x6e, 0x22, 0x1f, 0x0a, 0x1d, 0x55, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x45, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x2e, 0x0a, 0x1a, 0x4c, 0x69, 0x6e, 0x6b, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x6f, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6f,
	0x72, 0x6e, 0x22, 0x1d, 0x0a, 0x1b, 0x4c, 0x69, 0x6e, 0x6b, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x25, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x72, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6f, 0x72, 0x6e, 0x22, 0x50, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x6f, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6f, 0x72, 0x6e, 0x12,
	0x29, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x61, 0x69, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x8d, 0x01, 0x0a, 0x11, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x6f, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6f,
	0x72, 0x6e, 0x12, 0x29, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x61, 0x69, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x3b, 0x0a,
	0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x22, 0x3f, 0x0a, 0x12, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x29, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x61, 0x69, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x3f, 0x0a, 0x12, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x29, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x61, 0x69, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x2d, 0x0a, 0x13,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x4f, 0x72, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x03, 0x6f, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x04, 0xe2, 0x41, 0x01, 0x02, 0x52, 0x03, 0x6f, 0x72, 0x6e, 0x22, 0x41, 0x0a, 0x14, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x4f, 0x72, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x61, 0x69, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x32, 0xee,
	0x05, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6e,
	0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x22, 0x2e, 0x61,
	0x69, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x23, 0x2e, 0x61, 0x69, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x22, 0x09, 0x2f,
	0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x3a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x76,
	0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x22, 0x2e, 0x61,
	0x69, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x23, 0x2e, 0x61, 0x69, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x32, 0x11, 0x2f,
	0x76, 0x31, 0x2f, 0x7b, 0x6f, 0x72, 0x6e, 0x3d, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x2a, 0x7d,
	0x3a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x63, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x22, 0x2e, 0x61, 0x69, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x2a, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6f,
	0x72, 0x6e, 0x3d, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x2a, 0x7d, 0x12, 0x76, 0x0a, 0x0c, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x4f, 0x72, 0x6e, 0x12, 0x24, 0x2e, 0x61, 0x69,
	0x6f, 0x63, 0x65, 0x61, 0x6e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x4f, 0x72, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x25, 0x2e, 0x61, 0x69, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x4f, 0x72, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13,
	0x12, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6f, 0x72, 0x6e, 0x3d, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x2f, 0x2a, 0x7d, 0x12, 0x72, 0x0a, 0x13, 0x4c, 0x69, 0x6e, 0x6b, 0x45, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2b, 0x2e, 0x61, 0x69, 0x6f,
	0x63, 0x65, 0x61, 0x6e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x6e,
	0x6b, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x61, 0x69, 0x6f, 0x63, 0x65, 0x61,
	0x6e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x45, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xa5, 0x01, 0x0a, 0x15, 0x55, 0x6e, 0x6c, 0x69,
	0x6e, 0x6b, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x2d, 0x2e, 0x61, 0x69, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2e, 0x2e, 0x61, 0x69, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x2d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x2a, 0x25, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6f,
	0x72, 0x6e, 0x3d, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x2a, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x2a, 0x7d, 0x42,
	0x31, 0x5a, 0x2f, 0x70, 0x6b, 0x67, 0x2e, 0x61, 0x69, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x2e, 0x64,
	0x65, 0x76, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f,
	0x76, 0x31, 0x3b, 0x61, 0x69, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aiocean_user_v1_user_service_proto_rawDescOnce sync.Once
	file_aiocean_user_v1_user_service_proto_rawDescData = file_aiocean_user_v1_user_service_proto_rawDesc
)

func file_aiocean_user_v1_user_service_proto_rawDescGZIP() []byte {
	file_aiocean_user_v1_user_service_proto_rawDescOnce.Do(func() {
		file_aiocean_user_v1_user_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_aiocean_user_v1_user_service_proto_rawDescData)
	})
	return file_aiocean_user_v1_user_service_proto_rawDescData
}

var file_aiocean_user_v1_user_service_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_aiocean_user_v1_user_service_proto_goTypes = []interface{}{
	(*UnlinkExternalAccountRequest)(nil),  // 0: aiocean.user.v1.UnlinkExternalAccountRequest
	(*UnlinkExternalAccountResponse)(nil), // 1: aiocean.user.v1.UnlinkExternalAccountResponse
	(*LinkExternalAccountRequest)(nil),    // 2: aiocean.user.v1.LinkExternalAccountRequest
	(*LinkExternalAccountResponse)(nil),   // 3: aiocean.user.v1.LinkExternalAccountResponse
	(*DeleteUserRequest)(nil),             // 4: aiocean.user.v1.DeleteUserRequest
	(*CreateUserRequest)(nil),             // 5: aiocean.user.v1.CreateUserRequest
	(*UpdateUserRequest)(nil),             // 6: aiocean.user.v1.UpdateUserRequest
	(*UpdateUserResponse)(nil),            // 7: aiocean.user.v1.UpdateUserResponse
	(*CreateUserResponse)(nil),            // 8: aiocean.user.v1.CreateUserResponse
	(*GetUserByOrnRequest)(nil),           // 9: aiocean.user.v1.GetUserByOrnRequest
	(*GetUserByOrnResponse)(nil),          // 10: aiocean.user.v1.GetUserByOrnResponse
	(*User)(nil),                          // 11: aiocean.user.v1.User
	(*fieldmaskpb.FieldMask)(nil),         // 12: google.protobuf.FieldMask
	(*emptypb.Empty)(nil),                 // 13: google.protobuf.Empty
}
var file_aiocean_user_v1_user_service_proto_depIdxs = []int32{
	11, // 0: aiocean.user.v1.CreateUserRequest.user:type_name -> aiocean.user.v1.User
	11, // 1: aiocean.user.v1.UpdateUserRequest.user:type_name -> aiocean.user.v1.User
	12, // 2: aiocean.user.v1.UpdateUserRequest.update_mask:type_name -> google.protobuf.FieldMask
	11, // 3: aiocean.user.v1.UpdateUserResponse.user:type_name -> aiocean.user.v1.User
	11, // 4: aiocean.user.v1.CreateUserResponse.user:type_name -> aiocean.user.v1.User
	11, // 5: aiocean.user.v1.GetUserByOrnResponse.user:type_name -> aiocean.user.v1.User
	5,  // 6: aiocean.user.v1.UserService.CreateUser:input_type -> aiocean.user.v1.CreateUserRequest
	6,  // 7: aiocean.user.v1.UserService.UpdateUser:input_type -> aiocean.user.v1.UpdateUserRequest
	4,  // 8: aiocean.user.v1.UserService.DeleteUser:input_type -> aiocean.user.v1.DeleteUserRequest
	9,  // 9: aiocean.user.v1.UserService.GetUserByOrn:input_type -> aiocean.user.v1.GetUserByOrnRequest
	2,  // 10: aiocean.user.v1.UserService.LinkExternalAccount:input_type -> aiocean.user.v1.LinkExternalAccountRequest
	0,  // 11: aiocean.user.v1.UserService.UnlinkExternalAccount:input_type -> aiocean.user.v1.UnlinkExternalAccountRequest
	8,  // 12: aiocean.user.v1.UserService.CreateUser:output_type -> aiocean.user.v1.CreateUserResponse
	7,  // 13: aiocean.user.v1.UserService.UpdateUser:output_type -> aiocean.user.v1.UpdateUserResponse
	13, // 14: aiocean.user.v1.UserService.DeleteUser:output_type -> google.protobuf.Empty
	10, // 15: aiocean.user.v1.UserService.GetUserByOrn:output_type -> aiocean.user.v1.GetUserByOrnResponse
	3,  // 16: aiocean.user.v1.UserService.LinkExternalAccount:output_type -> aiocean.user.v1.LinkExternalAccountResponse
	1,  // 17: aiocean.user.v1.UserService.UnlinkExternalAccount:output_type -> aiocean.user.v1.UnlinkExternalAccountResponse
	12, // [12:18] is the sub-list for method output_type
	6,  // [6:12] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_aiocean_user_v1_user_service_proto_init() }
func file_aiocean_user_v1_user_service_proto_init() {
	if File_aiocean_user_v1_user_service_proto != nil {
		return
	}
	file_aiocean_user_v1_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_aiocean_user_v1_user_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnlinkExternalAccountRequest); i {
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
		file_aiocean_user_v1_user_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnlinkExternalAccountResponse); i {
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
		file_aiocean_user_v1_user_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LinkExternalAccountRequest); i {
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
		file_aiocean_user_v1_user_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LinkExternalAccountResponse); i {
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
		file_aiocean_user_v1_user_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserRequest); i {
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
		file_aiocean_user_v1_user_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserRequest); i {
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
		file_aiocean_user_v1_user_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserRequest); i {
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
		file_aiocean_user_v1_user_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserResponse); i {
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
		file_aiocean_user_v1_user_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserResponse); i {
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
		file_aiocean_user_v1_user_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserByOrnRequest); i {
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
		file_aiocean_user_v1_user_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserByOrnResponse); i {
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
			RawDescriptor: file_aiocean_user_v1_user_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_aiocean_user_v1_user_service_proto_goTypes,
		DependencyIndexes: file_aiocean_user_v1_user_service_proto_depIdxs,
		MessageInfos:      file_aiocean_user_v1_user_service_proto_msgTypes,
	}.Build()
	File_aiocean_user_v1_user_service_proto = out.File
	file_aiocean_user_v1_user_service_proto_rawDesc = nil
	file_aiocean_user_v1_user_service_proto_goTypes = nil
	file_aiocean_user_v1_user_service_proto_depIdxs = nil
}
