// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Qot_GetTradeDate/Qot_GetTradeDate.proto

package Qot_GetTradeDate

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/jerryharbour/futugg/pb/Common"
import _ "github.com/jerryharbour/futugg/pb/Qot_Common"

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
	Market               *int32   `protobuf:"varint,1,req,name=market" json:"market,omitempty"`
	BeginTime            *string  `protobuf:"bytes,2,req,name=beginTime" json:"beginTime,omitempty"`
	EndTime              *string  `protobuf:"bytes,3,req,name=endTime" json:"endTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2S) Reset()         { *m = C2S{} }
func (m *C2S) String() string { return proto.CompactTextString(m) }
func (*C2S) ProtoMessage()    {}
func (*C2S) Descriptor() ([]byte, []int) {
	return fileDescriptor_Qot_GetTradeDate_faeaffaa81cf0d97, []int{0}
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

func (m *C2S) GetMarket() int32 {
	if m != nil && m.Market != nil {
		return *m.Market
	}
	return 0
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

type TradeDate struct {
	Time                 *string  `protobuf:"bytes,1,req,name=time" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TradeDate) Reset()         { *m = TradeDate{} }
func (m *TradeDate) String() string { return proto.CompactTextString(m) }
func (*TradeDate) ProtoMessage()    {}
func (*TradeDate) Descriptor() ([]byte, []int) {
	return fileDescriptor_Qot_GetTradeDate_faeaffaa81cf0d97, []int{1}
}
func (m *TradeDate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TradeDate.Unmarshal(m, b)
}
func (m *TradeDate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TradeDate.Marshal(b, m, deterministic)
}
func (dst *TradeDate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TradeDate.Merge(dst, src)
}
func (m *TradeDate) XXX_Size() int {
	return xxx_messageInfo_TradeDate.Size(m)
}
func (m *TradeDate) XXX_DiscardUnknown() {
	xxx_messageInfo_TradeDate.DiscardUnknown(m)
}

var xxx_messageInfo_TradeDate proto.InternalMessageInfo

func (m *TradeDate) GetTime() string {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return ""
}

type S2C struct {
	TradeDateList        []*TradeDate `protobuf:"bytes,1,rep,name=tradeDateList" json:"tradeDateList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *S2C) Reset()         { *m = S2C{} }
func (m *S2C) String() string { return proto.CompactTextString(m) }
func (*S2C) ProtoMessage()    {}
func (*S2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_Qot_GetTradeDate_faeaffaa81cf0d97, []int{2}
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

func (m *S2C) GetTradeDateList() []*TradeDate {
	if m != nil {
		return m.TradeDateList
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
	return fileDescriptor_Qot_GetTradeDate_faeaffaa81cf0d97, []int{3}
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
	return fileDescriptor_Qot_GetTradeDate_faeaffaa81cf0d97, []int{4}
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
	proto.RegisterType((*C2S)(nil), "Qot_GetTradeDate.C2S")
	proto.RegisterType((*TradeDate)(nil), "Qot_GetTradeDate.TradeDate")
	proto.RegisterType((*S2C)(nil), "Qot_GetTradeDate.S2C")
	proto.RegisterType((*Request)(nil), "Qot_GetTradeDate.Request")
	proto.RegisterType((*Response)(nil), "Qot_GetTradeDate.Response")
}

func init() {
	proto.RegisterFile("Qot_GetTradeDate/Qot_GetTradeDate.proto", fileDescriptor_Qot_GetTradeDate_faeaffaa81cf0d97)
}

var fileDescriptor_Qot_GetTradeDate_faeaffaa81cf0d97 = []byte{
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x41, 0x4f, 0xbb, 0x40,
	0x10, 0xc5, 0xb3, 0xd0, 0xfe, 0xf9, 0x33, 0x8d, 0x89, 0x59, 0xa3, 0x21, 0xad, 0x51, 0xc2, 0xa5,
	0x5c, 0x6c, 0x9b, 0x8d, 0x27, 0x6f, 0x66, 0x4d, 0xf4, 0xa0, 0x07, 0x17, 0x3c, 0x1b, 0x2c, 0x93,
	0x86, 0x18, 0x58, 0xdc, 0x5d, 0x0f, 0x7e, 0x00, 0xbf, 0xb7, 0xd9, 0x05, 0xda, 0xd8, 0x7a, 0x62,
	0xde, 0x7b, 0x3f, 0x98, 0x37, 0xc0, 0xfc, 0x59, 0x9a, 0xd7, 0x7b, 0x34, 0xb9, 0x2a, 0x4a, 0xbc,
	0x2b, 0x0c, 0x2e, 0xf7, 0x8d, 0x45, 0xab, 0xa4, 0x91, 0xf4, 0x78, 0xdf, 0x9f, 0x9e, 0x70, 0x59,
	0xd7, 0xb2, 0x59, 0x76, 0x8f, 0x0e, 0x9b, 0xce, 0x2c, 0xd6, 0x07, 0xbb, 0xb1, 0x0b, 0x93, 0x17,
	0xf0, 0x39, 0xcb, 0xe8, 0x19, 0xfc, 0xab, 0x0b, 0xf5, 0x8e, 0x26, 0x22, 0xb1, 0x97, 0x8e, 0x45,
	0xaf, 0xe8, 0x39, 0x84, 0x6f, 0xb8, 0xa9, 0x9a, 0xbc, 0xaa, 0x31, 0xf2, 0x62, 0x2f, 0x0d, 0xc5,
	0xce, 0xa0, 0x11, 0x04, 0xd8, 0x94, 0x2e, 0xf3, 0x5d, 0x36, 0xc8, 0xe4, 0x12, 0xc2, 0x6d, 0x2b,
	0x4a, 0x61, 0x64, 0x2c, 0x43, 0x1c, 0xe3, 0xe6, 0xe4, 0x01, 0xfc, 0x8c, 0x71, 0x7a, 0x0b, 0x47,
	0x66, 0xe0, 0x1e, 0x2b, 0x6d, 0xd7, 0xfb, 0xe9, 0x84, 0xcd, 0x16, 0x07, 0x27, 0x6f, 0x27, 0xf1,
	0xfb, 0x8d, 0x84, 0x41, 0x20, 0xf0, 0xe3, 0x13, 0xb5, 0xa1, 0x73, 0xf0, 0xd7, 0x4c, 0xbb, 0x3d,
	0x13, 0x76, 0x7a, 0xf8, 0x0d, 0xce, 0x32, 0x61, 0x89, 0xe4, 0x9b, 0xc0, 0x7f, 0x81, 0xba, 0x95,
	0x8d, 0x46, 0x7a, 0x01, 0x81, 0x42, 0x93, 0x7f, 0xb5, 0x5d, 0xc3, 0xf1, 0xcd, 0xe8, 0xea, 0x7a,
	0xb5, 0x12, 0x83, 0x69, 0xff, 0x8d, 0x42, 0xf3, 0xa4, 0x37, 0x91, 0x17, 0x93, 0x34, 0x14, 0xbd,
	0x72, 0xd7, 0x2b, 0xc5, 0x65, 0x69, 0xaf, 0x27, 0xe9, 0x58, 0x0c, 0xd2, 0xf6, 0xd0, 0x6c, 0x1d,
	0x8d, 0x62, 0xf2, 0x77, 0x8f, 0x8c, 0x71, 0x61, 0x89, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x10,
	0xd8, 0x77, 0x2b, 0xeb, 0x01, 0x00, 0x00,
}
