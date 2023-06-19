// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.0
// source: rating_service/rating_service.proto

package rating_service

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

// RatingServiceClient is the client API for RatingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RatingServiceClient interface {
	CreateHostRating(ctx context.Context, in *CreateHostRatingRequest, opts ...grpc.CallOption) (*CreateHostRatingResponse, error)
	CreateAccommodationRating(ctx context.Context, in *CreateAccommodationRatingRequest, opts ...grpc.CallOption) (*CreateAccommodationRatingResponse, error)
	UpdateHostRating(ctx context.Context, in *DatiGaDam, opts ...grpc.CallOption) (*EmptyMessage, error)
	UpdateAccommodationRating(ctx context.Context, in *DatiGaDam, opts ...grpc.CallOption) (*EmptyMessage, error)
	GetUserHostRatings(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*GetUserHostRatingsResponse, error)
	GetUserAccommodationRatings(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*GetUserAccommodationRatingsResponse, error)
	IsHostMarked(ctx context.Context, in *MarkedHostRequest, opts ...grpc.CallOption) (*MarkedHostResponse, error)
	CreateMarkedHost(ctx context.Context, in *MarkedHostRequest, opts ...grpc.CallOption) (*MarkedHost, error)
}

type ratingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRatingServiceClient(cc grpc.ClientConnInterface) RatingServiceClient {
	return &ratingServiceClient{cc}
}

func (c *ratingServiceClient) CreateHostRating(ctx context.Context, in *CreateHostRatingRequest, opts ...grpc.CallOption) (*CreateHostRatingResponse, error) {
	out := new(CreateHostRatingResponse)
	err := c.cc.Invoke(ctx, "/rating_service.RatingService/CreateHostRating", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) CreateAccommodationRating(ctx context.Context, in *CreateAccommodationRatingRequest, opts ...grpc.CallOption) (*CreateAccommodationRatingResponse, error) {
	out := new(CreateAccommodationRatingResponse)
	err := c.cc.Invoke(ctx, "/rating_service.RatingService/CreateAccommodationRating", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) UpdateHostRating(ctx context.Context, in *DatiGaDam, opts ...grpc.CallOption) (*EmptyMessage, error) {
	out := new(EmptyMessage)
	err := c.cc.Invoke(ctx, "/rating_service.RatingService/UpdateHostRating", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) UpdateAccommodationRating(ctx context.Context, in *DatiGaDam, opts ...grpc.CallOption) (*EmptyMessage, error) {
	out := new(EmptyMessage)
	err := c.cc.Invoke(ctx, "/rating_service.RatingService/UpdateAccommodationRating", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) GetUserHostRatings(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*GetUserHostRatingsResponse, error) {
	out := new(GetUserHostRatingsResponse)
	err := c.cc.Invoke(ctx, "/rating_service.RatingService/GetUserHostRatings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) GetUserAccommodationRatings(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*GetUserAccommodationRatingsResponse, error) {
	out := new(GetUserAccommodationRatingsResponse)
	err := c.cc.Invoke(ctx, "/rating_service.RatingService/GetUserAccommodationRatings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) IsHostMarked(ctx context.Context, in *MarkedHostRequest, opts ...grpc.CallOption) (*MarkedHostResponse, error) {
	out := new(MarkedHostResponse)
	err := c.cc.Invoke(ctx, "/rating_service.RatingService/IsHostMarked", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) CreateMarkedHost(ctx context.Context, in *MarkedHostRequest, opts ...grpc.CallOption) (*MarkedHost, error) {
	out := new(MarkedHost)
	err := c.cc.Invoke(ctx, "/rating_service.RatingService/CreateMarkedHost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RatingServiceServer is the server API for RatingService service.
// All implementations must embed UnimplementedRatingServiceServer
// for forward compatibility
type RatingServiceServer interface {
	CreateHostRating(context.Context, *CreateHostRatingRequest) (*CreateHostRatingResponse, error)
	CreateAccommodationRating(context.Context, *CreateAccommodationRatingRequest) (*CreateAccommodationRatingResponse, error)
	UpdateHostRating(context.Context, *DatiGaDam) (*EmptyMessage, error)
	UpdateAccommodationRating(context.Context, *DatiGaDam) (*EmptyMessage, error)
	GetUserHostRatings(context.Context, *EmptyMessage) (*GetUserHostRatingsResponse, error)
	GetUserAccommodationRatings(context.Context, *EmptyMessage) (*GetUserAccommodationRatingsResponse, error)
	IsHostMarked(context.Context, *MarkedHostRequest) (*MarkedHostResponse, error)
	CreateMarkedHost(context.Context, *MarkedHostRequest) (*MarkedHost, error)
	mustEmbedUnimplementedRatingServiceServer()
}

// UnimplementedRatingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRatingServiceServer struct {
}

func (UnimplementedRatingServiceServer) CreateHostRating(context.Context, *CreateHostRatingRequest) (*CreateHostRatingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHostRating not implemented")
}
func (UnimplementedRatingServiceServer) CreateAccommodationRating(context.Context, *CreateAccommodationRatingRequest) (*CreateAccommodationRatingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccommodationRating not implemented")
}
func (UnimplementedRatingServiceServer) UpdateHostRating(context.Context, *DatiGaDam) (*EmptyMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHostRating not implemented")
}
func (UnimplementedRatingServiceServer) UpdateAccommodationRating(context.Context, *DatiGaDam) (*EmptyMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccommodationRating not implemented")
}
func (UnimplementedRatingServiceServer) GetUserHostRatings(context.Context, *EmptyMessage) (*GetUserHostRatingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserHostRatings not implemented")
}
func (UnimplementedRatingServiceServer) GetUserAccommodationRatings(context.Context, *EmptyMessage) (*GetUserAccommodationRatingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserAccommodationRatings not implemented")
}
func (UnimplementedRatingServiceServer) IsHostMarked(context.Context, *MarkedHostRequest) (*MarkedHostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsHostMarked not implemented")
}
func (UnimplementedRatingServiceServer) CreateMarkedHost(context.Context, *MarkedHostRequest) (*MarkedHost, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMarkedHost not implemented")
}
func (UnimplementedRatingServiceServer) mustEmbedUnimplementedRatingServiceServer() {}

// UnsafeRatingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RatingServiceServer will
// result in compilation errors.
type UnsafeRatingServiceServer interface {
	mustEmbedUnimplementedRatingServiceServer()
}

func RegisterRatingServiceServer(s grpc.ServiceRegistrar, srv RatingServiceServer) {
	s.RegisterService(&RatingService_ServiceDesc, srv)
}

func _RatingService_CreateHostRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateHostRatingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).CreateHostRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rating_service.RatingService/CreateHostRating",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).CreateHostRating(ctx, req.(*CreateHostRatingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_CreateAccommodationRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccommodationRatingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).CreateAccommodationRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rating_service.RatingService/CreateAccommodationRating",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).CreateAccommodationRating(ctx, req.(*CreateAccommodationRatingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_UpdateHostRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DatiGaDam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).UpdateHostRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rating_service.RatingService/UpdateHostRating",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).UpdateHostRating(ctx, req.(*DatiGaDam))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_UpdateAccommodationRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DatiGaDam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).UpdateAccommodationRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rating_service.RatingService/UpdateAccommodationRating",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).UpdateAccommodationRating(ctx, req.(*DatiGaDam))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_GetUserHostRatings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).GetUserHostRatings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rating_service.RatingService/GetUserHostRatings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).GetUserHostRatings(ctx, req.(*EmptyMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_GetUserAccommodationRatings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).GetUserAccommodationRatings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rating_service.RatingService/GetUserAccommodationRatings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).GetUserAccommodationRatings(ctx, req.(*EmptyMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_IsHostMarked_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarkedHostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).IsHostMarked(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rating_service.RatingService/IsHostMarked",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).IsHostMarked(ctx, req.(*MarkedHostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_CreateMarkedHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarkedHostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).CreateMarkedHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rating_service.RatingService/CreateMarkedHost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).CreateMarkedHost(ctx, req.(*MarkedHostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RatingService_ServiceDesc is the grpc.ServiceDesc for RatingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RatingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rating_service.RatingService",
	HandlerType: (*RatingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateHostRating",
			Handler:    _RatingService_CreateHostRating_Handler,
		},
		{
			MethodName: "CreateAccommodationRating",
			Handler:    _RatingService_CreateAccommodationRating_Handler,
		},
		{
			MethodName: "UpdateHostRating",
			Handler:    _RatingService_UpdateHostRating_Handler,
		},
		{
			MethodName: "UpdateAccommodationRating",
			Handler:    _RatingService_UpdateAccommodationRating_Handler,
		},
		{
			MethodName: "GetUserHostRatings",
			Handler:    _RatingService_GetUserHostRatings_Handler,
		},
		{
			MethodName: "GetUserAccommodationRatings",
			Handler:    _RatingService_GetUserAccommodationRatings_Handler,
		},
		{
			MethodName: "IsHostMarked",
			Handler:    _RatingService_IsHostMarked_Handler,
		},
		{
			MethodName: "CreateMarkedHost",
			Handler:    _RatingService_CreateMarkedHost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rating_service/rating_service.proto",
}