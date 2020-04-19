// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

//protoc --proto_path=. --micro_out=. --go_out=. ./api.proto

package micro_srv_auth

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type RegisterRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	Age                  int32    `protobuf:"varint,3,opt,name=Age,proto3" json:"Age,omitempty"`
	Addr                 string   `protobuf:"bytes,4,opt,name=Addr,proto3" json:"Addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (m *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(m, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *RegisterRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *RegisterRequest) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *RegisterRequest) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

type RegisterRespond struct {
	Code                 int32    `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Uid                  string   `protobuf:"bytes,2,opt,name=Uid,proto3" json:"Uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRespond) Reset()         { *m = RegisterRespond{} }
func (m *RegisterRespond) String() string { return proto.CompactTextString(m) }
func (*RegisterRespond) ProtoMessage()    {}
func (*RegisterRespond) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *RegisterRespond) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRespond.Unmarshal(m, b)
}
func (m *RegisterRespond) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRespond.Marshal(b, m, deterministic)
}
func (m *RegisterRespond) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRespond.Merge(m, src)
}
func (m *RegisterRespond) XXX_Size() int {
	return xxx_messageInfo_RegisterRespond.Size(m)
}
func (m *RegisterRespond) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRespond.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRespond proto.InternalMessageInfo

func (m *RegisterRespond) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RegisterRespond) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

type LoginRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{2}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginRespond struct {
	Code                 int32    `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=Token,proto3" json:"Token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRespond) Reset()         { *m = LoginRespond{} }
func (m *LoginRespond) String() string { return proto.CompactTextString(m) }
func (*LoginRespond) ProtoMessage()    {}
func (*LoginRespond) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{3}
}

func (m *LoginRespond) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRespond.Unmarshal(m, b)
}
func (m *LoginRespond) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRespond.Marshal(b, m, deterministic)
}
func (m *LoginRespond) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRespond.Merge(m, src)
}
func (m *LoginRespond) XXX_Size() int {
	return xxx_messageInfo_LoginRespond.Size(m)
}
func (m *LoginRespond) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRespond.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRespond proto.InternalMessageInfo

func (m *LoginRespond) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *LoginRespond) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type GetRequest struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=Uid,proto3" json:"Uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{4}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

type GetRespond struct {
	Code                 int32    `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=Username,proto3" json:"Username,omitempty"`
	Age                  int32    `protobuf:"varint,3,opt,name=Age,proto3" json:"Age,omitempty"`
	Addr                 string   `protobuf:"bytes,4,opt,name=Addr,proto3" json:"Addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRespond) Reset()         { *m = GetRespond{} }
func (m *GetRespond) String() string { return proto.CompactTextString(m) }
func (*GetRespond) ProtoMessage()    {}
func (*GetRespond) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{5}
}

func (m *GetRespond) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRespond.Unmarshal(m, b)
}
func (m *GetRespond) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRespond.Marshal(b, m, deterministic)
}
func (m *GetRespond) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRespond.Merge(m, src)
}
func (m *GetRespond) XXX_Size() int {
	return xxx_messageInfo_GetRespond.Size(m)
}
func (m *GetRespond) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRespond.DiscardUnknown(m)
}

var xxx_messageInfo_GetRespond proto.InternalMessageInfo

func (m *GetRespond) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetRespond) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *GetRespond) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *GetRespond) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

type ListRequest struct {
	Page                 int32    `protobuf:"varint,1,opt,name=Page,proto3" json:"Page,omitempty"`
	PageSize             int32    `protobuf:"varint,2,opt,name=PageSize,proto3" json:"PageSize,omitempty"`
	Username             string   `protobuf:"bytes,3,opt,name=Username,proto3" json:"Username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{6}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ListRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type One struct {
	Username             string   `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	Age                  int32    `protobuf:"varint,2,opt,name=Age,proto3" json:"Age,omitempty"`
	Addr                 string   `protobuf:"bytes,3,opt,name=Addr,proto3" json:"Addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *One) Reset()         { *m = One{} }
func (m *One) String() string { return proto.CompactTextString(m) }
func (*One) ProtoMessage()    {}
func (*One) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{7}
}

func (m *One) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_One.Unmarshal(m, b)
}
func (m *One) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_One.Marshal(b, m, deterministic)
}
func (m *One) XXX_Merge(src proto.Message) {
	xxx_messageInfo_One.Merge(m, src)
}
func (m *One) XXX_Size() int {
	return xxx_messageInfo_One.Size(m)
}
func (m *One) XXX_DiscardUnknown() {
	xxx_messageInfo_One.DiscardUnknown(m)
}

var xxx_messageInfo_One proto.InternalMessageInfo

func (m *One) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *One) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *One) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

type ListRespond struct {
	Code                 int32          `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Count                int32          `protobuf:"varint,2,opt,name=Count,proto3" json:"Count,omitempty"`
	All                  map[int32]*One `protobuf:"bytes,3,rep,name=All,proto3" json:"All,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ListRespond) Reset()         { *m = ListRespond{} }
func (m *ListRespond) String() string { return proto.CompactTextString(m) }
func (*ListRespond) ProtoMessage()    {}
func (*ListRespond) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{8}
}

func (m *ListRespond) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRespond.Unmarshal(m, b)
}
func (m *ListRespond) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRespond.Marshal(b, m, deterministic)
}
func (m *ListRespond) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRespond.Merge(m, src)
}
func (m *ListRespond) XXX_Size() int {
	return xxx_messageInfo_ListRespond.Size(m)
}
func (m *ListRespond) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRespond.DiscardUnknown(m)
}

var xxx_messageInfo_ListRespond proto.InternalMessageInfo

func (m *ListRespond) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ListRespond) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ListRespond) GetAll() map[int32]*One {
	if m != nil {
		return m.All
	}
	return nil
}

func init() {
	proto.RegisterType((*RegisterRequest)(nil), "micro.srv.auth.RegisterRequest")
	proto.RegisterType((*RegisterRespond)(nil), "micro.srv.auth.RegisterRespond")
	proto.RegisterType((*LoginRequest)(nil), "micro.srv.auth.LoginRequest")
	proto.RegisterType((*LoginRespond)(nil), "micro.srv.auth.LoginRespond")
	proto.RegisterType((*GetRequest)(nil), "micro.srv.auth.GetRequest")
	proto.RegisterType((*GetRespond)(nil), "micro.srv.auth.GetRespond")
	proto.RegisterType((*ListRequest)(nil), "micro.srv.auth.ListRequest")
	proto.RegisterType((*One)(nil), "micro.srv.auth.One")
	proto.RegisterType((*ListRespond)(nil), "micro.srv.auth.ListRespond")
	proto.RegisterMapType((map[int32]*One)(nil), "micro.srv.auth.ListRespond.AllEntry")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 438 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x97, 0x38, 0x41, 0xdd, 0x2b, 0x02, 0x64, 0x38, 0x44, 0x19, 0x82, 0xca, 0xe2, 0x50,
	0x2e, 0x39, 0x14, 0x09, 0x26, 0x2e, 0x28, 0x4c, 0xa3, 0x07, 0x26, 0x36, 0x19, 0x76, 0xe0, 0x98,
	0x11, 0x2b, 0x44, 0xcb, 0xec, 0x62, 0x3b, 0x43, 0xe3, 0xe3, 0xf0, 0x39, 0xf8, 0x70, 0xc8, 0x71,
	0x9c, 0xa6, 0x69, 0x1b, 0x21, 0x71, 0x7b, 0xaf, 0xef, 0xf9, 0xff, 0xff, 0xbd, 0xa7, 0xd7, 0xc0,
	0x61, 0xb6, 0x2a, 0x93, 0x95, 0x14, 0x5a, 0xe0, 0x07, 0x37, 0xe5, 0x37, 0x29, 0x12, 0x25, 0x6f,
	0x93, 0xac, 0xd6, 0xdf, 0x89, 0x80, 0x87, 0x94, 0x15, 0xa5, 0xd2, 0x4c, 0x52, 0xf6, 0xa3, 0x66,
	0x4a, 0xe3, 0x18, 0x26, 0x97, 0x8a, 0x49, 0x9e, 0xdd, 0xb0, 0xc8, 0x9b, 0x79, 0xf3, 0x43, 0xda,
	0xe5, 0xa6, 0x76, 0x91, 0x29, 0xf5, 0x53, 0xc8, 0x3c, 0xf2, 0x6d, 0xcd, 0xe5, 0xf8, 0x11, 0xa0,
	0xb4, 0x60, 0x11, 0x9a, 0x79, 0xf3, 0x90, 0x9a, 0x10, 0x63, 0x08, 0xd2, 0x3c, 0x97, 0x51, 0xd0,
	0x74, 0x36, 0x31, 0x79, 0xd3, 0x37, 0x54, 0x2b, 0xc1, 0x73, 0xd3, 0x76, 0x22, 0x72, 0x6b, 0x16,
	0xd2, 0x26, 0x36, 0x62, 0x97, 0xa5, 0xf3, 0x30, 0x21, 0xf9, 0x00, 0xf7, 0xcf, 0x44, 0x51, 0xf2,
	0xff, 0xc4, 0x24, 0xc7, 0x9d, 0xce, 0x7e, 0xf7, 0x27, 0x10, 0x7e, 0x11, 0xd7, 0x8c, 0xb7, 0x8f,
	0x6d, 0x42, 0x9e, 0x01, 0x2c, 0x99, 0x76, 0xfe, 0x2d, 0xa1, 0xb7, 0x26, 0xbc, 0x6a, 0xeb, 0xfb,
	0x75, 0xfb, 0xcc, 0xfe, 0x80, 0xf9, 0xdf, 0xd6, 0xf7, 0x15, 0xa6, 0x67, 0xa5, 0xea, 0x20, 0x30,
	0x04, 0x17, 0x59, 0xd1, 0x99, 0x98, 0xd8, 0x0e, 0x5f, 0xb0, 0xcf, 0xe5, 0x2f, 0x6b, 0x12, 0xd2,
	0x2e, 0xdf, 0x00, 0x40, 0x9b, 0x00, 0x64, 0x09, 0xe8, 0x9c, 0xb3, 0xd1, 0xbd, 0xb6, 0x8c, 0xfe,
	0x36, 0x23, 0xea, 0x31, 0xfe, 0xf1, 0x1c, 0xe4, 0xe8, 0x86, 0x4f, 0x44, 0xcd, 0x75, 0xab, 0x65,
	0x13, 0xfc, 0x1a, 0x50, 0x5a, 0x55, 0x11, 0x9a, 0xa1, 0xf9, 0x74, 0xf1, 0x22, 0xd9, 0xbc, 0xd5,
	0xa4, 0xa7, 0x99, 0xa4, 0x55, 0x75, 0xca, 0xb5, 0xbc, 0xa3, 0xe6, 0x41, 0xfc, 0x11, 0x26, 0xee,
	0x07, 0xc3, 0x78, 0xcd, 0xee, 0x5a, 0x33, 0x13, 0xe2, 0x97, 0x10, 0xde, 0x66, 0x55, 0x6d, 0xb9,
	0xa7, 0x8b, 0xc7, 0x43, 0xdd, 0x73, 0xce, 0xa8, 0xed, 0x78, 0xeb, 0x1f, 0x7b, 0x8b, 0xdf, 0x3e,
	0x04, 0x66, 0x62, 0xfc, 0x09, 0x26, 0xee, 0x54, 0xf1, 0xf3, 0xe1, 0xa3, 0xc1, 0xbf, 0x26, 0x1e,
	0x69, 0x68, 0x88, 0xc9, 0x01, 0x3e, 0x85, 0xb0, 0xb9, 0x3c, 0xfc, 0x74, 0x6b, 0xb2, 0xde, 0x61,
	0xc7, 0xfb, 0xaa, 0x4e, 0xe6, 0x1d, 0xa0, 0x25, 0xd3, 0x38, 0x1e, 0xb6, 0xad, 0x6f, 0x33, 0xde,
	0x5d, 0x73, 0x02, 0xef, 0x21, 0x30, 0xab, 0xc4, 0x47, 0xbb, 0x17, 0x6c, 0x25, 0x8e, 0x46, 0xb6,
	0x4f, 0x0e, 0xae, 0xee, 0x35, 0x9f, 0x93, 0x57, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x7e, 0x23,
	0xf5, 0xf4, 0x5b, 0x04, 0x00, 0x00,
}