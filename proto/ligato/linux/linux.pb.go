// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v6.30.2
// source: ligato/linux/linux.proto

package linux

import (
	interfaces "go.ligato.io/vpp-agent/v3/proto/ligato/linux/interfaces"
	l3 "go.ligato.io/vpp-agent/v3/proto/ligato/linux/l3"
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

type ConfigData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Interfaces []*interfaces.Interface `protobuf:"bytes,10,rep,name=interfaces,proto3" json:"interfaces,omitempty"`
	ArpEntries []*l3.ARPEntry          `protobuf:"bytes,20,rep,name=arp_entries,json=arpEntries,proto3" json:"arp_entries,omitempty"`
	Routes     []*l3.Route             `protobuf:"bytes,21,rep,name=routes,proto3" json:"routes,omitempty"`
}

func (x *ConfigData) Reset() {
	*x = ConfigData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ligato_linux_linux_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigData) ProtoMessage() {}

func (x *ConfigData) ProtoReflect() protoreflect.Message {
	mi := &file_ligato_linux_linux_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigData.ProtoReflect.Descriptor instead.
func (*ConfigData) Descriptor() ([]byte, []int) {
	return file_ligato_linux_linux_proto_rawDescGZIP(), []int{0}
}

func (x *ConfigData) GetInterfaces() []*interfaces.Interface {
	if x != nil {
		return x.Interfaces
	}
	return nil
}

func (x *ConfigData) GetArpEntries() []*l3.ARPEntry {
	if x != nil {
		return x.ArpEntries
	}
	return nil
}

func (x *ConfigData) GetRoutes() []*l3.Route {
	if x != nil {
		return x.Routes
	}
	return nil
}

type Notification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Interface *interfaces.InterfaceNotification `protobuf:"bytes,1,opt,name=interface,proto3" json:"interface,omitempty"`
}

func (x *Notification) Reset() {
	*x = Notification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ligato_linux_linux_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Notification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notification) ProtoMessage() {}

func (x *Notification) ProtoReflect() protoreflect.Message {
	mi := &file_ligato_linux_linux_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notification.ProtoReflect.Descriptor instead.
func (*Notification) Descriptor() ([]byte, []int) {
	return file_ligato_linux_linux_proto_rawDescGZIP(), []int{1}
}

func (x *Notification) GetInterface() *interfaces.InterfaceNotification {
	if x != nil {
		return x.Interface
	}
	return nil
}

var File_ligato_linux_linux_proto protoreflect.FileDescriptor

var file_ligato_linux_linux_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6c, 0x69, 0x67, 0x61, 0x74, 0x6f, 0x2f, 0x6c, 0x69, 0x6e, 0x75, 0x78, 0x2f, 0x6c,
	0x69, 0x6e, 0x75, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6c, 0x69, 0x67, 0x61,
	0x74, 0x6f, 0x2e, 0x6c, 0x69, 0x6e, 0x75, 0x78, 0x1a, 0x27, 0x6c, 0x69, 0x67, 0x61, 0x74, 0x6f,
	0x2f, 0x6c, 0x69, 0x6e, 0x75, 0x78, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x23, 0x6c, 0x69, 0x67, 0x61, 0x74, 0x6f, 0x2f, 0x6c, 0x69, 0x6e, 0x75, 0x78, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x6c, 0x69, 0x67, 0x61, 0x74, 0x6f, 0x2f, 0x6c,
	0x69, 0x6e, 0x75, 0x78, 0x2f, 0x6c, 0x33, 0x2f, 0x61, 0x72, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x6c, 0x69, 0x67, 0x61, 0x74, 0x6f, 0x2f, 0x6c, 0x69, 0x6e, 0x75, 0x78, 0x2f,
	0x6c, 0x33, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbc,
	0x01, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x44, 0x61, 0x74, 0x61, 0x12, 0x42, 0x0a,
	0x0a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x6c, 0x69, 0x67, 0x61, 0x74, 0x6f, 0x2e, 0x6c, 0x69, 0x6e, 0x75, 0x78,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x0a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x73, 0x12, 0x3a, 0x0a, 0x0b, 0x61, 0x72, 0x70, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x18, 0x14, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6c, 0x69, 0x67, 0x61, 0x74, 0x6f, 0x2e,
	0x6c, 0x69, 0x6e, 0x75, 0x78, 0x2e, 0x6c, 0x33, 0x2e, 0x41, 0x52, 0x50, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x0a, 0x61, 0x72, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x2e, 0x0a,
	0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x15, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x6c, 0x69, 0x67, 0x61, 0x74, 0x6f, 0x2e, 0x6c, 0x69, 0x6e, 0x75, 0x78, 0x2e, 0x6c, 0x33, 0x2e,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x22, 0x5c, 0x0a,
	0x0c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4c, 0x0a,
	0x09, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2e, 0x2e, 0x6c, 0x69, 0x67, 0x61, 0x74, 0x6f, 0x2e, 0x6c, 0x69, 0x6e, 0x75, 0x78, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x09, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x42, 0x34, 0x5a, 0x32, 0x67,
	0x6f, 0x2e, 0x6c, 0x69, 0x67, 0x61, 0x74, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x70, 0x70, 0x2d,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x33, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c,
	0x69, 0x67, 0x61, 0x74, 0x6f, 0x2f, 0x6c, 0x69, 0x6e, 0x75, 0x78, 0x3b, 0x6c, 0x69, 0x6e, 0x75,
	0x78, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ligato_linux_linux_proto_rawDescOnce sync.Once
	file_ligato_linux_linux_proto_rawDescData = file_ligato_linux_linux_proto_rawDesc
)

func file_ligato_linux_linux_proto_rawDescGZIP() []byte {
	file_ligato_linux_linux_proto_rawDescOnce.Do(func() {
		file_ligato_linux_linux_proto_rawDescData = protoimpl.X.CompressGZIP(file_ligato_linux_linux_proto_rawDescData)
	})
	return file_ligato_linux_linux_proto_rawDescData
}

var file_ligato_linux_linux_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ligato_linux_linux_proto_goTypes = []interface{}{
	(*ConfigData)(nil),                       // 0: ligato.linux.ConfigData
	(*Notification)(nil),                     // 1: ligato.linux.Notification
	(*interfaces.Interface)(nil),             // 2: ligato.linux.interfaces.Interface
	(*l3.ARPEntry)(nil),                      // 3: ligato.linux.l3.ARPEntry
	(*l3.Route)(nil),                         // 4: ligato.linux.l3.Route
	(*interfaces.InterfaceNotification)(nil), // 5: ligato.linux.interfaces.InterfaceNotification
}
var file_ligato_linux_linux_proto_depIdxs = []int32{
	2, // 0: ligato.linux.ConfigData.interfaces:type_name -> ligato.linux.interfaces.Interface
	3, // 1: ligato.linux.ConfigData.arp_entries:type_name -> ligato.linux.l3.ARPEntry
	4, // 2: ligato.linux.ConfigData.routes:type_name -> ligato.linux.l3.Route
	5, // 3: ligato.linux.Notification.interface:type_name -> ligato.linux.interfaces.InterfaceNotification
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_ligato_linux_linux_proto_init() }
func file_ligato_linux_linux_proto_init() {
	if File_ligato_linux_linux_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ligato_linux_linux_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigData); i {
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
		file_ligato_linux_linux_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Notification); i {
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
			RawDescriptor: file_ligato_linux_linux_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ligato_linux_linux_proto_goTypes,
		DependencyIndexes: file_ligato_linux_linux_proto_depIdxs,
		MessageInfos:      file_ligato_linux_linux_proto_msgTypes,
	}.Build()
	File_ligato_linux_linux_proto = out.File
	file_ligato_linux_linux_proto_rawDesc = nil
	file_ligato_linux_linux_proto_goTypes = nil
	file_ligato_linux_linux_proto_depIdxs = nil
}
