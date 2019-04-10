// Code generated by protoc-gen-go. DO NOT EDIT.
// source: stateUpdate.proto

/*
Package stateUpdate is a generated protocol buffer package.

It is generated from these files:
	stateUpdate.proto

It has these top-level messages:
	ProjectRequest
	ProjectResponse
	SearchRequest
	SearchResponse
*/
package stateUpdate

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type ProjectRequest struct {
	Id     string `protobuf:"bytes,1,opt,name=Id" json:"Id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	Status string `protobuf:"bytes,3,opt,name=Status" json:"Status,omitempty"`
}

func (m *ProjectRequest) Reset()                    { *m = ProjectRequest{} }
func (m *ProjectRequest) String() string            { return proto.CompactTextString(m) }
func (*ProjectRequest) ProtoMessage()               {}
func (*ProjectRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ProjectRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ProjectRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ProjectRequest) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type ProjectResponse struct {
	Success bool   `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *ProjectResponse) Reset()                    { *m = ProjectResponse{} }
func (m *ProjectResponse) String() string            { return proto.CompactTextString(m) }
func (*ProjectResponse) ProtoMessage()               {}
func (*ProjectResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ProjectResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *ProjectResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type SearchRequest struct {
	Name string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
}

func (m *SearchRequest) Reset()                    { *m = SearchRequest{} }
func (m *SearchRequest) String() string            { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()               {}
func (*SearchRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SearchRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type SearchResponse struct {
	Find    bool   `protobuf:"varint,1,opt,name=Find" json:"Find,omitempty"`
	Id      string `protobuf:"bytes,2,opt,name=Id" json:"Id,omitempty"`
	Name    string `protobuf:"bytes,3,opt,name=Name" json:"Name,omitempty"`
	Status  string `protobuf:"bytes,4,opt,name=Status" json:"Status,omitempty"`
	Message string `protobuf:"bytes,5,opt,name=Message" json:"Message,omitempty"`
}

func (m *SearchResponse) Reset()                    { *m = SearchResponse{} }
func (m *SearchResponse) String() string            { return proto.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()               {}
func (*SearchResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SearchResponse) GetFind() bool {
	if m != nil {
		return m.Find
	}
	return false
}

func (m *SearchResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SearchResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SearchResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *SearchResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*ProjectRequest)(nil), "stateUpdate.ProjectRequest")
	proto.RegisterType((*ProjectResponse)(nil), "stateUpdate.ProjectResponse")
	proto.RegisterType((*SearchRequest)(nil), "stateUpdate.SearchRequest")
	proto.RegisterType((*SearchResponse)(nil), "stateUpdate.SearchResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Greeter service

type GreeterClient interface {
	// Método para criar projeto
	CreateProject(ctx context.Context, in *ProjectRequest, opts ...grpc.CallOption) (*ProjectResponse, error)
	// Método para buscar projetos
	GetProject(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) CreateProject(ctx context.Context, in *ProjectRequest, opts ...grpc.CallOption) (*ProjectResponse, error) {
	out := new(ProjectResponse)
	err := grpc.Invoke(ctx, "/stateUpdate.Greeter/CreateProject", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) GetProject(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := grpc.Invoke(ctx, "/stateUpdate.Greeter/GetProject", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Greeter service

type GreeterServer interface {
	// Método para criar projeto
	CreateProject(context.Context, *ProjectRequest) (*ProjectResponse, error)
	// Método para buscar projetos
	GetProject(context.Context, *SearchRequest) (*SearchResponse, error)
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_CreateProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).CreateProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stateUpdate.Greeter/CreateProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).CreateProject(ctx, req.(*ProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_GetProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).GetProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stateUpdate.Greeter/GetProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).GetProject(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stateUpdate.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProject",
			Handler:    _Greeter_CreateProject_Handler,
		},
		{
			MethodName: "GetProject",
			Handler:    _Greeter_GetProject_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stateUpdate.proto",
}

func init() { proto.RegisterFile("stateUpdate.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcf, 0x4e, 0xc3, 0x30,
	0x0c, 0xc6, 0x69, 0x57, 0x36, 0x30, 0x5a, 0x11, 0x3e, 0xa0, 0xaa, 0xe3, 0x80, 0xc2, 0x85, 0xd3,
	0x0e, 0xf0, 0x08, 0x08, 0xaa, 0x49, 0x03, 0xa1, 0x4e, 0x3c, 0x40, 0x68, 0x2d, 0xfe, 0x48, 0x5b,
	0x4a, 0xec, 0x5e, 0x78, 0x1a, 0x1e, 0x15, 0x2d, 0x4b, 0xa0, 0xd5, 0x7a, 0xcb, 0x67, 0x5b, 0x9f,
	0xbf, 0x9f, 0x03, 0x67, 0x2c, 0x5a, 0xe8, 0xa5, 0xa9, 0xb5, 0xd0, 0xbc, 0xb1, 0x46, 0x0c, 0x9e,
	0x74, 0x4a, 0x6a, 0x09, 0xe9, 0xb3, 0x35, 0x9f, 0x54, 0x49, 0x49, 0x5f, 0x2d, 0xb1, 0x60, 0x0a,
	0xf1, 0xa2, 0xce, 0xa2, 0xcb, 0xe8, 0xfa, 0xb8, 0x8c, 0x17, 0x35, 0x22, 0x24, 0x4f, 0x7a, 0x4d,
	0x59, 0xec, 0x2a, 0xee, 0x8d, 0xe7, 0x30, 0x5e, 0x89, 0x96, 0x96, 0xb3, 0x91, 0xab, 0x7a, 0xa5,
	0xee, 0xe1, 0xf4, 0xcf, 0x8d, 0x1b, 0xb3, 0x61, 0xc2, 0x0c, 0x26, 0xdc, 0x56, 0x15, 0x31, 0x3b,
	0xcf, 0xa3, 0x32, 0xc8, 0x6d, 0x67, 0x4d, 0xcc, 0xfa, 0x2d, 0x78, 0x07, 0xa9, 0xae, 0x60, 0xba,
	0x22, 0x6d, 0xab, 0xf7, 0x90, 0x29, 0x64, 0x88, 0xfe, 0x33, 0xa8, 0x6f, 0x48, 0xc3, 0x90, 0x5f,
	0x85, 0x90, 0x3c, 0x7c, 0x6c, 0x6a, 0xbf, 0xc7, 0xbd, 0x3d, 0x4d, 0xbc, 0x47, 0x33, 0x1a, 0xa4,
	0x49, 0xba, 0x34, 0xdb, 0x80, 0x8f, 0x3e, 0xe0, 0xe1, 0x2e, 0xa0, 0x97, 0x37, 0x3f, 0x11, 0x4c,
	0x0a, 0x4b, 0x24, 0x64, 0x71, 0x09, 0xd3, 0x3b, 0x4b, 0x5a, 0xc8, 0x93, 0xe3, 0x6c, 0xde, 0xbd,
	0x79, 0xff, 0xba, 0xf9, 0xc5, 0x70, 0x73, 0x47, 0xa0, 0x0e, 0xb0, 0x00, 0x28, 0x48, 0x82, 0x55,
	0xde, 0x9b, 0xee, 0xdd, 0x24, 0x9f, 0x0d, 0xf6, 0x82, 0xd1, 0xeb, 0xd8, 0x7d, 0xf6, 0xed, 0x6f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x88, 0xc9, 0xbc, 0xae, 0x01, 0x02, 0x00, 0x00,
}
