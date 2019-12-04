// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bean/tyto.proto

package bean

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

//链路
type Trace struct {
	//链路标识
	TraceId string `protobuf:"bytes,1,opt,name=traceId,proto3" json:"traceId,omitempty"`
	//对应的请求或者入口的标识
	SecondId string `protobuf:"bytes,2,opt,name=secondId,proto3" json:"secondId,omitempty"`
	//链路生成的时间
	TraceSTime int64 `protobuf:"varint,3,opt,name=traceSTime,proto3" json:"traceSTime,omitempty"`
	//链路完成的时间
	TraceETime int64 `protobuf:"varint,4,opt,name=traceETime,proto3" json:"traceETime,omitempty"`
	//所有链路生命周期内的相关节点集合
	Spans map[string]*Span `protobuf:"bytes,5,rep,name=spans,proto3" json:"spans,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	//用户标识
	UserId string `protobuf:"bytes,6,opt,name=userId,proto3" json:"userId,omitempty"`
	//用户名称
	UserName string `protobuf:"bytes,7,opt,name=userName,proto3" json:"userName,omitempty"`
	//生命周期的时长
	Duration int64 `protobuf:"varint,8,opt,name=duration,proto3" json:"duration,omitempty"`
	//是否记录到日志中
	Logging bool `protobuf:"varint,9,opt,name=logging,proto3" json:"logging,omitempty"`
	//描述
	Desc string `protobuf:"bytes,10,opt,name=desc,proto3" json:"desc,omitempty"`
	//平台标识
	Platform string `protobuf:"bytes,11,opt,name=platform,proto3" json:"platform,omitempty"`
	//超时时间 单位为秒
	Timeout int64 `protobuf:"varint,12,opt,name=timeout,proto3" json:"timeout,omitempty"`
	//flush时需要记录的此链路中的所有的span标识和tag标识
	Subs                 *Subs    `protobuf:"bytes,13,opt,name=subs,proto3" json:"subs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Trace) Reset()         { *m = Trace{} }
func (m *Trace) String() string { return proto.CompactTextString(m) }
func (*Trace) ProtoMessage()    {}
func (*Trace) Descriptor() ([]byte, []int) {
	return fileDescriptor_65eadb3a877d1afd, []int{0}
}

func (m *Trace) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Trace.Unmarshal(m, b)
}
func (m *Trace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Trace.Marshal(b, m, deterministic)
}
func (m *Trace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Trace.Merge(m, src)
}
func (m *Trace) XXX_Size() int {
	return xxx_messageInfo_Trace.Size(m)
}
func (m *Trace) XXX_DiscardUnknown() {
	xxx_messageInfo_Trace.DiscardUnknown(m)
}

var xxx_messageInfo_Trace proto.InternalMessageInfo

func (m *Trace) GetTraceId() string {
	if m != nil {
		return m.TraceId
	}
	return ""
}

func (m *Trace) GetSecondId() string {
	if m != nil {
		return m.SecondId
	}
	return ""
}

func (m *Trace) GetTraceSTime() int64 {
	if m != nil {
		return m.TraceSTime
	}
	return 0
}

func (m *Trace) GetTraceETime() int64 {
	if m != nil {
		return m.TraceETime
	}
	return 0
}

func (m *Trace) GetSpans() map[string]*Span {
	if m != nil {
		return m.Spans
	}
	return nil
}

func (m *Trace) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Trace) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *Trace) GetDuration() int64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *Trace) GetLogging() bool {
	if m != nil {
		return m.Logging
	}
	return false
}

func (m *Trace) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *Trace) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

func (m *Trace) GetTimeout() int64 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *Trace) GetSubs() *Subs {
	if m != nil {
		return m.Subs
	}
	return nil
}

type Subs struct {
	Spans                map[string]bool `protobuf:"bytes,1,rep,name=spans,proto3" json:"spans,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Tags                 map[string]bool `protobuf:"bytes,2,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Subs) Reset()         { *m = Subs{} }
func (m *Subs) String() string { return proto.CompactTextString(m) }
func (*Subs) ProtoMessage()    {}
func (*Subs) Descriptor() ([]byte, []int) {
	return fileDescriptor_65eadb3a877d1afd, []int{1}
}

func (m *Subs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Subs.Unmarshal(m, b)
}
func (m *Subs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Subs.Marshal(b, m, deterministic)
}
func (m *Subs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Subs.Merge(m, src)
}
func (m *Subs) XXX_Size() int {
	return xxx_messageInfo_Subs.Size(m)
}
func (m *Subs) XXX_DiscardUnknown() {
	xxx_messageInfo_Subs.DiscardUnknown(m)
}

var xxx_messageInfo_Subs proto.InternalMessageInfo

func (m *Subs) GetSpans() map[string]bool {
	if m != nil {
		return m.Spans
	}
	return nil
}

func (m *Subs) GetTags() map[string]bool {
	if m != nil {
		return m.Tags
	}
	return nil
}

type BaseResponse struct {
	//返回码
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	//返回信息
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BaseResponse) Reset()         { *m = BaseResponse{} }
func (m *BaseResponse) String() string { return proto.CompactTextString(m) }
func (*BaseResponse) ProtoMessage()    {}
func (*BaseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_65eadb3a877d1afd, []int{2}
}

func (m *BaseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BaseResponse.Unmarshal(m, b)
}
func (m *BaseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BaseResponse.Marshal(b, m, deterministic)
}
func (m *BaseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseResponse.Merge(m, src)
}
func (m *BaseResponse) XXX_Size() int {
	return xxx_messageInfo_BaseResponse.Size(m)
}
func (m *BaseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BaseResponse proto.InternalMessageInfo

func (m *BaseResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *BaseResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

//节点
type Span struct {
	//节点标识
	SpanId string `protobuf:"bytes,1,opt,name=spanId,proto3" json:"spanId,omitempty"`
	//对应的方法或者代码块的标识
	SecondId string `protobuf:"bytes,2,opt,name=secondId,proto3" json:"secondId,omitempty"`
	//节点创建的时间
	SpanSTime int64 `protobuf:"varint,3,opt,name=spanSTime,proto3" json:"spanSTime,omitempty"`
	//节点结束的时间
	SpanETime int64 `protobuf:"varint,4,opt,name=spanETime,proto3" json:"spanETime,omitempty"`
	//此节点生命周期内描述事件详情的标记集合
	Tags []*Tag `protobuf:"bytes,5,rep,name=Tags,proto3" json:"Tags,omitempty"`
	//关联的链路标识
	TraceId string `protobuf:"bytes,6,opt,name=TraceId,proto3" json:"TraceId,omitempty"`
	//生命周期的时长
	Duration             int64    `protobuf:"varint,7,opt,name=duration,proto3" json:"duration,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Span) Reset()         { *m = Span{} }
func (m *Span) String() string { return proto.CompactTextString(m) }
func (*Span) ProtoMessage()    {}
func (*Span) Descriptor() ([]byte, []int) {
	return fileDescriptor_65eadb3a877d1afd, []int{3}
}

func (m *Span) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Span.Unmarshal(m, b)
}
func (m *Span) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Span.Marshal(b, m, deterministic)
}
func (m *Span) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Span.Merge(m, src)
}
func (m *Span) XXX_Size() int {
	return xxx_messageInfo_Span.Size(m)
}
func (m *Span) XXX_DiscardUnknown() {
	xxx_messageInfo_Span.DiscardUnknown(m)
}

var xxx_messageInfo_Span proto.InternalMessageInfo

func (m *Span) GetSpanId() string {
	if m != nil {
		return m.SpanId
	}
	return ""
}

func (m *Span) GetSecondId() string {
	if m != nil {
		return m.SecondId
	}
	return ""
}

func (m *Span) GetSpanSTime() int64 {
	if m != nil {
		return m.SpanSTime
	}
	return 0
}

func (m *Span) GetSpanETime() int64 {
	if m != nil {
		return m.SpanETime
	}
	return 0
}

func (m *Span) GetTags() []*Tag {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Span) GetTraceId() string {
	if m != nil {
		return m.TraceId
	}
	return ""
}

func (m *Span) GetDuration() int64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

//用来描述节点生命周期内任何事件的标记
type Tag struct {
	TraceId              string   `protobuf:"bytes,1,opt,name=traceId,proto3" json:"traceId,omitempty"`
	SpanId               string   `protobuf:"bytes,2,opt,name=spanId,proto3" json:"spanId,omitempty"`
	TagId                string   `protobuf:"bytes,3,opt,name=tagId,proto3" json:"tagId,omitempty"`
	Desc                 string   `protobuf:"bytes,4,opt,name=desc,proto3" json:"desc,omitempty"`
	Time                 int64    `protobuf:"varint,5,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tag) Reset()         { *m = Tag{} }
func (m *Tag) String() string { return proto.CompactTextString(m) }
func (*Tag) ProtoMessage()    {}
func (*Tag) Descriptor() ([]byte, []int) {
	return fileDescriptor_65eadb3a877d1afd, []int{4}
}

func (m *Tag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tag.Unmarshal(m, b)
}
func (m *Tag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tag.Marshal(b, m, deterministic)
}
func (m *Tag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tag.Merge(m, src)
}
func (m *Tag) XXX_Size() int {
	return xxx_messageInfo_Tag.Size(m)
}
func (m *Tag) XXX_DiscardUnknown() {
	xxx_messageInfo_Tag.DiscardUnknown(m)
}

var xxx_messageInfo_Tag proto.InternalMessageInfo

func (m *Tag) GetTraceId() string {
	if m != nil {
		return m.TraceId
	}
	return ""
}

func (m *Tag) GetSpanId() string {
	if m != nil {
		return m.SpanId
	}
	return ""
}

func (m *Tag) GetTagId() string {
	if m != nil {
		return m.TagId
	}
	return ""
}

func (m *Tag) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *Tag) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

type Ping struct {
	Ping                 bool     `protobuf:"varint,1,opt,name=ping,proto3" json:"ping,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ping) Reset()         { *m = Ping{} }
func (m *Ping) String() string { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()    {}
func (*Ping) Descriptor() ([]byte, []int) {
	return fileDescriptor_65eadb3a877d1afd, []int{5}
}

func (m *Ping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ping.Unmarshal(m, b)
}
func (m *Ping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ping.Marshal(b, m, deterministic)
}
func (m *Ping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ping.Merge(m, src)
}
func (m *Ping) XXX_Size() int {
	return xxx_messageInfo_Ping.Size(m)
}
func (m *Ping) XXX_DiscardUnknown() {
	xxx_messageInfo_Ping.DiscardUnknown(m)
}

var xxx_messageInfo_Ping proto.InternalMessageInfo

func (m *Ping) GetPing() bool {
	if m != nil {
		return m.Ping
	}
	return false
}

type Pong struct {
	Pong                 bool     `protobuf:"varint,1,opt,name=pong,proto3" json:"pong,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pong) Reset()         { *m = Pong{} }
func (m *Pong) String() string { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()    {}
func (*Pong) Descriptor() ([]byte, []int) {
	return fileDescriptor_65eadb3a877d1afd, []int{6}
}

func (m *Pong) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pong.Unmarshal(m, b)
}
func (m *Pong) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pong.Marshal(b, m, deterministic)
}
func (m *Pong) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pong.Merge(m, src)
}
func (m *Pong) XXX_Size() int {
	return xxx_messageInfo_Pong.Size(m)
}
func (m *Pong) XXX_DiscardUnknown() {
	xxx_messageInfo_Pong.DiscardUnknown(m)
}

var xxx_messageInfo_Pong proto.InternalMessageInfo

func (m *Pong) GetPong() bool {
	if m != nil {
		return m.Pong
	}
	return false
}

func init() {
	proto.RegisterType((*Trace)(nil), "bean.Trace")
	proto.RegisterMapType((map[string]*Span)(nil), "bean.Trace.SpansEntry")
	proto.RegisterType((*Subs)(nil), "bean.subs")
	proto.RegisterMapType((map[string]bool)(nil), "bean.subs.SpansEntry")
	proto.RegisterMapType((map[string]bool)(nil), "bean.subs.TagsEntry")
	proto.RegisterType((*BaseResponse)(nil), "bean.baseResponse")
	proto.RegisterType((*Span)(nil), "bean.Span")
	proto.RegisterType((*Tag)(nil), "bean.Tag")
	proto.RegisterType((*Ping)(nil), "bean.ping")
	proto.RegisterType((*Pong)(nil), "bean.pong")
}

func init() { proto.RegisterFile("bean/tyto.proto", fileDescriptor_65eadb3a877d1afd) }

var fileDescriptor_65eadb3a877d1afd = []byte{
	// 622 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0xae, 0x13, 0x3b, 0x89, 0x27, 0xfd, 0xe9, 0x87, 0x56, 0xa5, 0x5a, 0x59, 0x50, 0x59, 0x3e,
	0x05, 0x95, 0xa6, 0x52, 0x38, 0x50, 0x21, 0x2e, 0x95, 0x28, 0x22, 0x17, 0x84, 0xdc, 0xbc, 0xc0,
	0x26, 0xbb, 0x2c, 0x51, 0x13, 0xaf, 0xe5, 0xb5, 0x2b, 0xf5, 0x3d, 0x78, 0x25, 0x4e, 0x5c, 0x78,
	0x24, 0x34, 0xe3, 0x7f, 0x5b, 0x89, 0x96, 0x72, 0xf2, 0x7c, 0x33, 0xb3, 0x3b, 0xdf, 0xcc, 0x7c,
	0x6b, 0xf8, 0x7f, 0xad, 0x44, 0x76, 0x5e, 0xde, 0x95, 0x66, 0x9e, 0x17, 0xa6, 0x34, 0xcc, 0x47,
	0x47, 0xf2, 0x63, 0x08, 0xc1, 0xaa, 0x10, 0x1b, 0xc5, 0x38, 0x8c, 0x4b, 0x34, 0x96, 0x92, 0x7b,
	0xb1, 0x37, 0x0b, 0xd3, 0x16, 0xb2, 0x08, 0x26, 0x56, 0x6d, 0x4c, 0x26, 0x97, 0x92, 0x0f, 0x28,
	0xd4, 0x61, 0x76, 0x02, 0x40, 0x69, 0xd7, 0xab, 0xed, 0x5e, 0xf1, 0x61, 0xec, 0xcd, 0x86, 0xa9,
	0xe3, 0xe9, 0xe2, 0x57, 0x14, 0xf7, 0x9d, 0x38, 0x79, 0xd8, 0x6b, 0x08, 0x6c, 0x2e, 0x32, 0xcb,
	0x83, 0x78, 0x38, 0x9b, 0x2e, 0x8e, 0xe7, 0xc8, 0x6a, 0x4e, 0x8c, 0xe6, 0xd7, 0x18, 0xb8, 0xca,
	0xca, 0xe2, 0x2e, 0xad, 0x93, 0xd8, 0x31, 0x8c, 0x2a, 0xab, 0x8a, 0xa5, 0xe4, 0x23, 0xe2, 0xd1,
	0x20, 0x64, 0x88, 0xd6, 0x67, 0xb1, 0x57, 0x7c, 0x5c, 0x33, 0x6c, 0x31, 0xc6, 0x64, 0x55, 0x88,
	0x72, 0x6b, 0x32, 0x3e, 0xa1, 0xfa, 0x1d, 0xc6, 0x9e, 0x77, 0x46, 0xeb, 0x6d, 0xa6, 0x79, 0x18,
	0x7b, 0xb3, 0x49, 0xda, 0x42, 0xc6, 0xc0, 0x97, 0xca, 0x6e, 0x38, 0xd0, 0x6d, 0x64, 0xe3, 0x4d,
	0xf9, 0x4e, 0x94, 0x5f, 0x4d, 0xb1, 0xe7, 0xd3, 0xba, 0x4a, 0x8b, 0x69, 0x7a, 0xdb, 0xbd, 0x32,
	0x55, 0xc9, 0x0f, 0xa9, 0x48, 0x0b, 0xd9, 0x09, 0xf8, 0xb6, 0x5a, 0x5b, 0xfe, 0x5f, 0xec, 0xcd,
	0xa6, 0x0b, 0xa8, 0x1b, 0x44, 0x4f, 0x4a, 0xfe, 0xe8, 0x03, 0x40, 0xdf, 0x28, 0x7b, 0x06, 0xc3,
	0x1b, 0x75, 0xd7, 0x6c, 0x00, 0x4d, 0x16, 0x43, 0x70, 0x2b, 0x76, 0x95, 0xa2, 0xd1, 0x77, 0x17,
	0xe0, 0x91, 0xb4, 0x0e, 0xbc, 0x1b, 0x5c, 0x78, 0xc9, 0x2f, 0xaf, 0x2e, 0xc3, 0x4e, 0xdb, 0x81,
	0x7a, 0x34, 0xd0, 0xe7, 0x7d, 0xbd, 0x3f, 0xcc, 0x73, 0x06, 0x7e, 0x29, 0xb4, 0xe5, 0x03, 0xca,
	0x3d, 0x72, 0x72, 0x57, 0x42, 0x37, 0xa9, 0x94, 0x11, 0x5d, 0xfc, 0x85, 0xe5, 0x91, 0xcb, 0x72,
	0xe2, 0x30, 0x8b, 0xde, 0x42, 0xd8, 0x5d, 0xf6, 0x2f, 0x07, 0x93, 0xf7, 0x70, 0xb8, 0x16, 0x56,
	0xa5, 0xca, 0xe6, 0x26, 0xb3, 0x0a, 0x57, 0xb2, 0x31, 0x52, 0xd1, 0xe1, 0x20, 0x25, 0x1b, 0xc7,
	0xbe, 0x57, 0xd6, 0x0a, 0xad, 0x1a, 0x65, 0xb6, 0x30, 0xf9, 0xe9, 0x81, 0x8f, 0x8c, 0x51, 0x33,
	0xd8, 0x6c, 0x27, 0xeb, 0x06, 0x3d, 0xaa, 0xea, 0x17, 0x10, 0x62, 0x96, 0x2b, 0xea, 0xde, 0xd1,
	0x46, 0x5d, 0x49, 0xf7, 0x0e, 0xf6, 0x12, 0x7c, 0xec, 0xb7, 0x11, 0x74, 0xd8, 0x08, 0x5a, 0xe8,
	0x94, 0xdc, 0xc8, 0x78, 0xd5, 0x3c, 0xb3, 0x5a, 0xc3, 0x2d, 0xbc, 0x27, 0xd4, 0xf1, 0x7d, 0xa1,
	0x26, 0x15, 0x0c, 0x57, 0x42, 0x3f, 0xf2, 0x46, 0xfb, 0x2e, 0x07, 0xf7, 0xba, 0x3c, 0x82, 0xa0,
	0x14, 0x7a, 0x29, 0xa9, 0x8b, 0x30, 0xad, 0x41, 0xa7, 0x6e, 0xdf, 0x51, 0x37, 0x03, 0x1f, 0x25,
	0xcb, 0x03, 0x2a, 0x4d, 0x76, 0x12, 0x81, 0x9f, 0x37, 0xaf, 0x01, 0xbf, 0x54, 0x74, 0x92, 0x92,
	0x4d, 0x31, 0xd3, 0xc4, 0x8c, 0x13, 0x33, 0x99, 0x5e, 0x64, 0x70, 0x48, 0x5d, 0x7d, 0x12, 0x99,
	0xdc, 0xa9, 0x82, 0x9d, 0xc1, 0xe4, 0x52, 0xca, 0xfa, 0x3f, 0x33, 0x75, 0x9e, 0x78, 0xc4, 0x6a,
	0xe0, 0xee, 0x39, 0x39, 0x60, 0xe7, 0x00, 0x1f, 0x77, 0x95, 0xfd, 0xf6, 0xd4, 0x03, 0x8b, 0xef,
	0x1e, 0x4c, 0x71, 0xd9, 0x6d, 0xbd, 0x53, 0x18, 0x5f, 0x4a, 0x49, 0xeb, 0x77, 0xde, 0xcb, 0x03,
	0xd5, 0xce, 0x20, 0xa4, 0x6a, 0x4f, 0x4c, 0x7f, 0x05, 0x23, 0xec, 0x45, 0x68, 0xd6, 0xef, 0xf6,
	0x01, 0x5a, 0x67, 0x10, 0xde, 0x28, 0x95, 0x5f, 0xee, 0xb6, 0xb7, 0x8a, 0xc5, 0xe0, 0x7f, 0xc1,
	0x59, 0x36, 0x15, 0x70, 0x86, 0x51, 0x6b, 0x9b, 0x4c, 0x27, 0x07, 0xeb, 0x11, 0xfd, 0x98, 0xdf,
	0xfc, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x8b, 0x42, 0xc7, 0x04, 0xab, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TraceHandlerClient is the client API for TraceHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TraceHandlerClient interface {
	AddTrace(ctx context.Context, in *Trace, opts ...grpc.CallOption) (*BaseResponse, error)
	FlushTrace(ctx context.Context, in *Trace, opts ...grpc.CallOption) (*BaseResponse, error)
}

type traceHandlerClient struct {
	cc *grpc.ClientConn
}

func NewTraceHandlerClient(cc *grpc.ClientConn) TraceHandlerClient {
	return &traceHandlerClient{cc}
}

func (c *traceHandlerClient) AddTrace(ctx context.Context, in *Trace, opts ...grpc.CallOption) (*BaseResponse, error) {
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, "/bean.TraceHandler/AddTrace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *traceHandlerClient) FlushTrace(ctx context.Context, in *Trace, opts ...grpc.CallOption) (*BaseResponse, error) {
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, "/bean.TraceHandler/FlushTrace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TraceHandlerServer is the server API for TraceHandler service.
type TraceHandlerServer interface {
	AddTrace(context.Context, *Trace) (*BaseResponse, error)
	FlushTrace(context.Context, *Trace) (*BaseResponse, error)
}

func RegisterTraceHandlerServer(s *grpc.Server, srv TraceHandlerServer) {
	s.RegisterService(&_TraceHandler_serviceDesc, srv)
}

func _TraceHandler_AddTrace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Trace)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TraceHandlerServer).AddTrace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bean.TraceHandler/AddTrace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TraceHandlerServer).AddTrace(ctx, req.(*Trace))
	}
	return interceptor(ctx, in, info, handler)
}

func _TraceHandler_FlushTrace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Trace)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TraceHandlerServer).FlushTrace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bean.TraceHandler/FlushTrace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TraceHandlerServer).FlushTrace(ctx, req.(*Trace))
	}
	return interceptor(ctx, in, info, handler)
}

var _TraceHandler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bean.TraceHandler",
	HandlerType: (*TraceHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddTrace",
			Handler:    _TraceHandler_AddTrace_Handler,
		},
		{
			MethodName: "FlushTrace",
			Handler:    _TraceHandler_FlushTrace_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bean/tyto.proto",
}

// SpanHandlerClient is the client API for SpanHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SpanHandlerClient interface {
	AddSpan(ctx context.Context, in *Span, opts ...grpc.CallOption) (*BaseResponse, error)
	FlushSpan(ctx context.Context, in *Span, opts ...grpc.CallOption) (*BaseResponse, error)
	AddTag(ctx context.Context, in *Tag, opts ...grpc.CallOption) (*BaseResponse, error)
}

type spanHandlerClient struct {
	cc *grpc.ClientConn
}

func NewSpanHandlerClient(cc *grpc.ClientConn) SpanHandlerClient {
	return &spanHandlerClient{cc}
}

func (c *spanHandlerClient) AddSpan(ctx context.Context, in *Span, opts ...grpc.CallOption) (*BaseResponse, error) {
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, "/bean.SpanHandler/AddSpan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spanHandlerClient) FlushSpan(ctx context.Context, in *Span, opts ...grpc.CallOption) (*BaseResponse, error) {
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, "/bean.SpanHandler/FlushSpan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spanHandlerClient) AddTag(ctx context.Context, in *Tag, opts ...grpc.CallOption) (*BaseResponse, error) {
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, "/bean.SpanHandler/AddTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SpanHandlerServer is the server API for SpanHandler service.
type SpanHandlerServer interface {
	AddSpan(context.Context, *Span) (*BaseResponse, error)
	FlushSpan(context.Context, *Span) (*BaseResponse, error)
	AddTag(context.Context, *Tag) (*BaseResponse, error)
}

func RegisterSpanHandlerServer(s *grpc.Server, srv SpanHandlerServer) {
	s.RegisterService(&_SpanHandler_serviceDesc, srv)
}

func _SpanHandler_AddSpan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Span)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpanHandlerServer).AddSpan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bean.SpanHandler/AddSpan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpanHandlerServer).AddSpan(ctx, req.(*Span))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpanHandler_FlushSpan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Span)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpanHandlerServer).FlushSpan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bean.SpanHandler/FlushSpan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpanHandlerServer).FlushSpan(ctx, req.(*Span))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpanHandler_AddTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Tag)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpanHandlerServer).AddTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bean.SpanHandler/AddTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpanHandlerServer).AddTag(ctx, req.(*Tag))
	}
	return interceptor(ctx, in, info, handler)
}

var _SpanHandler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bean.SpanHandler",
	HandlerType: (*SpanHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddSpan",
			Handler:    _SpanHandler_AddSpan_Handler,
		},
		{
			MethodName: "FlushSpan",
			Handler:    _SpanHandler_FlushSpan_Handler,
		},
		{
			MethodName: "AddTag",
			Handler:    _SpanHandler_AddTag_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bean/tyto.proto",
}

// KeepAliveClient is the client API for KeepAlive service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KeepAliveClient interface {
	Ping(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*Pong, error)
}

type keepAliveClient struct {
	cc *grpc.ClientConn
}

func NewKeepAliveClient(cc *grpc.ClientConn) KeepAliveClient {
	return &keepAliveClient{cc}
}

func (c *keepAliveClient) Ping(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*Pong, error) {
	out := new(Pong)
	err := c.cc.Invoke(ctx, "/bean.keepAlive/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeepAliveServer is the server API for KeepAlive service.
type KeepAliveServer interface {
	Ping(context.Context, *Ping) (*Pong, error)
}

func RegisterKeepAliveServer(s *grpc.Server, srv KeepAliveServer) {
	s.RegisterService(&_KeepAlive_serviceDesc, srv)
}

func _KeepAlive_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ping)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeepAliveServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bean.keepAlive/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeepAliveServer).Ping(ctx, req.(*Ping))
	}
	return interceptor(ctx, in, info, handler)
}

var _KeepAlive_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bean.keepAlive",
	HandlerType: (*KeepAliveServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _KeepAlive_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bean/tyto.proto",
}
