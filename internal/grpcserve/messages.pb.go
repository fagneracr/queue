// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messages/messages.proto

package grpcserve

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

type Varables struct {
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Varables) Reset()         { *m = Varables{} }
func (m *Varables) String() string { return proto.CompactTextString(m) }
func (*Varables) ProtoMessage()    {}
func (*Varables) Descriptor() ([]byte, []int) {
	return fileDescriptor_83994550f81e9f35, []int{0}
}

func (m *Varables) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Varables.Unmarshal(m, b)
}
func (m *Varables) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Varables.Marshal(b, m, deterministic)
}
func (m *Varables) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Varables.Merge(m, src)
}
func (m *Varables) XXX_Size() int {
	return xxx_messageInfo_Varables.Size(m)
}
func (m *Varables) XXX_DiscardUnknown() {
	xxx_messageInfo_Varables.DiscardUnknown(m)
}

var xxx_messageInfo_Varables proto.InternalMessageInfo

func (m *Varables) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Varables) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type QueueMSG struct {
	Name                 string      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Persistent           string      `protobuf:"bytes,2,opt,name=persistent,proto3" json:"persistent,omitempty"`
	Maxsize              int64       `protobuf:"varint,3,opt,name=maxsize,proto3" json:"maxsize,omitempty"`
	Ttl                  int32       `protobuf:"varint,4,opt,name=ttl,proto3" json:"ttl,omitempty"`
	Variables            []*Varables `protobuf:"bytes,5,rep,name=variables,proto3" json:"variables,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *QueueMSG) Reset()         { *m = QueueMSG{} }
func (m *QueueMSG) String() string { return proto.CompactTextString(m) }
func (*QueueMSG) ProtoMessage()    {}
func (*QueueMSG) Descriptor() ([]byte, []int) {
	return fileDescriptor_83994550f81e9f35, []int{1}
}

func (m *QueueMSG) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueueMSG.Unmarshal(m, b)
}
func (m *QueueMSG) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueueMSG.Marshal(b, m, deterministic)
}
func (m *QueueMSG) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueueMSG.Merge(m, src)
}
func (m *QueueMSG) XXX_Size() int {
	return xxx_messageInfo_QueueMSG.Size(m)
}
func (m *QueueMSG) XXX_DiscardUnknown() {
	xxx_messageInfo_QueueMSG.DiscardUnknown(m)
}

var xxx_messageInfo_QueueMSG proto.InternalMessageInfo

func (m *QueueMSG) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *QueueMSG) GetPersistent() string {
	if m != nil {
		return m.Persistent
	}
	return ""
}

func (m *QueueMSG) GetMaxsize() int64 {
	if m != nil {
		return m.Maxsize
	}
	return 0
}

func (m *QueueMSG) GetTtl() int32 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

func (m *QueueMSG) GetVariables() []*Varables {
	if m != nil {
		return m.Variables
	}
	return nil
}

type NewQRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewQRequest) Reset()         { *m = NewQRequest{} }
func (m *NewQRequest) String() string { return proto.CompactTextString(m) }
func (*NewQRequest) ProtoMessage()    {}
func (*NewQRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_83994550f81e9f35, []int{2}
}

func (m *NewQRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewQRequest.Unmarshal(m, b)
}
func (m *NewQRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewQRequest.Marshal(b, m, deterministic)
}
func (m *NewQRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewQRequest.Merge(m, src)
}
func (m *NewQRequest) XXX_Size() int {
	return xxx_messageInfo_NewQRequest.Size(m)
}
func (m *NewQRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NewQRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NewQRequest proto.InternalMessageInfo

type NewQResponse struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewQResponse) Reset()         { *m = NewQResponse{} }
func (m *NewQResponse) String() string { return proto.CompactTextString(m) }
func (*NewQResponse) ProtoMessage()    {}
func (*NewQResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_83994550f81e9f35, []int{3}
}

func (m *NewQResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewQResponse.Unmarshal(m, b)
}
func (m *NewQResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewQResponse.Marshal(b, m, deterministic)
}
func (m *NewQResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewQResponse.Merge(m, src)
}
func (m *NewQResponse) XXX_Size() int {
	return xxx_messageInfo_NewQResponse.Size(m)
}
func (m *NewQResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NewQResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NewQResponse proto.InternalMessageInfo

func (m *NewQResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Varables)(nil), "Varables")
	proto.RegisterType((*QueueMSG)(nil), "QueueMSG")
	proto.RegisterType((*NewQRequest)(nil), "NewQRequest")
	proto.RegisterType((*NewQResponse)(nil), "NewQResponse")
}

func init() { proto.RegisterFile("messages/messages.proto", fileDescriptor_83994550f81e9f35) }

var fileDescriptor_83994550f81e9f35 = []byte{
	// 268 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xcd, 0x4a, 0xc4, 0x30,
	0x10, 0xc7, 0xad, 0x6d, 0xb5, 0x9d, 0x75, 0x41, 0x82, 0x60, 0x50, 0x94, 0x92, 0x8b, 0xbd, 0xd8,
	0x42, 0x7d, 0x03, 0x2f, 0x9e, 0x14, 0xda, 0x05, 0x0f, 0xde, 0xb2, 0xcb, 0x50, 0x8a, 0xfd, 0xda,
	0x4c, 0x5a, 0x3f, 0x5e, 0xc3, 0x17, 0x96, 0x64, 0x37, 0xec, 0xde, 0x7e, 0xff, 0x99, 0xc9, 0xf0,
	0xcb, 0xc0, 0x75, 0x87, 0x44, 0xb2, 0x46, 0xca, 0x1d, 0x64, 0xa3, 0x1a, 0xf4, 0x20, 0x0a, 0x88,
	0xde, 0xa5, 0x92, 0xeb, 0x16, 0x89, 0x5d, 0x82, 0xff, 0x89, 0x3f, 0xfc, 0x34, 0xf1, 0xd2, 0xb8,
	0x32, 0xc8, 0xae, 0x20, 0x9c, 0x65, 0x3b, 0x21, 0xf7, 0x6c, 0x6d, 0x17, 0xc4, 0x9f, 0x07, 0x51,
	0x39, 0xe1, 0x84, 0xaf, 0xab, 0x17, 0xc6, 0x20, 0xe8, 0x65, 0xe7, 0x26, 0x2c, 0xb3, 0x7b, 0x80,
	0x11, 0x15, 0x35, 0xa4, 0xb1, 0xd7, 0xfb, 0x7d, 0x47, 0x15, 0xc6, 0xe1, 0xbc, 0x93, 0xdf, 0xd4,
	0xfc, 0x22, 0xf7, 0x13, 0x2f, 0xf5, 0x2b, 0x17, 0x8d, 0x82, 0xd6, 0x2d, 0x0f, 0x12, 0x2f, 0x0d,
	0x2b, 0x83, 0xec, 0x01, 0xe2, 0x59, 0xaa, 0xc6, 0x1a, 0xf2, 0x30, 0xf1, 0xd3, 0x45, 0x11, 0x67,
	0x4e, 0xb9, 0x3a, 0xf4, 0xc4, 0x12, 0x16, 0x6f, 0xf8, 0x55, 0x56, 0xb8, 0x9d, 0x90, 0xb4, 0x48,
	0xe0, 0x62, 0x17, 0x69, 0x1c, 0x7a, 0xb2, 0x9b, 0x3b, 0xaa, 0xf7, 0x9a, 0x06, 0x8b, 0x0c, 0xa2,
	0x72, 0x85, 0x6a, 0x6e, 0x36, 0xc8, 0x04, 0x04, 0x66, 0x9a, 0xc5, 0x99, 0xfb, 0xd8, 0xcd, 0x32,
	0x3b, 0x7e, 0x2f, 0x4e, 0x9e, 0xef, 0x3e, 0x6e, 0xeb, 0xe1, 0x71, 0x6b, 0xfa, 0x79, 0xd3, 0x6b,
	0x54, 0xbd, 0x6c, 0xf3, 0x5a, 0x8d, 0x1b, 0x42, 0x35, 0xe3, 0xfa, 0xcc, 0x1e, 0xf4, 0xe9, 0x3f,
	0x00, 0x00, 0xff, 0xff, 0xbe, 0xda, 0x8b, 0xa9, 0x6b, 0x01, 0x00, 0x00,
}
