// Code generated by protoc-gen-go. DO NOT EDIT.
// source: srv/proto/common/common.service.proto

package common_service_srv

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	grpc "google.golang.org/grpc"
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

type CallRequest struct {
	Action               string              `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
	Header               map[string]string   `protobuf:"bytes,2,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Body                 map[string]*any.Any `protobuf:"bytes,3,rep,name=body,proto3" json:"body,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *CallRequest) Reset()         { *m = CallRequest{} }
func (m *CallRequest) String() string { return proto.CompactTextString(m) }
func (*CallRequest) ProtoMessage()    {}
func (*CallRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e745c4b5f95f8fb7, []int{0}
}

func (m *CallRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallRequest.Unmarshal(m, b)
}
func (m *CallRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallRequest.Marshal(b, m, deterministic)
}
func (m *CallRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallRequest.Merge(m, src)
}
func (m *CallRequest) XXX_Size() int {
	return xxx_messageInfo_CallRequest.Size(m)
}
func (m *CallRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CallRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CallRequest proto.InternalMessageInfo

func (m *CallRequest) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *CallRequest) GetHeader() map[string]string {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *CallRequest) GetBody() map[string]*any.Any {
	if m != nil {
		return m.Body
	}
	return nil
}

type CallResponse struct {
	Status               int32               `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg                  string              `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Content              map[string]*any.Any `protobuf:"bytes,3,rep,name=content,proto3" json:"content,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *CallResponse) Reset()         { *m = CallResponse{} }
func (m *CallResponse) String() string { return proto.CompactTextString(m) }
func (*CallResponse) ProtoMessage()    {}
func (*CallResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e745c4b5f95f8fb7, []int{1}
}

func (m *CallResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallResponse.Unmarshal(m, b)
}
func (m *CallResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallResponse.Marshal(b, m, deterministic)
}
func (m *CallResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallResponse.Merge(m, src)
}
func (m *CallResponse) XXX_Size() int {
	return xxx_messageInfo_CallResponse.Size(m)
}
func (m *CallResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CallResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CallResponse proto.InternalMessageInfo

func (m *CallResponse) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *CallResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *CallResponse) GetContent() map[string]*any.Any {
	if m != nil {
		return m.Content
	}
	return nil
}

type TypeString struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TypeString) Reset()         { *m = TypeString{} }
func (m *TypeString) String() string { return proto.CompactTextString(m) }
func (*TypeString) ProtoMessage()    {}
func (*TypeString) Descriptor() ([]byte, []int) {
	return fileDescriptor_e745c4b5f95f8fb7, []int{2}
}

func (m *TypeString) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TypeString.Unmarshal(m, b)
}
func (m *TypeString) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TypeString.Marshal(b, m, deterministic)
}
func (m *TypeString) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TypeString.Merge(m, src)
}
func (m *TypeString) XXX_Size() int {
	return xxx_messageInfo_TypeString.Size(m)
}
func (m *TypeString) XXX_DiscardUnknown() {
	xxx_messageInfo_TypeString.DiscardUnknown(m)
}

var xxx_messageInfo_TypeString proto.InternalMessageInfo

func (m *TypeString) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type TypeBool struct {
	Value                bool     `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TypeBool) Reset()         { *m = TypeBool{} }
func (m *TypeBool) String() string { return proto.CompactTextString(m) }
func (*TypeBool) ProtoMessage()    {}
func (*TypeBool) Descriptor() ([]byte, []int) {
	return fileDescriptor_e745c4b5f95f8fb7, []int{3}
}

func (m *TypeBool) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TypeBool.Unmarshal(m, b)
}
func (m *TypeBool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TypeBool.Marshal(b, m, deterministic)
}
func (m *TypeBool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TypeBool.Merge(m, src)
}
func (m *TypeBool) XXX_Size() int {
	return xxx_messageInfo_TypeBool.Size(m)
}
func (m *TypeBool) XXX_DiscardUnknown() {
	xxx_messageInfo_TypeBool.DiscardUnknown(m)
}

var xxx_messageInfo_TypeBool proto.InternalMessageInfo

func (m *TypeBool) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

type TypeInt64 struct {
	Value                int64    `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TypeInt64) Reset()         { *m = TypeInt64{} }
func (m *TypeInt64) String() string { return proto.CompactTextString(m) }
func (*TypeInt64) ProtoMessage()    {}
func (*TypeInt64) Descriptor() ([]byte, []int) {
	return fileDescriptor_e745c4b5f95f8fb7, []int{4}
}

func (m *TypeInt64) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TypeInt64.Unmarshal(m, b)
}
func (m *TypeInt64) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TypeInt64.Marshal(b, m, deterministic)
}
func (m *TypeInt64) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TypeInt64.Merge(m, src)
}
func (m *TypeInt64) XXX_Size() int {
	return xxx_messageInfo_TypeInt64.Size(m)
}
func (m *TypeInt64) XXX_DiscardUnknown() {
	xxx_messageInfo_TypeInt64.DiscardUnknown(m)
}

var xxx_messageInfo_TypeInt64 proto.InternalMessageInfo

func (m *TypeInt64) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type TypeUint64 struct {
	Value                uint64   `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TypeUint64) Reset()         { *m = TypeUint64{} }
func (m *TypeUint64) String() string { return proto.CompactTextString(m) }
func (*TypeUint64) ProtoMessage()    {}
func (*TypeUint64) Descriptor() ([]byte, []int) {
	return fileDescriptor_e745c4b5f95f8fb7, []int{5}
}

func (m *TypeUint64) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TypeUint64.Unmarshal(m, b)
}
func (m *TypeUint64) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TypeUint64.Marshal(b, m, deterministic)
}
func (m *TypeUint64) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TypeUint64.Merge(m, src)
}
func (m *TypeUint64) XXX_Size() int {
	return xxx_messageInfo_TypeUint64.Size(m)
}
func (m *TypeUint64) XXX_DiscardUnknown() {
	xxx_messageInfo_TypeUint64.DiscardUnknown(m)
}

var xxx_messageInfo_TypeUint64 proto.InternalMessageInfo

func (m *TypeUint64) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type TypeSint64 struct {
	Value                int64    `protobuf:"zigzag64,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TypeSint64) Reset()         { *m = TypeSint64{} }
func (m *TypeSint64) String() string { return proto.CompactTextString(m) }
func (*TypeSint64) ProtoMessage()    {}
func (*TypeSint64) Descriptor() ([]byte, []int) {
	return fileDescriptor_e745c4b5f95f8fb7, []int{6}
}

func (m *TypeSint64) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TypeSint64.Unmarshal(m, b)
}
func (m *TypeSint64) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TypeSint64.Marshal(b, m, deterministic)
}
func (m *TypeSint64) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TypeSint64.Merge(m, src)
}
func (m *TypeSint64) XXX_Size() int {
	return xxx_messageInfo_TypeSint64.Size(m)
}
func (m *TypeSint64) XXX_DiscardUnknown() {
	xxx_messageInfo_TypeSint64.DiscardUnknown(m)
}

var xxx_messageInfo_TypeSint64 proto.InternalMessageInfo

func (m *TypeSint64) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type TypeSint32 struct {
	Value                int32    `protobuf:"zigzag32,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TypeSint32) Reset()         { *m = TypeSint32{} }
func (m *TypeSint32) String() string { return proto.CompactTextString(m) }
func (*TypeSint32) ProtoMessage()    {}
func (*TypeSint32) Descriptor() ([]byte, []int) {
	return fileDescriptor_e745c4b5f95f8fb7, []int{7}
}

func (m *TypeSint32) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TypeSint32.Unmarshal(m, b)
}
func (m *TypeSint32) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TypeSint32.Marshal(b, m, deterministic)
}
func (m *TypeSint32) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TypeSint32.Merge(m, src)
}
func (m *TypeSint32) XXX_Size() int {
	return xxx_messageInfo_TypeSint32.Size(m)
}
func (m *TypeSint32) XXX_DiscardUnknown() {
	xxx_messageInfo_TypeSint32.DiscardUnknown(m)
}

var xxx_messageInfo_TypeSint32 proto.InternalMessageInfo

func (m *TypeSint32) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type TypeInt32 struct {
	Value                int32    `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TypeInt32) Reset()         { *m = TypeInt32{} }
func (m *TypeInt32) String() string { return proto.CompactTextString(m) }
func (*TypeInt32) ProtoMessage()    {}
func (*TypeInt32) Descriptor() ([]byte, []int) {
	return fileDescriptor_e745c4b5f95f8fb7, []int{8}
}

func (m *TypeInt32) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TypeInt32.Unmarshal(m, b)
}
func (m *TypeInt32) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TypeInt32.Marshal(b, m, deterministic)
}
func (m *TypeInt32) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TypeInt32.Merge(m, src)
}
func (m *TypeInt32) XXX_Size() int {
	return xxx_messageInfo_TypeInt32.Size(m)
}
func (m *TypeInt32) XXX_DiscardUnknown() {
	xxx_messageInfo_TypeInt32.DiscardUnknown(m)
}

var xxx_messageInfo_TypeInt32 proto.InternalMessageInfo

func (m *TypeInt32) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type TypeUint32 struct {
	Value                uint32   `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TypeUint32) Reset()         { *m = TypeUint32{} }
func (m *TypeUint32) String() string { return proto.CompactTextString(m) }
func (*TypeUint32) ProtoMessage()    {}
func (*TypeUint32) Descriptor() ([]byte, []int) {
	return fileDescriptor_e745c4b5f95f8fb7, []int{9}
}

func (m *TypeUint32) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TypeUint32.Unmarshal(m, b)
}
func (m *TypeUint32) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TypeUint32.Marshal(b, m, deterministic)
}
func (m *TypeUint32) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TypeUint32.Merge(m, src)
}
func (m *TypeUint32) XXX_Size() int {
	return xxx_messageInfo_TypeUint32.Size(m)
}
func (m *TypeUint32) XXX_DiscardUnknown() {
	xxx_messageInfo_TypeUint32.DiscardUnknown(m)
}

var xxx_messageInfo_TypeUint32 proto.InternalMessageInfo

func (m *TypeUint32) GetValue() uint32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type TypeDouble struct {
	Value                float64  `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TypeDouble) Reset()         { *m = TypeDouble{} }
func (m *TypeDouble) String() string { return proto.CompactTextString(m) }
func (*TypeDouble) ProtoMessage()    {}
func (*TypeDouble) Descriptor() ([]byte, []int) {
	return fileDescriptor_e745c4b5f95f8fb7, []int{10}
}

func (m *TypeDouble) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TypeDouble.Unmarshal(m, b)
}
func (m *TypeDouble) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TypeDouble.Marshal(b, m, deterministic)
}
func (m *TypeDouble) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TypeDouble.Merge(m, src)
}
func (m *TypeDouble) XXX_Size() int {
	return xxx_messageInfo_TypeDouble.Size(m)
}
func (m *TypeDouble) XXX_DiscardUnknown() {
	xxx_messageInfo_TypeDouble.DiscardUnknown(m)
}

var xxx_messageInfo_TypeDouble proto.InternalMessageInfo

func (m *TypeDouble) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type TypeFloat struct {
	Value                float32  `protobuf:"fixed32,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TypeFloat) Reset()         { *m = TypeFloat{} }
func (m *TypeFloat) String() string { return proto.CompactTextString(m) }
func (*TypeFloat) ProtoMessage()    {}
func (*TypeFloat) Descriptor() ([]byte, []int) {
	return fileDescriptor_e745c4b5f95f8fb7, []int{11}
}

func (m *TypeFloat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TypeFloat.Unmarshal(m, b)
}
func (m *TypeFloat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TypeFloat.Marshal(b, m, deterministic)
}
func (m *TypeFloat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TypeFloat.Merge(m, src)
}
func (m *TypeFloat) XXX_Size() int {
	return xxx_messageInfo_TypeFloat.Size(m)
}
func (m *TypeFloat) XXX_DiscardUnknown() {
	xxx_messageInfo_TypeFloat.DiscardUnknown(m)
}

var xxx_messageInfo_TypeFloat proto.InternalMessageInfo

func (m *TypeFloat) GetValue() float32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type TypeFixed32 struct {
	Value                uint32   `protobuf:"fixed32,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TypeFixed32) Reset()         { *m = TypeFixed32{} }
func (m *TypeFixed32) String() string { return proto.CompactTextString(m) }
func (*TypeFixed32) ProtoMessage()    {}
func (*TypeFixed32) Descriptor() ([]byte, []int) {
	return fileDescriptor_e745c4b5f95f8fb7, []int{12}
}

func (m *TypeFixed32) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TypeFixed32.Unmarshal(m, b)
}
func (m *TypeFixed32) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TypeFixed32.Marshal(b, m, deterministic)
}
func (m *TypeFixed32) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TypeFixed32.Merge(m, src)
}
func (m *TypeFixed32) XXX_Size() int {
	return xxx_messageInfo_TypeFixed32.Size(m)
}
func (m *TypeFixed32) XXX_DiscardUnknown() {
	xxx_messageInfo_TypeFixed32.DiscardUnknown(m)
}

var xxx_messageInfo_TypeFixed32 proto.InternalMessageInfo

func (m *TypeFixed32) GetValue() uint32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type TypeFixed64 struct {
	Value                uint64   `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TypeFixed64) Reset()         { *m = TypeFixed64{} }
func (m *TypeFixed64) String() string { return proto.CompactTextString(m) }
func (*TypeFixed64) ProtoMessage()    {}
func (*TypeFixed64) Descriptor() ([]byte, []int) {
	return fileDescriptor_e745c4b5f95f8fb7, []int{13}
}

func (m *TypeFixed64) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TypeFixed64.Unmarshal(m, b)
}
func (m *TypeFixed64) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TypeFixed64.Marshal(b, m, deterministic)
}
func (m *TypeFixed64) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TypeFixed64.Merge(m, src)
}
func (m *TypeFixed64) XXX_Size() int {
	return xxx_messageInfo_TypeFixed64.Size(m)
}
func (m *TypeFixed64) XXX_DiscardUnknown() {
	xxx_messageInfo_TypeFixed64.DiscardUnknown(m)
}

var xxx_messageInfo_TypeFixed64 proto.InternalMessageInfo

func (m *TypeFixed64) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type TypeSfixed32 struct {
	Value                int32    `protobuf:"fixed32,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TypeSfixed32) Reset()         { *m = TypeSfixed32{} }
func (m *TypeSfixed32) String() string { return proto.CompactTextString(m) }
func (*TypeSfixed32) ProtoMessage()    {}
func (*TypeSfixed32) Descriptor() ([]byte, []int) {
	return fileDescriptor_e745c4b5f95f8fb7, []int{14}
}

func (m *TypeSfixed32) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TypeSfixed32.Unmarshal(m, b)
}
func (m *TypeSfixed32) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TypeSfixed32.Marshal(b, m, deterministic)
}
func (m *TypeSfixed32) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TypeSfixed32.Merge(m, src)
}
func (m *TypeSfixed32) XXX_Size() int {
	return xxx_messageInfo_TypeSfixed32.Size(m)
}
func (m *TypeSfixed32) XXX_DiscardUnknown() {
	xxx_messageInfo_TypeSfixed32.DiscardUnknown(m)
}

var xxx_messageInfo_TypeSfixed32 proto.InternalMessageInfo

func (m *TypeSfixed32) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type TypeSfixed64 struct {
	Value                int64    `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TypeSfixed64) Reset()         { *m = TypeSfixed64{} }
func (m *TypeSfixed64) String() string { return proto.CompactTextString(m) }
func (*TypeSfixed64) ProtoMessage()    {}
func (*TypeSfixed64) Descriptor() ([]byte, []int) {
	return fileDescriptor_e745c4b5f95f8fb7, []int{15}
}

func (m *TypeSfixed64) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TypeSfixed64.Unmarshal(m, b)
}
func (m *TypeSfixed64) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TypeSfixed64.Marshal(b, m, deterministic)
}
func (m *TypeSfixed64) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TypeSfixed64.Merge(m, src)
}
func (m *TypeSfixed64) XXX_Size() int {
	return xxx_messageInfo_TypeSfixed64.Size(m)
}
func (m *TypeSfixed64) XXX_DiscardUnknown() {
	xxx_messageInfo_TypeSfixed64.DiscardUnknown(m)
}

var xxx_messageInfo_TypeSfixed64 proto.InternalMessageInfo

func (m *TypeSfixed64) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type TypeBytes struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TypeBytes) Reset()         { *m = TypeBytes{} }
func (m *TypeBytes) String() string { return proto.CompactTextString(m) }
func (*TypeBytes) ProtoMessage()    {}
func (*TypeBytes) Descriptor() ([]byte, []int) {
	return fileDescriptor_e745c4b5f95f8fb7, []int{16}
}

func (m *TypeBytes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TypeBytes.Unmarshal(m, b)
}
func (m *TypeBytes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TypeBytes.Marshal(b, m, deterministic)
}
func (m *TypeBytes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TypeBytes.Merge(m, src)
}
func (m *TypeBytes) XXX_Size() int {
	return xxx_messageInfo_TypeBytes.Size(m)
}
func (m *TypeBytes) XXX_DiscardUnknown() {
	xxx_messageInfo_TypeBytes.DiscardUnknown(m)
}

var xxx_messageInfo_TypeBytes proto.InternalMessageInfo

func (m *TypeBytes) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*CallRequest)(nil), "common.service.srv.CallRequest")
	proto.RegisterMapType((map[string]*any.Any)(nil), "common.service.srv.CallRequest.BodyEntry")
	proto.RegisterMapType((map[string]string)(nil), "common.service.srv.CallRequest.HeaderEntry")
	proto.RegisterType((*CallResponse)(nil), "common.service.srv.CallResponse")
	proto.RegisterMapType((map[string]*any.Any)(nil), "common.service.srv.CallResponse.ContentEntry")
	proto.RegisterType((*TypeString)(nil), "common.service.srv.TypeString")
	proto.RegisterType((*TypeBool)(nil), "common.service.srv.TypeBool")
	proto.RegisterType((*TypeInt64)(nil), "common.service.srv.TypeInt64")
	proto.RegisterType((*TypeUint64)(nil), "common.service.srv.TypeUint64")
	proto.RegisterType((*TypeSint64)(nil), "common.service.srv.TypeSint64")
	proto.RegisterType((*TypeSint32)(nil), "common.service.srv.TypeSint32")
	proto.RegisterType((*TypeInt32)(nil), "common.service.srv.TypeInt32")
	proto.RegisterType((*TypeUint32)(nil), "common.service.srv.TypeUint32")
	proto.RegisterType((*TypeDouble)(nil), "common.service.srv.TypeDouble")
	proto.RegisterType((*TypeFloat)(nil), "common.service.srv.TypeFloat")
	proto.RegisterType((*TypeFixed32)(nil), "common.service.srv.TypeFixed32")
	proto.RegisterType((*TypeFixed64)(nil), "common.service.srv.TypeFixed64")
	proto.RegisterType((*TypeSfixed32)(nil), "common.service.srv.TypeSfixed32")
	proto.RegisterType((*TypeSfixed64)(nil), "common.service.srv.TypeSfixed64")
	proto.RegisterType((*TypeBytes)(nil), "common.service.srv.TypeBytes")
}

func init() {
	proto.RegisterFile("srv/proto/common/common.service.proto", fileDescriptor_e745c4b5f95f8fb7)
}

var fileDescriptor_e745c4b5f95f8fb7 = []byte{
	// 468 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xcd, 0x8e, 0xd3, 0x30,
	0x14, 0x85, 0x49, 0xda, 0x66, 0xa6, 0x37, 0x45, 0x0c, 0xd6, 0x08, 0x95, 0x6c, 0x28, 0x01, 0xa4,
	0x01, 0x84, 0x2b, 0x75, 0x10, 0x02, 0x24, 0x16, 0xd3, 0xf2, 0xbb, 0x18, 0x09, 0x75, 0x98, 0x07,
	0x48, 0x1b, 0x37, 0x44, 0xa4, 0x76, 0x89, 0x9d, 0x08, 0x3f, 0x17, 0x8f, 0xc2, 0x0b, 0xa1, 0xd8,
	0x0e, 0x24, 0x6e, 0x50, 0x37, 0xac, 0xea, 0x9f, 0xef, 0x9e, 0x7b, 0x7c, 0x74, 0x1b, 0x78, 0xc4,
	0xf3, 0x72, 0xba, 0xcb, 0x99, 0x60, 0xd3, 0x35, 0xdb, 0x6e, 0x19, 0x35, 0x3f, 0x98, 0x93, 0xbc,
	0x4c, 0xd7, 0x04, 0xab, 0x3b, 0x84, 0xac, 0x53, 0x9e, 0x97, 0xc1, 0xdd, 0x84, 0xb1, 0x24, 0x23,
	0xba, 0x7a, 0x55, 0x6c, 0xa6, 0x11, 0x95, 0x1a, 0x0f, 0x7f, 0xba, 0xe0, 0x2f, 0xa2, 0x2c, 0x5b,
	0x92, 0xef, 0x05, 0xe1, 0x02, 0xdd, 0x01, 0x2f, 0x5a, 0x8b, 0x94, 0xd1, 0xb1, 0x33, 0x71, 0xce,
	0x86, 0x4b, 0xb3, 0x43, 0x0b, 0xf0, 0xbe, 0x92, 0x28, 0x26, 0xf9, 0xd8, 0x9d, 0xf4, 0xce, 0xfc,
	0xd9, 0x53, 0xbc, 0xdf, 0x07, 0x37, 0x84, 0xf0, 0x47, 0x45, 0xbf, 0xa3, 0x22, 0x97, 0x4b, 0x53,
	0x8a, 0xde, 0x40, 0x7f, 0xc5, 0x62, 0x39, 0xee, 0x29, 0x89, 0xc7, 0x87, 0x24, 0xe6, 0x2c, 0x96,
	0x5a, 0x40, 0x95, 0x05, 0xaf, 0xc0, 0x6f, 0xa8, 0xa2, 0x13, 0xe8, 0x7d, 0x23, 0xd2, 0xf8, 0xac,
	0x96, 0xe8, 0x14, 0x06, 0x65, 0x94, 0x15, 0x64, 0xec, 0xaa, 0x33, 0xbd, 0x79, 0xed, 0xbe, 0x74,
	0x82, 0x4b, 0x18, 0xfe, 0x51, 0xeb, 0x28, 0x7c, 0xd2, 0x2c, 0xf4, 0x67, 0xa7, 0x58, 0x07, 0x86,
	0xeb, 0xc0, 0xf0, 0x05, 0x95, 0x0d, 0xb9, 0xf0, 0x97, 0x03, 0x23, 0xed, 0x94, 0xef, 0x18, 0xe5,
	0xa4, 0x8a, 0x8d, 0x8b, 0x48, 0x14, 0x5c, 0xa9, 0x0e, 0x96, 0x66, 0x57, 0xb5, 0xda, 0xf2, 0xc4,
	0xf8, 0xa9, 0x96, 0xe8, 0x03, 0x1c, 0xad, 0x19, 0x15, 0x84, 0x0a, 0x13, 0xc3, 0xb3, 0x7f, 0xc7,
	0xa0, 0xc5, 0xf1, 0x42, 0xf3, 0x3a, 0x8a, 0xba, 0x3a, 0xf8, 0x0c, 0xa3, 0xe6, 0xc5, 0x7f, 0x78,
	0x55, 0x08, 0xf0, 0x45, 0xee, 0xc8, 0x95, 0xc8, 0x53, 0x9a, 0xfc, 0x0d, 0xd3, 0x69, 0x84, 0x19,
	0x4e, 0xe0, 0xb8, 0x62, 0xe6, 0x8c, 0x65, 0x6d, 0xe2, 0xb8, 0x26, 0xee, 0xc3, 0xb0, 0x22, 0x3e,
	0x51, 0xf1, 0xe2, 0x79, 0x1b, 0xe9, 0xd5, 0x88, 0x69, 0x74, 0x9d, 0xee, 0x33, 0x7d, 0x8b, 0xb9,
	0xea, 0x60, 0x50, 0x07, 0x73, 0x3e, 0x6b, 0x33, 0xb7, 0xf7, 0xed, 0xd8, 0xc8, 0xa0, 0xc3, 0x8e,
	0xcd, 0xdc, 0xb4, 0x98, 0xb7, 0xac, 0x58, 0x65, 0xa4, 0xcd, 0x38, 0x56, 0xab, 0xf7, 0x19, 0x8b,
	0x44, 0x1b, 0x71, 0x6b, 0xe4, 0x01, 0xf8, 0x0a, 0x49, 0x7f, 0x90, 0xd8, 0xee, 0x75, 0xd4, 0x05,
	0xd9, 0x6f, 0xf7, 0x6a, 0xe8, 0x21, 0x8c, 0xd4, 0xdb, 0x37, 0x5d, 0x52, 0xb7, 0x3a, 0x29, 0x5b,
	0xeb, 0xc4, 0x32, 0x3e, 0x97, 0x82, 0xf0, 0x36, 0x32, 0x32, 0xc8, 0xec, 0x1a, 0xfa, 0xd5, 0x4c,
	0xa2, 0x4b, 0xf0, 0x2e, 0xf4, 0x17, 0xe1, 0xde, 0x81, 0xbf, 0x6f, 0x30, 0x39, 0x34, 0xd8, 0xe1,
	0x8d, 0x95, 0xa7, 0x66, 0xf1, 0xfc, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xea, 0x16, 0xf6, 0xdc,
	0xdd, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CallClient is the client API for Call service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CallClient interface {
	Action(ctx context.Context, in *CallRequest, opts ...grpc.CallOption) (*CallResponse, error)
}

type callClient struct {
	cc *grpc.ClientConn
}

func NewCallClient(cc *grpc.ClientConn) CallClient {
	return &callClient{cc}
}

func (c *callClient) Action(ctx context.Context, in *CallRequest, opts ...grpc.CallOption) (*CallResponse, error) {
	out := new(CallResponse)
	err := c.cc.Invoke(ctx, "/common.service.srv.Call/Action", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CallServer is the server API for Call service.
type CallServer interface {
	Action(context.Context, *CallRequest) (*CallResponse, error)
}

func RegisterCallServer(s *grpc.Server, srv CallServer) {
	s.RegisterService(&_Call_serviceDesc, srv)
}

func _Call_Action_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CallServer).Action(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/common.service.srv.Call/Action",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CallServer).Action(ctx, req.(*CallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Call_serviceDesc = grpc.ServiceDesc{
	ServiceName: "common.service.srv.Call",
	HandlerType: (*CallServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Action",
			Handler:    _Call_Action_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "srv/proto/common/common.service.proto",
}