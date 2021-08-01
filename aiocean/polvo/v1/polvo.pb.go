// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.2
// source: aiocean/polvo/v1/polvo.proto

package v1

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

type Application struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid    string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Config string `protobuf:"bytes,3,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *Application) Reset() {
	*x = Application{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aiocean_polvo_v1_polvo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Application) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Application) ProtoMessage() {}

func (x *Application) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Application.ProtoReflect.Descriptor instead.
func (*Application) Descriptor() ([]byte, []int) {
	return file_aiocean_polvo_v1_polvo_proto_rawDescGZIP(), []int{0}
}

func (x *Application) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *Application) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Application) GetConfig() string {
	if x != nil {
		return x.Config
	}
	return ""
}

type Package struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid        string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Maintainer string `protobuf:"bytes,4,opt,name=maintainer,proto3" json:"maintainer,omitempty"`
}

func (x *Package) Reset() {
	*x = Package{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aiocean_polvo_v1_polvo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Package) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Package) ProtoMessage() {}

func (x *Package) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Package.ProtoReflect.Descriptor instead.
func (*Package) Descriptor() ([]byte, []int) {
	return file_aiocean_polvo_v1_polvo_proto_rawDescGZIP(), []int{1}
}

func (x *Package) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *Package) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
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

	Uid         string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"` // latest, v1.0.0, v1.0.0+customer-1
	ManifestUrl string `protobuf:"bytes,7,opt,name=manifest_url,json=manifestUrl,proto3" json:"manifest_url,omitempty"`
	Weight      int32  `protobuf:"varint,9,opt,name=weight,proto3" json:"weight,omitempty"`
}

func (x *Version) Reset() {
	*x = Version{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aiocean_polvo_v1_polvo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Version) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Version) ProtoMessage() {}

func (x *Version) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Version.ProtoReflect.Descriptor instead.
func (*Version) Descriptor() ([]byte, []int) {
	return file_aiocean_polvo_v1_polvo_proto_rawDescGZIP(), []int{2}
}

func (x *Version) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *Version) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Version) GetManifestUrl() string {
	if x != nil {
		return x.ManifestUrl
	}
	return ""
}

func (x *Version) GetWeight() int32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

var File_aiocean_polvo_v1_polvo_proto protoreflect.FileDescriptor

var file_aiocean_polvo_v1_polvo_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x69, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x2f, 0x70, 0x6f, 0x6c, 0x76, 0x6f, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x6f, 0x6c, 0x76, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10,
	0x61, 0x69, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x2e, 0x70, 0x6f, 0x6c, 0x76, 0x6f, 0x2e, 0x76, 0x31,
	0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x54, 0x0a, 0x0b, 0x41, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10,
	0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22,
	0x63, 0x0a, 0x07, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72,
	0x02, 0x10, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x0a, 0x6d, 0x61, 0x69,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa,
	0x42, 0x06, 0x72, 0x04, 0x10, 0x01, 0x60, 0x01, 0x52, 0x0a, 0x6d, 0x61, 0x69, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x22, 0x85, 0x01, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x69, 0x64, 0x12, 0x1e, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x0a, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x01, 0xd0, 0x01, 0x01, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x30, 0x0a, 0x0c, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0xfa, 0x42, 0x0a, 0x72, 0x08, 0x10,
	0x01, 0x88, 0x01, 0x01, 0xd0, 0x01, 0x01, 0x52, 0x0b, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73,
	0x74, 0x55, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x42, 0x1c, 0x5a, 0x1a,
	0x70, 0x6b, 0x67, 0x2e, 0x61, 0x69, 0x6f, 0x63, 0x65, 0x61, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2f,
	0x70, 0x6f, 0x6c, 0x76, 0x6f, 0x67, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
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

var file_aiocean_polvo_v1_polvo_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_aiocean_polvo_v1_polvo_proto_goTypes = []interface{}{
	(*Application)(nil), // 0: aiocean.polvo.v1.Application
	(*Package)(nil),     // 1: aiocean.polvo.v1.Package
	(*Version)(nil),     // 2: aiocean.polvo.v1.Version
}
var file_aiocean_polvo_v1_polvo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_aiocean_polvo_v1_polvo_proto_init() }
func file_aiocean_polvo_v1_polvo_proto_init() {
	if File_aiocean_polvo_v1_polvo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_aiocean_polvo_v1_polvo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Application); i {
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
		file_aiocean_polvo_v1_polvo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_aiocean_polvo_v1_polvo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_aiocean_polvo_v1_polvo_proto_goTypes,
		DependencyIndexes: file_aiocean_polvo_v1_polvo_proto_depIdxs,
		MessageInfos:      file_aiocean_polvo_v1_polvo_proto_msgTypes,
	}.Build()
	File_aiocean_polvo_v1_polvo_proto = out.File
	file_aiocean_polvo_v1_polvo_proto_rawDesc = nil
	file_aiocean_polvo_v1_polvo_proto_goTypes = nil
	file_aiocean_polvo_v1_polvo_proto_depIdxs = nil
}
