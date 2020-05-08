// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

//protoc --proto_path=. --micro_out=. --go_out=. ./api.proto

package micro_srv_cheetah_worker

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

type StartJobRequest struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Type                 int32    `protobuf:"varint,3,opt,name=Type,proto3" json:"Type,omitempty"`
	Content              string   `protobuf:"bytes,4,opt,name=Content,proto3" json:"Content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartJobRequest) Reset()         { *m = StartJobRequest{} }
func (m *StartJobRequest) String() string { return proto.CompactTextString(m) }
func (*StartJobRequest) ProtoMessage()    {}
func (*StartJobRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *StartJobRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartJobRequest.Unmarshal(m, b)
}
func (m *StartJobRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartJobRequest.Marshal(b, m, deterministic)
}
func (m *StartJobRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartJobRequest.Merge(m, src)
}
func (m *StartJobRequest) XXX_Size() int {
	return xxx_messageInfo_StartJobRequest.Size(m)
}
func (m *StartJobRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StartJobRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StartJobRequest proto.InternalMessageInfo

func (m *StartJobRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *StartJobRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StartJobRequest) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *StartJobRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type StartJobRespond struct {
	Code                 int32    `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartJobRespond) Reset()         { *m = StartJobRespond{} }
func (m *StartJobRespond) String() string { return proto.CompactTextString(m) }
func (*StartJobRespond) ProtoMessage()    {}
func (*StartJobRespond) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *StartJobRespond) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartJobRespond.Unmarshal(m, b)
}
func (m *StartJobRespond) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartJobRespond.Marshal(b, m, deterministic)
}
func (m *StartJobRespond) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartJobRespond.Merge(m, src)
}
func (m *StartJobRespond) XXX_Size() int {
	return xxx_messageInfo_StartJobRespond.Size(m)
}
func (m *StartJobRespond) XXX_DiscardUnknown() {
	xxx_messageInfo_StartJobRespond.DiscardUnknown(m)
}

var xxx_messageInfo_StartJobRespond proto.InternalMessageInfo

func (m *StartJobRespond) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

type StopJobRequest struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopJobRequest) Reset()         { *m = StopJobRequest{} }
func (m *StopJobRequest) String() string { return proto.CompactTextString(m) }
func (*StopJobRequest) ProtoMessage()    {}
func (*StopJobRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{2}
}

func (m *StopJobRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopJobRequest.Unmarshal(m, b)
}
func (m *StopJobRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopJobRequest.Marshal(b, m, deterministic)
}
func (m *StopJobRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopJobRequest.Merge(m, src)
}
func (m *StopJobRequest) XXX_Size() int {
	return xxx_messageInfo_StopJobRequest.Size(m)
}
func (m *StopJobRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StopJobRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StopJobRequest proto.InternalMessageInfo

func (m *StopJobRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type StopJobRespond struct {
	Code                 int32    `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopJobRespond) Reset()         { *m = StopJobRespond{} }
func (m *StopJobRespond) String() string { return proto.CompactTextString(m) }
func (*StopJobRespond) ProtoMessage()    {}
func (*StopJobRespond) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{3}
}

func (m *StopJobRespond) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopJobRespond.Unmarshal(m, b)
}
func (m *StopJobRespond) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopJobRespond.Marshal(b, m, deterministic)
}
func (m *StopJobRespond) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopJobRespond.Merge(m, src)
}
func (m *StopJobRespond) XXX_Size() int {
	return xxx_messageInfo_StopJobRespond.Size(m)
}
func (m *StopJobRespond) XXX_DiscardUnknown() {
	xxx_messageInfo_StopJobRespond.DiscardUnknown(m)
}

var xxx_messageInfo_StopJobRespond proto.InternalMessageInfo

func (m *StopJobRespond) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

type RestartJobRequest struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RestartJobRequest) Reset()         { *m = RestartJobRequest{} }
func (m *RestartJobRequest) String() string { return proto.CompactTextString(m) }
func (*RestartJobRequest) ProtoMessage()    {}
func (*RestartJobRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{4}
}

func (m *RestartJobRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RestartJobRequest.Unmarshal(m, b)
}
func (m *RestartJobRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RestartJobRequest.Marshal(b, m, deterministic)
}
func (m *RestartJobRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RestartJobRequest.Merge(m, src)
}
func (m *RestartJobRequest) XXX_Size() int {
	return xxx_messageInfo_RestartJobRequest.Size(m)
}
func (m *RestartJobRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RestartJobRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RestartJobRequest proto.InternalMessageInfo

func (m *RestartJobRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type RestartRespond struct {
	Code                 int32    `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RestartRespond) Reset()         { *m = RestartRespond{} }
func (m *RestartRespond) String() string { return proto.CompactTextString(m) }
func (*RestartRespond) ProtoMessage()    {}
func (*RestartRespond) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{5}
}

func (m *RestartRespond) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RestartRespond.Unmarshal(m, b)
}
func (m *RestartRespond) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RestartRespond.Marshal(b, m, deterministic)
}
func (m *RestartRespond) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RestartRespond.Merge(m, src)
}
func (m *RestartRespond) XXX_Size() int {
	return xxx_messageInfo_RestartRespond.Size(m)
}
func (m *RestartRespond) XXX_DiscardUnknown() {
	xxx_messageInfo_RestartRespond.DiscardUnknown(m)
}

var xxx_messageInfo_RestartRespond proto.InternalMessageInfo

func (m *RestartRespond) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

type JobProgressRequest struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JobProgressRequest) Reset()         { *m = JobProgressRequest{} }
func (m *JobProgressRequest) String() string { return proto.CompactTextString(m) }
func (*JobProgressRequest) ProtoMessage()    {}
func (*JobProgressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{6}
}

func (m *JobProgressRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobProgressRequest.Unmarshal(m, b)
}
func (m *JobProgressRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobProgressRequest.Marshal(b, m, deterministic)
}
func (m *JobProgressRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobProgressRequest.Merge(m, src)
}
func (m *JobProgressRequest) XXX_Size() int {
	return xxx_messageInfo_JobProgressRequest.Size(m)
}
func (m *JobProgressRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_JobProgressRequest.DiscardUnknown(m)
}

var xxx_messageInfo_JobProgressRequest proto.InternalMessageInfo

func (m *JobProgressRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type JobProgressRespond struct {
	Code                 int32    `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Progress             int32    `protobuf:"varint,2,opt,name=Progress,proto3" json:"Progress,omitempty"`
	Status               int32    `protobuf:"varint,3,opt,name=Status,proto3" json:"Status,omitempty"`
	Msg                  string   `protobuf:"bytes,4,opt,name=Msg,proto3" json:"Msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JobProgressRespond) Reset()         { *m = JobProgressRespond{} }
func (m *JobProgressRespond) String() string { return proto.CompactTextString(m) }
func (*JobProgressRespond) ProtoMessage()    {}
func (*JobProgressRespond) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{7}
}

func (m *JobProgressRespond) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobProgressRespond.Unmarshal(m, b)
}
func (m *JobProgressRespond) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobProgressRespond.Marshal(b, m, deterministic)
}
func (m *JobProgressRespond) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobProgressRespond.Merge(m, src)
}
func (m *JobProgressRespond) XXX_Size() int {
	return xxx_messageInfo_JobProgressRespond.Size(m)
}
func (m *JobProgressRespond) XXX_DiscardUnknown() {
	xxx_messageInfo_JobProgressRespond.DiscardUnknown(m)
}

var xxx_messageInfo_JobProgressRespond proto.InternalMessageInfo

func (m *JobProgressRespond) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *JobProgressRespond) GetProgress() int32 {
	if m != nil {
		return m.Progress
	}
	return 0
}

func (m *JobProgressRespond) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *JobProgressRespond) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*StartJobRequest)(nil), "micro.srv.cheetah.worker.StartJobRequest")
	proto.RegisterType((*StartJobRespond)(nil), "micro.srv.cheetah.worker.StartJobRespond")
	proto.RegisterType((*StopJobRequest)(nil), "micro.srv.cheetah.worker.StopJobRequest")
	proto.RegisterType((*StopJobRespond)(nil), "micro.srv.cheetah.worker.StopJobRespond")
	proto.RegisterType((*RestartJobRequest)(nil), "micro.srv.cheetah.worker.RestartJobRequest")
	proto.RegisterType((*RestartRespond)(nil), "micro.srv.cheetah.worker.RestartRespond")
	proto.RegisterType((*JobProgressRequest)(nil), "micro.srv.cheetah.worker.JobProgressRequest")
	proto.RegisterType((*JobProgressRespond)(nil), "micro.srv.cheetah.worker.JobProgressRespond")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x5f, 0x4b, 0xf3, 0x30,
	0x14, 0xc6, 0xd7, 0xfd, 0xdf, 0x79, 0x61, 0xaf, 0x9e, 0x0b, 0x29, 0xbd, 0x1a, 0x71, 0xc2, 0x44,
	0x09, 0xa2, 0x1f, 0x61, 0xbb, 0xd9, 0x40, 0x91, 0x4c, 0xf0, 0x52, 0xda, 0xed, 0xb0, 0x0d, 0x59,
	0x53, 0x93, 0x4c, 0xf1, 0xf3, 0xf8, 0x45, 0xa5, 0x21, 0xd5, 0xba, 0x91, 0xb2, 0xbb, 0x93, 0xf0,
	0x9c, 0xe7, 0x69, 0x9e, 0x1f, 0x85, 0x5e, 0x9c, 0x6d, 0x78, 0xa6, 0xa4, 0x91, 0x18, 0x6e, 0x37,
	0x0b, 0x25, 0xb9, 0x56, 0xef, 0x7c, 0xb1, 0x26, 0x32, 0xf1, 0x9a, 0x7f, 0x48, 0xf5, 0x4a, 0x8a,
	0x2d, 0xe0, 0xff, 0xdc, 0xc4, 0xca, 0xcc, 0x64, 0x22, 0xe8, 0x6d, 0x47, 0xda, 0x60, 0x1f, 0xea,
	0xd3, 0x49, 0x18, 0x0c, 0x82, 0x51, 0x4f, 0xd4, 0xa7, 0x13, 0x44, 0x68, 0x3e, 0xc4, 0x5b, 0x0a,
	0xeb, 0xf6, 0xc6, 0xce, 0xf9, 0xdd, 0xd3, 0x67, 0x46, 0x61, 0x63, 0x10, 0x8c, 0x5a, 0xc2, 0xce,
	0x18, 0x42, 0x67, 0x2c, 0x53, 0x43, 0xa9, 0x09, 0x9b, 0x56, 0x5a, 0x1c, 0xd9, 0x45, 0x39, 0x44,
	0x67, 0x32, 0x5d, 0xe6, 0x06, 0x63, 0xb9, 0x24, 0x1b, 0xd3, 0x12, 0x76, 0x66, 0x03, 0xe8, 0xcf,
	0x8d, 0xcc, 0xfc, 0x9f, 0xc2, 0x86, 0x25, 0x85, 0xdf, 0xe7, 0x1c, 0x4e, 0x05, 0xe9, 0xea, 0x57,
	0xe5, 0x56, 0x4e, 0x54, 0x65, 0x35, 0x04, 0x9c, 0xc9, 0xe4, 0x51, 0xc9, 0x95, 0x22, 0xad, 0x7d,
	0x5e, 0xe9, 0x9e, 0xca, 0xeb, 0x87, 0x11, 0x74, 0x0b, 0x99, 0xed, 0xb3, 0x25, 0x7e, 0xce, 0x78,
	0x06, 0xed, 0xb9, 0x89, 0xcd, 0x4e, 0xbb, 0x56, 0xdd, 0x09, 0x4f, 0xa0, 0x71, 0xaf, 0x57, 0xae,
	0xd3, 0x7c, 0xbc, 0xfd, 0x6a, 0x40, 0xfb, 0xd9, 0xf2, 0xc3, 0x04, 0xba, 0x45, 0xb5, 0x78, 0xc9,
	0x7d, 0x98, 0xf9, 0x1e, 0xe3, 0xe8, 0x28, 0xa9, 0x7d, 0x06, 0xab, 0xe1, 0x0b, 0x74, 0x5c, 0xeb,
	0x38, 0xaa, 0xda, 0x2b, 0xa3, 0x8b, 0x8e, 0x51, 0x16, 0x01, 0x04, 0xf0, 0x0b, 0x0c, 0xaf, 0xfc,
	0x9b, 0x07, 0x58, 0xab, 0x62, 0xfe, 0xe2, 0x65, 0x35, 0xdc, 0xc2, 0xbf, 0x12, 0x26, 0xbc, 0xf6,
	0xaf, 0x1e, 0x32, 0x8f, 0x8e, 0x55, 0xbb, 0xb0, 0x9b, 0x20, 0x69, 0xdb, 0x7f, 0xef, 0xee, 0x3b,
	0x00, 0x00, 0xff, 0xff, 0x4c, 0x09, 0xc2, 0x50, 0x88, 0x03, 0x00, 0x00,
}