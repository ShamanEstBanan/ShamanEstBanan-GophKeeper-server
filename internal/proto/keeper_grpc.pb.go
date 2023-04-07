// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: proto/keeper.proto

package __

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

const (
	KeeperService_SignUp_FullMethodName           = "/gophKeeper.proto.KeeperService/SignUp"
	KeeperService_LogIn_FullMethodName            = "/gophKeeper.proto.KeeperService/LogIn"
	KeeperService_GetAllRecords_FullMethodName    = "/gophKeeper.proto.KeeperService/GetAllRecords"
	KeeperService_GetRecordsByType_FullMethodName = "/gophKeeper.proto.KeeperService/GetRecordsByType"
	KeeperService_CreateRecord_FullMethodName     = "/gophKeeper.proto.KeeperService/CreateRecord"
	KeeperService_GetRecord_FullMethodName        = "/gophKeeper.proto.KeeperService/GetRecord"
	KeeperService_EditRecord_FullMethodName       = "/gophKeeper.proto.KeeperService/EditRecord"
	KeeperService_DeleteRecord_FullMethodName     = "/gophKeeper.proto.KeeperService/DeleteRecord"
)

// KeeperServiceClient is the client API for KeeperService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KeeperServiceClient interface {
	SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponse, error)
	LogIn(ctx context.Context, in *LogInRequest, opts ...grpc.CallOption) (*LogInResponse, error)
	GetAllRecords(ctx context.Context, in *GetAllRecordsRequest, opts ...grpc.CallOption) (*GetAllRecordsResponse, error)
	GetRecordsByType(ctx context.Context, in *GetRecordsByTypeRequest, opts ...grpc.CallOption) (*GetRecordsByTypeResponse, error)
	CreateRecord(ctx context.Context, in *CreateRecordRequest, opts ...grpc.CallOption) (*CreateRecordResponse, error)
	GetRecord(ctx context.Context, in *GetRecordRequest, opts ...grpc.CallOption) (*GetRecordResponse, error)
	EditRecord(ctx context.Context, in *EditRecordRequest, opts ...grpc.CallOption) (*EditRecordResponse, error)
	DeleteRecord(ctx context.Context, in *DeleteRecordRequest, opts ...grpc.CallOption) (*DeleteRecordResponse, error)
}

type keeperServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKeeperServiceClient(cc grpc.ClientConnInterface) KeeperServiceClient {
	return &keeperServiceClient{cc}
}

func (c *keeperServiceClient) SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponse, error) {
	out := new(SignUpResponse)
	err := c.cc.Invoke(ctx, KeeperService_SignUp_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keeperServiceClient) LogIn(ctx context.Context, in *LogInRequest, opts ...grpc.CallOption) (*LogInResponse, error) {
	out := new(LogInResponse)
	err := c.cc.Invoke(ctx, KeeperService_LogIn_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keeperServiceClient) GetAllRecords(ctx context.Context, in *GetAllRecordsRequest, opts ...grpc.CallOption) (*GetAllRecordsResponse, error) {
	out := new(GetAllRecordsResponse)
	err := c.cc.Invoke(ctx, KeeperService_GetAllRecords_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keeperServiceClient) GetRecordsByType(ctx context.Context, in *GetRecordsByTypeRequest, opts ...grpc.CallOption) (*GetRecordsByTypeResponse, error) {
	out := new(GetRecordsByTypeResponse)
	err := c.cc.Invoke(ctx, KeeperService_GetRecordsByType_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keeperServiceClient) CreateRecord(ctx context.Context, in *CreateRecordRequest, opts ...grpc.CallOption) (*CreateRecordResponse, error) {
	out := new(CreateRecordResponse)
	err := c.cc.Invoke(ctx, KeeperService_CreateRecord_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keeperServiceClient) GetRecord(ctx context.Context, in *GetRecordRequest, opts ...grpc.CallOption) (*GetRecordResponse, error) {
	out := new(GetRecordResponse)
	err := c.cc.Invoke(ctx, KeeperService_GetRecord_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keeperServiceClient) EditRecord(ctx context.Context, in *EditRecordRequest, opts ...grpc.CallOption) (*EditRecordResponse, error) {
	out := new(EditRecordResponse)
	err := c.cc.Invoke(ctx, KeeperService_EditRecord_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keeperServiceClient) DeleteRecord(ctx context.Context, in *DeleteRecordRequest, opts ...grpc.CallOption) (*DeleteRecordResponse, error) {
	out := new(DeleteRecordResponse)
	err := c.cc.Invoke(ctx, KeeperService_DeleteRecord_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeeperServiceServer is the server API for KeeperService service.
// All implementations must embed UnimplementedKeeperServiceServer
// for forward compatibility
type KeeperServiceServer interface {
	SignUp(context.Context, *SignUpRequest) (*SignUpResponse, error)
	LogIn(context.Context, *LogInRequest) (*LogInResponse, error)
	GetAllRecords(context.Context, *GetAllRecordsRequest) (*GetAllRecordsResponse, error)
	GetRecordsByType(context.Context, *GetRecordsByTypeRequest) (*GetRecordsByTypeResponse, error)
	CreateRecord(context.Context, *CreateRecordRequest) (*CreateRecordResponse, error)
	GetRecord(context.Context, *GetRecordRequest) (*GetRecordResponse, error)
	EditRecord(context.Context, *EditRecordRequest) (*EditRecordResponse, error)
	DeleteRecord(context.Context, *DeleteRecordRequest) (*DeleteRecordResponse, error)
	mustEmbedUnimplementedKeeperServiceServer()
}

// UnimplementedKeeperServiceServer must be embedded to have forward compatible implementations.
type UnimplementedKeeperServiceServer struct {
}

func (UnimplementedKeeperServiceServer) SignUp(context.Context, *SignUpRequest) (*SignUpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}
func (UnimplementedKeeperServiceServer) LogIn(context.Context, *LogInRequest) (*LogInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogIn not implemented")
}
func (UnimplementedKeeperServiceServer) GetAllRecords(context.Context, *GetAllRecordsRequest) (*GetAllRecordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllRecords not implemented")
}
func (UnimplementedKeeperServiceServer) GetRecordsByType(context.Context, *GetRecordsByTypeRequest) (*GetRecordsByTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecordsByType not implemented")
}
func (UnimplementedKeeperServiceServer) CreateRecord(context.Context, *CreateRecordRequest) (*CreateRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRecord not implemented")
}
func (UnimplementedKeeperServiceServer) GetRecord(context.Context, *GetRecordRequest) (*GetRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecord not implemented")
}
func (UnimplementedKeeperServiceServer) EditRecord(context.Context, *EditRecordRequest) (*EditRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditRecord not implemented")
}
func (UnimplementedKeeperServiceServer) DeleteRecord(context.Context, *DeleteRecordRequest) (*DeleteRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRecord not implemented")
}
func (UnimplementedKeeperServiceServer) mustEmbedUnimplementedKeeperServiceServer() {}

// UnsafeKeeperServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KeeperServiceServer will
// result in compilation errors.
type UnsafeKeeperServiceServer interface {
	mustEmbedUnimplementedKeeperServiceServer()
}

func RegisterKeeperServiceServer(s grpc.ServiceRegistrar, srv KeeperServiceServer) {
	s.RegisterService(&KeeperService_ServiceDesc, srv)
}

func _KeeperService_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeeperServiceServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeeperService_SignUp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeeperServiceServer).SignUp(ctx, req.(*SignUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeeperService_LogIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeeperServiceServer).LogIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeeperService_LogIn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeeperServiceServer).LogIn(ctx, req.(*LogInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeeperService_GetAllRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRecordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeeperServiceServer).GetAllRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeeperService_GetAllRecords_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeeperServiceServer).GetAllRecords(ctx, req.(*GetAllRecordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeeperService_GetRecordsByType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecordsByTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeeperServiceServer).GetRecordsByType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeeperService_GetRecordsByType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeeperServiceServer).GetRecordsByType(ctx, req.(*GetRecordsByTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeeperService_CreateRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeeperServiceServer).CreateRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeeperService_CreateRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeeperServiceServer).CreateRecord(ctx, req.(*CreateRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeeperService_GetRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeeperServiceServer).GetRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeeperService_GetRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeeperServiceServer).GetRecord(ctx, req.(*GetRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeeperService_EditRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeeperServiceServer).EditRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeeperService_EditRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeeperServiceServer).EditRecord(ctx, req.(*EditRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeeperService_DeleteRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeeperServiceServer).DeleteRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeeperService_DeleteRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeeperServiceServer).DeleteRecord(ctx, req.(*DeleteRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KeeperService_ServiceDesc is the grpc.ServiceDesc for KeeperService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KeeperService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gophKeeper.proto.KeeperService",
	HandlerType: (*KeeperServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignUp",
			Handler:    _KeeperService_SignUp_Handler,
		},
		{
			MethodName: "LogIn",
			Handler:    _KeeperService_LogIn_Handler,
		},
		{
			MethodName: "GetAllRecords",
			Handler:    _KeeperService_GetAllRecords_Handler,
		},
		{
			MethodName: "GetRecordsByType",
			Handler:    _KeeperService_GetRecordsByType_Handler,
		},
		{
			MethodName: "CreateRecord",
			Handler:    _KeeperService_CreateRecord_Handler,
		},
		{
			MethodName: "GetRecord",
			Handler:    _KeeperService_GetRecord_Handler,
		},
		{
			MethodName: "EditRecord",
			Handler:    _KeeperService_EditRecord_Handler,
		},
		{
			MethodName: "DeleteRecord",
			Handler:    _KeeperService_DeleteRecord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/keeper.proto",
}
