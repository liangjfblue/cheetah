// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

//protoc --proto_path=. --micro_out=. --go_out=. ./api.proto

package micro_srv_cheetah_web

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

type AuthRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthRequest) Reset()         { *m = AuthRequest{} }
func (m *AuthRequest) String() string { return proto.CompactTextString(m) }
func (*AuthRequest) ProtoMessage()    {}
func (*AuthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{9}
}

func (m *AuthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthRequest.Unmarshal(m, b)
}
func (m *AuthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthRequest.Marshal(b, m, deterministic)
}
func (m *AuthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthRequest.Merge(m, src)
}
func (m *AuthRequest) XXX_Size() int {
	return xxx_messageInfo_AuthRequest.Size(m)
}
func (m *AuthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthRequest proto.InternalMessageInfo

func (m *AuthRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type AuthRespond struct {
	Code                 int32    `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=Username,proto3" json:"Username,omitempty"`
	Uid                  string   `protobuf:"bytes,3,opt,name=Uid,proto3" json:"Uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthRespond) Reset()         { *m = AuthRespond{} }
func (m *AuthRespond) String() string { return proto.CompactTextString(m) }
func (*AuthRespond) ProtoMessage()    {}
func (*AuthRespond) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{10}
}

func (m *AuthRespond) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthRespond.Unmarshal(m, b)
}
func (m *AuthRespond) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthRespond.Marshal(b, m, deterministic)
}
func (m *AuthRespond) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthRespond.Merge(m, src)
}
func (m *AuthRespond) XXX_Size() int {
	return xxx_messageInfo_AuthRespond.Size(m)
}
func (m *AuthRespond) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthRespond.DiscardUnknown(m)
}

var xxx_messageInfo_AuthRespond proto.InternalMessageInfo

func (m *AuthRespond) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *AuthRespond) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AuthRespond) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

type PrivilegeMidRequest struct {
	Sub                  string   `protobuf:"bytes,1,opt,name=sub,proto3" json:"sub,omitempty"`
	Obj                  string   `protobuf:"bytes,2,opt,name=obj,proto3" json:"obj,omitempty"`
	Act                  string   `protobuf:"bytes,3,opt,name=act,proto3" json:"act,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrivilegeMidRequest) Reset()         { *m = PrivilegeMidRequest{} }
func (m *PrivilegeMidRequest) String() string { return proto.CompactTextString(m) }
func (*PrivilegeMidRequest) ProtoMessage()    {}
func (*PrivilegeMidRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{11}
}

func (m *PrivilegeMidRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrivilegeMidRequest.Unmarshal(m, b)
}
func (m *PrivilegeMidRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrivilegeMidRequest.Marshal(b, m, deterministic)
}
func (m *PrivilegeMidRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrivilegeMidRequest.Merge(m, src)
}
func (m *PrivilegeMidRequest) XXX_Size() int {
	return xxx_messageInfo_PrivilegeMidRequest.Size(m)
}
func (m *PrivilegeMidRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PrivilegeMidRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PrivilegeMidRequest proto.InternalMessageInfo

func (m *PrivilegeMidRequest) GetSub() string {
	if m != nil {
		return m.Sub
	}
	return ""
}

func (m *PrivilegeMidRequest) GetObj() string {
	if m != nil {
		return m.Obj
	}
	return ""
}

func (m *PrivilegeMidRequest) GetAct() string {
	if m != nil {
		return m.Act
	}
	return ""
}

type PrivilegeMidRespond struct {
	Code                 int32    `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrivilegeMidRespond) Reset()         { *m = PrivilegeMidRespond{} }
func (m *PrivilegeMidRespond) String() string { return proto.CompactTextString(m) }
func (*PrivilegeMidRespond) ProtoMessage()    {}
func (*PrivilegeMidRespond) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{12}
}

func (m *PrivilegeMidRespond) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrivilegeMidRespond.Unmarshal(m, b)
}
func (m *PrivilegeMidRespond) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrivilegeMidRespond.Marshal(b, m, deterministic)
}
func (m *PrivilegeMidRespond) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrivilegeMidRespond.Merge(m, src)
}
func (m *PrivilegeMidRespond) XXX_Size() int {
	return xxx_messageInfo_PrivilegeMidRespond.Size(m)
}
func (m *PrivilegeMidRespond) XXX_DiscardUnknown() {
	xxx_messageInfo_PrivilegeMidRespond.DiscardUnknown(m)
}

var xxx_messageInfo_PrivilegeMidRespond proto.InternalMessageInfo

func (m *PrivilegeMidRespond) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func init() {
	proto.RegisterType((*RegisterRequest)(nil), "micro.srv.cheetah.web.RegisterRequest")
	proto.RegisterType((*RegisterRespond)(nil), "micro.srv.cheetah.web.RegisterRespond")
	proto.RegisterType((*LoginRequest)(nil), "micro.srv.cheetah.web.LoginRequest")
	proto.RegisterType((*LoginRespond)(nil), "micro.srv.cheetah.web.LoginRespond")
	proto.RegisterType((*GetRequest)(nil), "micro.srv.cheetah.web.GetRequest")
	proto.RegisterType((*GetRespond)(nil), "micro.srv.cheetah.web.GetRespond")
	proto.RegisterType((*ListRequest)(nil), "micro.srv.cheetah.web.ListRequest")
	proto.RegisterType((*One)(nil), "micro.srv.cheetah.web.One")
	proto.RegisterType((*ListRespond)(nil), "micro.srv.cheetah.web.ListRespond")
	proto.RegisterMapType((map[int32]*One)(nil), "micro.srv.cheetah.web.ListRespond.AllEntry")
	proto.RegisterType((*AuthRequest)(nil), "micro.srv.cheetah.web.AuthRequest")
	proto.RegisterType((*AuthRespond)(nil), "micro.srv.cheetah.web.AuthRespond")
	proto.RegisterType((*PrivilegeMidRequest)(nil), "micro.srv.cheetah.web.PrivilegeMidRequest")
	proto.RegisterType((*PrivilegeMidRespond)(nil), "micro.srv.cheetah.web.PrivilegeMidRespond")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 540 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0x41, 0x6f, 0xd3, 0x30,
	0x18, 0x25, 0x75, 0x82, 0xb6, 0xaf, 0x93, 0x40, 0x06, 0xa4, 0x28, 0x07, 0x54, 0x3c, 0x09, 0x15,
	0x90, 0x22, 0x54, 0x0e, 0x4c, 0x48, 0x1c, 0xa2, 0x09, 0x7a, 0x60, 0x53, 0xab, 0x8c, 0x1d, 0x90,
	0xb8, 0x24, 0x8d, 0x95, 0x7a, 0xcb, 0xe2, 0x12, 0x3b, 0x9d, 0xc6, 0x0f, 0xe1, 0x47, 0xf1, 0xab,
	0x90, 0xe3, 0x38, 0xcd, 0xca, 0x92, 0x4e, 0x70, 0xfb, 0x6c, 0x3f, 0xbf, 0xf7, 0xfa, 0xf9, 0x7d,
	0x0d, 0xec, 0x47, 0x2b, 0xe6, 0xaf, 0x0a, 0x2e, 0x39, 0x7e, 0x76, 0xc5, 0x16, 0x05, 0xf7, 0x45,
	0xb1, 0xf6, 0x17, 0x4b, 0x4a, 0x65, 0xb4, 0xf4, 0xaf, 0x69, 0x4c, 0x38, 0x3c, 0x0a, 0x69, 0xca,
	0x84, 0xa4, 0x45, 0x48, 0x7f, 0x94, 0x54, 0x48, 0xec, 0xc1, 0xde, 0xb9, 0xa0, 0x45, 0x1e, 0x5d,
	0x51, 0xd7, 0x1a, 0x59, 0xe3, 0xfd, 0xb0, 0x59, 0xab, 0xb3, 0x79, 0x24, 0xc4, 0x35, 0x2f, 0x12,
	0x77, 0xa0, 0xcf, 0xcc, 0x1a, 0x3f, 0x06, 0x14, 0xa4, 0xd4, 0x45, 0x23, 0x6b, 0xec, 0x84, 0xaa,
	0xc4, 0x18, 0xec, 0x20, 0x49, 0x0a, 0xd7, 0xae, 0x90, 0x55, 0x4d, 0xde, 0xb7, 0x05, 0xc5, 0x8a,
	0xe7, 0x89, 0x82, 0x1d, 0xf3, 0x44, 0x8b, 0x39, 0x61, 0x55, 0x2b, 0xb2, 0x73, 0x66, 0x34, 0x54,
	0x49, 0x3e, 0xc3, 0xc1, 0x09, 0x4f, 0x59, 0xfe, 0x9f, 0x36, 0xc9, 0x51, 0xc3, 0xd3, 0xad, 0xfe,
	0x14, 0x9c, 0xaf, 0xfc, 0x92, 0xe6, 0xf5, 0x65, 0xbd, 0x20, 0xcf, 0x01, 0xa6, 0x54, 0x1a, 0xfd,
	0xda, 0xa1, 0xb5, 0x71, 0x18, 0xd7, 0xe7, 0xdd, 0xbc, 0x6d, 0xcf, 0x83, 0x2d, 0xcf, 0xf7, 0x6b,
	0xdf, 0x37, 0x18, 0x9e, 0x30, 0xd1, 0x98, 0xc0, 0x60, 0xcf, 0xa3, 0xb4, 0x11, 0x51, 0xb5, 0xfe,
	0xf1, 0x29, 0x3d, 0x63, 0x3f, 0xb5, 0x88, 0x13, 0x36, 0xeb, 0x5b, 0x06, 0xd0, 0x6d, 0x03, 0x64,
	0x0a, 0x68, 0x96, 0xd3, 0xde, 0xbe, 0xd6, 0x1e, 0x07, 0x7f, 0x7b, 0x44, 0x2d, 0x8f, 0xbf, 0x2d,
	0x63, 0xb2, 0xb7, 0xc3, 0xc7, 0xbc, 0xcc, 0x65, 0xcd, 0xa5, 0x17, 0xf8, 0x23, 0xa0, 0x20, 0xcb,
	0x5c, 0x34, 0x42, 0xe3, 0xe1, 0xe4, 0x8d, 0x7f, 0x67, 0x64, 0xfd, 0x16, 0xb5, 0x1f, 0x64, 0xd9,
	0xa7, 0x5c, 0x16, 0x37, 0xa1, 0xba, 0xe7, 0x85, 0xb0, 0x67, 0x36, 0x94, 0xd5, 0x4b, 0x7a, 0x53,
	0x6b, 0xaa, 0x12, 0xbf, 0x05, 0x67, 0x1d, 0x65, 0xa5, 0xb6, 0x3f, 0x9c, 0x78, 0x1d, 0xf4, 0xb3,
	0x9c, 0x86, 0x1a, 0xf8, 0x61, 0x70, 0x64, 0x91, 0x43, 0x18, 0x06, 0xa5, 0x5c, 0x9a, 0x86, 0x37,
	0xc9, 0xb0, 0xda, 0xc9, 0x98, 0x19, 0xd0, 0x3f, 0x3f, 0xbd, 0x8a, 0x12, 0xda, 0x44, 0xe9, 0x0b,
	0x3c, 0x99, 0x17, 0x6c, 0xcd, 0x32, 0x9a, 0xd2, 0x53, 0x96, 0xb4, 0x32, 0x27, 0xca, 0xd8, 0x64,
	0x4e, 0x94, 0xb1, 0xda, 0xe1, 0xf1, 0x85, 0x99, 0x13, 0x1e, 0x5f, 0xa8, 0x9d, 0x68, 0x21, 0x0d,
	0x59, 0xb4, 0x90, 0xe4, 0xd5, 0x36, 0x59, 0xa7, 0xcb, 0xc9, 0x2f, 0x1b, 0x6c, 0x65, 0x0b, 0x7f,
	0x87, 0x3d, 0x33, 0xa6, 0xf8, 0x65, 0x47, 0xa7, 0xb6, 0xfe, 0x38, 0xbc, 0xdd, 0xb8, 0x4a, 0x98,
	0x3c, 0xc0, 0x67, 0xe0, 0x54, 0x33, 0x88, 0x0f, 0xbb, 0xde, 0xb8, 0x35, 0xe9, 0xde, 0x0e, 0x90,
	0x21, 0x3d, 0x05, 0x34, 0xa5, 0x12, 0xbf, 0xe8, 0x40, 0x6f, 0x46, 0xd7, 0xeb, 0x85, 0x18, 0xba,
	0x39, 0xd8, 0x2a, 0x69, 0x98, 0xf4, 0xc6, 0x50, 0x13, 0x92, 0xdd, 0x51, 0xd5, 0x8c, 0x2a, 0x25,
	0x9d, 0x8c, 0xad, 0x9c, 0x79, 0xfd, 0x18, 0xc3, 0xb8, 0x84, 0x83, 0xf6, 0xcb, 0xe2, 0xd7, 0x1d,
	0xb7, 0xee, 0xc8, 0x92, 0x77, 0x3f, 0x6c, 0xad, 0x14, 0x3f, 0xac, 0xbe, 0x22, 0xef, 0xfe, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x1a, 0x70, 0xda, 0x4f, 0x52, 0x06, 0x00, 0x00,
}
