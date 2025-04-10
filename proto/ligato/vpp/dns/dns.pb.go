// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v6.30.2
// source: ligato/vpp/dns/dns.proto

package vpp_dns

import (
	_ "go.ligato.io/vpp-agent/v3/proto/ligato"
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

// DNSCache configuration models VPP's DNS cache server functionality. The main goal of this functionality is
// to cache DNS records and minimize external DNS traffic.
// The presence of this configuration enables the VPP DNS functionality and VPP start to acts as DNS cache Server.
// It responds on standard DNS port(53) to DNS requests. Removing of this configuration disables the VPP DNS
// functionality.
type DNSCache struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of upstream DNS servers that are contacted by VPP when unknown domain name needs to be resolved.
	// The results are cached and there should be no further upstream DNS server request for the same domain
	// name until cached DNS record expiration.
	UpstreamDnsServers []string `protobuf:"bytes,1,rep,name=upstream_dns_servers,json=upstreamDnsServers,proto3" json:"upstream_dns_servers,omitempty"`
}

func (x *DNSCache) Reset() {
	*x = DNSCache{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ligato_vpp_dns_dns_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DNSCache) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DNSCache) ProtoMessage() {}

func (x *DNSCache) ProtoReflect() protoreflect.Message {
	mi := &file_ligato_vpp_dns_dns_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DNSCache.ProtoReflect.Descriptor instead.
func (*DNSCache) Descriptor() ([]byte, []int) {
	return file_ligato_vpp_dns_dns_proto_rawDescGZIP(), []int{0}
}

func (x *DNSCache) GetUpstreamDnsServers() []string {
	if x != nil {
		return x.UpstreamDnsServers
	}
	return nil
}

var File_ligato_vpp_dns_dns_proto protoreflect.FileDescriptor

var file_ligato_vpp_dns_dns_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6c, 0x69, 0x67, 0x61, 0x74, 0x6f, 0x2f, 0x76, 0x70, 0x70, 0x2f, 0x64, 0x6e, 0x73,
	0x2f, 0x64, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6c, 0x69, 0x67, 0x61,
	0x74, 0x6f, 0x2e, 0x76, 0x70, 0x70, 0x2e, 0x64, 0x6e, 0x73, 0x1a, 0x18, 0x6c, 0x69, 0x67, 0x61,
	0x74, 0x6f, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x43, 0x0a, 0x08, 0x44, 0x4e, 0x53, 0x43, 0x61, 0x63, 0x68, 0x65,
	0x12, 0x37, 0x0a, 0x14, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x64, 0x6e, 0x73,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x42, 0x05,
	0x82, 0x7d, 0x02, 0x08, 0x01, 0x52, 0x12, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x44,
	0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x6f, 0x2e,
	0x6c, 0x69, 0x67, 0x61, 0x74, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x70, 0x70, 0x2d, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x2f, 0x76, 0x33, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x69, 0x67,
	0x61, 0x74, 0x6f, 0x2f, 0x76, 0x70, 0x70, 0x2f, 0x64, 0x6e, 0x73, 0x3b, 0x76, 0x70, 0x70, 0x5f,
	0x64, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ligato_vpp_dns_dns_proto_rawDescOnce sync.Once
	file_ligato_vpp_dns_dns_proto_rawDescData = file_ligato_vpp_dns_dns_proto_rawDesc
)

func file_ligato_vpp_dns_dns_proto_rawDescGZIP() []byte {
	file_ligato_vpp_dns_dns_proto_rawDescOnce.Do(func() {
		file_ligato_vpp_dns_dns_proto_rawDescData = protoimpl.X.CompressGZIP(file_ligato_vpp_dns_dns_proto_rawDescData)
	})
	return file_ligato_vpp_dns_dns_proto_rawDescData
}

var file_ligato_vpp_dns_dns_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ligato_vpp_dns_dns_proto_goTypes = []interface{}{
	(*DNSCache)(nil), // 0: ligato.vpp.dns.DNSCache
}
var file_ligato_vpp_dns_dns_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ligato_vpp_dns_dns_proto_init() }
func file_ligato_vpp_dns_dns_proto_init() {
	if File_ligato_vpp_dns_dns_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ligato_vpp_dns_dns_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DNSCache); i {
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
			RawDescriptor: file_ligato_vpp_dns_dns_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ligato_vpp_dns_dns_proto_goTypes,
		DependencyIndexes: file_ligato_vpp_dns_dns_proto_depIdxs,
		MessageInfos:      file_ligato_vpp_dns_dns_proto_msgTypes,
	}.Build()
	File_ligato_vpp_dns_dns_proto = out.File
	file_ligato_vpp_dns_dns_proto_rawDesc = nil
	file_ligato_vpp_dns_dns_proto_goTypes = nil
	file_ligato_vpp_dns_dns_proto_depIdxs = nil
}
