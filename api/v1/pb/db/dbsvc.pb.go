// Code generated by protoc-gen-go. DO NOT EDIT.
// source: db/dbsvc.proto

/*
Package db is a generated protocol buffer package.

It is generated from these files:
	db/dbsvc.proto

It has these top-level messages:
	Document
	GetRequest
	GetReply
	UpdateRequest
	UpdateReply
	RemoveRequest
	RemoveReply
	AddRequest
	AddReply
	ServiceStatusRequest
	ServiceStatusReply
*/
package db

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

type Document struct {
	Content   string `protobuf:"bytes,1,opt,name=content" json:"content,omitempty"`
	Title     string `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	Author    string `protobuf:"bytes,3,opt,name=author" json:"author,omitempty"`
	Topic     string `protobuf:"bytes,4,opt,name=topic" json:"topic,omitempty"`
	Watermark string `protobuf:"bytes,5,opt,name=watermark" json:"watermark,omitempty"`
}

func (m *Document) Reset()                    { *m = Document{} }
func (m *Document) String() string            { return proto.CompactTextString(m) }
func (*Document) ProtoMessage()               {}
func (*Document) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Document) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Document) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Document) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *Document) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *Document) GetWatermark() string {
	if m != nil {
		return m.Watermark
	}
	return ""
}

type GetRequest struct {
	Filters []*GetRequest_Filters `protobuf:"bytes,1,rep,name=filters" json:"filters,omitempty"`
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetRequest) GetFilters() []*GetRequest_Filters {
	if m != nil {
		return m.Filters
	}
	return nil
}

type GetRequest_Filters struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *GetRequest_Filters) Reset()                    { *m = GetRequest_Filters{} }
func (m *GetRequest_Filters) String() string            { return proto.CompactTextString(m) }
func (*GetRequest_Filters) ProtoMessage()               {}
func (*GetRequest_Filters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *GetRequest_Filters) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *GetRequest_Filters) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type GetReply struct {
	Documents []*Document `protobuf:"bytes,1,rep,name=documents" json:"documents,omitempty"`
	Err       string      `protobuf:"bytes,2,opt,name=Err,json=err" json:"Err,omitempty"`
}

func (m *GetReply) Reset()                    { *m = GetReply{} }
func (m *GetReply) String() string            { return proto.CompactTextString(m) }
func (*GetReply) ProtoMessage()               {}
func (*GetReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetReply) GetDocuments() []*Document {
	if m != nil {
		return m.Documents
	}
	return nil
}

func (m *GetReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type UpdateRequest struct {
	TicketID string    `protobuf:"bytes,1,opt,name=ticketID" json:"ticketID,omitempty"`
	Document *Document `protobuf:"bytes,2,opt,name=document" json:"document,omitempty"`
}

func (m *UpdateRequest) Reset()                    { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()               {}
func (*UpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *UpdateRequest) GetTicketID() string {
	if m != nil {
		return m.TicketID
	}
	return ""
}

func (m *UpdateRequest) GetDocument() *Document {
	if m != nil {
		return m.Document
	}
	return nil
}

type UpdateReply struct {
	Code int64  `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Err  string `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
}

func (m *UpdateReply) Reset()                    { *m = UpdateReply{} }
func (m *UpdateReply) String() string            { return proto.CompactTextString(m) }
func (*UpdateReply) ProtoMessage()               {}
func (*UpdateReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UpdateReply) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *UpdateReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type RemoveRequest struct {
	TicketID string `protobuf:"bytes,1,opt,name=ticketID" json:"ticketID,omitempty"`
}

func (m *RemoveRequest) Reset()                    { *m = RemoveRequest{} }
func (m *RemoveRequest) String() string            { return proto.CompactTextString(m) }
func (*RemoveRequest) ProtoMessage()               {}
func (*RemoveRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *RemoveRequest) GetTicketID() string {
	if m != nil {
		return m.TicketID
	}
	return ""
}

type RemoveReply struct {
	Code int64  `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Err  string `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
}

func (m *RemoveReply) Reset()                    { *m = RemoveReply{} }
func (m *RemoveReply) String() string            { return proto.CompactTextString(m) }
func (*RemoveReply) ProtoMessage()               {}
func (*RemoveReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *RemoveReply) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RemoveReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type AddRequest struct {
	Document *Document `protobuf:"bytes,1,opt,name=document" json:"document,omitempty"`
}

func (m *AddRequest) Reset()                    { *m = AddRequest{} }
func (m *AddRequest) String() string            { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()               {}
func (*AddRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *AddRequest) GetDocument() *Document {
	if m != nil {
		return m.Document
	}
	return nil
}

type AddReply struct {
	TicketID string `protobuf:"bytes,1,opt,name=ticketID" json:"ticketID,omitempty"`
	Err      string `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
}

func (m *AddReply) Reset()                    { *m = AddReply{} }
func (m *AddReply) String() string            { return proto.CompactTextString(m) }
func (*AddReply) ProtoMessage()               {}
func (*AddReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *AddReply) GetTicketID() string {
	if m != nil {
		return m.TicketID
	}
	return ""
}

func (m *AddReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type ServiceStatusRequest struct {
}

func (m *ServiceStatusRequest) Reset()                    { *m = ServiceStatusRequest{} }
func (m *ServiceStatusRequest) String() string            { return proto.CompactTextString(m) }
func (*ServiceStatusRequest) ProtoMessage()               {}
func (*ServiceStatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type ServiceStatusReply struct {
	Code int64  `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Err  string `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
}

func (m *ServiceStatusReply) Reset()                    { *m = ServiceStatusReply{} }
func (m *ServiceStatusReply) String() string            { return proto.CompactTextString(m) }
func (*ServiceStatusReply) ProtoMessage()               {}
func (*ServiceStatusReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *ServiceStatusReply) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ServiceStatusReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*Document)(nil), "db.Document")
	proto.RegisterType((*GetRequest)(nil), "db.GetRequest")
	proto.RegisterType((*GetRequest_Filters)(nil), "db.GetRequest.Filters")
	proto.RegisterType((*GetReply)(nil), "db.GetReply")
	proto.RegisterType((*UpdateRequest)(nil), "db.UpdateRequest")
	proto.RegisterType((*UpdateReply)(nil), "db.UpdateReply")
	proto.RegisterType((*RemoveRequest)(nil), "db.RemoveRequest")
	proto.RegisterType((*RemoveReply)(nil), "db.RemoveReply")
	proto.RegisterType((*AddRequest)(nil), "db.AddRequest")
	proto.RegisterType((*AddReply)(nil), "db.AddReply")
	proto.RegisterType((*ServiceStatusRequest)(nil), "db.ServiceStatusRequest")
	proto.RegisterType((*ServiceStatusReply)(nil), "db.ServiceStatusReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Database service

type DatabaseClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error)
	Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveReply, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateReply, error)
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddReply, error)
	ServiceStatus(ctx context.Context, in *ServiceStatusRequest, opts ...grpc.CallOption) (*ServiceStatusReply, error)
}

type databaseClient struct {
	cc *grpc.ClientConn
}

func NewDatabaseClient(cc *grpc.ClientConn) DatabaseClient {
	return &databaseClient{cc}
}

func (c *databaseClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error) {
	out := new(GetReply)
	err := grpc.Invoke(ctx, "/db.database/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseClient) Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveReply, error) {
	out := new(RemoveReply)
	err := grpc.Invoke(ctx, "/db.database/Remove", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateReply, error) {
	out := new(UpdateReply)
	err := grpc.Invoke(ctx, "/db.database/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddReply, error) {
	out := new(AddReply)
	err := grpc.Invoke(ctx, "/db.database/Add", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseClient) ServiceStatus(ctx context.Context, in *ServiceStatusRequest, opts ...grpc.CallOption) (*ServiceStatusReply, error) {
	out := new(ServiceStatusReply)
	err := grpc.Invoke(ctx, "/db.database/ServiceStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Database service

type DatabaseServer interface {
	Get(context.Context, *GetRequest) (*GetReply, error)
	Remove(context.Context, *RemoveRequest) (*RemoveReply, error)
	Update(context.Context, *UpdateRequest) (*UpdateReply, error)
	Add(context.Context, *AddRequest) (*AddReply, error)
	ServiceStatus(context.Context, *ServiceStatusRequest) (*ServiceStatusReply, error)
}

func RegisterDatabaseServer(s *grpc.Server, srv DatabaseServer) {
	s.RegisterService(&_Database_serviceDesc, srv)
}

func _Database_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.database/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Database_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.database/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).Remove(ctx, req.(*RemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Database_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.database/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Database_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.database/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Database_ServiceStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).ServiceStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.database/ServiceStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).ServiceStatus(ctx, req.(*ServiceStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Database_serviceDesc = grpc.ServiceDesc{
	ServiceName: "db.database",
	HandlerType: (*DatabaseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Database_Get_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _Database_Remove_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Database_Update_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _Database_Add_Handler,
		},
		{
			MethodName: "ServiceStatus",
			Handler:    _Database_ServiceStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "db/dbsvc.proto",
}

func init() { proto.RegisterFile("db/dbsvc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 452 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xd1, 0x6a, 0xdb, 0x40,
	0x10, 0xb4, 0xa2, 0xc4, 0x96, 0x37, 0x71, 0xda, 0x2e, 0xc1, 0x08, 0xd3, 0x87, 0x70, 0x50, 0x30,
	0x2d, 0xa8, 0x6d, 0x02, 0xa5, 0xf4, 0x2d, 0x34, 0x6d, 0xda, 0x57, 0x85, 0x7c, 0xc0, 0x49, 0xb7,
	0xa5, 0xc2, 0xb2, 0xa5, 0x9c, 0x56, 0x2e, 0xfe, 0x81, 0xfe, 0x75, 0xa1, 0xdc, 0x49, 0x27, 0x5b,
	0x49, 0x08, 0x7e, 0xbb, 0xdd, 0x1d, 0xcd, 0xcc, 0xce, 0x1d, 0x82, 0x53, 0x95, 0xbc, 0x57, 0x49,
	0xb5, 0x4e, 0xa3, 0x52, 0x17, 0x5c, 0xe0, 0x81, 0x4a, 0xc4, 0x5f, 0x0f, 0x82, 0xeb, 0x22, 0xad,
	0x97, 0xb4, 0x62, 0x0c, 0x61, 0x94, 0x16, 0x2b, 0xa6, 0x15, 0x87, 0xde, 0xb9, 0x37, 0x1f, 0xc7,
	0xae, 0xc4, 0x33, 0x38, 0xe2, 0x8c, 0x73, 0x0a, 0x0f, 0x6c, 0xbf, 0x29, 0x70, 0x0a, 0x43, 0x59,
	0xf3, 0xef, 0x42, 0x87, 0xbe, 0x6d, 0xb7, 0x95, 0x45, 0x17, 0x65, 0x96, 0x86, 0x87, 0x2d, 0xda,
	0x14, 0xf8, 0x1a, 0xc6, 0x7f, 0x24, 0x93, 0x5e, 0x4a, 0xbd, 0x08, 0x8f, 0xec, 0x64, 0xdb, 0x10,
	0xf7, 0x00, 0x37, 0xc4, 0x31, 0xdd, 0xd7, 0x54, 0x31, 0x7e, 0x80, 0xd1, 0xaf, 0x2c, 0x67, 0xd2,
	0x55, 0xe8, 0x9d, 0xfb, 0xf3, 0xe3, 0x8b, 0x69, 0xa4, 0x92, 0x68, 0x0b, 0x88, 0xbe, 0x37, 0xd3,
	0xd8, 0xc1, 0x66, 0x1f, 0x61, 0xd4, 0xf6, 0xf0, 0x25, 0xf8, 0x0b, 0xda, 0xb4, 0x2b, 0x98, 0xa3,
	0x31, 0xb4, 0x96, 0x79, 0xdd, 0xd9, 0xb7, 0x85, 0xf8, 0x01, 0x81, 0x65, 0x2c, 0xf3, 0x0d, 0xbe,
	0x85, 0xb1, 0x6a, 0x63, 0x70, 0x92, 0x27, 0x46, 0xd2, 0x65, 0x13, 0x6f, 0xc7, 0x86, 0xff, 0x9b,
	0xd6, 0x2d, 0x97, 0x4f, 0x5a, 0x8b, 0x3b, 0x98, 0xdc, 0x95, 0x4a, 0x32, 0x39, 0xff, 0x33, 0x08,
	0x38, 0x4b, 0x17, 0xc4, 0x3f, 0xaf, 0x5b, 0x1f, 0x5d, 0x8d, 0x73, 0x08, 0x1c, 0x97, 0xe5, 0x78,
	0xa8, 0xd4, 0x4d, 0xc5, 0x25, 0x1c, 0x3b, 0x5a, 0xe3, 0x11, 0xe1, 0x30, 0x2d, 0x14, 0x59, 0x42,
	0x3f, 0xb6, 0x67, 0xe3, 0x85, 0xfa, 0x5e, 0xde, 0xc1, 0x24, 0xa6, 0x65, 0xb1, 0xde, 0xc7, 0x8b,
	0x51, 0x70, 0xe0, 0xfd, 0x15, 0x3e, 0x01, 0x5c, 0x29, 0xe5, 0xe8, 0x77, 0xd7, 0xf1, 0x9e, 0x5d,
	0xe7, 0x33, 0x04, 0xf6, 0x3b, 0xa3, 0xf4, 0x5c, 0x40, 0x8f, 0x15, 0xa7, 0x70, 0x76, 0x4b, 0x7a,
	0x9d, 0xa5, 0x74, 0xcb, 0x92, 0xeb, 0xaa, 0xd5, 0x16, 0x5f, 0x00, 0x1f, 0xf4, 0xf7, 0xde, 0xe2,
	0xe2, 0x9f, 0x07, 0x81, 0x92, 0x2c, 0x13, 0x59, 0x11, 0xbe, 0x01, 0xff, 0x86, 0x18, 0x4f, 0xfb,
	0xaf, 0x6c, 0x76, 0xd2, 0xd5, 0x65, 0xbe, 0x11, 0x03, 0x8c, 0x60, 0xd8, 0xc4, 0x85, 0xaf, 0xcc,
	0xa4, 0x97, 0xf3, 0xec, 0xc5, 0x6e, 0xab, 0xc3, 0x37, 0x17, 0xd8, 0xe0, 0x7b, 0x6f, 0xa4, 0xc1,
	0xef, 0xdc, 0xaf, 0x18, 0x18, 0x1b, 0x57, 0x4a, 0x35, 0x36, 0xb6, 0x11, 0x37, 0x36, 0x5c, 0x74,
	0x62, 0x80, 0x5f, 0x61, 0xd2, 0x5b, 0x1b, 0x43, 0x03, 0x78, 0x2a, 0xa1, 0xd9, 0xf4, 0x89, 0x89,
	0x25, 0x49, 0x86, 0xf6, 0x27, 0x70, 0xf9, 0x3f, 0x00, 0x00, 0xff, 0xff, 0x67, 0x73, 0x72, 0xc6,
	0x16, 0x04, 0x00, 0x00,
}