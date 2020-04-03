// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/account/account.proto

package _go

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
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

type User struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	PasswordHash         []byte   `protobuf:"bytes,3,opt,name=password_hash,json=passwordHash,proto3" json:"password_hash,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	CompanyId            uint64   `protobuf:"varint,5,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d492a0187472a3b, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPasswordHash() []byte {
	if m != nil {
		return m.PasswordHash
	}
	return nil
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetCompanyId() uint64 {
	if m != nil {
		return m.CompanyId
	}
	return 0
}

type CreateUserRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             []byte   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d492a0187472a3b, []int{1}
}

func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (m *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(m, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *CreateUserRequest) GetPassword() []byte {
	if m != nil {
		return m.Password
	}
	return nil
}

type CreateUserResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserResponse) Reset()         { *m = CreateUserResponse{} }
func (m *CreateUserResponse) String() string { return proto.CompactTextString(m) }
func (*CreateUserResponse) ProtoMessage()    {}
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d492a0187472a3b, []int{2}
}

func (m *CreateUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserResponse.Unmarshal(m, b)
}
func (m *CreateUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserResponse.Marshal(b, m, deterministic)
}
func (m *CreateUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserResponse.Merge(m, src)
}
func (m *CreateUserResponse) XXX_Size() int {
	return xxx_messageInfo_CreateUserResponse.Size(m)
}
func (m *CreateUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserResponse proto.InternalMessageInfo

func (m *CreateUserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "account.User")
	proto.RegisterType((*CreateUserRequest)(nil), "account.CreateUserRequest")
	proto.RegisterType((*CreateUserResponse)(nil), "account.CreateUserResponse")
}

func init() { proto.RegisterFile("proto/account/account.proto", fileDescriptor_5d492a0187472a3b) }

var fileDescriptor_5d492a0187472a3b = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x5d, 0xeb, 0xd3, 0x30,
	0x14, 0xc6, 0xc9, 0xec, 0x36, 0x77, 0xb6, 0x09, 0x06, 0x2f, 0xca, 0x86, 0xec, 0x45, 0x91, 0xea,
	0x68, 0xab, 0x15, 0x95, 0xdd, 0x49, 0x65, 0xa0, 0x78, 0x57, 0x15, 0x86, 0x4e, 0x47, 0xd6, 0x86,
	0x36, 0xb8, 0x36, 0x35, 0x49, 0x57, 0xbd, 0xf7, 0xc2, 0x6f, 0x39, 0xd8, 0x27, 0x91, 0xa5, 0xeb,
	0xdc, 0xf0, 0x7f, 0xd5, 0x3c, 0xcf, 0x79, 0xfb, 0x1d, 0x4e, 0x61, 0x98, 0x0b, 0xae, 0xb8, 0x4b,
	0xc2, 0x90, 0x17, 0x99, 0xaa, 0xbf, 0x8e, 0x76, 0x71, 0xfb, 0x24, 0x07, 0x2f, 0x63, 0xa6, 0x92,
	0x62, 0xe3, 0x84, 0x3c, 0x75, 0xd3, 0x92, 0xa9, 0xef, 0xbc, 0x74, 0x63, 0x6e, 0xeb, 0x2c, 0x7b,
	0x47, 0xb6, 0x2c, 0x22, 0x8a, 0x0b, 0xe9, 0x9e, 0x9f, 0x55, 0x83, 0xe9, 0x6f, 0x04, 0xc6, 0x27,
	0x49, 0x05, 0xbe, 0x03, 0x0d, 0x16, 0x99, 0x68, 0x8c, 0x2c, 0x23, 0x68, 0xb0, 0x08, 0xdf, 0x83,
	0x26, 0x4d, 0x09, 0xdb, 0x9a, 0x8d, 0x31, 0xb2, 0x3a, 0x41, 0x25, 0xf0, 0x03, 0xe8, 0xe7, 0x44,
	0xca, 0x92, 0x8b, 0x68, 0x9d, 0x10, 0x99, 0x98, 0xb7, 0xc6, 0xc8, 0xea, 0x05, 0xbd, 0xda, 0x7c,
	0x4b, 0x64, 0x82, 0x31, 0x18, 0x19, 0x49, 0xa9, 0x69, 0xe8, 0x4a, 0xfd, 0xc6, 0xf7, 0x01, 0x42,
	0x9e, 0xe6, 0x24, 0xfb, 0xb5, 0x66, 0x91, 0xd9, 0xd4, 0x63, 0x3a, 0x27, 0xe7, 0x5d, 0x34, 0xfd,
	0x83, 0xe0, 0xee, 0x1b, 0x41, 0x89, 0xa2, 0x47, 0x98, 0x80, 0xfe, 0x28, 0xa8, 0x54, 0xf8, 0x7d,
	0xcd, 0x70, 0xc4, 0xea, 0xf8, 0x2f, 0x0e, 0xfb, 0xd1, 0x33, 0xb0, 0xbf, 0xad, 0xca, 0x99, 0xf5,
	0xc5, 0x9e, 0x39, 0x5f, 0x57, 0xe5, 0xec, 0xf1, 0x93, 0xd7, 0x95, 0x3c, 0xa9, 0x95, 0x73, 0x25,
	0x1f, 0x2e, 0x51, 0x8d, 0xfe, 0x08, 0x6e, 0xd7, 0x94, 0x7a, 0xa7, 0x9e, 0x0f, 0x87, 0xfd, 0xa8,
	0xb5, 0x44, 0x79, 0xfb, 0xe7, 0x24, 0x38, 0xc7, 0xa6, 0xaf, 0x00, 0x5f, 0x92, 0xc8, 0x9c, 0x67,
	0x92, 0xe2, 0x09, 0x18, 0x85, 0xa4, 0x42, 0x93, 0x74, 0xbd, 0xbe, 0x53, 0x9f, 0x41, 0x27, 0xe9,
	0x90, 0xf7, 0x11, 0xba, 0x47, 0xf5, 0x81, 0x8a, 0x1d, 0x0b, 0x29, 0x5e, 0x00, 0xfc, 0xeb, 0x83,
	0x07, 0xe7, 0x8a, 0xff, 0xd6, 0x1c, 0x0c, 0x6f, 0x8c, 0x55, 0x83, 0x7d, 0xef, 0xf3, 0xd3, 0x8b,
	0xd3, 0x72, 0xc1, 0x62, 0x96, 0x6d, 0x68, 0x96, 0x29, 0x5e, 0xb8, 0xde, 0xdc, 0x9b, 0xfb, 0x0b,
	0xf7, 0xfa, 0x07, 0x89, 0xf9, 0xa6, 0xa5, 0x9d, 0xe7, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x22,
	0x24, 0x11, 0x17, 0x3b, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/account.UserService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.UserService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "account.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/account/account.proto",
}
