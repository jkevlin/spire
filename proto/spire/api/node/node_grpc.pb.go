// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package node

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// NodeClient is the client API for Node service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NodeClient interface {
	// Attest the node, get base node SVID.
	Attest(ctx context.Context, opts ...grpc.CallOption) (Node_AttestClient, error)
	// Get Workload, Node Agent certs and CA trust bundles. Also used for rotation
	// Base Node SVID or the Registered Node SVID used for this call)
	// List can be empty to allow Node Agent cache refresh).
	FetchX509SVID(ctx context.Context, opts ...grpc.CallOption) (Node_FetchX509SVIDClient, error)
	// Fetches a signed JWT-SVID for a workload intended for a specific audience.
	FetchJWTSVID(ctx context.Context, in *FetchJWTSVIDRequest, opts ...grpc.CallOption) (*FetchJWTSVIDResponse, error)
	// Fetches an X509 CA SVID for a downstream SPIRE server.
	FetchX509CASVID(ctx context.Context, in *FetchX509CASVIDRequest, opts ...grpc.CallOption) (*FetchX509CASVIDResponse, error)
	// PushJWTKeyUpstream pushes new public JWKs to upstream SPIRE Server, unless this
	// is the root server, in which case it stores the JWK in its bundle. Returns an
	// up-to-date list of the JWT signing keys stored in the bundle.
	PushJWTKeyUpstream(ctx context.Context, in *PushJWTKeyUpstreamRequest, opts ...grpc.CallOption) (*PushJWTKeyUpstreamResponse, error)
	// FetchBundle fetches the bundle of the local trust domain
	FetchBundle(ctx context.Context, in *FetchBundleRequest, opts ...grpc.CallOption) (*FetchBundleResponse, error)
}

type nodeClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeClient(cc grpc.ClientConnInterface) NodeClient {
	return &nodeClient{cc}
}

func (c *nodeClient) Attest(ctx context.Context, opts ...grpc.CallOption) (Node_AttestClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Node_serviceDesc.Streams[0], "/spire.api.node.Node/Attest", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeAttestClient{stream}
	return x, nil
}

type Node_AttestClient interface {
	Send(*AttestRequest) error
	Recv() (*AttestResponse, error)
	grpc.ClientStream
}

type nodeAttestClient struct {
	grpc.ClientStream
}

func (x *nodeAttestClient) Send(m *AttestRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *nodeAttestClient) Recv() (*AttestResponse, error) {
	m := new(AttestResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeClient) FetchX509SVID(ctx context.Context, opts ...grpc.CallOption) (Node_FetchX509SVIDClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Node_serviceDesc.Streams[1], "/spire.api.node.Node/FetchX509SVID", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeFetchX509SVIDClient{stream}
	return x, nil
}

type Node_FetchX509SVIDClient interface {
	Send(*FetchX509SVIDRequest) error
	Recv() (*FetchX509SVIDResponse, error)
	grpc.ClientStream
}

type nodeFetchX509SVIDClient struct {
	grpc.ClientStream
}

func (x *nodeFetchX509SVIDClient) Send(m *FetchX509SVIDRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *nodeFetchX509SVIDClient) Recv() (*FetchX509SVIDResponse, error) {
	m := new(FetchX509SVIDResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeClient) FetchJWTSVID(ctx context.Context, in *FetchJWTSVIDRequest, opts ...grpc.CallOption) (*FetchJWTSVIDResponse, error) {
	out := new(FetchJWTSVIDResponse)
	err := c.cc.Invoke(ctx, "/spire.api.node.Node/FetchJWTSVID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) FetchX509CASVID(ctx context.Context, in *FetchX509CASVIDRequest, opts ...grpc.CallOption) (*FetchX509CASVIDResponse, error) {
	out := new(FetchX509CASVIDResponse)
	err := c.cc.Invoke(ctx, "/spire.api.node.Node/FetchX509CASVID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) PushJWTKeyUpstream(ctx context.Context, in *PushJWTKeyUpstreamRequest, opts ...grpc.CallOption) (*PushJWTKeyUpstreamResponse, error) {
	out := new(PushJWTKeyUpstreamResponse)
	err := c.cc.Invoke(ctx, "/spire.api.node.Node/PushJWTKeyUpstream", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) FetchBundle(ctx context.Context, in *FetchBundleRequest, opts ...grpc.CallOption) (*FetchBundleResponse, error) {
	out := new(FetchBundleResponse)
	err := c.cc.Invoke(ctx, "/spire.api.node.Node/FetchBundle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeServer is the server API for Node service.
// All implementations must embed UnimplementedNodeServer
// for forward compatibility
type NodeServer interface {
	// Attest the node, get base node SVID.
	Attest(Node_AttestServer) error
	// Get Workload, Node Agent certs and CA trust bundles. Also used for rotation
	// Base Node SVID or the Registered Node SVID used for this call)
	// List can be empty to allow Node Agent cache refresh).
	FetchX509SVID(Node_FetchX509SVIDServer) error
	// Fetches a signed JWT-SVID for a workload intended for a specific audience.
	FetchJWTSVID(context.Context, *FetchJWTSVIDRequest) (*FetchJWTSVIDResponse, error)
	// Fetches an X509 CA SVID for a downstream SPIRE server.
	FetchX509CASVID(context.Context, *FetchX509CASVIDRequest) (*FetchX509CASVIDResponse, error)
	// PushJWTKeyUpstream pushes new public JWKs to upstream SPIRE Server, unless this
	// is the root server, in which case it stores the JWK in its bundle. Returns an
	// up-to-date list of the JWT signing keys stored in the bundle.
	PushJWTKeyUpstream(context.Context, *PushJWTKeyUpstreamRequest) (*PushJWTKeyUpstreamResponse, error)
	// FetchBundle fetches the bundle of the local trust domain
	FetchBundle(context.Context, *FetchBundleRequest) (*FetchBundleResponse, error)
	mustEmbedUnimplementedNodeServer()
}

// UnimplementedNodeServer must be embedded to have forward compatible implementations.
type UnimplementedNodeServer struct {
}

func (UnimplementedNodeServer) Attest(Node_AttestServer) error {
	return status.Errorf(codes.Unimplemented, "method Attest not implemented")
}
func (UnimplementedNodeServer) FetchX509SVID(Node_FetchX509SVIDServer) error {
	return status.Errorf(codes.Unimplemented, "method FetchX509SVID not implemented")
}
func (UnimplementedNodeServer) FetchJWTSVID(context.Context, *FetchJWTSVIDRequest) (*FetchJWTSVIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchJWTSVID not implemented")
}
func (UnimplementedNodeServer) FetchX509CASVID(context.Context, *FetchX509CASVIDRequest) (*FetchX509CASVIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchX509CASVID not implemented")
}
func (UnimplementedNodeServer) PushJWTKeyUpstream(context.Context, *PushJWTKeyUpstreamRequest) (*PushJWTKeyUpstreamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushJWTKeyUpstream not implemented")
}
func (UnimplementedNodeServer) FetchBundle(context.Context, *FetchBundleRequest) (*FetchBundleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchBundle not implemented")
}
func (UnimplementedNodeServer) mustEmbedUnimplementedNodeServer() {}

// UnsafeNodeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NodeServer will
// result in compilation errors.
type UnsafeNodeServer interface {
	mustEmbedUnimplementedNodeServer()
}

func RegisterNodeServer(s grpc.ServiceRegistrar, srv NodeServer) {
	s.RegisterService(&_Node_serviceDesc, srv)
}

func _Node_Attest_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NodeServer).Attest(&nodeAttestServer{stream})
}

type Node_AttestServer interface {
	Send(*AttestResponse) error
	Recv() (*AttestRequest, error)
	grpc.ServerStream
}

type nodeAttestServer struct {
	grpc.ServerStream
}

func (x *nodeAttestServer) Send(m *AttestResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *nodeAttestServer) Recv() (*AttestRequest, error) {
	m := new(AttestRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Node_FetchX509SVID_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NodeServer).FetchX509SVID(&nodeFetchX509SVIDServer{stream})
}

type Node_FetchX509SVIDServer interface {
	Send(*FetchX509SVIDResponse) error
	Recv() (*FetchX509SVIDRequest, error)
	grpc.ServerStream
}

type nodeFetchX509SVIDServer struct {
	grpc.ServerStream
}

func (x *nodeFetchX509SVIDServer) Send(m *FetchX509SVIDResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *nodeFetchX509SVIDServer) Recv() (*FetchX509SVIDRequest, error) {
	m := new(FetchX509SVIDRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Node_FetchJWTSVID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchJWTSVIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).FetchJWTSVID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.node.Node/FetchJWTSVID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).FetchJWTSVID(ctx, req.(*FetchJWTSVIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_FetchX509CASVID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchX509CASVIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).FetchX509CASVID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.node.Node/FetchX509CASVID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).FetchX509CASVID(ctx, req.(*FetchX509CASVIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_PushJWTKeyUpstream_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushJWTKeyUpstreamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).PushJWTKeyUpstream(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.node.Node/PushJWTKeyUpstream",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).PushJWTKeyUpstream(ctx, req.(*PushJWTKeyUpstreamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_FetchBundle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchBundleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).FetchBundle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.api.node.Node/FetchBundle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).FetchBundle(ctx, req.(*FetchBundleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Node_serviceDesc = grpc.ServiceDesc{
	ServiceName: "spire.api.node.Node",
	HandlerType: (*NodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchJWTSVID",
			Handler:    _Node_FetchJWTSVID_Handler,
		},
		{
			MethodName: "FetchX509CASVID",
			Handler:    _Node_FetchX509CASVID_Handler,
		},
		{
			MethodName: "PushJWTKeyUpstream",
			Handler:    _Node_PushJWTKeyUpstream_Handler,
		},
		{
			MethodName: "FetchBundle",
			Handler:    _Node_FetchBundle_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Attest",
			Handler:       _Node_Attest_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "FetchX509SVID",
			Handler:       _Node_FetchX509SVID_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "spire/api/node/node.proto",
}
