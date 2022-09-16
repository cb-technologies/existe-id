// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.9.2
// source: exist_crud_service.proto

package pb

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

// ExistCRUDServiceClient is the client API for ExistCRUDService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExistCRUDServiceClient interface {
	AddNewPerson(ctx context.Context, in *PersonInfo, opts ...grpc.CallOption) (*Response, error)
	EditPersonInfo(ctx context.Context, in *EditPersonInfoParameters, opts ...grpc.CallOption) (*Response, error)
	FindPersonInfo(ctx context.Context, in *UUID, opts ...grpc.CallOption) (*PersonInfo, error)
}

type existCRUDServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExistCRUDServiceClient(cc grpc.ClientConnInterface) ExistCRUDServiceClient {
	return &existCRUDServiceClient{cc}
}

func (c *existCRUDServiceClient) AddNewPerson(ctx context.Context, in *PersonInfo, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/pb.ExistCRUDService/AddNewPerson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *existCRUDServiceClient) EditPersonInfo(ctx context.Context, in *EditPersonInfoParameters, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/pb.ExistCRUDService/EditPersonInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *existCRUDServiceClient) FindPersonInfo(ctx context.Context, in *UUID, opts ...grpc.CallOption) (*PersonInfo, error) {
	out := new(PersonInfo)
	err := c.cc.Invoke(ctx, "/pb.ExistCRUDService/FindPersonInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExistCRUDServiceServer is the server API for ExistCRUDService service.
// All implementations must embed UnimplementedExistCRUDServiceServer
// for forward compatibility
type ExistCRUDServiceServer interface {
	AddNewPerson(context.Context, *PersonInfo) (*Response, error)
	EditPersonInfo(context.Context, *EditPersonInfoParameters) (*Response, error)
	FindPersonInfo(context.Context, *UUID) (*PersonInfo, error)
	mustEmbedUnimplementedExistCRUDServiceServer()
}

// UnimplementedExistCRUDServiceServer must be embedded to have forward compatible implementations.
type UnimplementedExistCRUDServiceServer struct {
}

func (UnimplementedExistCRUDServiceServer) AddNewPerson(context.Context, *PersonInfo) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddNewPerson not implemented")
}
func (UnimplementedExistCRUDServiceServer) EditPersonInfo(context.Context, *EditPersonInfoParameters) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditPersonInfo not implemented")
}
func (UnimplementedExistCRUDServiceServer) FindPersonInfo(context.Context, *UUID) (*PersonInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindPersonInfo not implemented")
}
func (UnimplementedExistCRUDServiceServer) mustEmbedUnimplementedExistCRUDServiceServer() {}

// UnsafeExistCRUDServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExistCRUDServiceServer will
// result in compilation errors.
type UnsafeExistCRUDServiceServer interface {
	mustEmbedUnimplementedExistCRUDServiceServer()
}

func RegisterExistCRUDServiceServer(s grpc.ServiceRegistrar, srv ExistCRUDServiceServer) {
	s.RegisterService(&ExistCRUDService_ServiceDesc, srv)
}

func _ExistCRUDService_AddNewPerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PersonInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExistCRUDServiceServer).AddNewPerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ExistCRUDService/AddNewPerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExistCRUDServiceServer).AddNewPerson(ctx, req.(*PersonInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExistCRUDService_EditPersonInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditPersonInfoParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExistCRUDServiceServer).EditPersonInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ExistCRUDService/EditPersonInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExistCRUDServiceServer).EditPersonInfo(ctx, req.(*EditPersonInfoParameters))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExistCRUDService_FindPersonInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExistCRUDServiceServer).FindPersonInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ExistCRUDService/FindPersonInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExistCRUDServiceServer).FindPersonInfo(ctx, req.(*UUID))
	}
	return interceptor(ctx, in, info, handler)
}

// ExistCRUDService_ServiceDesc is the grpc.ServiceDesc for ExistCRUDService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExistCRUDService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ExistCRUDService",
	HandlerType: (*ExistCRUDServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddNewPerson",
			Handler:    _ExistCRUDService_AddNewPerson_Handler,
		},
		{
			MethodName: "EditPersonInfo",
			Handler:    _ExistCRUDService_EditPersonInfo_Handler,
		},
		{
			MethodName: "FindPersonInfo",
			Handler:    _ExistCRUDService_FindPersonInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "exist_crud_service.proto",
}