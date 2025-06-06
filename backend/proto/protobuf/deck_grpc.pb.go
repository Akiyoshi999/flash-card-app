// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.29.3
// source: deck.proto

package protobuf

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
	DeckService_CreateDeck_FullMethodName = "/deck.DeckService/CreateDeck"
	DeckService_GetDeck_FullMethodName    = "/deck.DeckService/GetDeck"
	DeckService_ListDecks_FullMethodName  = "/deck.DeckService/ListDecks"
	DeckService_UpdateDeck_FullMethodName = "/deck.DeckService/UpdateDeck"
	DeckService_DeleteDeck_FullMethodName = "/deck.DeckService/DeleteDeck"
)

// DeckServiceClient is the client API for DeckService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeckServiceClient interface {
	CreateDeck(ctx context.Context, in *CreateDeckRequest, opts ...grpc.CallOption) (*Deck, error)
	GetDeck(ctx context.Context, in *GetDeckRequest, opts ...grpc.CallOption) (*Deck, error)
	ListDecks(ctx context.Context, in *ListDecksRequest, opts ...grpc.CallOption) (*ListDecksResponse, error)
	UpdateDeck(ctx context.Context, in *UpdateDeckRequest, opts ...grpc.CallOption) (*Deck, error)
	DeleteDeck(ctx context.Context, in *DeleteDeckRequest, opts ...grpc.CallOption) (*DeleteDeckResponse, error)
}

type deckServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDeckServiceClient(cc grpc.ClientConnInterface) DeckServiceClient {
	return &deckServiceClient{cc}
}

func (c *deckServiceClient) CreateDeck(ctx context.Context, in *CreateDeckRequest, opts ...grpc.CallOption) (*Deck, error) {
	out := new(Deck)
	err := c.cc.Invoke(ctx, DeckService_CreateDeck_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deckServiceClient) GetDeck(ctx context.Context, in *GetDeckRequest, opts ...grpc.CallOption) (*Deck, error) {
	out := new(Deck)
	err := c.cc.Invoke(ctx, DeckService_GetDeck_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deckServiceClient) ListDecks(ctx context.Context, in *ListDecksRequest, opts ...grpc.CallOption) (*ListDecksResponse, error) {
	out := new(ListDecksResponse)
	err := c.cc.Invoke(ctx, DeckService_ListDecks_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deckServiceClient) UpdateDeck(ctx context.Context, in *UpdateDeckRequest, opts ...grpc.CallOption) (*Deck, error) {
	out := new(Deck)
	err := c.cc.Invoke(ctx, DeckService_UpdateDeck_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deckServiceClient) DeleteDeck(ctx context.Context, in *DeleteDeckRequest, opts ...grpc.CallOption) (*DeleteDeckResponse, error) {
	out := new(DeleteDeckResponse)
	err := c.cc.Invoke(ctx, DeckService_DeleteDeck_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeckServiceServer is the server API for DeckService service.
// All implementations must embed UnimplementedDeckServiceServer
// for forward compatibility
type DeckServiceServer interface {
	CreateDeck(context.Context, *CreateDeckRequest) (*Deck, error)
	GetDeck(context.Context, *GetDeckRequest) (*Deck, error)
	ListDecks(context.Context, *ListDecksRequest) (*ListDecksResponse, error)
	UpdateDeck(context.Context, *UpdateDeckRequest) (*Deck, error)
	DeleteDeck(context.Context, *DeleteDeckRequest) (*DeleteDeckResponse, error)
	mustEmbedUnimplementedDeckServiceServer()
}

// UnimplementedDeckServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDeckServiceServer struct {
}

func (UnimplementedDeckServiceServer) CreateDeck(context.Context, *CreateDeckRequest) (*Deck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDeck not implemented")
}
func (UnimplementedDeckServiceServer) GetDeck(context.Context, *GetDeckRequest) (*Deck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeck not implemented")
}
func (UnimplementedDeckServiceServer) ListDecks(context.Context, *ListDecksRequest) (*ListDecksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDecks not implemented")
}
func (UnimplementedDeckServiceServer) UpdateDeck(context.Context, *UpdateDeckRequest) (*Deck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDeck not implemented")
}
func (UnimplementedDeckServiceServer) DeleteDeck(context.Context, *DeleteDeckRequest) (*DeleteDeckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDeck not implemented")
}
func (UnimplementedDeckServiceServer) mustEmbedUnimplementedDeckServiceServer() {}

// UnsafeDeckServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeckServiceServer will
// result in compilation errors.
type UnsafeDeckServiceServer interface {
	mustEmbedUnimplementedDeckServiceServer()
}

func RegisterDeckServiceServer(s grpc.ServiceRegistrar, srv DeckServiceServer) {
	s.RegisterService(&DeckService_ServiceDesc, srv)
}

func _DeckService_CreateDeck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDeckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeckServiceServer).CreateDeck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeckService_CreateDeck_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeckServiceServer).CreateDeck(ctx, req.(*CreateDeckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeckService_GetDeck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeckServiceServer).GetDeck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeckService_GetDeck_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeckServiceServer).GetDeck(ctx, req.(*GetDeckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeckService_ListDecks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDecksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeckServiceServer).ListDecks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeckService_ListDecks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeckServiceServer).ListDecks(ctx, req.(*ListDecksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeckService_UpdateDeck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDeckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeckServiceServer).UpdateDeck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeckService_UpdateDeck_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeckServiceServer).UpdateDeck(ctx, req.(*UpdateDeckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeckService_DeleteDeck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDeckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeckServiceServer).DeleteDeck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeckService_DeleteDeck_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeckServiceServer).DeleteDeck(ctx, req.(*DeleteDeckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DeckService_ServiceDesc is the grpc.ServiceDesc for DeckService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DeckService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "deck.DeckService",
	HandlerType: (*DeckServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDeck",
			Handler:    _DeckService_CreateDeck_Handler,
		},
		{
			MethodName: "GetDeck",
			Handler:    _DeckService_GetDeck_Handler,
		},
		{
			MethodName: "ListDecks",
			Handler:    _DeckService_ListDecks_Handler,
		},
		{
			MethodName: "UpdateDeck",
			Handler:    _DeckService_UpdateDeck_Handler,
		},
		{
			MethodName: "DeleteDeck",
			Handler:    _DeckService_DeleteDeck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "deck.proto",
}
