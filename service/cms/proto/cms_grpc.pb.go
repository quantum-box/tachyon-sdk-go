// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package cmspb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CmsApiClient is the client API for CmsApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CmsApiClient interface {
	GetById(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	FindAll(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*FindResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
}

type cmsApiClient struct {
	cc grpc.ClientConnInterface
}

func NewCmsApiClient(cc grpc.ClientConnInterface) CmsApiClient {
	return &cmsApiClient{cc}
}

func (c *cmsApiClient) GetById(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/cms.CmsApi/GetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsApiClient) FindAll(ctx context.Context, in *FindRequest, opts ...grpc.CallOption) (*FindResponse, error) {
	out := new(FindResponse)
	err := c.cc.Invoke(ctx, "/cms.CmsApi/FindAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsApiClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/cms.CmsApi/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmsApiClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/cms.CmsApi/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CmsApiServer is the server API for CmsApi service.
// All implementations must embed UnimplementedCmsApiServer
// for forward compatibility
type CmsApiServer interface {
	GetById(context.Context, *GetRequest) (*GetResponse, error)
	FindAll(context.Context, *FindRequest) (*FindResponse, error)
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	mustEmbedUnimplementedCmsApiServer()
}

// UnimplementedCmsApiServer must be embedded to have forward compatible implementations.
type UnimplementedCmsApiServer struct {
}

func (UnimplementedCmsApiServer) GetById(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedCmsApiServer) FindAll(context.Context, *FindRequest) (*FindResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAll not implemented")
}
func (UnimplementedCmsApiServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedCmsApiServer) Update(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedCmsApiServer) mustEmbedUnimplementedCmsApiServer() {}

// UnsafeCmsApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CmsApiServer will
// result in compilation errors.
type UnsafeCmsApiServer interface {
	mustEmbedUnimplementedCmsApiServer()
}

func RegisterCmsApiServer(s grpc.ServiceRegistrar, srv CmsApiServer) {
	s.RegisterService(&CmsApi_ServiceDesc, srv)
}

func _CmsApi_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmsApiServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cms.CmsApi/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmsApiServer).GetById(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CmsApi_FindAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmsApiServer).FindAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cms.CmsApi/FindAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmsApiServer).FindAll(ctx, req.(*FindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CmsApi_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmsApiServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cms.CmsApi/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmsApiServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CmsApi_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmsApiServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cms.CmsApi/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmsApiServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CmsApi_ServiceDesc is the grpc.ServiceDesc for CmsApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CmsApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cms.CmsApi",
	HandlerType: (*CmsApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetById",
			Handler:    _CmsApi_GetById_Handler,
		},
		{
			MethodName: "FindAll",
			Handler:    _CmsApi_FindAll_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _CmsApi_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _CmsApi_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/cms.proto",
}

// QueryApiClient is the client API for QueryApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryApiClient interface {
	QueryById(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error)
}

type queryApiClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryApiClient(cc grpc.ClientConnInterface) QueryApiClient {
	return &queryApiClient{cc}
}

func (c *queryApiClient) QueryById(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, "/cms.QueryApi/QueryById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryApiServer is the server API for QueryApi service.
// All implementations must embed UnimplementedQueryApiServer
// for forward compatibility
type QueryApiServer interface {
	QueryById(context.Context, *QueryRequest) (*QueryResponse, error)
	mustEmbedUnimplementedQueryApiServer()
}

// UnimplementedQueryApiServer must be embedded to have forward compatible implementations.
type UnimplementedQueryApiServer struct {
}

func (UnimplementedQueryApiServer) QueryById(context.Context, *QueryRequest) (*QueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryById not implemented")
}
func (UnimplementedQueryApiServer) mustEmbedUnimplementedQueryApiServer() {}

// UnsafeQueryApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryApiServer will
// result in compilation errors.
type UnsafeQueryApiServer interface {
	mustEmbedUnimplementedQueryApiServer()
}

func RegisterQueryApiServer(s grpc.ServiceRegistrar, srv QueryApiServer) {
	s.RegisterService(&QueryApi_ServiceDesc, srv)
}

func _QueryApi_QueryById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryApiServer).QueryById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cms.QueryApi/QueryById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryApiServer).QueryById(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// QueryApi_ServiceDesc is the grpc.ServiceDesc for QueryApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QueryApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cms.QueryApi",
	HandlerType: (*QueryApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryById",
			Handler:    _QueryApi_QueryById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/cms.proto",
}
