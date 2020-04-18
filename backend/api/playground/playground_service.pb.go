// Code generated by protoc-gen-go. DO NOT EDIT.
// source: playground_service.proto

package playground

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type CodeRequest struct {
	Body                 string   `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CodeRequest) Reset()         { *m = CodeRequest{} }
func (m *CodeRequest) String() string { return proto.CompactTextString(m) }
func (*CodeRequest) ProtoMessage()    {}
func (*CodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f8460c736b806de, []int{0}
}

func (m *CodeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CodeRequest.Unmarshal(m, b)
}
func (m *CodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CodeRequest.Marshal(b, m, deterministic)
}
func (m *CodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CodeRequest.Merge(m, src)
}
func (m *CodeRequest) XXX_Size() int {
	return xxx_messageInfo_CodeRequest.Size(m)
}
func (m *CodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CodeRequest proto.InternalMessageInfo

func (m *CodeRequest) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

type CommonRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommonRequest) Reset()         { *m = CommonRequest{} }
func (m *CommonRequest) String() string { return proto.CompactTextString(m) }
func (*CommonRequest) ProtoMessage()    {}
func (*CommonRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f8460c736b806de, []int{1}
}

func (m *CommonRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommonRequest.Unmarshal(m, b)
}
func (m *CommonRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommonRequest.Marshal(b, m, deterministic)
}
func (m *CommonRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonRequest.Merge(m, src)
}
func (m *CommonRequest) XXX_Size() int {
	return xxx_messageInfo_CommonRequest.Size(m)
}
func (m *CommonRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CommonRequest proto.InternalMessageInfo

func (m *CommonRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type FormatCodeResponse struct {
	FormattedCode        string   `protobuf:"bytes,1,opt,name=formatted_code,json=formattedCode,proto3" json:"formatted_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FormatCodeResponse) Reset()         { *m = FormatCodeResponse{} }
func (m *FormatCodeResponse) String() string { return proto.CompactTextString(m) }
func (*FormatCodeResponse) ProtoMessage()    {}
func (*FormatCodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f8460c736b806de, []int{2}
}

func (m *FormatCodeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FormatCodeResponse.Unmarshal(m, b)
}
func (m *FormatCodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FormatCodeResponse.Marshal(b, m, deterministic)
}
func (m *FormatCodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FormatCodeResponse.Merge(m, src)
}
func (m *FormatCodeResponse) XXX_Size() int {
	return xxx_messageInfo_FormatCodeResponse.Size(m)
}
func (m *FormatCodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FormatCodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FormatCodeResponse proto.InternalMessageInfo

func (m *FormatCodeResponse) GetFormattedCode() string {
	if m != nil {
		return m.FormattedCode
	}
	return ""
}

type ShareCodeResponse struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShareCodeResponse) Reset()         { *m = ShareCodeResponse{} }
func (m *ShareCodeResponse) String() string { return proto.CompactTextString(m) }
func (*ShareCodeResponse) ProtoMessage()    {}
func (*ShareCodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f8460c736b806de, []int{3}
}

func (m *ShareCodeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShareCodeResponse.Unmarshal(m, b)
}
func (m *ShareCodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShareCodeResponse.Marshal(b, m, deterministic)
}
func (m *ShareCodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShareCodeResponse.Merge(m, src)
}
func (m *ShareCodeResponse) XXX_Size() int {
	return xxx_messageInfo_ShareCodeResponse.Size(m)
}
func (m *ShareCodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ShareCodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ShareCodeResponse proto.InternalMessageInfo

func (m *ShareCodeResponse) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type RunResponse struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Output               string   `protobuf:"bytes,2,opt,name=output,proto3" json:"output,omitempty"`
	Error                string   `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	RunTime              float64  `protobuf:"fixed64,4,opt,name=run_time,json=runTime,proto3" json:"run_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RunResponse) Reset()         { *m = RunResponse{} }
func (m *RunResponse) String() string { return proto.CompactTextString(m) }
func (*RunResponse) ProtoMessage()    {}
func (*RunResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f8460c736b806de, []int{4}
}

func (m *RunResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunResponse.Unmarshal(m, b)
}
func (m *RunResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunResponse.Marshal(b, m, deterministic)
}
func (m *RunResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunResponse.Merge(m, src)
}
func (m *RunResponse) XXX_Size() int {
	return xxx_messageInfo_RunResponse.Size(m)
}
func (m *RunResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RunResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RunResponse proto.InternalMessageInfo

func (m *RunResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *RunResponse) GetOutput() string {
	if m != nil {
		return m.Output
	}
	return ""
}

func (m *RunResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *RunResponse) GetRunTime() float64 {
	if m != nil {
		return m.RunTime
	}
	return 0
}

type GetCodeByShareResponse struct {
	Body                 string   `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	ShareCode            string   `protobuf:"bytes,2,opt,name=share_code,json=shareCode,proto3" json:"share_code,omitempty"`
	CreatedAt            string   `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCodeByShareResponse) Reset()         { *m = GetCodeByShareResponse{} }
func (m *GetCodeByShareResponse) String() string { return proto.CompactTextString(m) }
func (*GetCodeByShareResponse) ProtoMessage()    {}
func (*GetCodeByShareResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f8460c736b806de, []int{5}
}

func (m *GetCodeByShareResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCodeByShareResponse.Unmarshal(m, b)
}
func (m *GetCodeByShareResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCodeByShareResponse.Marshal(b, m, deterministic)
}
func (m *GetCodeByShareResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCodeByShareResponse.Merge(m, src)
}
func (m *GetCodeByShareResponse) XXX_Size() int {
	return xxx_messageInfo_GetCodeByShareResponse.Size(m)
}
func (m *GetCodeByShareResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCodeByShareResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetCodeByShareResponse proto.InternalMessageInfo

func (m *GetCodeByShareResponse) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *GetCodeByShareResponse) GetShareCode() string {
	if m != nil {
		return m.ShareCode
	}
	return ""
}

func (m *GetCodeByShareResponse) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *GetCodeByShareResponse) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type EmptyRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyRequest) Reset()         { *m = EmptyRequest{} }
func (m *EmptyRequest) String() string { return proto.CompactTextString(m) }
func (*EmptyRequest) ProtoMessage()    {}
func (*EmptyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f8460c736b806de, []int{6}
}

func (m *EmptyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyRequest.Unmarshal(m, b)
}
func (m *EmptyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyRequest.Marshal(b, m, deterministic)
}
func (m *EmptyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyRequest.Merge(m, src)
}
func (m *EmptyRequest) XXX_Size() int {
	return xxx_messageInfo_EmptyRequest.Size(m)
}
func (m *EmptyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyRequest proto.InternalMessageInfo

type PingResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingResponse) Reset()         { *m = PingResponse{} }
func (m *PingResponse) String() string { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()    {}
func (*PingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f8460c736b806de, []int{7}
}

func (m *PingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingResponse.Unmarshal(m, b)
}
func (m *PingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingResponse.Marshal(b, m, deterministic)
}
func (m *PingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingResponse.Merge(m, src)
}
func (m *PingResponse) XXX_Size() int {
	return xxx_messageInfo_PingResponse.Size(m)
}
func (m *PingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PingResponse proto.InternalMessageInfo

func (m *PingResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*CodeRequest)(nil), "playground.CodeRequest")
	proto.RegisterType((*CommonRequest)(nil), "playground.CommonRequest")
	proto.RegisterType((*FormatCodeResponse)(nil), "playground.FormatCodeResponse")
	proto.RegisterType((*ShareCodeResponse)(nil), "playground.ShareCodeResponse")
	proto.RegisterType((*RunResponse)(nil), "playground.RunResponse")
	proto.RegisterType((*GetCodeByShareResponse)(nil), "playground.GetCodeByShareResponse")
	proto.RegisterType((*EmptyRequest)(nil), "playground.EmptyRequest")
	proto.RegisterType((*PingResponse)(nil), "playground.PingResponse")
}

func init() { proto.RegisterFile("playground_service.proto", fileDescriptor_2f8460c736b806de) }

var fileDescriptor_2f8460c736b806de = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0x4d, 0x8f, 0xda, 0x30,
	0x10, 0x55, 0xd2, 0x14, 0x9a, 0x01, 0x22, 0xd5, 0xaa, 0x68, 0x40, 0xa2, 0xa5, 0x96, 0xaa, 0x72,
	0xe2, 0xd0, 0xde, 0xca, 0x89, 0x22, 0xda, 0x63, 0x51, 0xda, 0x3b, 0x0a, 0xd8, 0xcb, 0x46, 0xda,
	0xd8, 0x59, 0x7f, 0xac, 0x94, 0x9f, 0xb0, 0x7f, 0x73, 0x7f, 0xc9, 0x2a, 0x8e, 0x13, 0x9c, 0x5d,
	0xb8, 0x65, 0xde, 0x7b, 0x33, 0x9e, 0x8f, 0x17, 0x88, 0x8b, 0xbb, 0xb4, 0x3c, 0x09, 0xae, 0x19,
	0xd9, 0x4b, 0x2a, 0x1e, 0xb2, 0x23, 0x5d, 0x16, 0x82, 0x2b, 0x8e, 0xe0, 0xcc, 0xe0, 0x2f, 0x30,
	0xd8, 0x70, 0x42, 0x13, 0x7a, 0xaf, 0xa9, 0x54, 0x08, 0x41, 0x70, 0xe0, 0xa4, 0x8c, 0xbd, 0xb9,
	0xb7, 0x08, 0x13, 0xf3, 0x8d, 0x3f, 0xc3, 0x68, 0xc3, 0xf3, 0x9c, 0xb3, 0x46, 0x14, 0x81, 0x9f,
	0x11, 0x2b, 0xf1, 0x33, 0x82, 0x57, 0x80, 0x7e, 0x73, 0x91, 0xa7, 0xaa, 0xae, 0x24, 0x0b, 0xce,
	0x24, 0x45, 0x5f, 0x21, 0xba, 0x31, 0xa8, 0xa2, 0x64, 0x7f, 0xe4, 0x84, 0xda, 0x8c, 0x51, 0x8b,
	0x56, 0x72, 0xfc, 0x0d, 0xde, 0xff, 0xbb, 0x4d, 0x05, 0xed, 0xe4, 0x22, 0x08, 0x9c, 0x0c, 0xf3,
	0x8d, 0x19, 0x0c, 0x12, 0xcd, 0x5a, 0xc9, 0x18, 0x7a, 0x52, 0xa5, 0x4a, 0x4b, 0x2b, 0xb2, 0x51,
	0x85, 0x73, 0xad, 0x0a, 0xad, 0x62, 0xbf, 0xc6, 0xeb, 0x08, 0x7d, 0x80, 0xb7, 0x54, 0x08, 0x2e,
	0xe2, 0x37, 0x06, 0xae, 0x03, 0x34, 0x81, 0x77, 0x42, 0xb3, 0xbd, 0xca, 0x72, 0x1a, 0x07, 0x73,
	0x6f, 0xe1, 0x25, 0x7d, 0xa1, 0xd9, 0xff, 0x2c, 0xa7, 0xf8, 0xd1, 0x83, 0xf1, 0x1f, 0x6a, 0x66,
	0xfa, 0x55, 0x9a, 0x16, 0xdd, 0xf6, 0x5e, 0x6e, 0x09, 0xcd, 0x00, 0x64, 0x25, 0xaa, 0x47, 0xad,
	0xdf, 0x0e, 0x65, 0x33, 0x59, 0x45, 0x1f, 0x05, 0x4d, 0xab, 0x5d, 0xa4, 0xca, 0xf6, 0x10, 0x5a,
	0x64, 0xad, 0x2a, 0x5a, 0x17, 0xa4, 0xa1, 0x83, 0x9a, 0xb6, 0xc8, 0x5a, 0xe1, 0x08, 0x86, 0xdb,
	0xbc, 0x50, 0xa5, 0xbd, 0x00, 0x5e, 0xc0, 0x70, 0x97, 0xb1, 0x53, 0xdb, 0x50, 0x0c, 0xfd, 0x9c,
	0x4a, 0x99, 0x9e, 0x9a, 0x95, 0x35, 0xe1, 0xf7, 0x27, 0x1f, 0x60, 0xd7, 0x9e, 0x1b, 0xfd, 0x84,
	0xa0, 0x4a, 0x44, 0xf1, 0xf2, 0xec, 0x81, 0xa5, 0x5b, 0x7a, 0xda, 0x61, 0x3a, 0x8f, 0xac, 0xa0,
	0x9f, 0x68, 0x66, 0xa6, 0xf9, 0xe8, 0x8a, 0x1c, 0xff, 0x4c, 0x3b, 0x84, 0x7b, 0xae, 0x2d, 0xc0,
	0xd9, 0x23, 0xd7, 0xf3, 0x3f, 0xb9, 0xc4, 0x05, 0x53, 0x6d, 0x20, 0x6c, 0xdd, 0x72, 0xbd, 0xca,
	0xcc, 0x25, 0x5e, 0xbb, 0xeb, 0x2f, 0x44, 0xdd, 0xc3, 0xa2, 0x49, 0xb7, 0x92, 0x63, 0xf6, 0x29,
	0x76, 0xa9, 0xcb, 0x7e, 0x38, 0xf4, 0xcc, 0x7f, 0xf5, 0xe3, 0x39, 0x00, 0x00, 0xff, 0xff, 0x12,
	0x8e, 0x3c, 0xb9, 0x73, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PlaygroundClient is the client API for Playground service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PlaygroundClient interface {
	Ping(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*PingResponse, error)
	RunCode(ctx context.Context, in *CodeRequest, opts ...grpc.CallOption) (*RunResponse, error)
	FormatCode(ctx context.Context, in *CodeRequest, opts ...grpc.CallOption) (*FormatCodeResponse, error)
	ShareCode(ctx context.Context, in *CodeRequest, opts ...grpc.CallOption) (*ShareCodeResponse, error)
	GetCodeByShare(ctx context.Context, in *CommonRequest, opts ...grpc.CallOption) (*GetCodeByShareResponse, error)
}

type playgroundClient struct {
	cc *grpc.ClientConn
}

func NewPlaygroundClient(cc *grpc.ClientConn) PlaygroundClient {
	return &playgroundClient{cc}
}

func (c *playgroundClient) Ping(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/playground.Playground/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playgroundClient) RunCode(ctx context.Context, in *CodeRequest, opts ...grpc.CallOption) (*RunResponse, error) {
	out := new(RunResponse)
	err := c.cc.Invoke(ctx, "/playground.Playground/RunCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playgroundClient) FormatCode(ctx context.Context, in *CodeRequest, opts ...grpc.CallOption) (*FormatCodeResponse, error) {
	out := new(FormatCodeResponse)
	err := c.cc.Invoke(ctx, "/playground.Playground/FormatCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playgroundClient) ShareCode(ctx context.Context, in *CodeRequest, opts ...grpc.CallOption) (*ShareCodeResponse, error) {
	out := new(ShareCodeResponse)
	err := c.cc.Invoke(ctx, "/playground.Playground/ShareCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playgroundClient) GetCodeByShare(ctx context.Context, in *CommonRequest, opts ...grpc.CallOption) (*GetCodeByShareResponse, error) {
	out := new(GetCodeByShareResponse)
	err := c.cc.Invoke(ctx, "/playground.Playground/GetCodeByShare", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlaygroundServer is the server API for Playground service.
type PlaygroundServer interface {
	Ping(context.Context, *EmptyRequest) (*PingResponse, error)
	RunCode(context.Context, *CodeRequest) (*RunResponse, error)
	FormatCode(context.Context, *CodeRequest) (*FormatCodeResponse, error)
	ShareCode(context.Context, *CodeRequest) (*ShareCodeResponse, error)
	GetCodeByShare(context.Context, *CommonRequest) (*GetCodeByShareResponse, error)
}

// UnimplementedPlaygroundServer can be embedded to have forward compatible implementations.
type UnimplementedPlaygroundServer struct {
}

func (*UnimplementedPlaygroundServer) Ping(ctx context.Context, req *EmptyRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedPlaygroundServer) RunCode(ctx context.Context, req *CodeRequest) (*RunResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunCode not implemented")
}
func (*UnimplementedPlaygroundServer) FormatCode(ctx context.Context, req *CodeRequest) (*FormatCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FormatCode not implemented")
}
func (*UnimplementedPlaygroundServer) ShareCode(ctx context.Context, req *CodeRequest) (*ShareCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShareCode not implemented")
}
func (*UnimplementedPlaygroundServer) GetCodeByShare(ctx context.Context, req *CommonRequest) (*GetCodeByShareResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCodeByShare not implemented")
}

func RegisterPlaygroundServer(s *grpc.Server, srv PlaygroundServer) {
	s.RegisterService(&_Playground_serviceDesc, srv)
}

func _Playground_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaygroundServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/playground.Playground/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaygroundServer).Ping(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Playground_RunCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaygroundServer).RunCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/playground.Playground/RunCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaygroundServer).RunCode(ctx, req.(*CodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Playground_FormatCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaygroundServer).FormatCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/playground.Playground/FormatCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaygroundServer).FormatCode(ctx, req.(*CodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Playground_ShareCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaygroundServer).ShareCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/playground.Playground/ShareCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaygroundServer).ShareCode(ctx, req.(*CodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Playground_GetCodeByShare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlaygroundServer).GetCodeByShare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/playground.Playground/GetCodeByShare",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlaygroundServer).GetCodeByShare(ctx, req.(*CommonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Playground_serviceDesc = grpc.ServiceDesc{
	ServiceName: "playground.Playground",
	HandlerType: (*PlaygroundServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Playground_Ping_Handler,
		},
		{
			MethodName: "RunCode",
			Handler:    _Playground_RunCode_Handler,
		},
		{
			MethodName: "FormatCode",
			Handler:    _Playground_FormatCode_Handler,
		},
		{
			MethodName: "ShareCode",
			Handler:    _Playground_ShareCode_Handler,
		},
		{
			MethodName: "GetCodeByShare",
			Handler:    _Playground_GetCodeByShare_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "playground_service.proto",
}
