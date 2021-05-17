// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.2
// source: aiocean/polvo/v1/polvo.proto

package aiocean_polvo_v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type Package_Status int32

const (
	Package_STATUS_ACTIVE_UNSPECIFIED Package_Status = 0
	Package_STATUS_INACTIVE           Package_Status = 1
)

// Enum value maps for Package_Status.
var (
	Package_Status_name = map[int32]string{
		0: "STATUS_ACTIVE_UNSPECIFIED",
		1: "STATUS_INACTIVE",
	}
	Package_Status_value = map[string]int32{
		"STATUS_ACTIVE_UNSPECIFIED": 0,
		"STATUS_INACTIVE":           1,
	}
)

func (x Package_Status) Enum() *Package_Status {
	p := new(Package_Status)
	*p = x
	return p
}

func (x Package_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Package_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_aiocean_polvo_v1_polvo_proto_enumTypes[0].Descriptor()
}

func (Package_Status) Type() protoreflect.EnumType {
	return &file_aiocean_polvo_v1_polvo_proto_enumTypes[0]
}

func (x Package_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Package_Status.Descriptor instead.
func (Package_Status) EnumDescriptor() ([]byte, []int) {
	return file_aiocean_polvo_v1_polvo_proto_rawDescGZIP(), []int{0, 0}
}

type Build_Status int32

const (
	Build_STATUS_STARTED_UNSPECIFIED Build_Status = 0
	Build_STATUS_PROCESSING          Build_Status = 1
	Build_STATUS_STOPPED             Build_Status = 2
	Build_STATUS_FAILED              Build_Status = 3
	Build_STATUS_SUCCESS             Build_Status = 4
)

// Enum value maps for Build_Status.
var (
	Build_Status_name = map[int32]string{
		0: "STATUS_STARTED_UNSPECIFIED",
		1: "STATUS_PROCESSING",
		2: "STATUS_STOPPED",
		3: "STATUS_FAILED",
		4: "STATUS_SUCCESS",
	}
	Build_Status_value = map[string]int32{
		"STATUS_STARTED_UNSPECIFIED": 0,
		"STATUS_PROCESSING":          1,
		"STATUS_STOPPED":             2,
		"STATUS_FAILED":              3,
		"STATUS_SUCCESS":             4,
	}
)

func (x Build_Status) Enum() *Build_Status {
	p := new(Build_Status)
	*p = x
	return p
}

func (x Build_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Build_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_aiocean_polvo_v1_polvo_proto_enumTypes[1].Descriptor()
}

func (Build_Status) Type() protoreflect.EnumType {
	return &file_aiocean_polvo_v1_polvo_proto_enumTypes[1]
}

func (x Build_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Build_Status.Descriptor instead.
func (Build_Status) EnumDescriptor() ([]byte, []int) {
	return file_aiocean_polvo_v1_polvo_proto_rawDescGZIP(), []int{3, 0}
}

type Package struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orn         string         `protobuf:"bytes,1,opt,name=orn,proto3" json:"orn,omitempty"`
	DisplayName string         `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"` // Dashboard
	Status      Package_Status `protobuf:"varint,3,opt,name=status,proto3,enum=aiocean.polvo.v1.Package_Status" json:"status,omitempty"`
	Maintainer  string         `protobuf:"bytes,4,opt,name=maintainer,proto3" json:"maintainer,omitempty"`
}

func (x *Package) Reset() {
	*x = Package{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aiocean_polvo_v1_polvo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Package) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Package) ProtoMessage() {}

func (x *Package) ProtoReflect() protoreflect.Message {
	mi := &file_aiocean_polvo_v1_polvo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Package.ProtoReflect.Descriptor instead.
func (*Package) Descriptor() ([]byte, []int) {
	return file_aiocean_polvo_v1_polvo_proto_rawDescGZIP(), []int{0}
}

func (x *Package) GetOrn() string {
	if x != nil {
		return x.Orn
	}
	return ""
}

func (x *Package) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Package) GetStatus() Package_Status {
	if x != nil {
		return x.Status
	}
	return Package_STATUS_ACTIVE_UNSPECIFIED
}

func (x *Package) GetMaintainer() string {
	if x != nil {
		return x.Maintainer
	}
	return ""
}

type Version struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orn           string    `protobuf:"bytes,1,opt,name=orn,proto3" json:"orn,omitempty"`
	DisplayName   string    `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"` // latest, v1.0.0, v1.0.0+customer-1
	EntryPointUrl string    `protobuf:"bytes,3,opt,name=entry_point_url,json=entryPointUrl,proto3" json:"entry_point_url,omitempty"`
	Modules       []*Module `protobuf:"bytes,4,rep,name=modules,proto3" json:"modules,omitempty"`
}

func (x *Version) Reset() {
	*x = Version{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aiocean_polvo_v1_polvo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Version) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Version) ProtoMessage() {}

func (x *Version) ProtoReflect() protoreflect.Message {
	mi := &file_aiocean_polvo_v1_polvo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Version.ProtoReflect.Descriptor instead.
func (*Version) Descriptor() ([]byte, []int) {
	return file_aiocean_polvo_v1_polvo_proto_rawDescGZIP(), []int{1}
}

func (x *Version) GetOrn() string {
	if x != nil {
		return x.Orn
	}
	return ""
}

func (x *Version) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Version) GetEntryPointUrl() string {
	if x != nil {
		return x.EntryPointUrl
	}
	return ""
}

func (x *Version) GetModules() []*Module {
	if x != nil {
		return x.Modules
	}
	return nil
}

type Module struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"` // DashboardView
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"` // ./DashboardView
}

func (x *Module) Reset() {
	*x = Module{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aiocean_polvo_v1_polvo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Module) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Module) ProtoMessage() {}

func (x *Module) ProtoReflect() protoreflect.Message {
	mi := &file_aiocean_polvo_v1_polvo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Module.ProtoReflect.Descriptor instead.
func (*Module) Descriptor() ([]byte, []int) {
	return file_aiocean_polvo_v1_polvo_proto_rawDescGZIP(), []int{2}
}

func (x *Module) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Module) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type Build struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orn    string       `protobuf:"bytes,1,opt,name=orn,proto3" json:"orn,omitempty"`
	Status Build_Status `protobuf:"varint,2,opt,name=status,proto3,enum=aiocean.polvo.v1.Build_Status" json:"status,omitempty"`
}

func (x *Build) Reset() {
	*x = Build{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aiocean_polvo_v1_polvo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Build) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Build) ProtoMessage() {}

func (x *Build) ProtoReflect() protoreflect.Message {
	mi := &file_aiocean_polvo_v1_polvo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Build.ProtoReflect.Descriptor instead.
func (*Build) Descriptor() ([]byte, []int) {
	return file_aiocean_polvo_v1_polvo_proto_rawDescGZIP(), []int{3}
}

func (x *Build) GetOrn() string {
	if x != nil {
		return x.Orn
	}
	return ""
}

func (x *Build) GetStatus() Build_Status {
	if x != nil {
		return x.Status
	}
	return Build_STATUS_STARTED_UNSPECIFIED
}

var File_aiocean_polvo_v1_polvo_proto protoreflect.FileDescriptor

var file_aiocean_polvo_v1_polvo_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x69, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x2f, 0x70, 0x6f, 0x6c, 0x76, 0x6f, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x6f, 0x6c, 0x76, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10,
	0x61, 0x69, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x2e, 0x70, 0x6f, 0x6c, 0x76, 0x6f, 0x2e, 0x76, 0x31,
	0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9d, 0x02, 0x0a, 0x07, 0x50, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x57, 0x0a, 0x03, 0x6f, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x45, 0xfa, 0x42, 0x42, 0x72, 0x40, 0x32, 0x3e, 0x28, 0x3f, 0x6d, 0x29, 0x5e,
	0x70, 0x6f, 0x6c, 0x76, 0x6f, 0x2e, 0x61, 0x69, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x5b, 0x5e, 0x2f, 0x5d, 0x2b, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x73, 0x2f, 0x5b, 0x5e, 0x2f, 0x5d, 0x2b, 0x24, 0x52, 0x03, 0x6f, 0x72, 0x6e, 0x12, 0x21,
	0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x38, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x20, 0x2e, 0x61, 0x69, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x2e, 0x70, 0x6f, 0x6c, 0x76,
	0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x6d,
	0x61, 0x69, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x6d, 0x61, 0x69, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x22, 0x3c, 0x0a, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x19, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x49,
	0x4e, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x01, 0x22, 0xd6, 0x02, 0x0a, 0x07, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x66, 0x0a, 0x03, 0x6f, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x54, 0xfa, 0x42, 0x51, 0x72, 0x4f, 0x32, 0x4d, 0x28, 0x3f, 0x6d, 0x29, 0x5e,
	0x70, 0x6f, 0x6c, 0x76, 0x6f, 0x2e, 0x61, 0x69, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x5b, 0x5e, 0x2f, 0x5d, 0x2b, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x73, 0x2f, 0x5b, 0x5e, 0x2f, 0x5d, 0x2b, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x5b, 0x5e, 0x2f, 0x5d, 0x2b, 0x24, 0x52, 0x03, 0x6f, 0x72, 0x6e, 0x12, 0x21, 0x0a,
	0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x8b, 0x01, 0x0a, 0x0f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x63, 0xfa, 0x42, 0x60, 0x72,
	0x5e, 0x32, 0x5c, 0x28, 0x3f, 0x6d, 0x29, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x70,
	0x6f, 0x6c, 0x76, 0x6f, 0x2e, 0x61, 0x69, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x5b, 0x5e, 0x2f, 0x5d, 0x2b, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x73, 0x2f, 0x5b, 0x5e, 0x2f, 0x5d, 0x2b, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x5b, 0x5e, 0x2f, 0x5d, 0x2b, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x6a, 0x73, 0x52,
	0x0d, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x32,
	0x0a, 0x07, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x61, 0x69, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x2e, 0x70, 0x6f, 0x6c, 0x76, 0x6f, 0x2e,
	0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x07, 0x6d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x73, 0x22, 0x30, 0x0a, 0x06, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x22, 0xb0, 0x02, 0x0a, 0x05, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x12, 0x73,
	0x0a, 0x03, 0x6f, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x61, 0xfa, 0x42, 0x5e,
	0x72, 0x5c, 0x32, 0x5a, 0x28, 0x3f, 0x6d, 0x29, 0x5e, 0x70, 0x6f, 0x6c, 0x76, 0x6f, 0x2e, 0x61,
	0x69, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x5b, 0x5e, 0x2f,
	0x5d, 0x2b, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x5b, 0x5e, 0x2f, 0x5d,
	0x2b, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x5b, 0x5e, 0x2f, 0x5d, 0x2b,
	0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x2f, 0x5b, 0x5e, 0x2f, 0x5d, 0x2b, 0x24, 0x52, 0x03,
	0x6f, 0x72, 0x6e, 0x12, 0x36, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x61, 0x69, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x2e, 0x70, 0x6f,
	0x6c, 0x76, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x7a, 0x0a, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x1a, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x53, 0x54, 0x41, 0x52, 0x54, 0x45, 0x44, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x54, 0x4f, 0x50, 0x50, 0x45, 0x44, 0x10, 0x02,
	0x12, 0x11, 0x0a, 0x0d, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45,
	0x44, 0x10, 0x03, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x55,
	0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x04, 0x42, 0x33, 0x5a, 0x31, 0x70, 0x6b, 0x67, 0x2e, 0x61,
	0x69, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x67, 0x6f, 0x2f, 0x70, 0x6f, 0x6c, 0x76, 0x6f, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x69, 0x6f, 0x63,
	0x65, 0x61, 0x6e, 0x5f, 0x70, 0x6f, 0x6c, 0x76, 0x6f, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aiocean_polvo_v1_polvo_proto_rawDescOnce sync.Once
	file_aiocean_polvo_v1_polvo_proto_rawDescData = file_aiocean_polvo_v1_polvo_proto_rawDesc
)

func file_aiocean_polvo_v1_polvo_proto_rawDescGZIP() []byte {
	file_aiocean_polvo_v1_polvo_proto_rawDescOnce.Do(func() {
		file_aiocean_polvo_v1_polvo_proto_rawDescData = protoimpl.X.CompressGZIP(file_aiocean_polvo_v1_polvo_proto_rawDescData)
	})
	return file_aiocean_polvo_v1_polvo_proto_rawDescData
}

var file_aiocean_polvo_v1_polvo_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_aiocean_polvo_v1_polvo_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_aiocean_polvo_v1_polvo_proto_goTypes = []interface{}{
	(Package_Status)(0), // 0: aiocean.polvo.v1.Package.Status
	(Build_Status)(0),   // 1: aiocean.polvo.v1.Build.Status
	(*Package)(nil),     // 2: aiocean.polvo.v1.Package
	(*Version)(nil),     // 3: aiocean.polvo.v1.Version
	(*Module)(nil),      // 4: aiocean.polvo.v1.Module
	(*Build)(nil),       // 5: aiocean.polvo.v1.Build
}
var file_aiocean_polvo_v1_polvo_proto_depIdxs = []int32{
	0, // 0: aiocean.polvo.v1.Package.status:type_name -> aiocean.polvo.v1.Package.Status
	4, // 1: aiocean.polvo.v1.Version.modules:type_name -> aiocean.polvo.v1.Module
	1, // 2: aiocean.polvo.v1.Build.status:type_name -> aiocean.polvo.v1.Build.Status
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_aiocean_polvo_v1_polvo_proto_init() }
func file_aiocean_polvo_v1_polvo_proto_init() {
	if File_aiocean_polvo_v1_polvo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_aiocean_polvo_v1_polvo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Package); i {
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
		file_aiocean_polvo_v1_polvo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Version); i {
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
		file_aiocean_polvo_v1_polvo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Module); i {
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
		file_aiocean_polvo_v1_polvo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Build); i {
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
			RawDescriptor: file_aiocean_polvo_v1_polvo_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_aiocean_polvo_v1_polvo_proto_goTypes,
		DependencyIndexes: file_aiocean_polvo_v1_polvo_proto_depIdxs,
		EnumInfos:         file_aiocean_polvo_v1_polvo_proto_enumTypes,
		MessageInfos:      file_aiocean_polvo_v1_polvo_proto_msgTypes,
	}.Build()
	File_aiocean_polvo_v1_polvo_proto = out.File
	file_aiocean_polvo_v1_polvo_proto_rawDesc = nil
	file_aiocean_polvo_v1_polvo_proto_goTypes = nil
	file_aiocean_polvo_v1_polvo_proto_depIdxs = nil
}
