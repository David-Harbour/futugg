// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Qot_GetSuspend/Qot_GetSuspend.proto

package Qot_GetSuspend

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/jerryharbour/futugg/pb/Common"
import Qot_Common "github.com/jerryharbour/futugg/pb/Qot_Common"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type C2S struct {
	SecurityList         []*Qot_Common.Security `protobuf:"bytes,1,rep,name=securityList" json:"securityList,omitempty"`
	BeginTime            *string                `protobuf:"bytes,2,req,name=beginTime" json:"beginTime,omitempty"`
	EndTime              *string                `protobuf:"bytes,3,req,name=endTime" json:"endTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *C2S) Reset()         { *m = C2S{} }
func (m *C2S) String() string { return proto.CompactTextString(m) }
func (*C2S) ProtoMessage()    {}
func (*C2S) Descriptor() ([]byte, []int) {
	return fileDescriptor_Qot_GetSuspend_6e3c07775f957e58, []int{0}
}
func (m *C2S) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2S.Unmarshal(m, b)
}
func (m *C2S) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2S.Marshal(b, m, deterministic)
}
func (dst *C2S) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2S.Merge(dst, src)
}
func (m *C2S) XXX_Size() int {
	return xxx_messageInfo_C2S.Size(m)
}
func (m *C2S) XXX_DiscardUnknown() {
	xxx_messageInfo_C2S.DiscardUnknown(m)
}

var xxx_messageInfo_C2S proto.InternalMessageInfo

func (m *C2S) GetSecurityList() []*Qot_Common.Security {
	if m != nil {
		return m.SecurityList
	}
	return nil
}

func (m *C2S) GetBeginTime() string {
	if m != nil && m.BeginTime != nil {
		return *m.BeginTime
	}
	return ""
}

func (m *C2S) GetEndTime() string {
	if m != nil && m.EndTime != nil {
		return *m.EndTime
	}
	return ""
}

type Suspend struct {
	Time                 *string  `protobuf:"bytes,1,req,name=time" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Suspend) Reset()         { *m = Suspend{} }
func (m *Suspend) String() string { return proto.CompactTextString(m) }
func (*Suspend) ProtoMessage()    {}
func (*Suspend) Descriptor() ([]byte, []int) {
	return fileDescriptor_Qot_GetSuspend_6e3c07775f957e58, []int{1}
}
func (m *Suspend) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Suspend.Unmarshal(m, b)
}
func (m *Suspend) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Suspend.Marshal(b, m, deterministic)
}
func (dst *Suspend) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Suspend.Merge(dst, src)
}
func (m *Suspend) XXX_Size() int {
	return xxx_messageInfo_Suspend.Size(m)
}
func (m *Suspend) XXX_DiscardUnknown() {
	xxx_messageInfo_Suspend.DiscardUnknown(m)
}

var xxx_messageInfo_Suspend proto.InternalMessageInfo

func (m *Suspend) GetTime() string {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return ""
}

type SecuritySuspend struct {
	Security             *Qot_Common.Security `protobuf:"bytes,1,req,name=security" json:"security,omitempty"`
	SuspendList          []*Suspend           `protobuf:"bytes,2,rep,name=suspendList" json:"suspendList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SecuritySuspend) Reset()         { *m = SecuritySuspend{} }
func (m *SecuritySuspend) String() string { return proto.CompactTextString(m) }
func (*SecuritySuspend) ProtoMessage()    {}
func (*SecuritySuspend) Descriptor() ([]byte, []int) {
	return fileDescriptor_Qot_GetSuspend_6e3c07775f957e58, []int{2}
}
func (m *SecuritySuspend) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SecuritySuspend.Unmarshal(m, b)
}
func (m *SecuritySuspend) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SecuritySuspend.Marshal(b, m, deterministic)
}
func (dst *SecuritySuspend) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SecuritySuspend.Merge(dst, src)
}
func (m *SecuritySuspend) XXX_Size() int {
	return xxx_messageInfo_SecuritySuspend.Size(m)
}
func (m *SecuritySuspend) XXX_DiscardUnknown() {
	xxx_messageInfo_SecuritySuspend.DiscardUnknown(m)
}

var xxx_messageInfo_SecuritySuspend proto.InternalMessageInfo

func (m *SecuritySuspend) GetSecurity() *Qot_Common.Security {
	if m != nil {
		return m.Security
	}
	return nil
}

func (m *SecuritySuspend) GetSuspendList() []*Suspend {
	if m != nil {
		return m.SuspendList
	}
	return nil
}

type S2C struct {
	SecuritySuspendList  []*SecuritySuspend `protobuf:"bytes,1,rep,name=SecuritySuspendList" json:"SecuritySuspendList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *S2C) Reset()         { *m = S2C{} }
func (m *S2C) String() string { return proto.CompactTextString(m) }
func (*S2C) ProtoMessage()    {}
func (*S2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_Qot_GetSuspend_6e3c07775f957e58, []int{3}
}
func (m *S2C) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2C.Unmarshal(m, b)
}
func (m *S2C) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2C.Marshal(b, m, deterministic)
}
func (dst *S2C) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2C.Merge(dst, src)
}
func (m *S2C) XXX_Size() int {
	return xxx_messageInfo_S2C.Size(m)
}
func (m *S2C) XXX_DiscardUnknown() {
	xxx_messageInfo_S2C.DiscardUnknown(m)
}

var xxx_messageInfo_S2C proto.InternalMessageInfo

func (m *S2C) GetSecuritySuspendList() []*SecuritySuspend {
	if m != nil {
		return m.SecuritySuspendList
	}
	return nil
}

type Request struct {
	C2S                  *C2S     `protobuf:"bytes,1,req,name=c2s" json:"c2s,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_Qot_GetSuspend_6e3c07775f957e58, []int{4}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetC2S() *C2S {
	if m != nil {
		return m.C2S
	}
	return nil
}

type Response struct {
	RetType              *int32   `protobuf:"varint,1,req,name=retType,def=-400" json:"retType,omitempty"`
	RetMsg               *string  `protobuf:"bytes,2,opt,name=retMsg" json:"retMsg,omitempty"`
	ErrCode              *int32   `protobuf:"varint,3,opt,name=errCode" json:"errCode,omitempty"`
	S2C                  *S2C     `protobuf:"bytes,4,opt,name=s2c" json:"s2c,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_Qot_GetSuspend_6e3c07775f957e58, []int{5}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

const Default_Response_RetType int32 = -400

func (m *Response) GetRetType() int32 {
	if m != nil && m.RetType != nil {
		return *m.RetType
	}
	return Default_Response_RetType
}

func (m *Response) GetRetMsg() string {
	if m != nil && m.RetMsg != nil {
		return *m.RetMsg
	}
	return ""
}

func (m *Response) GetErrCode() int32 {
	if m != nil && m.ErrCode != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *Response) GetS2C() *S2C {
	if m != nil {
		return m.S2C
	}
	return nil
}

func init() {
	proto.RegisterType((*C2S)(nil), "Qot_GetSuspend.C2S")
	proto.RegisterType((*Suspend)(nil), "Qot_GetSuspend.Suspend")
	proto.RegisterType((*SecuritySuspend)(nil), "Qot_GetSuspend.SecuritySuspend")
	proto.RegisterType((*S2C)(nil), "Qot_GetSuspend.S2C")
	proto.RegisterType((*Request)(nil), "Qot_GetSuspend.Request")
	proto.RegisterType((*Response)(nil), "Qot_GetSuspend.Response")
}

func init() {
	proto.RegisterFile("Qot_GetSuspend/Qot_GetSuspend.proto", fileDescriptor_Qot_GetSuspend_6e3c07775f957e58)
}

var fileDescriptor_Qot_GetSuspend_6e3c07775f957e58 = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xbf, 0x4f, 0xfb, 0x30,
	0x10, 0xc5, 0x95, 0xa4, 0xfd, 0xa6, 0xbd, 0x7e, 0x05, 0x92, 0x8b, 0x20, 0x2a, 0xbf, 0xa2, 0x20,
	0xa4, 0x2c, 0xb4, 0x91, 0xc5, 0x00, 0xac, 0x19, 0x58, 0x60, 0xa8, 0xd3, 0x81, 0x0d, 0x89, 0xf6,
	0x54, 0x65, 0x68, 0x1c, 0x62, 0x57, 0xa8, 0x0b, 0x0b, 0xff, 0x38, 0xb2, 0x63, 0xd3, 0x26, 0x94,
	0xc9, 0x7e, 0xf7, 0x3e, 0xb6, 0xdf, 0x9d, 0xe1, 0x6a, 0xca, 0xe5, 0xeb, 0x23, 0xca, 0x6c, 0x2d,
	0x4a, 0x2c, 0x16, 0x93, 0xa6, 0x1c, 0x97, 0x15, 0x97, 0x9c, 0x1c, 0x34, 0xab, 0xa3, 0x61, 0xca,
	0x57, 0x2b, 0x5e, 0x4c, 0xea, 0xa5, 0x86, 0x46, 0xa7, 0x0a, 0x32, 0xc6, 0x76, 0x5b, 0x9b, 0xd1,
	0x07, 0x78, 0x29, 0xcd, 0xc8, 0x1d, 0xfc, 0x17, 0x38, 0x5f, 0x57, 0xb9, 0xdc, 0x3c, 0xe5, 0x42,
	0x06, 0x4e, 0xe8, 0xc5, 0x03, 0x7a, 0x34, 0xde, 0xe1, 0x33, 0xe3, 0xb3, 0x06, 0x49, 0xce, 0xa0,
	0xff, 0x86, 0xcb, 0xbc, 0x98, 0xe5, 0x2b, 0x0c, 0xdc, 0xd0, 0x8d, 0xfb, 0x6c, 0x5b, 0x20, 0x01,
	0xf8, 0x58, 0x2c, 0xb4, 0xe7, 0x69, 0xcf, 0xca, 0xe8, 0x1c, 0x7c, 0x93, 0x9a, 0x10, 0xe8, 0x48,
	0x45, 0x38, 0x9a, 0xd0, 0xfb, 0xe8, 0x13, 0x0e, 0xed, 0x83, 0x16, 0x4b, 0xa0, 0x67, 0x5f, 0xd6,
	0xe8, 0x5f, 0xf9, 0x7e, 0x28, 0x72, 0x0f, 0x03, 0x51, 0x1f, 0xd6, 0x4d, 0xb9, 0xba, 0xa9, 0x93,
	0x71, 0x6b, 0x94, 0x66, 0x65, 0xbb, 0x6c, 0xf4, 0x02, 0x5e, 0x46, 0x53, 0x32, 0x85, 0x61, 0x2b,
	0xc6, 0xce, 0x78, 0x2e, 0x7f, 0xdd, 0xd4, 0x44, 0xd9, 0xbe, 0xb3, 0x51, 0x02, 0x3e, 0xc3, 0xf7,
	0x35, 0x0a, 0x49, 0xae, 0xc1, 0x9b, 0x53, 0x61, 0x9a, 0x19, 0xb6, 0x6f, 0x4b, 0x69, 0xc6, 0x94,
	0x1f, 0x7d, 0x39, 0xd0, 0x63, 0x28, 0x4a, 0x5e, 0x08, 0x24, 0x17, 0xe0, 0x57, 0x28, 0x67, 0x9b,
	0xb2, 0x9e, 0x57, 0xf7, 0xa1, 0x73, 0x73, 0x9b, 0x24, 0xcc, 0x16, 0xc9, 0x31, 0xfc, 0xab, 0x50,
	0x3e, 0x8b, 0x65, 0xe0, 0x86, 0x4e, 0xdc, 0x67, 0x46, 0xe9, 0x9f, 0xa8, 0xaa, 0x94, 0x2f, 0xd4,
	0x4f, 0x38, 0x71, 0x97, 0x59, 0xa9, 0x52, 0x08, 0x3a, 0x0f, 0x3a, 0xa1, 0xb3, 0x2f, 0x45, 0x46,
	0x53, 0xa6, 0xfc, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x25, 0xf2, 0xd0, 0x25, 0x91, 0x02, 0x00,
	0x00,
}
