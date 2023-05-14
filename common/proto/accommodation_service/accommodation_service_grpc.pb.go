// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.0
// source: accommodation_service/accommodation_service.proto

package accommodation_service

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

// AccommodationServiceClient is the client API for AccommodationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccommodationServiceClient interface {
	CreateAcc(ctx context.Context, in *CreateAccommodationRequest, opts ...grpc.CallOption) (*CreateAccommodationResponse, error)
	GetAccommodationById(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*CreateAccommodationResponse, error)
	SearchAccommodations(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
	CreateAvailableDate(ctx context.Context, in *CreateAvailableDateRequest, opts ...grpc.CallOption) (*CreateAvailableDateResponse, error)
	GetAvailableDateById(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*CreateAvailableDateResponse, error)
	UpdateAvailableDate(ctx context.Context, in *UpdateAvailableDateRequest, opts ...grpc.CallOption) (*UpdateAvailableDateResponse, error)
	TimeSlotAvailableForAccommodation(ctx context.Context, in *TimeSlotAvailableRequest, opts ...grpc.CallOption) (*TimeSlotAvailableResponse, error)
	GetAutomaticAcceptById(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*GetAutomaticAcceptByIdResponse, error)
}

type accommodationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccommodationServiceClient(cc grpc.ClientConnInterface) AccommodationServiceClient {
	return &accommodationServiceClient{cc}
}

func (c *accommodationServiceClient) CreateAcc(ctx context.Context, in *CreateAccommodationRequest, opts ...grpc.CallOption) (*CreateAccommodationResponse, error) {
	out := new(CreateAccommodationResponse)
	err := c.cc.Invoke(ctx, "/accommodation_service.AccommodationService/CreateAcc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) GetAccommodationById(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*CreateAccommodationResponse, error) {
	out := new(CreateAccommodationResponse)
	err := c.cc.Invoke(ctx, "/accommodation_service.AccommodationService/GetAccommodationById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) SearchAccommodations(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/accommodation_service.AccommodationService/SearchAccommodations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) CreateAvailableDate(ctx context.Context, in *CreateAvailableDateRequest, opts ...grpc.CallOption) (*CreateAvailableDateResponse, error) {
	out := new(CreateAvailableDateResponse)
	err := c.cc.Invoke(ctx, "/accommodation_service.AccommodationService/CreateAvailableDate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) GetAvailableDateById(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*CreateAvailableDateResponse, error) {
	out := new(CreateAvailableDateResponse)
	err := c.cc.Invoke(ctx, "/accommodation_service.AccommodationService/GetAvailableDateById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) UpdateAvailableDate(ctx context.Context, in *UpdateAvailableDateRequest, opts ...grpc.CallOption) (*UpdateAvailableDateResponse, error) {
	out := new(UpdateAvailableDateResponse)
	err := c.cc.Invoke(ctx, "/accommodation_service.AccommodationService/UpdateAvailableDate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) TimeSlotAvailableForAccommodation(ctx context.Context, in *TimeSlotAvailableRequest, opts ...grpc.CallOption) (*TimeSlotAvailableResponse, error) {
	out := new(TimeSlotAvailableResponse)
	err := c.cc.Invoke(ctx, "/accommodation_service.AccommodationService/TimeSlotAvailableForAccommodation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) GetAutomaticAcceptById(ctx context.Context, in *GetByIdRequest, opts ...grpc.CallOption) (*GetAutomaticAcceptByIdResponse, error) {
	out := new(GetAutomaticAcceptByIdResponse)
	err := c.cc.Invoke(ctx, "/accommodation_service.AccommodationService/GetAutomaticAcceptById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccommodationServiceServer is the server API for AccommodationService service.
// All implementations must embed UnimplementedAccommodationServiceServer
// for forward compatibility
type AccommodationServiceServer interface {
	CreateAcc(context.Context, *CreateAccommodationRequest) (*CreateAccommodationResponse, error)
	GetAccommodationById(context.Context, *GetByIdRequest) (*CreateAccommodationResponse, error)
	SearchAccommodations(context.Context, *SearchRequest) (*SearchResponse, error)
	CreateAvailableDate(context.Context, *CreateAvailableDateRequest) (*CreateAvailableDateResponse, error)
	GetAvailableDateById(context.Context, *GetByIdRequest) (*CreateAvailableDateResponse, error)
	UpdateAvailableDate(context.Context, *UpdateAvailableDateRequest) (*UpdateAvailableDateResponse, error)
	TimeSlotAvailableForAccommodation(context.Context, *TimeSlotAvailableRequest) (*TimeSlotAvailableResponse, error)
	GetAutomaticAcceptById(context.Context, *GetByIdRequest) (*GetAutomaticAcceptByIdResponse, error)
	mustEmbedUnimplementedAccommodationServiceServer()
}

// UnimplementedAccommodationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAccommodationServiceServer struct {
}

func (UnimplementedAccommodationServiceServer) CreateAcc(context.Context, *CreateAccommodationRequest) (*CreateAccommodationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAcc not implemented")
}
func (UnimplementedAccommodationServiceServer) GetAccommodationById(context.Context, *GetByIdRequest) (*CreateAccommodationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccommodationById not implemented")
}
func (UnimplementedAccommodationServiceServer) SearchAccommodations(context.Context, *SearchRequest) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchAccommodations not implemented")
}
func (UnimplementedAccommodationServiceServer) CreateAvailableDate(context.Context, *CreateAvailableDateRequest) (*CreateAvailableDateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAvailableDate not implemented")
}
func (UnimplementedAccommodationServiceServer) GetAvailableDateById(context.Context, *GetByIdRequest) (*CreateAvailableDateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAvailableDateById not implemented")
}
func (UnimplementedAccommodationServiceServer) UpdateAvailableDate(context.Context, *UpdateAvailableDateRequest) (*UpdateAvailableDateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAvailableDate not implemented")
}
func (UnimplementedAccommodationServiceServer) TimeSlotAvailableForAccommodation(context.Context, *TimeSlotAvailableRequest) (*TimeSlotAvailableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TimeSlotAvailableForAccommodation not implemented")
}
func (UnimplementedAccommodationServiceServer) GetAutomaticAcceptById(context.Context, *GetByIdRequest) (*GetAutomaticAcceptByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAutomaticAcceptById not implemented")
}
func (UnimplementedAccommodationServiceServer) mustEmbedUnimplementedAccommodationServiceServer() {}

// UnsafeAccommodationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccommodationServiceServer will
// result in compilation errors.
type UnsafeAccommodationServiceServer interface {
	mustEmbedUnimplementedAccommodationServiceServer()
}

func RegisterAccommodationServiceServer(s grpc.ServiceRegistrar, srv AccommodationServiceServer) {
	s.RegisterService(&AccommodationService_ServiceDesc, srv)
}

func _AccommodationService_CreateAcc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccommodationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).CreateAcc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accommodation_service.AccommodationService/CreateAcc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).CreateAcc(ctx, req.(*CreateAccommodationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_GetAccommodationById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).GetAccommodationById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accommodation_service.AccommodationService/GetAccommodationById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).GetAccommodationById(ctx, req.(*GetByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_SearchAccommodations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).SearchAccommodations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accommodation_service.AccommodationService/SearchAccommodations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).SearchAccommodations(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_CreateAvailableDate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAvailableDateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).CreateAvailableDate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accommodation_service.AccommodationService/CreateAvailableDate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).CreateAvailableDate(ctx, req.(*CreateAvailableDateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_GetAvailableDateById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).GetAvailableDateById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accommodation_service.AccommodationService/GetAvailableDateById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).GetAvailableDateById(ctx, req.(*GetByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_UpdateAvailableDate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAvailableDateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).UpdateAvailableDate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accommodation_service.AccommodationService/UpdateAvailableDate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).UpdateAvailableDate(ctx, req.(*UpdateAvailableDateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_TimeSlotAvailableForAccommodation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TimeSlotAvailableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).TimeSlotAvailableForAccommodation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accommodation_service.AccommodationService/TimeSlotAvailableForAccommodation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).TimeSlotAvailableForAccommodation(ctx, req.(*TimeSlotAvailableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_GetAutomaticAcceptById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).GetAutomaticAcceptById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accommodation_service.AccommodationService/GetAutomaticAcceptById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).GetAutomaticAcceptById(ctx, req.(*GetByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AccommodationService_ServiceDesc is the grpc.ServiceDesc for AccommodationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccommodationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "accommodation_service.AccommodationService",
	HandlerType: (*AccommodationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAcc",
			Handler:    _AccommodationService_CreateAcc_Handler,
		},
		{
			MethodName: "GetAccommodationById",
			Handler:    _AccommodationService_GetAccommodationById_Handler,
		},
		{
			MethodName: "SearchAccommodations",
			Handler:    _AccommodationService_SearchAccommodations_Handler,
		},
		{
			MethodName: "CreateAvailableDate",
			Handler:    _AccommodationService_CreateAvailableDate_Handler,
		},
		{
			MethodName: "GetAvailableDateById",
			Handler:    _AccommodationService_GetAvailableDateById_Handler,
		},
		{
			MethodName: "UpdateAvailableDate",
			Handler:    _AccommodationService_UpdateAvailableDate_Handler,
		},
		{
			MethodName: "TimeSlotAvailableForAccommodation",
			Handler:    _AccommodationService_TimeSlotAvailableForAccommodation_Handler,
		},
		{
			MethodName: "GetAutomaticAcceptById",
			Handler:    _AccommodationService_GetAutomaticAcceptById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "accommodation_service/accommodation_service.proto",
}