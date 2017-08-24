// Code generated by protoc-gen-go. DO NOT EDIT.
// source: node_attestor.proto

/*
Package cpnodeattestor is a generated protocol buffer package.

It is generated from these files:
	node_attestor.proto

It has these top-level messages:
	AttestRequest
	AttestResponse
*/
package cpnodeattestor

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import sriplugin "github.com/spiffe/sri/common/plugin"
import common "github.com/spiffe/sri/common"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// ConfigureRequest from public import github.com/spiffe/sri/common/plugin/plugin.proto
type ConfigureRequest sriplugin.ConfigureRequest

func (m *ConfigureRequest) Reset()         { (*sriplugin.ConfigureRequest)(m).Reset() }
func (m *ConfigureRequest) String() string { return (*sriplugin.ConfigureRequest)(m).String() }
func (*ConfigureRequest) ProtoMessage()    {}
func (m *ConfigureRequest) GetConfiguration() string {
	return (*sriplugin.ConfigureRequest)(m).GetConfiguration()
}

// ConfigureResponse from public import github.com/spiffe/sri/common/plugin/plugin.proto
type ConfigureResponse sriplugin.ConfigureResponse

func (m *ConfigureResponse) Reset()         { (*sriplugin.ConfigureResponse)(m).Reset() }
func (m *ConfigureResponse) String() string { return (*sriplugin.ConfigureResponse)(m).String() }
func (*ConfigureResponse) ProtoMessage()    {}
func (m *ConfigureResponse) GetErrorList() []string {
	return (*sriplugin.ConfigureResponse)(m).GetErrorList()
}

// GetPluginInfoRequest from public import github.com/spiffe/sri/common/plugin/plugin.proto
type GetPluginInfoRequest sriplugin.GetPluginInfoRequest

func (m *GetPluginInfoRequest) Reset()         { (*sriplugin.GetPluginInfoRequest)(m).Reset() }
func (m *GetPluginInfoRequest) String() string { return (*sriplugin.GetPluginInfoRequest)(m).String() }
func (*GetPluginInfoRequest) ProtoMessage()    {}

// GetPluginInfoResponse from public import github.com/spiffe/sri/common/plugin/plugin.proto
type GetPluginInfoResponse sriplugin.GetPluginInfoResponse

func (m *GetPluginInfoResponse) Reset()         { (*sriplugin.GetPluginInfoResponse)(m).Reset() }
func (m *GetPluginInfoResponse) String() string { return (*sriplugin.GetPluginInfoResponse)(m).String() }
func (*GetPluginInfoResponse) ProtoMessage()    {}
func (m *GetPluginInfoResponse) GetName() string {
	return (*sriplugin.GetPluginInfoResponse)(m).GetName()
}
func (m *GetPluginInfoResponse) GetCategory() string {
	return (*sriplugin.GetPluginInfoResponse)(m).GetCategory()
}
func (m *GetPluginInfoResponse) GetType() string {
	return (*sriplugin.GetPluginInfoResponse)(m).GetType()
}
func (m *GetPluginInfoResponse) GetDescription() string {
	return (*sriplugin.GetPluginInfoResponse)(m).GetDescription()
}
func (m *GetPluginInfoResponse) GetDateCreated() string {
	return (*sriplugin.GetPluginInfoResponse)(m).GetDateCreated()
}
func (m *GetPluginInfoResponse) GetLocation() string {
	return (*sriplugin.GetPluginInfoResponse)(m).GetLocation()
}
func (m *GetPluginInfoResponse) GetVersion() string {
	return (*sriplugin.GetPluginInfoResponse)(m).GetVersion()
}
func (m *GetPluginInfoResponse) GetAuthor() string {
	return (*sriplugin.GetPluginInfoResponse)(m).GetAuthor()
}
func (m *GetPluginInfoResponse) GetCompany() string {
	return (*sriplugin.GetPluginInfoResponse)(m).GetCompany()
}

// PluginInfoRequest from public import github.com/spiffe/sri/common/plugin/plugin.proto
type PluginInfoRequest sriplugin.PluginInfoRequest

func (m *PluginInfoRequest) Reset()         { (*sriplugin.PluginInfoRequest)(m).Reset() }
func (m *PluginInfoRequest) String() string { return (*sriplugin.PluginInfoRequest)(m).String() }
func (*PluginInfoRequest) ProtoMessage()    {}

// PluginInfoReply from public import github.com/spiffe/sri/common/plugin/plugin.proto
type PluginInfoReply sriplugin.PluginInfoReply

func (m *PluginInfoReply) Reset()         { (*sriplugin.PluginInfoReply)(m).Reset() }
func (m *PluginInfoReply) String() string { return (*sriplugin.PluginInfoReply)(m).String() }
func (*PluginInfoReply) ProtoMessage()    {}
func (m *PluginInfoReply) GetPluginInfo() []*GetPluginInfoResponse {
	o := (*sriplugin.PluginInfoReply)(m).GetPluginInfo()
	if o == nil {
		return nil
	}
	s := make([]*GetPluginInfoResponse, len(o))
	for i, x := range o {
		s[i] = (*GetPluginInfoResponse)(x)
	}
	return s
}

// StopRequest from public import github.com/spiffe/sri/common/plugin/plugin.proto
type StopRequest sriplugin.StopRequest

func (m *StopRequest) Reset()         { (*sriplugin.StopRequest)(m).Reset() }
func (m *StopRequest) String() string { return (*sriplugin.StopRequest)(m).String() }
func (*StopRequest) ProtoMessage()    {}

// StopReply from public import github.com/spiffe/sri/common/plugin/plugin.proto
type StopReply sriplugin.StopReply

func (m *StopReply) Reset()         { (*sriplugin.StopReply)(m).Reset() }
func (m *StopReply) String() string { return (*sriplugin.StopReply)(m).String() }
func (*StopReply) ProtoMessage()    {}

// Empty from public import github.com/spiffe/sri/common/common.proto
type Empty common.Empty

func (m *Empty) Reset()         { (*common.Empty)(m).Reset() }
func (m *Empty) String() string { return (*common.Empty)(m).String() }
func (*Empty) ProtoMessage()    {}

// AttestedData from public import github.com/spiffe/sri/common/common.proto
type AttestedData common.AttestedData

func (m *AttestedData) Reset()          { (*common.AttestedData)(m).Reset() }
func (m *AttestedData) String() string  { return (*common.AttestedData)(m).String() }
func (*AttestedData) ProtoMessage()     {}
func (m *AttestedData) GetType() string { return (*common.AttestedData)(m).GetType() }
func (m *AttestedData) GetData() string { return (*common.AttestedData)(m).GetData() }

// Selector from public import github.com/spiffe/sri/common/common.proto
type Selector common.Selector

func (m *Selector) Reset()           { (*common.Selector)(m).Reset() }
func (m *Selector) String() string   { return (*common.Selector)(m).String() }
func (*Selector) ProtoMessage()      {}
func (m *Selector) GetType() string  { return (*common.Selector)(m).GetType() }
func (m *Selector) GetValue() string { return (*common.Selector)(m).GetValue() }

// Selectors from public import github.com/spiffe/sri/common/common.proto
type Selectors common.Selectors

func (m *Selectors) Reset()         { (*common.Selectors)(m).Reset() }
func (m *Selectors) String() string { return (*common.Selectors)(m).String() }
func (*Selectors) ProtoMessage()    {}
func (m *Selectors) GetEntries() []*Selector {
	o := (*common.Selectors)(m).GetEntries()
	if o == nil {
		return nil
	}
	s := make([]*Selector, len(o))
	for i, x := range o {
		s[i] = (*Selector)(x)
	}
	return s
}

// RegistrationEntry from public import github.com/spiffe/sri/common/common.proto
type RegistrationEntry common.RegistrationEntry

func (m *RegistrationEntry) Reset()         { (*common.RegistrationEntry)(m).Reset() }
func (m *RegistrationEntry) String() string { return (*common.RegistrationEntry)(m).String() }
func (*RegistrationEntry) ProtoMessage()    {}
func (m *RegistrationEntry) GetSelectors() []*Selector {
	o := (*common.RegistrationEntry)(m).GetSelectors()
	if o == nil {
		return nil
	}
	s := make([]*Selector, len(o))
	for i, x := range o {
		s[i] = (*Selector)(x)
	}
	return s
}
func (m *RegistrationEntry) GetParentId() string { return (*common.RegistrationEntry)(m).GetParentId() }
func (m *RegistrationEntry) GetSpiffeId() string { return (*common.RegistrationEntry)(m).GetSpiffeId() }
func (m *RegistrationEntry) GetTtl() int32       { return (*common.RegistrationEntry)(m).GetTtl() }
func (m *RegistrationEntry) GetFbSpiffeIds() []string {
	return (*common.RegistrationEntry)(m).GetFbSpiffeIds()
}

// RegistrationEntries from public import github.com/spiffe/sri/common/common.proto
type RegistrationEntries common.RegistrationEntries

func (m *RegistrationEntries) Reset()         { (*common.RegistrationEntries)(m).Reset() }
func (m *RegistrationEntries) String() string { return (*common.RegistrationEntries)(m).String() }
func (*RegistrationEntries) ProtoMessage()    {}
func (m *RegistrationEntries) GetEntries() []*RegistrationEntry {
	o := (*common.RegistrationEntries)(m).GetEntries()
	if o == nil {
		return nil
	}
	s := make([]*RegistrationEntry, len(o))
	for i, x := range o {
		s[i] = (*RegistrationEntry)(x)
	}
	return s
}

// *Represents a request to attest a node.
type AttestRequest struct {
	AttestedData   *common.AttestedData `protobuf:"bytes,1,opt,name=attestedData" json:"attestedData,omitempty"`
	AttestedBefore bool                 `protobuf:"varint,2,opt,name=attestedBefore" json:"attestedBefore,omitempty"`
}

func (m *AttestRequest) Reset()                    { *m = AttestRequest{} }
func (m *AttestRequest) String() string            { return proto.CompactTextString(m) }
func (*AttestRequest) ProtoMessage()               {}
func (*AttestRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AttestRequest) GetAttestedData() *common.AttestedData {
	if m != nil {
		return m.AttestedData
	}
	return nil
}

func (m *AttestRequest) GetAttestedBefore() bool {
	if m != nil {
		return m.AttestedBefore
	}
	return false
}

// *Represents a response when attesting a node.
type AttestResponse struct {
	Valid        bool   `protobuf:"varint,1,opt,name=valid" json:"valid,omitempty"`
	BaseSPIFFEID string `protobuf:"bytes,2,opt,name=baseSPIFFEID" json:"baseSPIFFEID,omitempty"`
}

func (m *AttestResponse) Reset()                    { *m = AttestResponse{} }
func (m *AttestResponse) String() string            { return proto.CompactTextString(m) }
func (*AttestResponse) ProtoMessage()               {}
func (*AttestResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AttestResponse) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *AttestResponse) GetBaseSPIFFEID() string {
	if m != nil {
		return m.BaseSPIFFEID
	}
	return ""
}

func init() {
	proto.RegisterType((*AttestRequest)(nil), "cpnodeattestor.AttestRequest")
	proto.RegisterType((*AttestResponse)(nil), "cpnodeattestor.AttestResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for NodeAttestor service

type NodeAttestorClient interface {
	// *Responsible for configuration of the plugin.
	Configure(ctx context.Context, in *sriplugin.ConfigureRequest, opts ...grpc.CallOption) (*sriplugin.ConfigureResponse, error)
	// *Returns the  version and related metadata of the installed plugin.
	GetPluginInfo(ctx context.Context, in *sriplugin.GetPluginInfoRequest, opts ...grpc.CallOption) (*sriplugin.GetPluginInfoResponse, error)
	// *Attesta a node.
	Attest(ctx context.Context, in *AttestRequest, opts ...grpc.CallOption) (*AttestResponse, error)
}

type nodeAttestorClient struct {
	cc *grpc.ClientConn
}

func NewNodeAttestorClient(cc *grpc.ClientConn) NodeAttestorClient {
	return &nodeAttestorClient{cc}
}

func (c *nodeAttestorClient) Configure(ctx context.Context, in *sriplugin.ConfigureRequest, opts ...grpc.CallOption) (*sriplugin.ConfigureResponse, error) {
	out := new(sriplugin.ConfigureResponse)
	err := grpc.Invoke(ctx, "/cpnodeattestor.NodeAttestor/Configure", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeAttestorClient) GetPluginInfo(ctx context.Context, in *sriplugin.GetPluginInfoRequest, opts ...grpc.CallOption) (*sriplugin.GetPluginInfoResponse, error) {
	out := new(sriplugin.GetPluginInfoResponse)
	err := grpc.Invoke(ctx, "/cpnodeattestor.NodeAttestor/GetPluginInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeAttestorClient) Attest(ctx context.Context, in *AttestRequest, opts ...grpc.CallOption) (*AttestResponse, error) {
	out := new(AttestResponse)
	err := grpc.Invoke(ctx, "/cpnodeattestor.NodeAttestor/Attest", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NodeAttestor service

type NodeAttestorServer interface {
	// *Responsible for configuration of the plugin.
	Configure(context.Context, *sriplugin.ConfigureRequest) (*sriplugin.ConfigureResponse, error)
	// *Returns the  version and related metadata of the installed plugin.
	GetPluginInfo(context.Context, *sriplugin.GetPluginInfoRequest) (*sriplugin.GetPluginInfoResponse, error)
	// *Attesta a node.
	Attest(context.Context, *AttestRequest) (*AttestResponse, error)
}

func RegisterNodeAttestorServer(s *grpc.Server, srv NodeAttestorServer) {
	s.RegisterService(&_NodeAttestor_serviceDesc, srv)
}

func _NodeAttestor_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(sriplugin.ConfigureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeAttestorServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cpnodeattestor.NodeAttestor/Configure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeAttestorServer).Configure(ctx, req.(*sriplugin.ConfigureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeAttestor_GetPluginInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(sriplugin.GetPluginInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeAttestorServer).GetPluginInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cpnodeattestor.NodeAttestor/GetPluginInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeAttestorServer).GetPluginInfo(ctx, req.(*sriplugin.GetPluginInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeAttestor_Attest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeAttestorServer).Attest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cpnodeattestor.NodeAttestor/Attest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeAttestorServer).Attest(ctx, req.(*AttestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NodeAttestor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cpnodeattestor.NodeAttestor",
	HandlerType: (*NodeAttestorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Configure",
			Handler:    _NodeAttestor_Configure_Handler,
		},
		{
			MethodName: "GetPluginInfo",
			Handler:    _NodeAttestor_GetPluginInfo_Handler,
		},
		{
			MethodName: "Attest",
			Handler:    _NodeAttestor_Attest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "node_attestor.proto",
}

func init() { proto.RegisterFile("node_attestor.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xcd, 0x4e, 0xf3, 0x30,
	0x10, 0xfc, 0xf2, 0x49, 0x54, 0xad, 0x69, 0x7b, 0x30, 0x3d, 0x54, 0xe1, 0xaf, 0xea, 0x01, 0x95,
	0x8b, 0x83, 0xca, 0x85, 0x6b, 0xa1, 0xa4, 0x0a, 0x07, 0x14, 0x99, 0x07, 0x40, 0xf9, 0xd9, 0x04,
	0x4b, 0x4d, 0x36, 0xb5, 0x1d, 0x5e, 0x9c, 0x17, 0x40, 0xc4, 0x0e, 0x6a, 0x10, 0xe5, 0x64, 0xed,
	0xcc, 0xec, 0xcc, 0x7a, 0x97, 0x9c, 0x94, 0x98, 0xc2, 0x6b, 0xa4, 0x35, 0x28, 0x8d, 0x92, 0x55,
	0x12, 0x35, 0xd2, 0x71, 0x52, 0x7d, 0xc1, 0x2d, 0xea, 0xde, 0xe4, 0x42, 0xbf, 0xd5, 0x31, 0x4b,
	0xb0, 0xf0, 0x54, 0x25, 0xb2, 0x0c, 0x3c, 0x25, 0x85, 0x97, 0x60, 0x51, 0x60, 0xe9, 0x55, 0xdb,
	0x3a, 0x17, 0xed, 0x63, 0x1c, 0xdc, 0xeb, 0x3f, 0x3b, 0xcc, 0x63, 0xa4, 0xf3, 0x1d, 0x19, 0xad,
	0x9a, 0x20, 0x0e, 0xbb, 0x1a, 0x94, 0xa6, 0x77, 0x64, 0x68, 0x92, 0x21, 0x5d, 0x47, 0x3a, 0x9a,
	0x3a, 0x33, 0x67, 0x71, 0xbc, 0x9c, 0x30, 0xdb, 0xb5, 0xda, 0xe3, 0x78, 0x47, 0x49, 0xaf, 0xc8,
	0xb8, 0xad, 0xef, 0x21, 0x43, 0x09, 0xd3, 0xff, 0x33, 0x67, 0xd1, 0xe7, 0x3f, 0xd0, 0xf9, 0x13,
	0x19, 0xb7, 0x91, 0xaa, 0xc2, 0x52, 0x01, 0x9d, 0x90, 0xa3, 0xf7, 0x68, 0x2b, 0xd2, 0x26, 0xac,
	0xcf, 0x4d, 0x41, 0xe7, 0x64, 0x18, 0x47, 0x0a, 0x5e, 0xc2, 0xc0, 0xf7, 0x1f, 0x83, 0x75, 0xe3,
	0x36, 0xe0, 0x1d, 0x6c, 0xf9, 0xe1, 0x90, 0xe1, 0x33, 0xa6, 0xb0, 0xb2, 0xcb, 0xa2, 0x3e, 0x19,
	0x3c, 0x60, 0x99, 0x89, 0xbc, 0x96, 0x40, 0x4f, 0x99, 0x92, 0xc2, 0x6e, 0xe6, 0x1b, 0xb5, 0x1f,
	0x75, 0xcf, 0x7e, 0x27, 0xed, 0x48, 0x9c, 0x8c, 0x36, 0xa0, 0xc3, 0x86, 0x0e, 0xca, 0x0c, 0xe9,
	0xe5, 0x9e, 0xbc, 0xc3, 0xb4, 0x7e, 0xb3, 0xc3, 0x02, 0xeb, 0xb9, 0x21, 0x3d, 0x33, 0x27, 0x3d,
	0x67, 0xdd, 0x1b, 0xb3, 0xce, 0x0d, 0xdc, 0x8b, 0x43, 0xb4, 0x31, 0x0a, 0xff, 0x85, 0x4e, 0xdc,
	0x6b, 0xee, 0x77, 0xfb, 0x19, 0x00, 0x00, 0xff, 0xff, 0x95, 0x32, 0xf6, 0xb5, 0x43, 0x02, 0x00,
	0x00,
}