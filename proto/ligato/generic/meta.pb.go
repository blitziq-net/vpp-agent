// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v6.30.2
// source: ligato/generic/meta.proto

package generic

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type KnownModelsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Class string `protobuf:"bytes,1,opt,name=class,proto3" json:"class,omitempty"`
}

func (x *KnownModelsRequest) Reset() {
	*x = KnownModelsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ligato_generic_meta_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KnownModelsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KnownModelsRequest) ProtoMessage() {}

func (x *KnownModelsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ligato_generic_meta_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KnownModelsRequest.ProtoReflect.Descriptor instead.
func (*KnownModelsRequest) Descriptor() ([]byte, []int) {
	return file_ligato_generic_meta_proto_rawDescGZIP(), []int{0}
}

func (x *KnownModelsRequest) GetClass() string {
	if x != nil {
		return x.Class
	}
	return ""
}

type ProtoFileDescriptorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// full_proto_file_name is full name of proto file that is needed to identify it. It has the form
	// "<proto package name ('.' replaced with '/')>/<simple file name>" (i.e. for this proto model
	// it is "ligato/generic/meta.proto").
	// If you are using rpc ProtoFileDescriptor for additional information retrieve for known models from
	// rpc KnownModels call, you can use usually present ModelDetail's generic.ModelDetail_Option for
	// key "protoFile" that is containing full proto file name in correct format.
	FullProtoFileName string `protobuf:"bytes,1,opt,name=full_proto_file_name,json=fullProtoFileName,proto3" json:"full_proto_file_name,omitempty"`
}

func (x *ProtoFileDescriptorRequest) Reset() {
	*x = ProtoFileDescriptorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ligato_generic_meta_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoFileDescriptorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoFileDescriptorRequest) ProtoMessage() {}

func (x *ProtoFileDescriptorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ligato_generic_meta_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoFileDescriptorRequest.ProtoReflect.Descriptor instead.
func (*ProtoFileDescriptorRequest) Descriptor() ([]byte, []int) {
	return file_ligato_generic_meta_proto_rawDescGZIP(), []int{1}
}

func (x *ProtoFileDescriptorRequest) GetFullProtoFileName() string {
	if x != nil {
		return x.FullProtoFileName
	}
	return ""
}

type KnownModelsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KnownModels   []*ModelDetail `protobuf:"bytes,1,rep,name=known_models,json=knownModels,proto3" json:"known_models,omitempty"`
	ActiveModules []string       `protobuf:"bytes,2,rep,name=active_modules,json=activeModules,proto3" json:"active_modules,omitempty"`
}

func (x *KnownModelsResponse) Reset() {
	*x = KnownModelsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ligato_generic_meta_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KnownModelsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KnownModelsResponse) ProtoMessage() {}

func (x *KnownModelsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ligato_generic_meta_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KnownModelsResponse.ProtoReflect.Descriptor instead.
func (*KnownModelsResponse) Descriptor() ([]byte, []int) {
	return file_ligato_generic_meta_proto_rawDescGZIP(), []int{2}
}

func (x *KnownModelsResponse) GetKnownModels() []*ModelDetail {
	if x != nil {
		return x.KnownModels
	}
	return nil
}

func (x *KnownModelsResponse) GetActiveModules() []string {
	if x != nil {
		return x.ActiveModules
	}
	return nil
}

type ProtoFileDescriptorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// file_descriptor is proto message representing proto file descriptor
	FileDescriptor *descriptorpb.FileDescriptorProto `protobuf:"bytes,1,opt,name=file_descriptor,json=fileDescriptor,proto3" json:"file_descriptor,omitempty"`
	// file_import_descriptors is set of file descriptors that the file_descriptor is using as import. This
	// is needed when converting file descriptor proto to protoreflect.FileDescriptor (using
	// "google.golang.org/protobuf/reflect/protodesc".NewFile(...) )
	FileImportDescriptors *descriptorpb.FileDescriptorSet `protobuf:"bytes,2,opt,name=file_import_descriptors,json=fileImportDescriptors,proto3" json:"file_import_descriptors,omitempty"`
}

func (x *ProtoFileDescriptorResponse) Reset() {
	*x = ProtoFileDescriptorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ligato_generic_meta_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoFileDescriptorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoFileDescriptorResponse) ProtoMessage() {}

func (x *ProtoFileDescriptorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ligato_generic_meta_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoFileDescriptorResponse.ProtoReflect.Descriptor instead.
func (*ProtoFileDescriptorResponse) Descriptor() ([]byte, []int) {
	return file_ligato_generic_meta_proto_rawDescGZIP(), []int{3}
}

func (x *ProtoFileDescriptorResponse) GetFileDescriptor() *descriptorpb.FileDescriptorProto {
	if x != nil {
		return x.FileDescriptor
	}
	return nil
}

func (x *ProtoFileDescriptorResponse) GetFileImportDescriptors() *descriptorpb.FileDescriptorSet {
	if x != nil {
		return x.FileImportDescriptors
	}
	return nil
}

var File_ligato_generic_meta_proto protoreflect.FileDescriptor

var file_ligato_generic_meta_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6c, 0x69, 0x67, 0x61, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63,
	0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6c, 0x69, 0x67,
	0x61, 0x74, 0x6f, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x1a, 0x20, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x6c,
	0x69, 0x67, 0x61, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2a, 0x0a, 0x12, 0x4b, 0x6e, 0x6f,
	0x77, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x22, 0x4d, 0x0a, 0x1a, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x46, 0x69,
	0x6c, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x14, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x11, 0x66, 0x75, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x46, 0x69, 0x6c, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0x7c, 0x0a, 0x13, 0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0c, 0x6b,
	0x6e, 0x6f, 0x77, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x6c, 0x69, 0x67, 0x61, 0x74, 0x6f, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x69, 0x63, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x0b,
	0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x4d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x73, 0x22, 0xc8, 0x01, 0x0a, 0x1b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x46, 0x69, 0x6c, 0x65,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x52, 0x0e, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f,
	0x72, 0x12, 0x5a, 0x0a, 0x17, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74,
	0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x52, 0x15, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x6d, 0x70, 0x6f,
	0x72, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x73, 0x32, 0xd5, 0x01,
	0x0a, 0x0b, 0x4d, 0x65, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x56, 0x0a,
	0x0b, 0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x12, 0x22, 0x2e, 0x6c,
	0x69, 0x67, 0x61, 0x74, 0x6f, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x2e, 0x4b, 0x6e,
	0x6f, 0x77, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x23, 0x2e, 0x6c, 0x69, 0x67, 0x61, 0x74, 0x6f, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69,
	0x63, 0x2e, 0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6e, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x46, 0x69,
	0x6c, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x2a, 0x2e, 0x6c,
	0x69, 0x67, 0x61, 0x74, 0x6f, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x2e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x6c, 0x69, 0x67, 0x61, 0x74,
	0x6f, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x46,
	0x69, 0x6c, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x6f, 0x2e, 0x6c, 0x69, 0x67, 0x61,
	0x74, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x70, 0x70, 0x2d, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f,
	0x76, 0x33, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x69, 0x67, 0x61, 0x74, 0x6f, 0x2f,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ligato_generic_meta_proto_rawDescOnce sync.Once
	file_ligato_generic_meta_proto_rawDescData = file_ligato_generic_meta_proto_rawDesc
)

func file_ligato_generic_meta_proto_rawDescGZIP() []byte {
	file_ligato_generic_meta_proto_rawDescOnce.Do(func() {
		file_ligato_generic_meta_proto_rawDescData = protoimpl.X.CompressGZIP(file_ligato_generic_meta_proto_rawDescData)
	})
	return file_ligato_generic_meta_proto_rawDescData
}

var file_ligato_generic_meta_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_ligato_generic_meta_proto_goTypes = []interface{}{
	(*KnownModelsRequest)(nil),               // 0: ligato.generic.KnownModelsRequest
	(*ProtoFileDescriptorRequest)(nil),       // 1: ligato.generic.ProtoFileDescriptorRequest
	(*KnownModelsResponse)(nil),              // 2: ligato.generic.KnownModelsResponse
	(*ProtoFileDescriptorResponse)(nil),      // 3: ligato.generic.ProtoFileDescriptorResponse
	(*ModelDetail)(nil),                      // 4: ligato.generic.ModelDetail
	(*descriptorpb.FileDescriptorProto)(nil), // 5: google.protobuf.FileDescriptorProto
	(*descriptorpb.FileDescriptorSet)(nil),   // 6: google.protobuf.FileDescriptorSet
}
var file_ligato_generic_meta_proto_depIdxs = []int32{
	4, // 0: ligato.generic.KnownModelsResponse.known_models:type_name -> ligato.generic.ModelDetail
	5, // 1: ligato.generic.ProtoFileDescriptorResponse.file_descriptor:type_name -> google.protobuf.FileDescriptorProto
	6, // 2: ligato.generic.ProtoFileDescriptorResponse.file_import_descriptors:type_name -> google.protobuf.FileDescriptorSet
	0, // 3: ligato.generic.MetaService.KnownModels:input_type -> ligato.generic.KnownModelsRequest
	1, // 4: ligato.generic.MetaService.ProtoFileDescriptor:input_type -> ligato.generic.ProtoFileDescriptorRequest
	2, // 5: ligato.generic.MetaService.KnownModels:output_type -> ligato.generic.KnownModelsResponse
	3, // 6: ligato.generic.MetaService.ProtoFileDescriptor:output_type -> ligato.generic.ProtoFileDescriptorResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_ligato_generic_meta_proto_init() }
func file_ligato_generic_meta_proto_init() {
	if File_ligato_generic_meta_proto != nil {
		return
	}
	file_ligato_generic_model_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ligato_generic_meta_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KnownModelsRequest); i {
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
		file_ligato_generic_meta_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoFileDescriptorRequest); i {
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
		file_ligato_generic_meta_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KnownModelsResponse); i {
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
		file_ligato_generic_meta_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoFileDescriptorResponse); i {
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
			RawDescriptor: file_ligato_generic_meta_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ligato_generic_meta_proto_goTypes,
		DependencyIndexes: file_ligato_generic_meta_proto_depIdxs,
		MessageInfos:      file_ligato_generic_meta_proto_msgTypes,
	}.Build()
	File_ligato_generic_meta_proto = out.File
	file_ligato_generic_meta_proto_rawDesc = nil
	file_ligato_generic_meta_proto_goTypes = nil
	file_ligato_generic_meta_proto_depIdxs = nil
}
