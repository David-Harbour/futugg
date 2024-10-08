// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Trd_GetAccList/Trd_GetAccList.proto

package Trd_GetAccList

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/jerryharbour/futugg/pb/Common"
import Trd_Common "github.com/jerryharbour/futugg/pb/Trd_Common"

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
	UserID               *uint64  `protobuf:"varint,1,req,name=userID" json:"userID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2S) Reset()         { *m = C2S{} }
func (m *C2S) String() string { return proto.CompactTextString(m) }
func (*C2S) ProtoMessage()    {}
func (*C2S) Descriptor() ([]byte, []int) {
	return fileDescriptor_Trd_GetAccList_dc0bd2d154a8826f, []int{0}
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

func (m *C2S) GetUserID() uint64 {
	if m != nil && m.UserID != nil {
		return *m.UserID
	}
	return 0
}

type S2C struct {
	AccList              []*Trd_Common.TrdAcc `protobuf:"bytes,1,rep,name=accList" json:"accList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *S2C) Reset()         { *m = S2C{} }
func (m *S2C) String() string { return proto.CompactTextString(m) }
func (*S2C) ProtoMessage()    {}
func (*S2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_Trd_GetAccList_dc0bd2d154a8826f, []int{1}
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

func (m *S2C) GetAccList() []*Trd_Common.TrdAcc {
	if m != nil {
		return m.AccList
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
	return fileDescriptor_Trd_GetAccList_dc0bd2d154a8826f, []int{2}
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
	// 以下3个字段每条协议都有，注释说明在InitConnect.proto中
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
	return fileDescriptor_Trd_GetAccList_dc0bd2d154a8826f, []int{3}
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
	proto.RegisterType((*C2S)(nil), "Trd_GetAccList.C2S")
	proto.RegisterType((*S2C)(nil), "Trd_GetAccList.S2C")
	proto.RegisterType((*Request)(nil), "Trd_GetAccList.Request")
	proto.RegisterType((*Response)(nil), "Trd_GetAccList.Response")
}

func init() {
	proto.RegisterFile("Trd_GetAccList/Trd_GetAccList.proto", fileDescriptor_Trd_GetAccList_dc0bd2d154a8826f)
}

var fileDescriptor_Trd_GetAccList_dc0bd2d154a8826f = []byte{
	// 253 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0xc9, 0xda, 0x59, 0x7d, 0x03, 0x0f, 0x19, 0x48, 0x98, 0x28, 0xa1, 0x22, 0xf4, 0xa0,
	0x5d, 0x89, 0x9e, 0xbc, 0x8d, 0x08, 0x22, 0xe8, 0xe5, 0xb5, 0x77, 0x91, 0xf4, 0x21, 0x1e, 0xb6,
	0xd4, 0x24, 0x3b, 0x78, 0xf6, 0x1f, 0x97, 0xac, 0x2d, 0x3a, 0xf1, 0x94, 0x7c, 0xef, 0xf7, 0xe5,
	0xcb, 0xf7, 0xe0, 0xa2, 0x71, 0xed, 0xcb, 0x03, 0x85, 0x95, 0x31, 0x4f, 0xef, 0x3e, 0x2c, 0xf7,
	0x65, 0xd9, 0x39, 0x1b, 0x2c, 0x3f, 0xde, 0x9f, 0x2e, 0xe6, 0xda, 0xae, 0xd7, 0x76, 0xb3, 0xec,
	0x8f, 0xde, 0xb4, 0x38, 0x8d, 0xa6, 0x01, 0xfc, 0x5c, 0x7b, 0x98, 0x9f, 0x41, 0xa2, 0x55, 0xcd,
	0x4f, 0xe0, 0x60, 0xeb, 0xc9, 0x3d, 0xde, 0x0b, 0x26, 0x27, 0x45, 0x8a, 0x83, 0xca, 0x6f, 0x20,
	0xa9, 0x95, 0xe6, 0x57, 0x90, 0xbd, 0xf6, 0x5f, 0x08, 0x26, 0x93, 0x62, 0xa6, 0x78, 0xf9, 0x2b,
	0xa9, 0x71, 0xed, 0xca, 0x18, 0x1c, 0x2d, 0x79, 0x05, 0x19, 0xd2, 0xc7, 0x96, 0x7c, 0xe0, 0x97,
	0x90, 0x18, 0xe5, 0x77, 0xa1, 0x33, 0x35, 0x2f, 0xff, 0x2c, 0xa1, 0x55, 0x8d, 0x91, 0xe7, 0x5f,
	0x0c, 0x0e, 0x91, 0x7c, 0x67, 0x37, 0x9e, 0xf8, 0x39, 0x64, 0x8e, 0x42, 0xf3, 0xd9, 0xd1, 0xee,
	0xdd, 0xf4, 0x2e, 0xbd, 0xbe, 0xad, 0x2a, 0x1c, 0x87, 0xb1, 0xab, 0xa3, 0xf0, 0xec, 0xdf, 0xc4,
	0x44, 0xb2, 0xe2, 0x08, 0x07, 0xc5, 0x05, 0x64, 0xe4, 0x9c, 0xb6, 0x2d, 0x89, 0x44, 0xb2, 0x62,
	0x8a, 0xa3, 0x8c, 0x2d, 0xbc, 0x32, 0x22, 0x95, 0xec, 0xbf, 0x16, 0xb5, 0xd2, 0x18, 0xf9, 0x77,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x74, 0xcd, 0x93, 0xa9, 0x73, 0x01, 0x00, 0x00,
}
