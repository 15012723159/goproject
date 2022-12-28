// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: envoy/extensions/filters/udp/dns_filter/v3/dns_filter.proto

package dns_filterv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/envoyproxy/go-control-plane/envoy/annotations"
	v31 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	v3 "github.com/envoyproxy/go-control-plane/envoy/data/dns/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	duration "github.com/golang/protobuf/ptypes/duration"
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

// Configuration for the DNS filter.
type DnsFilterConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The stat prefix used when emitting DNS filter statistics
	StatPrefix string `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// Server context configuration contains the data that the filter uses to respond
	// to DNS requests.
	ServerConfig *DnsFilterConfig_ServerContextConfig `protobuf:"bytes,2,opt,name=server_config,json=serverConfig,proto3" json:"server_config,omitempty"`
	// Client context configuration controls Envoy's behavior when it must use external
	// resolvers to answer a query. This object is optional and if omitted instructs
	// the filter to resolve queries from the data in the server_config
	ClientConfig *DnsFilterConfig_ClientContextConfig `protobuf:"bytes,3,opt,name=client_config,json=clientConfig,proto3" json:"client_config,omitempty"`
}

func (x *DnsFilterConfig) Reset() {
	*x = DnsFilterConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_udp_dns_filter_v3_dns_filter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DnsFilterConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DnsFilterConfig) ProtoMessage() {}

func (x *DnsFilterConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_udp_dns_filter_v3_dns_filter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DnsFilterConfig.ProtoReflect.Descriptor instead.
func (*DnsFilterConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_udp_dns_filter_v3_dns_filter_proto_rawDescGZIP(), []int{0}
}

func (x *DnsFilterConfig) GetStatPrefix() string {
	if x != nil {
		return x.StatPrefix
	}
	return ""
}

func (x *DnsFilterConfig) GetServerConfig() *DnsFilterConfig_ServerContextConfig {
	if x != nil {
		return x.ServerConfig
	}
	return nil
}

func (x *DnsFilterConfig) GetClientConfig() *DnsFilterConfig_ClientContextConfig {
	if x != nil {
		return x.ClientConfig
	}
	return nil
}

// This message contains the configuration for the DNS Filter operating
// in a server context. This message will contain the virtual hosts and
// associated addresses with which Envoy will respond to queries
type DnsFilterConfig_ServerContextConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to ConfigSource:
	//	*DnsFilterConfig_ServerContextConfig_InlineDnsTable
	//	*DnsFilterConfig_ServerContextConfig_ExternalDnsTable
	ConfigSource isDnsFilterConfig_ServerContextConfig_ConfigSource `protobuf_oneof:"config_source"`
}

func (x *DnsFilterConfig_ServerContextConfig) Reset() {
	*x = DnsFilterConfig_ServerContextConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_udp_dns_filter_v3_dns_filter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DnsFilterConfig_ServerContextConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DnsFilterConfig_ServerContextConfig) ProtoMessage() {}

func (x *DnsFilterConfig_ServerContextConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_udp_dns_filter_v3_dns_filter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DnsFilterConfig_ServerContextConfig.ProtoReflect.Descriptor instead.
func (*DnsFilterConfig_ServerContextConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_udp_dns_filter_v3_dns_filter_proto_rawDescGZIP(), []int{0, 0}
}

func (m *DnsFilterConfig_ServerContextConfig) GetConfigSource() isDnsFilterConfig_ServerContextConfig_ConfigSource {
	if m != nil {
		return m.ConfigSource
	}
	return nil
}

func (x *DnsFilterConfig_ServerContextConfig) GetInlineDnsTable() *v3.DnsTable {
	if x, ok := x.GetConfigSource().(*DnsFilterConfig_ServerContextConfig_InlineDnsTable); ok {
		return x.InlineDnsTable
	}
	return nil
}

func (x *DnsFilterConfig_ServerContextConfig) GetExternalDnsTable() *v31.DataSource {
	if x, ok := x.GetConfigSource().(*DnsFilterConfig_ServerContextConfig_ExternalDnsTable); ok {
		return x.ExternalDnsTable
	}
	return nil
}

type isDnsFilterConfig_ServerContextConfig_ConfigSource interface {
	isDnsFilterConfig_ServerContextConfig_ConfigSource()
}

type DnsFilterConfig_ServerContextConfig_InlineDnsTable struct {
	// Load the configuration specified from the control plane
	InlineDnsTable *v3.DnsTable `protobuf:"bytes,1,opt,name=inline_dns_table,json=inlineDnsTable,proto3,oneof"`
}

type DnsFilterConfig_ServerContextConfig_ExternalDnsTable struct {
	// Seed the filter configuration from an external path. This source
	// is a yaml formatted file that contains the DnsTable driving Envoy's
	// responses to DNS queries
	ExternalDnsTable *v31.DataSource `protobuf:"bytes,2,opt,name=external_dns_table,json=externalDnsTable,proto3,oneof"`
}

func (*DnsFilterConfig_ServerContextConfig_InlineDnsTable) isDnsFilterConfig_ServerContextConfig_ConfigSource() {
}

func (*DnsFilterConfig_ServerContextConfig_ExternalDnsTable) isDnsFilterConfig_ServerContextConfig_ConfigSource() {
}

// This message contains the configuration for the DNS Filter operating
// in a client context. This message will contain the timeouts, retry,
// and forwarding configuration for Envoy to make DNS requests to other
// resolvers
//
// [#next-free-field: 6]
type DnsFilterConfig_ClientContextConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Sets the maximum time we will wait for the upstream query to complete
	// We allow 5s for the upstream resolution to complete, so the minimum
	// value here is 1. Note that the total latency for a failed query is the
	// number of retries multiplied by the resolver_timeout.
	ResolverTimeout *duration.Duration `protobuf:"bytes,1,opt,name=resolver_timeout,json=resolverTimeout,proto3" json:"resolver_timeout,omitempty"`
	// This field was used for `dns_resolution_config` in Envoy 1.19.0 and
	// 1.19.1.
	// Control planes that need to set this field for Envoy 1.19.0 and
	// 1.19.1 clients should fork the protobufs and change the field type
	// to `DnsResolutionConfig`.
	// Control planes that need to simultaneously support Envoy 1.18.x and
	// Envoy 1.19.x should avoid Envoy 1.19.0 and 1.19.1.
	//
	// [#not-implemented-hide:]
	//
	// Deprecated: Do not use.
	UpstreamResolvers []*v31.Address `protobuf:"bytes,2,rep,name=upstream_resolvers,json=upstreamResolvers,proto3" json:"upstream_resolvers,omitempty"`
	// DNS resolution configuration which includes the underlying dns resolver addresses and options.
	// This field is deprecated in favor of
	// :ref:`typed_dns_resolver_config <envoy_v3_api_field_extensions.filters.udp.dns_filter.v3.DnsFilterConfig.ClientContextConfig.typed_dns_resolver_config>`.
	//
	// Deprecated: Do not use.
	DnsResolutionConfig *v31.DnsResolutionConfig `protobuf:"bytes,5,opt,name=dns_resolution_config,json=dnsResolutionConfig,proto3" json:"dns_resolution_config,omitempty"`
	// DNS resolver type configuration extension. This extension can be used to configure c-ares, apple,
	// or any other DNS resolver types and the related parameters.
	// For example, an object of
	// :ref:`CaresDnsResolverConfig <envoy_v3_api_msg_extensions.network.dns_resolver.cares.v3.CaresDnsResolverConfig>`
	// can be packed into this *typed_dns_resolver_config*. This configuration replaces the
	// :ref:`dns_resolution_config <envoy_v3_api_field_extensions.filters.udp.dns_filter.v3.DnsFilterConfig.ClientContextConfig.dns_resolution_config>`
	// configuration.
	// During the transition period when both *dns_resolution_config* and *typed_dns_resolver_config* exists,
	// when *typed_dns_resolver_config* is in place, Envoy will use it and ignore *dns_resolution_config*.
	// When *typed_dns_resolver_config* is missing, the default behavior is in place.
	// [#extension-category: envoy.network.dns_resolver]
	TypedDnsResolverConfig *v31.TypedExtensionConfig `protobuf:"bytes,4,opt,name=typed_dns_resolver_config,json=typedDnsResolverConfig,proto3" json:"typed_dns_resolver_config,omitempty"`
	// Controls how many outstanding external lookup contexts the filter tracks.
	// The context structure allows the filter to respond to every query even if the external
	// resolution times out or is otherwise unsuccessful
	MaxPendingLookups uint64 `protobuf:"varint,3,opt,name=max_pending_lookups,json=maxPendingLookups,proto3" json:"max_pending_lookups,omitempty"`
}

func (x *DnsFilterConfig_ClientContextConfig) Reset() {
	*x = DnsFilterConfig_ClientContextConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_udp_dns_filter_v3_dns_filter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DnsFilterConfig_ClientContextConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DnsFilterConfig_ClientContextConfig) ProtoMessage() {}

func (x *DnsFilterConfig_ClientContextConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_udp_dns_filter_v3_dns_filter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DnsFilterConfig_ClientContextConfig.ProtoReflect.Descriptor instead.
func (*DnsFilterConfig_ClientContextConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_udp_dns_filter_v3_dns_filter_proto_rawDescGZIP(), []int{0, 1}
}

func (x *DnsFilterConfig_ClientContextConfig) GetResolverTimeout() *duration.Duration {
	if x != nil {
		return x.ResolverTimeout
	}
	return nil
}

// Deprecated: Do not use.
func (x *DnsFilterConfig_ClientContextConfig) GetUpstreamResolvers() []*v31.Address {
	if x != nil {
		return x.UpstreamResolvers
	}
	return nil
}

// Deprecated: Do not use.
func (x *DnsFilterConfig_ClientContextConfig) GetDnsResolutionConfig() *v31.DnsResolutionConfig {
	if x != nil {
		return x.DnsResolutionConfig
	}
	return nil
}

func (x *DnsFilterConfig_ClientContextConfig) GetTypedDnsResolverConfig() *v31.TypedExtensionConfig {
	if x != nil {
		return x.TypedDnsResolverConfig
	}
	return nil
}

func (x *DnsFilterConfig_ClientContextConfig) GetMaxPendingLookups() uint64 {
	if x != nil {
		return x.MaxPendingLookups
	}
	return 0
}

var File_envoy_extensions_filters_udp_dns_filter_v3_dns_filter_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_udp_dns_filter_v3_dns_filter_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x75, 0x64, 0x70, 0x2f, 0x64,
	0x6e, 0x73, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x33, 0x2f, 0x64, 0x6e, 0x73,
	0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2a, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x75, 0x64, 0x70, 0x2e, 0x64, 0x6e, 0x73, 0x5f,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x1a, 0x22, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x76, 0x33, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x76, 0x33, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x6c,
	0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x64, 0x6e, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x64, 0x6e, 0x73,
	0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc1, 0x07, 0x0a, 0x0f, 0x44, 0x6e,
	0x73, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x28, 0x0a,
	0x0b, 0x73, 0x74, 0x61, 0x74, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0a, 0x73, 0x74, 0x61,
	0x74, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x74, 0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4f,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x75, 0x64, 0x70, 0x2e, 0x64, 0x6e,
	0x73, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x6e, 0x73, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x0c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x74, 0x0a,
	0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x4f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e,
	0x75, 0x64, 0x70, 0x2e, 0x64, 0x6e, 0x73, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x76,
	0x33, 0x2e, 0x44, 0x6e, 0x73, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x1a, 0xc6, 0x01, 0x0a, 0x13, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x47, 0x0a, 0x10, 0x69,
	0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x64, 0x6e, 0x73, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x64, 0x6e, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x6e, 0x73, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x48, 0x00, 0x52, 0x0e, 0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x44, 0x6e, 0x73, 0x54,
	0x61, 0x62, 0x6c, 0x65, 0x12, 0x50, 0x0a, 0x12, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x5f, 0x64, 0x6e, 0x73, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x48, 0x00, 0x52, 0x10, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x44, 0x6e,
	0x73, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x14, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x1a, 0xce, 0x03, 0x0a,
	0x13, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x50, 0x0a, 0x10, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0xaa, 0x01,
	0x04, 0x32, 0x02, 0x08, 0x01, 0x52, 0x0f, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x54,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x59, 0x0a, 0x12, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x42, 0x0b, 0x18, 0x01, 0x92, 0xc7, 0x86, 0xd8, 0x04, 0x03, 0x33, 0x2e, 0x30, 0x52, 0x11,
	0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72,
	0x73, 0x12, 0x6a, 0x0a, 0x15, 0x64, 0x6e, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x29, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x6f, 0x6c,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x0b, 0x18, 0x01, 0x92,
	0xc7, 0x86, 0xd8, 0x04, 0x03, 0x33, 0x2e, 0x30, 0x52, 0x13, 0x64, 0x6e, 0x73, 0x52, 0x65, 0x73,
	0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x65, 0x0a,
	0x19, 0x74, 0x79, 0x70, 0x65, 0x64, 0x5f, 0x64, 0x6e, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x6c,
	0x76, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2a, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x64, 0x45, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x16, 0x74, 0x79,
	0x70, 0x65, 0x64, 0x44, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x37, 0x0a, 0x13, 0x6d, 0x61, 0x78, 0x5f, 0x70, 0x65, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x5f, 0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x28, 0x01, 0x52, 0x11, 0x6d, 0x61, 0x78, 0x50,
	0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x73, 0x42, 0xb4, 0x01,
	0x0a, 0x38, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x75, 0x64, 0x70, 0x2e, 0x64, 0x6e, 0x73,
	0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x42, 0x0e, 0x44, 0x6e, 0x73, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70,
	0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x75, 0x64,
	0x70, 0x2f, 0x64, 0x6e, 0x73, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x33, 0x3b,
	0x64, 0x6e, 0x73, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1,
	0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_filters_udp_dns_filter_v3_dns_filter_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_udp_dns_filter_v3_dns_filter_proto_rawDescData = file_envoy_extensions_filters_udp_dns_filter_v3_dns_filter_proto_rawDesc
)

func file_envoy_extensions_filters_udp_dns_filter_v3_dns_filter_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_udp_dns_filter_v3_dns_filter_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_udp_dns_filter_v3_dns_filter_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_udp_dns_filter_v3_dns_filter_proto_rawDescData)
	})
	return file_envoy_extensions_filters_udp_dns_filter_v3_dns_filter_proto_rawDescData
}

var file_envoy_extensions_filters_udp_dns_filter_v3_dns_filter_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_envoy_extensions_filters_udp_dns_filter_v3_dns_filter_proto_goTypes = []interface{}{
	(*DnsFilterConfig)(nil),                     // 0: envoy.extensions.filters.udp.dns_filter.v3.DnsFilterConfig
	(*DnsFilterConfig_ServerContextConfig)(nil), // 1: envoy.extensions.filters.udp.dns_filter.v3.DnsFilterConfig.ServerContextConfig
	(*DnsFilterConfig_ClientContextConfig)(nil), // 2: envoy.extensions.filters.udp.dns_filter.v3.DnsFilterConfig.ClientContextConfig
	(*v3.DnsTable)(nil),                         // 3: envoy.data.dns.v3.DnsTable
	(*v31.DataSource)(nil),                      // 4: envoy.config.core.v3.DataSource
	(*duration.Duration)(nil),                   // 5: google.protobuf.Duration
	(*v31.Address)(nil),                         // 6: envoy.config.core.v3.Address
	(*v31.DnsResolutionConfig)(nil),             // 7: envoy.config.core.v3.DnsResolutionConfig
	(*v31.TypedExtensionConfig)(nil),            // 8: envoy.config.core.v3.TypedExtensionConfig
}
var file_envoy_extensions_filters_udp_dns_filter_v3_dns_filter_proto_depIdxs = []int32{
	1, // 0: envoy.extensions.filters.udp.dns_filter.v3.DnsFilterConfig.server_config:type_name -> envoy.extensions.filters.udp.dns_filter.v3.DnsFilterConfig.ServerContextConfig
	2, // 1: envoy.extensions.filters.udp.dns_filter.v3.DnsFilterConfig.client_config:type_name -> envoy.extensions.filters.udp.dns_filter.v3.DnsFilterConfig.ClientContextConfig
	3, // 2: envoy.extensions.filters.udp.dns_filter.v3.DnsFilterConfig.ServerContextConfig.inline_dns_table:type_name -> envoy.data.dns.v3.DnsTable
	4, // 3: envoy.extensions.filters.udp.dns_filter.v3.DnsFilterConfig.ServerContextConfig.external_dns_table:type_name -> envoy.config.core.v3.DataSource
	5, // 4: envoy.extensions.filters.udp.dns_filter.v3.DnsFilterConfig.ClientContextConfig.resolver_timeout:type_name -> google.protobuf.Duration
	6, // 5: envoy.extensions.filters.udp.dns_filter.v3.DnsFilterConfig.ClientContextConfig.upstream_resolvers:type_name -> envoy.config.core.v3.Address
	7, // 6: envoy.extensions.filters.udp.dns_filter.v3.DnsFilterConfig.ClientContextConfig.dns_resolution_config:type_name -> envoy.config.core.v3.DnsResolutionConfig
	8, // 7: envoy.extensions.filters.udp.dns_filter.v3.DnsFilterConfig.ClientContextConfig.typed_dns_resolver_config:type_name -> envoy.config.core.v3.TypedExtensionConfig
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_envoy_extensions_filters_udp_dns_filter_v3_dns_filter_proto_init() }
func file_envoy_extensions_filters_udp_dns_filter_v3_dns_filter_proto_init() {
	if File_envoy_extensions_filters_udp_dns_filter_v3_dns_filter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_udp_dns_filter_v3_dns_filter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DnsFilterConfig); i {
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
		file_envoy_extensions_filters_udp_dns_filter_v3_dns_filter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DnsFilterConfig_ServerContextConfig); i {
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
		file_envoy_extensions_filters_udp_dns_filter_v3_dns_filter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DnsFilterConfig_ClientContextConfig); i {
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
	file_envoy_extensions_filters_udp_dns_filter_v3_dns_filter_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*DnsFilterConfig_ServerContextConfig_InlineDnsTable)(nil),
		(*DnsFilterConfig_ServerContextConfig_ExternalDnsTable)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_extensions_filters_udp_dns_filter_v3_dns_filter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_udp_dns_filter_v3_dns_filter_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_udp_dns_filter_v3_dns_filter_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_filters_udp_dns_filter_v3_dns_filter_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_udp_dns_filter_v3_dns_filter_proto = out.File
	file_envoy_extensions_filters_udp_dns_filter_v3_dns_filter_proto_rawDesc = nil
	file_envoy_extensions_filters_udp_dns_filter_v3_dns_filter_proto_goTypes = nil
	file_envoy_extensions_filters_udp_dns_filter_v3_dns_filter_proto_depIdxs = nil
}
