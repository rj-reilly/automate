// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/config/cs_nginx/config_request.proto

package cs_nginx // import "github.com/chef/automate/api/config/cs_nginx"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import shared "github.com/chef/automate/api/config/shared"
import _ "github.com/chef/automate/components/automate-grpc/protoc-gen-a2-config/api/a2conf"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ConfigRequest struct {
	V1                   *ConfigRequest_V1 `protobuf:"bytes,1,opt,name=v1,proto3" json:"v1,omitempty" toml:"v1,omitempty" mapstructure:"v1,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte            `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32             `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest) Reset()         { *m = ConfigRequest{} }
func (m *ConfigRequest) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest) ProtoMessage()    {}
func (*ConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_request_2c948ffbf42cfc52, []int{0}
}
func (m *ConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest.Unmarshal(m, b)
}
func (m *ConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest.Marshal(b, m, deterministic)
}
func (dst *ConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest.Merge(dst, src)
}
func (m *ConfigRequest) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest.Size(m)
}
func (m *ConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest proto.InternalMessageInfo

func (m *ConfigRequest) GetV1() *ConfigRequest_V1 {
	if m != nil {
		return m.V1
	}
	return nil
}

type ConfigRequest_V1 struct {
	Sys                  *ConfigRequest_V1_System  `protobuf:"bytes,1,opt,name=sys,proto3" json:"sys,omitempty" toml:"sys,omitempty" mapstructure:"sys,omitempty"`
	Svc                  *ConfigRequest_V1_Service `protobuf:"bytes,2,opt,name=svc,proto3" json:"svc,omitempty" toml:"svc,omitempty" mapstructure:"svc,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                    `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                     `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1) Reset()         { *m = ConfigRequest_V1{} }
func (m *ConfigRequest_V1) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1) ProtoMessage()    {}
func (*ConfigRequest_V1) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_request_2c948ffbf42cfc52, []int{0, 0}
}
func (m *ConfigRequest_V1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1.Unmarshal(m, b)
}
func (m *ConfigRequest_V1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1.Marshal(b, m, deterministic)
}
func (dst *ConfigRequest_V1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1.Merge(dst, src)
}
func (m *ConfigRequest_V1) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1.Size(m)
}
func (m *ConfigRequest_V1) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1 proto.InternalMessageInfo

func (m *ConfigRequest_V1) GetSys() *ConfigRequest_V1_System {
	if m != nil {
		return m.Sys
	}
	return nil
}

func (m *ConfigRequest_V1) GetSvc() *ConfigRequest_V1_Service {
	if m != nil {
		return m.Svc
	}
	return nil
}

type ConfigRequest_V1_System struct {
	Mlsa                 *shared.Mlsa                            `protobuf:"bytes,1,opt,name=mlsa,proto3" json:"mlsa,omitempty" toml:"mlsa,omitempty" mapstructure:"mlsa,omitempty"`
	Log                  *ConfigRequest_V1_System_Log            `protobuf:"bytes,2,opt,name=log,proto3" json:"log,omitempty" toml:"log,omitempty" mapstructure:"log,omitempty"`
	Service              *ConfigRequest_V1_System_Service        `protobuf:"bytes,3,opt,name=service,proto3" json:"service,omitempty" toml:"service,omitempty" mapstructure:"service,omitempty"`
	Ngx                  *ConfigRequest_V1_System_Nginx          `protobuf:"bytes,4,opt,name=ngx,proto3" json:"ngx,omitempty" toml:"ngx,omitempty" mapstructure:"ngx,omitempty"`
	Tls                  *shared.TLSCredentials                  `protobuf:"bytes,5,opt,name=tls,proto3" json:"tls,omitempty" toml:"tls,omitempty" mapstructure:"tls,omitempty"`
	RequiredRecipe       *ConfigRequest_V1_System_RequiredRecipe `protobuf:"bytes,6,opt,name=required_recipe,json=requiredRecipe,proto3" json:"required_recipe,omitempty" toml:"required_recipe,omitempty" mapstructure:"required_recipe,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                                  `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                                   `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System) Reset()         { *m = ConfigRequest_V1_System{} }
func (m *ConfigRequest_V1_System) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System) ProtoMessage()    {}
func (*ConfigRequest_V1_System) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_request_2c948ffbf42cfc52, []int{0, 0, 0}
}
func (m *ConfigRequest_V1_System) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_System.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_System) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_System.Marshal(b, m, deterministic)
}
func (dst *ConfigRequest_V1_System) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_System.Merge(dst, src)
}
func (m *ConfigRequest_V1_System) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_System.Size(m)
}
func (m *ConfigRequest_V1_System) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_System.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_System proto.InternalMessageInfo

func (m *ConfigRequest_V1_System) GetMlsa() *shared.Mlsa {
	if m != nil {
		return m.Mlsa
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetLog() *ConfigRequest_V1_System_Log {
	if m != nil {
		return m.Log
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetService() *ConfigRequest_V1_System_Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetNgx() *ConfigRequest_V1_System_Nginx {
	if m != nil {
		return m.Ngx
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetTls() *shared.TLSCredentials {
	if m != nil {
		return m.Tls
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetRequiredRecipe() *ConfigRequest_V1_System_RequiredRecipe {
	if m != nil {
		return m.RequiredRecipe
	}
	return nil
}

type ConfigRequest_V1_System_Service struct {
	Host                 *wrappers.StringValue `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty" toml:"host,omitempty" mapstructure:"host,omitempty"`
	Port                 *wrappers.Int32Value  `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty" toml:"port,omitempty" mapstructure:"port,omitempty"`
	StatusPort           *wrappers.Int32Value  `protobuf:"bytes,3,opt,name=status_port,json=statusPort,proto3" json:"status_port,omitempty" toml:"status_port,omitempty" mapstructure:"status_port,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                 `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System_Service) Reset()         { *m = ConfigRequest_V1_System_Service{} }
func (m *ConfigRequest_V1_System_Service) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System_Service) ProtoMessage()    {}
func (*ConfigRequest_V1_System_Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_request_2c948ffbf42cfc52, []int{0, 0, 0, 0}
}
func (m *ConfigRequest_V1_System_Service) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_System_Service.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_System_Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_System_Service.Marshal(b, m, deterministic)
}
func (dst *ConfigRequest_V1_System_Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_System_Service.Merge(dst, src)
}
func (m *ConfigRequest_V1_System_Service) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_System_Service.Size(m)
}
func (m *ConfigRequest_V1_System_Service) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_System_Service.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_System_Service proto.InternalMessageInfo

func (m *ConfigRequest_V1_System_Service) GetHost() *wrappers.StringValue {
	if m != nil {
		return m.Host
	}
	return nil
}

func (m *ConfigRequest_V1_System_Service) GetPort() *wrappers.Int32Value {
	if m != nil {
		return m.Port
	}
	return nil
}

func (m *ConfigRequest_V1_System_Service) GetStatusPort() *wrappers.Int32Value {
	if m != nil {
		return m.StatusPort
	}
	return nil
}

type ConfigRequest_V1_System_Log struct {
	Level                *wrappers.StringValue `protobuf:"bytes,1,opt,name=level,proto3" json:"level,omitempty" toml:"level,omitempty" mapstructure:"level,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                 `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System_Log) Reset()         { *m = ConfigRequest_V1_System_Log{} }
func (m *ConfigRequest_V1_System_Log) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System_Log) ProtoMessage()    {}
func (*ConfigRequest_V1_System_Log) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_request_2c948ffbf42cfc52, []int{0, 0, 0, 1}
}
func (m *ConfigRequest_V1_System_Log) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_System_Log.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_System_Log) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_System_Log.Marshal(b, m, deterministic)
}
func (dst *ConfigRequest_V1_System_Log) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_System_Log.Merge(dst, src)
}
func (m *ConfigRequest_V1_System_Log) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_System_Log.Size(m)
}
func (m *ConfigRequest_V1_System_Log) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_System_Log.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_System_Log proto.InternalMessageInfo

func (m *ConfigRequest_V1_System_Log) GetLevel() *wrappers.StringValue {
	if m != nil {
		return m.Level
	}
	return nil
}

type ConfigRequest_V1_System_Nginx struct {
	Main                 *ConfigRequest_V1_System_Nginx_Main   `protobuf:"bytes,1,opt,name=main,proto3" json:"main,omitempty" toml:"main,omitempty" mapstructure:"main,omitempty"`
	Events               *ConfigRequest_V1_System_Nginx_Events `protobuf:"bytes,2,opt,name=events,proto3" json:"events,omitempty" toml:"events,omitempty" mapstructure:"events,omitempty"`
	Http                 *ConfigRequest_V1_System_Nginx_Http   `protobuf:"bytes,3,opt,name=http,proto3" json:"http,omitempty" toml:"http,omitempty" mapstructure:"http,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                                `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                                 `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System_Nginx) Reset()         { *m = ConfigRequest_V1_System_Nginx{} }
func (m *ConfigRequest_V1_System_Nginx) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System_Nginx) ProtoMessage()    {}
func (*ConfigRequest_V1_System_Nginx) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_request_2c948ffbf42cfc52, []int{0, 0, 0, 2}
}
func (m *ConfigRequest_V1_System_Nginx) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_System_Nginx.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_System_Nginx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_System_Nginx.Marshal(b, m, deterministic)
}
func (dst *ConfigRequest_V1_System_Nginx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_System_Nginx.Merge(dst, src)
}
func (m *ConfigRequest_V1_System_Nginx) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_System_Nginx.Size(m)
}
func (m *ConfigRequest_V1_System_Nginx) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_System_Nginx.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_System_Nginx proto.InternalMessageInfo

func (m *ConfigRequest_V1_System_Nginx) GetMain() *ConfigRequest_V1_System_Nginx_Main {
	if m != nil {
		return m.Main
	}
	return nil
}

func (m *ConfigRequest_V1_System_Nginx) GetEvents() *ConfigRequest_V1_System_Nginx_Events {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *ConfigRequest_V1_System_Nginx) GetHttp() *ConfigRequest_V1_System_Nginx_Http {
	if m != nil {
		return m.Http
	}
	return nil
}

type ConfigRequest_V1_System_Nginx_Main struct {
	WorkerProcesses      *wrappers.Int32Value `protobuf:"bytes,1,opt,name=worker_processes,json=workerProcesses,proto3" json:"worker_processes,omitempty" toml:"worker_processes,omitempty" mapstructure:"worker_processes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte               `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System_Nginx_Main) Reset()         { *m = ConfigRequest_V1_System_Nginx_Main{} }
func (m *ConfigRequest_V1_System_Nginx_Main) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System_Nginx_Main) ProtoMessage()    {}
func (*ConfigRequest_V1_System_Nginx_Main) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_request_2c948ffbf42cfc52, []int{0, 0, 0, 2, 0}
}
func (m *ConfigRequest_V1_System_Nginx_Main) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_System_Nginx_Main.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_System_Nginx_Main) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_System_Nginx_Main.Marshal(b, m, deterministic)
}
func (dst *ConfigRequest_V1_System_Nginx_Main) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_System_Nginx_Main.Merge(dst, src)
}
func (m *ConfigRequest_V1_System_Nginx_Main) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_System_Nginx_Main.Size(m)
}
func (m *ConfigRequest_V1_System_Nginx_Main) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_System_Nginx_Main.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_System_Nginx_Main proto.InternalMessageInfo

func (m *ConfigRequest_V1_System_Nginx_Main) GetWorkerProcesses() *wrappers.Int32Value {
	if m != nil {
		return m.WorkerProcesses
	}
	return nil
}

type ConfigRequest_V1_System_Nginx_Events struct {
	WorkerConnections    *wrappers.Int32Value `protobuf:"bytes,1,opt,name=worker_connections,json=workerConnections,proto3" json:"worker_connections,omitempty" toml:"worker_connections,omitempty" mapstructure:"worker_connections,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte               `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System_Nginx_Events) Reset()         { *m = ConfigRequest_V1_System_Nginx_Events{} }
func (m *ConfigRequest_V1_System_Nginx_Events) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System_Nginx_Events) ProtoMessage()    {}
func (*ConfigRequest_V1_System_Nginx_Events) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_request_2c948ffbf42cfc52, []int{0, 0, 0, 2, 1}
}
func (m *ConfigRequest_V1_System_Nginx_Events) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_System_Nginx_Events.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_System_Nginx_Events) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_System_Nginx_Events.Marshal(b, m, deterministic)
}
func (dst *ConfigRequest_V1_System_Nginx_Events) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_System_Nginx_Events.Merge(dst, src)
}
func (m *ConfigRequest_V1_System_Nginx_Events) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_System_Nginx_Events.Size(m)
}
func (m *ConfigRequest_V1_System_Nginx_Events) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_System_Nginx_Events.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_System_Nginx_Events proto.InternalMessageInfo

func (m *ConfigRequest_V1_System_Nginx_Events) GetWorkerConnections() *wrappers.Int32Value {
	if m != nil {
		return m.WorkerConnections
	}
	return nil
}

type ConfigRequest_V1_System_Nginx_Http struct {
	ClientMaxBodySize   *wrappers.StringValue `protobuf:"bytes,1,opt,name=client_max_body_size,json=clientMaxBodySize,proto3" json:"client_max_body_size,omitempty" toml:"client_max_body_size,omitempty" mapstructure:"client_max_body_size,omitempty"`
	ProxyConnectTimeout *wrappers.Int32Value  `protobuf:"bytes,2,opt,name=proxy_connect_timeout,json=proxyConnectTimeout,proto3" json:"proxy_connect_timeout,omitempty" toml:"proxy_connect_timeout,omitempty" mapstructure:"proxy_connect_timeout,omitempty"`
	KeepaliveTimeout    *wrappers.Int32Value  `protobuf:"bytes,3,opt,name=keepalive_timeout,json=keepaliveTimeout,proto3" json:"keepalive_timeout,omitempty" toml:"keepalive_timeout,omitempty" mapstructure:"keepalive_timeout,omitempty"`
	Gzip                *wrappers.StringValue `protobuf:"bytes,4,opt,name=gzip,proto3" json:"gzip,omitempty" toml:"gzip,omitempty" mapstructure:"gzip,omitempty"`
	GzipHttpVersion     *wrappers.StringValue `protobuf:"bytes,5,opt,name=gzip_http_version,json=gzipHttpVersion,proto3" json:"gzip_http_version,omitempty" toml:"gzip_http_version,omitempty" mapstructure:"gzip_http_version,omitempty"`
	// StringValue for consitency with other nginx configs
	GzipCompLevel             *wrappers.StringValue `protobuf:"bytes,6,opt,name=gzip_comp_level,json=gzipCompLevel,proto3" json:"gzip_comp_level,omitempty" toml:"gzip_comp_level,omitempty" mapstructure:"gzip_comp_level,omitempty"`
	GzipProxied               *wrappers.StringValue `protobuf:"bytes,7,opt,name=gzip_proxied,json=gzipProxied,proto3" json:"gzip_proxied,omitempty" toml:"gzip_proxied,omitempty" mapstructure:"gzip_proxied,omitempty"`
	GzipTypes                 *wrappers.StringValue `protobuf:"bytes,8,opt,name=gzip_types,json=gzipTypes,proto3" json:"gzip_types,omitempty" toml:"gzip_types,omitempty" mapstructure:"gzip_types,omitempty"`
	Sendfile                  *wrappers.StringValue `protobuf:"bytes,9,opt,name=sendfile,proto3" json:"sendfile,omitempty" toml:"sendfile,omitempty" mapstructure:"sendfile,omitempty"`
	TcpNodelay                *wrappers.StringValue `protobuf:"bytes,10,opt,name=tcp_nodelay,json=tcpNodelay,proto3" json:"tcp_nodelay,omitempty" toml:"tcp_nodelay,omitempty" mapstructure:"tcp_nodelay,omitempty"`
	TcpNopush                 *wrappers.StringValue `protobuf:"bytes,11,opt,name=tcp_nopush,json=tcpNopush,proto3" json:"tcp_nopush,omitempty" toml:"tcp_nopush,omitempty" mapstructure:"tcp_nopush,omitempty"`
	SslCiphers                *wrappers.StringValue `protobuf:"bytes,12,opt,name=ssl_ciphers,json=sslCiphers,proto3" json:"ssl_ciphers,omitempty" toml:"ssl_ciphers,omitempty" mapstructure:"ssl_ciphers,omitempty"`
	SslProtocols              *wrappers.StringValue `protobuf:"bytes,13,opt,name=ssl_protocols,json=sslProtocols,proto3" json:"ssl_protocols,omitempty" toml:"ssl_protocols,omitempty" mapstructure:"ssl_protocols,omitempty"`
	SslVerifyDepth            *wrappers.Int32Value  `protobuf:"bytes,15,opt,name=ssl_verify_depth,json=sslVerifyDepth,proto3" json:"ssl_verify_depth,omitempty" toml:"ssl_verify_depth,omitempty" mapstructure:"ssl_verify_depth,omitempty"`
	ServerNamesHashBucketSize *wrappers.Int32Value  `protobuf:"bytes,16,opt,name=server_names_hash_bucket_size,json=serverNamesHashBucketSize,proto3" json:"server_names_hash_bucket_size,omitempty" toml:"server_names_hash_bucket_size,omitempty" mapstructure:"server_names_hash_bucket_size,omitempty"`
	ClientBodyBufferSize      *wrappers.StringValue `protobuf:"bytes,17,opt,name=client_body_buffer_size,json=clientBodyBufferSize,proto3" json:"client_body_buffer_size,omitempty" toml:"client_body_buffer_size,omitempty" mapstructure:"client_body_buffer_size,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}              `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized          []byte                `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache             int32                 `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System_Nginx_Http) Reset()         { *m = ConfigRequest_V1_System_Nginx_Http{} }
func (m *ConfigRequest_V1_System_Nginx_Http) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System_Nginx_Http) ProtoMessage()    {}
func (*ConfigRequest_V1_System_Nginx_Http) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_request_2c948ffbf42cfc52, []int{0, 0, 0, 2, 2}
}
func (m *ConfigRequest_V1_System_Nginx_Http) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_System_Nginx_Http.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_System_Nginx_Http) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_System_Nginx_Http.Marshal(b, m, deterministic)
}
func (dst *ConfigRequest_V1_System_Nginx_Http) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_System_Nginx_Http.Merge(dst, src)
}
func (m *ConfigRequest_V1_System_Nginx_Http) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_System_Nginx_Http.Size(m)
}
func (m *ConfigRequest_V1_System_Nginx_Http) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_System_Nginx_Http.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_System_Nginx_Http proto.InternalMessageInfo

func (m *ConfigRequest_V1_System_Nginx_Http) GetClientMaxBodySize() *wrappers.StringValue {
	if m != nil {
		return m.ClientMaxBodySize
	}
	return nil
}

func (m *ConfigRequest_V1_System_Nginx_Http) GetProxyConnectTimeout() *wrappers.Int32Value {
	if m != nil {
		return m.ProxyConnectTimeout
	}
	return nil
}

func (m *ConfigRequest_V1_System_Nginx_Http) GetKeepaliveTimeout() *wrappers.Int32Value {
	if m != nil {
		return m.KeepaliveTimeout
	}
	return nil
}

func (m *ConfigRequest_V1_System_Nginx_Http) GetGzip() *wrappers.StringValue {
	if m != nil {
		return m.Gzip
	}
	return nil
}

func (m *ConfigRequest_V1_System_Nginx_Http) GetGzipHttpVersion() *wrappers.StringValue {
	if m != nil {
		return m.GzipHttpVersion
	}
	return nil
}

func (m *ConfigRequest_V1_System_Nginx_Http) GetGzipCompLevel() *wrappers.StringValue {
	if m != nil {
		return m.GzipCompLevel
	}
	return nil
}

func (m *ConfigRequest_V1_System_Nginx_Http) GetGzipProxied() *wrappers.StringValue {
	if m != nil {
		return m.GzipProxied
	}
	return nil
}

func (m *ConfigRequest_V1_System_Nginx_Http) GetGzipTypes() *wrappers.StringValue {
	if m != nil {
		return m.GzipTypes
	}
	return nil
}

func (m *ConfigRequest_V1_System_Nginx_Http) GetSendfile() *wrappers.StringValue {
	if m != nil {
		return m.Sendfile
	}
	return nil
}

func (m *ConfigRequest_V1_System_Nginx_Http) GetTcpNodelay() *wrappers.StringValue {
	if m != nil {
		return m.TcpNodelay
	}
	return nil
}

func (m *ConfigRequest_V1_System_Nginx_Http) GetTcpNopush() *wrappers.StringValue {
	if m != nil {
		return m.TcpNopush
	}
	return nil
}

func (m *ConfigRequest_V1_System_Nginx_Http) GetSslCiphers() *wrappers.StringValue {
	if m != nil {
		return m.SslCiphers
	}
	return nil
}

func (m *ConfigRequest_V1_System_Nginx_Http) GetSslProtocols() *wrappers.StringValue {
	if m != nil {
		return m.SslProtocols
	}
	return nil
}

func (m *ConfigRequest_V1_System_Nginx_Http) GetSslVerifyDepth() *wrappers.Int32Value {
	if m != nil {
		return m.SslVerifyDepth
	}
	return nil
}

func (m *ConfigRequest_V1_System_Nginx_Http) GetServerNamesHashBucketSize() *wrappers.Int32Value {
	if m != nil {
		return m.ServerNamesHashBucketSize
	}
	return nil
}

func (m *ConfigRequest_V1_System_Nginx_Http) GetClientBodyBufferSize() *wrappers.StringValue {
	if m != nil {
		return m.ClientBodyBufferSize
	}
	return nil
}

type ConfigRequest_V1_System_RequiredRecipe struct {
	Enabled *wrappers.BoolValue   `protobuf:"bytes,1,opt,name=enabled,proto3" json:"enabled,omitempty" toml:"enabled,omitempty" mapstructure:"enabled,omitempty"`
	Path    *wrappers.StringValue `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty" toml:"path,omitempty" mapstructure:"path,omitempty"`
	// Auto-calculated
	ContentMd5           *wrappers.StringValue `protobuf:"bytes,3,opt,name=content_md5,json=contentMd5,proto3" json:"content_md5,omitempty" toml:"content_md5,omitempty" mapstructure:"content_md5,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                 `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System_RequiredRecipe) Reset() {
	*m = ConfigRequest_V1_System_RequiredRecipe{}
}
func (m *ConfigRequest_V1_System_RequiredRecipe) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System_RequiredRecipe) ProtoMessage()    {}
func (*ConfigRequest_V1_System_RequiredRecipe) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_request_2c948ffbf42cfc52, []int{0, 0, 0, 3}
}
func (m *ConfigRequest_V1_System_RequiredRecipe) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_System_RequiredRecipe.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_System_RequiredRecipe) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_System_RequiredRecipe.Marshal(b, m, deterministic)
}
func (dst *ConfigRequest_V1_System_RequiredRecipe) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_System_RequiredRecipe.Merge(dst, src)
}
func (m *ConfigRequest_V1_System_RequiredRecipe) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_System_RequiredRecipe.Size(m)
}
func (m *ConfigRequest_V1_System_RequiredRecipe) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_System_RequiredRecipe.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_System_RequiredRecipe proto.InternalMessageInfo

func (m *ConfigRequest_V1_System_RequiredRecipe) GetEnabled() *wrappers.BoolValue {
	if m != nil {
		return m.Enabled
	}
	return nil
}

func (m *ConfigRequest_V1_System_RequiredRecipe) GetPath() *wrappers.StringValue {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *ConfigRequest_V1_System_RequiredRecipe) GetContentMd5() *wrappers.StringValue {
	if m != nil {
		return m.ContentMd5
	}
	return nil
}

type ConfigRequest_V1_Service struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_Service) Reset()         { *m = ConfigRequest_V1_Service{} }
func (m *ConfigRequest_V1_Service) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_Service) ProtoMessage()    {}
func (*ConfigRequest_V1_Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_request_2c948ffbf42cfc52, []int{0, 0, 1}
}
func (m *ConfigRequest_V1_Service) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_Service.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_Service.Marshal(b, m, deterministic)
}
func (dst *ConfigRequest_V1_Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_Service.Merge(dst, src)
}
func (m *ConfigRequest_V1_Service) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_Service.Size(m)
}
func (m *ConfigRequest_V1_Service) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_Service.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_Service proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ConfigRequest)(nil), "chef.automate.domain.cs_nginx.ConfigRequest")
	proto.RegisterType((*ConfigRequest_V1)(nil), "chef.automate.domain.cs_nginx.ConfigRequest.V1")
	proto.RegisterType((*ConfigRequest_V1_System)(nil), "chef.automate.domain.cs_nginx.ConfigRequest.V1.System")
	proto.RegisterType((*ConfigRequest_V1_System_Service)(nil), "chef.automate.domain.cs_nginx.ConfigRequest.V1.System.Service")
	proto.RegisterType((*ConfigRequest_V1_System_Log)(nil), "chef.automate.domain.cs_nginx.ConfigRequest.V1.System.Log")
	proto.RegisterType((*ConfigRequest_V1_System_Nginx)(nil), "chef.automate.domain.cs_nginx.ConfigRequest.V1.System.Nginx")
	proto.RegisterType((*ConfigRequest_V1_System_Nginx_Main)(nil), "chef.automate.domain.cs_nginx.ConfigRequest.V1.System.Nginx.Main")
	proto.RegisterType((*ConfigRequest_V1_System_Nginx_Events)(nil), "chef.automate.domain.cs_nginx.ConfigRequest.V1.System.Nginx.Events")
	proto.RegisterType((*ConfigRequest_V1_System_Nginx_Http)(nil), "chef.automate.domain.cs_nginx.ConfigRequest.V1.System.Nginx.Http")
	proto.RegisterType((*ConfigRequest_V1_System_RequiredRecipe)(nil), "chef.automate.domain.cs_nginx.ConfigRequest.V1.System.RequiredRecipe")
	proto.RegisterType((*ConfigRequest_V1_Service)(nil), "chef.automate.domain.cs_nginx.ConfigRequest.V1.Service")
}

func init() {
	proto.RegisterFile("api/config/cs_nginx/config_request.proto", fileDescriptor_config_request_2c948ffbf42cfc52)
}

var fileDescriptor_config_request_2c948ffbf42cfc52 = []byte{
	// 1118 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x96, 0xdd, 0x6e, 0xdb, 0x36,
	0x14, 0x80, 0x91, 0xd8, 0x71, 0x52, 0xe6, 0xcf, 0x66, 0x17, 0x54, 0x53, 0xd7, 0xa2, 0xd8, 0x55,
	0x31, 0x2c, 0x72, 0x93, 0xb6, 0xfb, 0xe9, 0xba, 0x15, 0xb5, 0xdb, 0x21, 0x2d, 0x92, 0x34, 0x90,
	0xb3, 0x60, 0xd8, 0x30, 0x08, 0xb4, 0x74, 0x2c, 0x11, 0x91, 0x49, 0x8d, 0xa4, 0xdd, 0x38, 0xd7,
	0x7b, 0x94, 0xbd, 0xc3, 0xee, 0xfb, 0x14, 0xc3, 0xb0, 0x3d, 0xc0, 0xee, 0x06, 0xf4, 0x05, 0x06,
	0xfe, 0xc8, 0x5b, 0x96, 0x21, 0x56, 0x93, 0x2b, 0x43, 0xd6, 0xf9, 0x3e, 0x1e, 0x1e, 0x92, 0x87,
	0x42, 0x77, 0x49, 0x41, 0xdb, 0x31, 0x67, 0x03, 0x9a, 0xb6, 0x63, 0x19, 0xb1, 0x94, 0xb2, 0x13,
	0xf7, 0x1c, 0x09, 0xf8, 0x71, 0x04, 0x52, 0x05, 0x85, 0xe0, 0x8a, 0xe3, 0x5b, 0x71, 0x06, 0x83,
	0x80, 0x8c, 0x14, 0x1f, 0x12, 0x05, 0x41, 0xc2, 0x87, 0x84, 0xb2, 0xa0, 0x64, 0xfc, 0xdb, 0xff,
	0x12, 0xc9, 0x8c, 0x08, 0x48, 0xda, 0x69, 0xce, 0xfb, 0x24, 0xb7, 0xb8, 0x7f, 0xf3, 0xfc, 0x7b,
	0x95, 0x4b, 0xf7, 0xf2, 0x65, 0xcc, 0x87, 0x05, 0x67, 0xc0, 0x94, 0x6c, 0x97, 0x23, 0x6c, 0xa6,
	0xa2, 0x88, 0xdb, 0xe6, 0x7d, 0xbc, 0x99, 0x02, 0xdb, 0x24, 0xdb, 0x9b, 0x8e, 0xd7, 0x2a, 0xb2,
	0xad, 0x1f, 0xda, 0x84, 0x31, 0xae, 0x88, 0xa2, 0x9c, 0x95, 0xae, 0xdb, 0x29, 0xe7, 0x69, 0x0e,
	0x96, 0xec, 0x8f, 0x06, 0xed, 0xd7, 0x82, 0x14, 0x05, 0x08, 0xf7, 0xfe, 0xc3, 0xbf, 0x36, 0xd0,
	0x6a, 0xd7, 0x78, 0x42, 0x3b, 0x3f, 0xfc, 0x04, 0xcd, 0x8f, 0xb7, 0xbc, 0xb9, 0x3b, 0x73, 0x77,
	0x97, 0xb7, 0xdb, 0xc1, 0x85, 0xd3, 0x0c, 0xce, 0x90, 0xc1, 0xd1, 0x56, 0x38, 0x3f, 0xde, 0xf2,
	0x7f, 0xda, 0x40, 0xf3, 0x47, 0x5b, 0x78, 0x07, 0xd5, 0xe4, 0x44, 0x3a, 0xd1, 0x27, 0xef, 0x28,
	0x0a, 0x7a, 0x13, 0xa9, 0x60, 0x18, 0x6a, 0x05, 0x7e, 0x81, 0x6a, 0x72, 0x1c, 0x7b, 0xf3, 0xc6,
	0xf4, 0xe9, 0x3b, 0x9b, 0x40, 0x8c, 0x69, 0x0c, 0xa1, 0x76, 0xf8, 0x7f, 0x5c, 0x47, 0x0d, 0xab,
	0xc6, 0x0f, 0x50, 0x7d, 0x98, 0x4b, 0xe2, 0x12, 0xbc, 0xf3, 0x1f, 0x2d, 0x65, 0x03, 0x41, 0x02,
	0x5b, 0xe1, 0x60, 0x2f, 0x97, 0x24, 0x34, 0xd1, 0x78, 0x17, 0xd5, 0x72, 0x9e, 0xba, 0x5c, 0x1e,
	0x5d, 0x6e, 0x56, 0xc1, 0x2e, 0x4f, 0x43, 0xad, 0xc1, 0xdf, 0xa2, 0x45, 0x69, 0xd3, 0xf3, 0x6a,
	0xc6, 0xf8, 0xd5, 0x25, 0x8d, 0xe5, 0x24, 0x4b, 0x1d, 0xde, 0x47, 0x35, 0x96, 0x9e, 0x78, 0x75,
	0x63, 0x7d, 0x7c, 0x49, 0xeb, 0xbe, 0x7e, 0x1b, 0x6a, 0x11, 0x7e, 0x8c, 0x6a, 0x2a, 0x97, 0xde,
	0x82, 0xf1, 0x7d, 0x74, 0x51, 0xb1, 0x0e, 0x77, 0x7b, 0x5d, 0x01, 0x09, 0x30, 0x45, 0x49, 0x2e,
	0x43, 0x8d, 0x61, 0x86, 0xd6, 0xf5, 0xf1, 0xa1, 0x02, 0x92, 0x48, 0x40, 0x4c, 0x0b, 0xf0, 0x1a,
	0xc6, 0xf4, 0xfc, 0x92, 0x99, 0x85, 0xce, 0x16, 0x1a, 0x59, 0xb8, 0x26, 0xce, 0x3c, 0xfb, 0xbf,
	0xcd, 0xa1, 0x45, 0x57, 0x12, 0x7c, 0x0f, 0xd5, 0x33, 0x2e, 0x95, 0x5b, 0xe7, 0x0f, 0x02, 0x7b,
	0x20, 0x82, 0xf2, 0x40, 0x04, 0x3d, 0x25, 0x28, 0x4b, 0x8f, 0x48, 0x3e, 0x82, 0xd0, 0x44, 0xe2,
	0x1d, 0x54, 0x2f, 0xb8, 0x50, 0x6e, 0x91, 0x6f, 0x9e, 0x23, 0x5e, 0x30, 0x75, 0x7f, 0xdb, 0x00,
	0x9d, 0x1b, 0x6f, 0xde, 0x7a, 0xd7, 0xa7, 0x8b, 0xd8, 0xfc, 0xf5, 0x95, 0xbf, 0x90, 0x29, 0x55,
	0xc8, 0xd0, 0x18, 0x70, 0x0f, 0x2d, 0x4b, 0x45, 0xd4, 0x48, 0x46, 0x46, 0x58, 0x9b, 0x2d, 0xdc,
	0x78, 0xf3, 0xd6, 0x6b, 0xa1, 0x86, 0xa5, 0x9a, 0xbf, 0xbf, 0xf2, 0xeb, 0xda, 0x17, 0x22, 0xfb,
	0xc7, 0x01, 0x17, 0xca, 0xff, 0x1c, 0xd5, 0x76, 0x79, 0x8a, 0xb7, 0xd1, 0x42, 0x0e, 0x63, 0xc8,
	0x2b, 0x4d, 0xcc, 0x86, 0xfa, 0x3f, 0xaf, 0xa0, 0x05, 0xb3, 0xa8, 0xf8, 0x1b, 0x54, 0xd7, 0x95,
	0x76, 0xf0, 0xd3, 0xab, 0x6c, 0x90, 0x60, 0x8f, 0x50, 0x16, 0x1a, 0x1d, 0xfe, 0x1e, 0x35, 0x60,
	0xac, 0x1b, 0x97, 0x2b, 0x5e, 0xf7, 0x4a, 0xe2, 0xe7, 0x46, 0x15, 0x3a, 0xa5, 0xce, 0x59, 0x17,
	0xc3, 0x95, 0xf1, 0x6a, 0x39, 0xef, 0xe8, 0xaa, 0x1a, 0x9d, 0xbf, 0x8f, 0xea, 0x7a, 0x06, 0xf8,
	0x6b, 0xd4, 0x7c, 0xcd, 0xc5, 0x31, 0x88, 0xa8, 0x10, 0x3c, 0x06, 0x29, 0xa1, 0xec, 0x5e, 0x17,
	0xad, 0x58, 0xb8, 0x6e, 0xa1, 0x83, 0x92, 0xf1, 0x0f, 0x51, 0xc3, 0x26, 0x8e, 0x5f, 0x22, 0xec,
	0x8c, 0x31, 0x67, 0x0c, 0x62, 0xd3, 0x98, 0xab, 0x38, 0x5b, 0x16, 0xeb, 0xfe, 0x43, 0xf9, 0x7f,
	0x2e, 0xa1, 0xba, 0x4e, 0x1a, 0xef, 0xa1, 0xf7, 0xe2, 0x9c, 0x02, 0x53, 0xd1, 0x90, 0x9c, 0x44,
	0x7d, 0x9e, 0x4c, 0x22, 0x49, 0x4f, 0xa1, 0xd2, 0x36, 0x68, 0x59, 0x72, 0x8f, 0x9c, 0x74, 0x78,
	0x32, 0xe9, 0xd1, 0x53, 0xc0, 0xaf, 0xd0, 0x46, 0x21, 0xf8, 0xc9, 0xa4, 0x4c, 0x31, 0x52, 0x74,
	0x08, 0x7c, 0x54, 0x65, 0xf7, 0x87, 0xd7, 0x0d, 0xe9, 0xb2, 0x3c, 0xb4, 0x1c, 0xde, 0x41, 0xad,
	0x63, 0x80, 0x82, 0xe4, 0x74, 0x0c, 0x53, 0xd9, 0xec, 0x9d, 0x1f, 0x36, 0xa7, 0x54, 0x69, 0xba,
	0x87, 0xea, 0xe9, 0x29, 0x2d, 0x5c, 0x13, 0x9b, 0x71, 0x72, 0x75, 0xa4, 0x1e, 0x5b, 0xff, 0x46,
	0x7a, 0x5d, 0xa3, 0x31, 0x08, 0x49, 0x39, 0x73, 0x3d, 0xeb, 0x62, 0x7c, 0x5d, 0x63, 0xba, 0xbc,
	0x47, 0x16, 0xc2, 0xcf, 0x90, 0xf9, 0x2b, 0xd2, 0x57, 0x71, 0x64, 0xcf, 0x59, 0xa3, 0x82, 0x67,
	0x55, 0x43, 0x5d, 0x3e, 0x2c, 0x76, 0x35, 0x82, 0x9f, 0xa0, 0x15, 0x63, 0xd1, 0x75, 0xa2, 0x90,
	0x78, 0x8b, 0x15, 0x14, 0xcb, 0x9a, 0x38, 0xb0, 0x00, 0xfe, 0x02, 0x21, 0x23, 0x50, 0x93, 0x02,
	0xa4, 0xb7, 0x54, 0x01, 0xbf, 0xa6, 0xe3, 0x0f, 0x75, 0x38, 0xfe, 0x0c, 0x2d, 0x49, 0x60, 0xc9,
	0x80, 0xe6, 0xe0, 0x5d, 0xab, 0x80, 0x4e, 0xa3, 0xf1, 0x97, 0x68, 0x59, 0xc5, 0x45, 0xc4, 0x78,
	0x02, 0x39, 0x99, 0x78, 0xa8, 0x02, 0x8c, 0x54, 0x5c, 0xec, 0xdb, 0x78, 0x9d, 0xb5, 0xc5, 0x8b,
	0x91, 0xcc, 0xbc, 0xe5, 0x2a, 0x59, 0x1b, 0x5a, 0x87, 0xeb, 0xb1, 0xa5, 0xcc, 0xa3, 0x98, 0x16,
	0x19, 0x08, 0xe9, 0xad, 0x54, 0x19, 0x5b, 0xca, 0xbc, 0x6b, 0xe3, 0xf1, 0x53, 0xb4, 0xaa, 0x71,
	0xfb, 0xa5, 0xc4, 0x73, 0xe9, 0xad, 0x56, 0x10, 0xac, 0x48, 0x99, 0x1f, 0x94, 0x04, 0x7e, 0x8e,
	0x9a, 0x5a, 0x31, 0x06, 0x41, 0x07, 0x93, 0x28, 0x81, 0x42, 0x65, 0xde, 0xfa, 0xec, 0x0d, 0xbc,
	0x26, 0x65, 0x7e, 0x64, 0x98, 0x67, 0x1a, 0xc1, 0x3f, 0xa0, 0x5b, 0xfa, 0x5e, 0x00, 0x11, 0x31,
	0x32, 0x04, 0x19, 0x65, 0x44, 0x66, 0x51, 0x7f, 0x14, 0x1f, 0x83, 0xb2, 0x27, 0xb6, 0x39, 0xdb,
	0xf9, 0xbe, 0x35, 0xec, 0x6b, 0xc1, 0x0e, 0x91, 0x59, 0xc7, 0xe0, 0xe6, 0xe0, 0xf6, 0xd0, 0x0d,
	0xd7, 0x07, 0x4c, 0x0f, 0xe8, 0x8f, 0x06, 0x03, 0x10, 0x56, 0xdc, 0xaa, 0x30, 0x65, 0xd7, 0x44,
	0x74, 0x1f, 0xe8, 0x18, 0x54, 0x4b, 0x5f, 0xd6, 0x97, 0xd6, 0x9a, 0xeb, 0xfe, 0x2f, 0x73, 0x68,
	0xed, 0xec, 0x0d, 0x8b, 0x1f, 0xa0, 0x45, 0x60, 0xa4, 0x9f, 0x43, 0xe2, 0x1a, 0x8d, 0x7f, 0xce,
	0xde, 0xe1, 0x3c, 0xb7, 0xee, 0x32, 0x54, 0x9f, 0xe0, 0x82, 0xa8, 0xcc, 0xf5, 0x92, 0x19, 0x27,
	0x58, 0x47, 0xea, 0xd5, 0x8f, 0x39, 0x53, 0xa6, 0xbd, 0x25, 0x0f, 0x5d, 0xdf, 0x98, 0xb1, 0xfa,
	0x0e, 0xd8, 0x4b, 0x1e, 0xfa, 0xd7, 0xa6, 0xf7, 0xfe, 0x23, 0x7b, 0x35, 0xb7, 0xa6, 0x1f, 0xd0,
	0xb1, 0xdc, 0x34, 0x17, 0x43, 0x27, 0xf8, 0xee, 0xe3, 0x94, 0xaa, 0x6c, 0xd4, 0x0f, 0x62, 0x3e,
	0x6c, 0xeb, 0x4b, 0x64, 0xfa, 0x95, 0xdd, 0xfe, 0x9f, 0xef, 0xff, 0x7e, 0xc3, 0x8c, 0x7a, 0xff,
	0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x25, 0x86, 0xbf, 0xb7, 0x1d, 0x0c, 0x00, 0x00,
}
