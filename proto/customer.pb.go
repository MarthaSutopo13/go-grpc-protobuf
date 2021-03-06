// Code generated by protoc-gen-go. DO NOT EDIT.
// source: customer.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Start of Customer Proto
type Customer struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Firstname            string   `protobuf:"bytes,2,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname             string   `protobuf:"bytes,3,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Customer) Reset()         { *m = Customer{} }
func (m *Customer) String() string { return proto.CompactTextString(m) }
func (*Customer) ProtoMessage()    {}
func (*Customer) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{0}
}

func (m *Customer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Customer.Unmarshal(m, b)
}
func (m *Customer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Customer.Marshal(b, m, deterministic)
}
func (m *Customer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Customer.Merge(m, src)
}
func (m *Customer) XXX_Size() int {
	return xxx_messageInfo_Customer.Size(m)
}
func (m *Customer) XXX_DiscardUnknown() {
	xxx_messageInfo_Customer.DiscardUnknown(m)
}

var xxx_messageInfo_Customer proto.InternalMessageInfo

func (m *Customer) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Customer) GetFirstname() string {
	if m != nil {
		return m.Firstname
	}
	return ""
}

func (m *Customer) GetLastname() string {
	if m != nil {
		return m.Lastname
	}
	return ""
}

func (m *Customer) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type CreateCustomerRequest struct {
	Customer             *Customer `protobuf:"bytes,1,opt,name=customer,proto3" json:"customer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CreateCustomerRequest) Reset()         { *m = CreateCustomerRequest{} }
func (m *CreateCustomerRequest) String() string { return proto.CompactTextString(m) }
func (*CreateCustomerRequest) ProtoMessage()    {}
func (*CreateCustomerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{1}
}

func (m *CreateCustomerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCustomerRequest.Unmarshal(m, b)
}
func (m *CreateCustomerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCustomerRequest.Marshal(b, m, deterministic)
}
func (m *CreateCustomerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCustomerRequest.Merge(m, src)
}
func (m *CreateCustomerRequest) XXX_Size() int {
	return xxx_messageInfo_CreateCustomerRequest.Size(m)
}
func (m *CreateCustomerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCustomerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCustomerRequest proto.InternalMessageInfo

func (m *CreateCustomerRequest) GetCustomer() *Customer {
	if m != nil {
		return m.Customer
	}
	return nil
}

type CreateCustomerResponse struct {
	Code                 string                               `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Description          string                               `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Data                 *CreateCustomerResponseCreatecusdata `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Customer             *Customer                            `protobuf:"bytes,4,opt,name=customer,proto3" json:"customer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *CreateCustomerResponse) Reset()         { *m = CreateCustomerResponse{} }
func (m *CreateCustomerResponse) String() string { return proto.CompactTextString(m) }
func (*CreateCustomerResponse) ProtoMessage()    {}
func (*CreateCustomerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{2}
}

func (m *CreateCustomerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCustomerResponse.Unmarshal(m, b)
}
func (m *CreateCustomerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCustomerResponse.Marshal(b, m, deterministic)
}
func (m *CreateCustomerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCustomerResponse.Merge(m, src)
}
func (m *CreateCustomerResponse) XXX_Size() int {
	return xxx_messageInfo_CreateCustomerResponse.Size(m)
}
func (m *CreateCustomerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCustomerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCustomerResponse proto.InternalMessageInfo

func (m *CreateCustomerResponse) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *CreateCustomerResponse) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CreateCustomerResponse) GetData() *CreateCustomerResponseCreatecusdata {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *CreateCustomerResponse) GetCustomer() *Customer {
	if m != nil {
		return m.Customer
	}
	return nil
}

type CreateCustomerResponseCreatecusdata struct {
	Cus                  string   `protobuf:"bytes,1,opt,name=cus,proto3" json:"cus,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateCustomerResponseCreatecusdata) Reset()         { *m = CreateCustomerResponseCreatecusdata{} }
func (m *CreateCustomerResponseCreatecusdata) String() string { return proto.CompactTextString(m) }
func (*CreateCustomerResponseCreatecusdata) ProtoMessage()    {}
func (*CreateCustomerResponseCreatecusdata) Descriptor() ([]byte, []int) {
	return fileDescriptor_9efa92dae3d6ec46, []int{2, 0}
}

func (m *CreateCustomerResponseCreatecusdata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCustomerResponseCreatecusdata.Unmarshal(m, b)
}
func (m *CreateCustomerResponseCreatecusdata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCustomerResponseCreatecusdata.Marshal(b, m, deterministic)
}
func (m *CreateCustomerResponseCreatecusdata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCustomerResponseCreatecusdata.Merge(m, src)
}
func (m *CreateCustomerResponseCreatecusdata) XXX_Size() int {
	return xxx_messageInfo_CreateCustomerResponseCreatecusdata.Size(m)
}
func (m *CreateCustomerResponseCreatecusdata) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCustomerResponseCreatecusdata.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCustomerResponseCreatecusdata proto.InternalMessageInfo

func (m *CreateCustomerResponseCreatecusdata) GetCus() string {
	if m != nil {
		return m.Cus
	}
	return ""
}

func init() {
	proto.RegisterType((*Customer)(nil), "proto.Customer")
	proto.RegisterType((*CreateCustomerRequest)(nil), "proto.CreateCustomerRequest")
	proto.RegisterType((*CreateCustomerResponse)(nil), "proto.CreateCustomerResponse")
	proto.RegisterType((*CreateCustomerResponseCreatecusdata)(nil), "proto.CreateCustomerResponse.createcusdata")
}

func init() { proto.RegisterFile("customer.proto", fileDescriptor_9efa92dae3d6ec46) }

var fileDescriptor_9efa92dae3d6ec46 = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xcf, 0x4e, 0x83, 0x40,
	0x10, 0xc6, 0x53, 0x4a, 0x0d, 0x0c, 0x11, 0xcd, 0x44, 0x0d, 0x21, 0x35, 0xa9, 0x9c, 0x4c, 0x9a,
	0x70, 0xc0, 0x07, 0xf0, 0x50, 0xaf, 0x5e, 0x38, 0x7a, 0x5b, 0x97, 0x31, 0x59, 0x53, 0x58, 0xdc,
	0x5d, 0x9e, 0xd7, 0x57, 0x31, 0x9d, 0x2e, 0x28, 0xc6, 0x3f, 0x27, 0x76, 0xbe, 0x6f, 0x66, 0xbe,
	0xdf, 0x00, 0xa9, 0x1c, 0xac, 0xd3, 0x2d, 0x99, 0xb2, 0x37, 0xda, 0x69, 0x5c, 0xf1, 0xa7, 0x78,
	0x85, 0x68, 0xe7, 0x0d, 0x4c, 0x21, 0x50, 0x4d, 0xb6, 0xd8, 0x2c, 0x6e, 0xe3, 0x3a, 0x50, 0x0d,
	0xae, 0x21, 0x7e, 0x51, 0xc6, 0xba, 0x4e, 0xb4, 0x94, 0x05, 0x2c, 0x7f, 0x0a, 0x98, 0x43, 0xb4,
	0x17, 0xde, 0x5c, 0xb2, 0x39, 0xd5, 0x78, 0x01, 0x2b, 0x6a, 0x85, 0xda, 0x67, 0x21, 0x1b, 0xc7,
	0xa2, 0x78, 0x80, 0xcb, 0x9d, 0x21, 0xe1, 0x68, 0x4c, 0xac, 0xe9, 0x6d, 0x20, 0xeb, 0x70, 0x0b,
	0xd1, 0x48, 0xc7, 0xf1, 0x49, 0x75, 0x76, 0xa4, 0x2c, 0xa7, 0xce, 0xa9, 0xa1, 0x78, 0x5f, 0xc0,
	0xd5, 0xf7, 0x35, 0xb6, 0xd7, 0x9d, 0x25, 0x44, 0x08, 0xa5, 0x6e, 0xc8, 0x9f, 0xc0, 0x6f, 0xdc,
	0x40, 0xd2, 0x90, 0x95, 0x46, 0xf5, 0x4e, 0xe9, 0xce, 0x9f, 0xf1, 0x55, 0xc2, 0x7b, 0x08, 0x1b,
	0xe1, 0x04, 0x1f, 0x91, 0x54, 0xdb, 0x31, 0xf9, 0xc7, 0x88, 0x52, 0xb2, 0x2c, 0x07, 0x7b, 0x18,
	0xa9, 0x79, 0x70, 0x86, 0x1f, 0xfe, 0x83, 0x9f, 0xdf, 0xc0, 0xe9, 0x6c, 0x07, 0x9e, 0xc3, 0x52,
	0x0e, 0xd6, 0x33, 0x1f, 0x9e, 0xd5, 0x13, 0xc4, 0xe3, 0xa0, 0xc5, 0x47, 0x48, 0xe7, 0x28, 0xb8,
	0xfe, 0x85, 0x90, 0xff, 0x65, 0x7e, 0xfd, 0x27, 0xff, 0xf3, 0x09, 0xbb, 0x77, 0x1f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x1f, 0x21, 0x6d, 0x8c, 0x0f, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CustomersClient is the client API for Customers service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CustomersClient interface {
	CreateCustomer(ctx context.Context, in *CreateCustomerRequest, opts ...grpc.CallOption) (*CreateCustomerResponse, error)
}

type customersClient struct {
	cc *grpc.ClientConn
}

func NewCustomersClient(cc *grpc.ClientConn) CustomersClient {
	return &customersClient{cc}
}

func (c *customersClient) CreateCustomer(ctx context.Context, in *CreateCustomerRequest, opts ...grpc.CallOption) (*CreateCustomerResponse, error) {
	out := new(CreateCustomerResponse)
	err := c.cc.Invoke(ctx, "/proto.Customers/CreateCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomersServer is the server API for Customers service.
type CustomersServer interface {
	CreateCustomer(context.Context, *CreateCustomerRequest) (*CreateCustomerResponse, error)
}

// UnimplementedCustomersServer can be embedded to have forward compatible implementations.
type UnimplementedCustomersServer struct {
}

func (*UnimplementedCustomersServer) CreateCustomer(ctx context.Context, req *CreateCustomerRequest) (*CreateCustomerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCustomer not implemented")
}

func RegisterCustomersServer(s *grpc.Server, srv CustomersServer) {
	s.RegisterService(&_Customers_serviceDesc, srv)
}

func _Customers_CreateCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomersServer).CreateCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Customers/CreateCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomersServer).CreateCustomer(ctx, req.(*CreateCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Customers_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Customers",
	HandlerType: (*CustomersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCustomer",
			Handler:    _Customers_CreateCustomer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "customer.proto",
}
