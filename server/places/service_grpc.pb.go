// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: service.proto

package places

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
	Places_UploadPlaceInfo_FullMethodName = "/places.Places/uploadPlaceInfo"
	Places_UpdatePlace_FullMethodName     = "/places.Places/updatePlace"
	Places_GetPlaceInfo_FullMethodName    = "/places.Places/getPlaceInfo"
	Places_SearchPlaces_FullMethodName    = "/places.Places/searchPlaces"
	Places_FilterPlaces_FullMethodName    = "/places.Places/filterPlaces"
	Places_RemovePlaces_FullMethodName    = "/places.Places/removePlaces"
)

// PlacesClient is the client API for Places service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PlacesClient interface {
	UploadPlaceInfo(ctx context.Context, in *Place, opts ...grpc.CallOption) (*Place, error)
	UpdatePlace(ctx context.Context, in *UpdatePlace, opts ...grpc.CallOption) (*Place, error)
	GetPlaceInfo(ctx context.Context, in *PlaceName, opts ...grpc.CallOption) (*Place, error)
	SearchPlaces(ctx context.Context, in *PlaceName, opts ...grpc.CallOption) (*PlaceList, error)
	FilterPlaces(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*PlaceList, error)
	RemovePlaces(ctx context.Context, in *PlaceName, opts ...grpc.CallOption) (*Empty, error)
}

type placesClient struct {
	cc grpc.ClientConnInterface
}

func NewPlacesClient(cc grpc.ClientConnInterface) PlacesClient {
	return &placesClient{cc}
}

func (c *placesClient) UploadPlaceInfo(ctx context.Context, in *Place, opts ...grpc.CallOption) (*Place, error) {
	out := new(Place)
	err := c.cc.Invoke(ctx, Places_UploadPlaceInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *placesClient) UpdatePlace(ctx context.Context, in *UpdatePlace, opts ...grpc.CallOption) (*Place, error) {
	out := new(Place)
	err := c.cc.Invoke(ctx, Places_UpdatePlace_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *placesClient) GetPlaceInfo(ctx context.Context, in *PlaceName, opts ...grpc.CallOption) (*Place, error) {
	out := new(Place)
	err := c.cc.Invoke(ctx, Places_GetPlaceInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *placesClient) SearchPlaces(ctx context.Context, in *PlaceName, opts ...grpc.CallOption) (*PlaceList, error) {
	out := new(PlaceList)
	err := c.cc.Invoke(ctx, Places_SearchPlaces_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *placesClient) FilterPlaces(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*PlaceList, error) {
	out := new(PlaceList)
	err := c.cc.Invoke(ctx, Places_FilterPlaces_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *placesClient) RemovePlaces(ctx context.Context, in *PlaceName, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Places_RemovePlaces_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlacesServer is the server API for Places service.
// All implementations must embed UnimplementedPlacesServer
// for forward compatibility
type PlacesServer interface {
	UploadPlaceInfo(context.Context, *Place) (*Place, error)
	UpdatePlace(context.Context, *UpdatePlace) (*Place, error)
	GetPlaceInfo(context.Context, *PlaceName) (*Place, error)
	SearchPlaces(context.Context, *PlaceName) (*PlaceList, error)
	FilterPlaces(context.Context, *Filter) (*PlaceList, error)
	RemovePlaces(context.Context, *PlaceName) (*Empty, error)
	mustEmbedUnimplementedPlacesServer()
}

// UnimplementedPlacesServer must be embedded to have forward compatible implementations.
type UnimplementedPlacesServer struct {
}

func (UnimplementedPlacesServer) UploadPlaceInfo(context.Context, *Place) (*Place, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadPlaceInfo not implemented")
}
func (UnimplementedPlacesServer) UpdatePlace(context.Context, *UpdatePlace) (*Place, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePlace not implemented")
}
func (UnimplementedPlacesServer) GetPlaceInfo(context.Context, *PlaceName) (*Place, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlaceInfo not implemented")
}
func (UnimplementedPlacesServer) SearchPlaces(context.Context, *PlaceName) (*PlaceList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchPlaces not implemented")
}
func (UnimplementedPlacesServer) FilterPlaces(context.Context, *Filter) (*PlaceList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FilterPlaces not implemented")
}
func (UnimplementedPlacesServer) RemovePlaces(context.Context, *PlaceName) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemovePlaces not implemented")
}
func (UnimplementedPlacesServer) mustEmbedUnimplementedPlacesServer() {}

// UnsafePlacesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PlacesServer will
// result in compilation errors.
type UnsafePlacesServer interface {
	mustEmbedUnimplementedPlacesServer()
}

func RegisterPlacesServer(s grpc.ServiceRegistrar, srv PlacesServer) {
	s.RegisterService(&Places_ServiceDesc, srv)
}

func _Places_UploadPlaceInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Place)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlacesServer).UploadPlaceInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Places_UploadPlaceInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlacesServer).UploadPlaceInfo(ctx, req.(*Place))
	}
	return interceptor(ctx, in, info, handler)
}

func _Places_UpdatePlace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePlace)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlacesServer).UpdatePlace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Places_UpdatePlace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlacesServer).UpdatePlace(ctx, req.(*UpdatePlace))
	}
	return interceptor(ctx, in, info, handler)
}

func _Places_GetPlaceInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlaceName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlacesServer).GetPlaceInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Places_GetPlaceInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlacesServer).GetPlaceInfo(ctx, req.(*PlaceName))
	}
	return interceptor(ctx, in, info, handler)
}

func _Places_SearchPlaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlaceName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlacesServer).SearchPlaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Places_SearchPlaces_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlacesServer).SearchPlaces(ctx, req.(*PlaceName))
	}
	return interceptor(ctx, in, info, handler)
}

func _Places_FilterPlaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Filter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlacesServer).FilterPlaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Places_FilterPlaces_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlacesServer).FilterPlaces(ctx, req.(*Filter))
	}
	return interceptor(ctx, in, info, handler)
}

func _Places_RemovePlaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlaceName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlacesServer).RemovePlaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Places_RemovePlaces_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlacesServer).RemovePlaces(ctx, req.(*PlaceName))
	}
	return interceptor(ctx, in, info, handler)
}

// Places_ServiceDesc is the grpc.ServiceDesc for Places service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Places_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "places.Places",
	HandlerType: (*PlacesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "uploadPlaceInfo",
			Handler:    _Places_UploadPlaceInfo_Handler,
		},
		{
			MethodName: "updatePlace",
			Handler:    _Places_UpdatePlace_Handler,
		},
		{
			MethodName: "getPlaceInfo",
			Handler:    _Places_GetPlaceInfo_Handler,
		},
		{
			MethodName: "searchPlaces",
			Handler:    _Places_SearchPlaces_Handler,
		},
		{
			MethodName: "filterPlaces",
			Handler:    _Places_FilterPlaces_Handler,
		},
		{
			MethodName: "removePlaces",
			Handler:    _Places_RemovePlaces_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
