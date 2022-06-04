// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: denom/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryParamsRequest is the request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_40684369d1fc796f, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is the response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params defines the parameters of the module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_40684369d1fc796f, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

type QueryDenomAuthorityMetadataRequest struct {
	Denom string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty" yaml:"denom"`
}

func (m *QueryDenomAuthorityMetadataRequest) Reset()         { *m = QueryDenomAuthorityMetadataRequest{} }
func (m *QueryDenomAuthorityMetadataRequest) String() string { return proto.CompactTextString(m) }
func (*QueryDenomAuthorityMetadataRequest) ProtoMessage()    {}
func (*QueryDenomAuthorityMetadataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_40684369d1fc796f, []int{2}
}
func (m *QueryDenomAuthorityMetadataRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDenomAuthorityMetadataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDenomAuthorityMetadataRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDenomAuthorityMetadataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDenomAuthorityMetadataRequest.Merge(m, src)
}
func (m *QueryDenomAuthorityMetadataRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryDenomAuthorityMetadataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDenomAuthorityMetadataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDenomAuthorityMetadataRequest proto.InternalMessageInfo

func (m *QueryDenomAuthorityMetadataRequest) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

type QueryDenomAuthorityMetadataResponse struct {
	AuthorityMetadata DenomAuthorityMetadata `protobuf:"bytes,1,opt,name=authority_metadata,json=authorityMetadata,proto3" json:"authority_metadata" yaml:"authority_metadata"`
}

func (m *QueryDenomAuthorityMetadataResponse) Reset()         { *m = QueryDenomAuthorityMetadataResponse{} }
func (m *QueryDenomAuthorityMetadataResponse) String() string { return proto.CompactTextString(m) }
func (*QueryDenomAuthorityMetadataResponse) ProtoMessage()    {}
func (*QueryDenomAuthorityMetadataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_40684369d1fc796f, []int{3}
}
func (m *QueryDenomAuthorityMetadataResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDenomAuthorityMetadataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDenomAuthorityMetadataResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDenomAuthorityMetadataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDenomAuthorityMetadataResponse.Merge(m, src)
}
func (m *QueryDenomAuthorityMetadataResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryDenomAuthorityMetadataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDenomAuthorityMetadataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDenomAuthorityMetadataResponse proto.InternalMessageInfo

func (m *QueryDenomAuthorityMetadataResponse) GetAuthorityMetadata() DenomAuthorityMetadata {
	if m != nil {
		return m.AuthorityMetadata
	}
	return DenomAuthorityMetadata{}
}

type QueryDenomsFromCreatorRequest struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty" yaml:"creator"`
}

func (m *QueryDenomsFromCreatorRequest) Reset()         { *m = QueryDenomsFromCreatorRequest{} }
func (m *QueryDenomsFromCreatorRequest) String() string { return proto.CompactTextString(m) }
func (*QueryDenomsFromCreatorRequest) ProtoMessage()    {}
func (*QueryDenomsFromCreatorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_40684369d1fc796f, []int{4}
}
func (m *QueryDenomsFromCreatorRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDenomsFromCreatorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDenomsFromCreatorRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDenomsFromCreatorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDenomsFromCreatorRequest.Merge(m, src)
}
func (m *QueryDenomsFromCreatorRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryDenomsFromCreatorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDenomsFromCreatorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDenomsFromCreatorRequest proto.InternalMessageInfo

func (m *QueryDenomsFromCreatorRequest) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

type QueryDenomsFromCreatorResponse struct {
	Denoms []string `protobuf:"bytes,1,rep,name=denoms,proto3" json:"denoms,omitempty" yaml:"denoms"`
}

func (m *QueryDenomsFromCreatorResponse) Reset()         { *m = QueryDenomsFromCreatorResponse{} }
func (m *QueryDenomsFromCreatorResponse) String() string { return proto.CompactTextString(m) }
func (*QueryDenomsFromCreatorResponse) ProtoMessage()    {}
func (*QueryDenomsFromCreatorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_40684369d1fc796f, []int{5}
}
func (m *QueryDenomsFromCreatorResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDenomsFromCreatorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDenomsFromCreatorResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDenomsFromCreatorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDenomsFromCreatorResponse.Merge(m, src)
}
func (m *QueryDenomsFromCreatorResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryDenomsFromCreatorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDenomsFromCreatorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDenomsFromCreatorResponse proto.InternalMessageInfo

func (m *QueryDenomsFromCreatorResponse) GetDenoms() []string {
	if m != nil {
		return m.Denoms
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "gert.denom.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "gert.denom.QueryParamsResponse")
	proto.RegisterType((*QueryDenomAuthorityMetadataRequest)(nil), "gert.denom.QueryDenomAuthorityMetadataRequest")
	proto.RegisterType((*QueryDenomAuthorityMetadataResponse)(nil), "gert.denom.QueryDenomAuthorityMetadataResponse")
	proto.RegisterType((*QueryDenomsFromCreatorRequest)(nil), "gert.denom.QueryDenomsFromCreatorRequest")
	proto.RegisterType((*QueryDenomsFromCreatorResponse)(nil), "gert.denom.QueryDenomsFromCreatorResponse")
}

func init() { proto.RegisterFile("denom/query.proto", fileDescriptor_40684369d1fc796f) }

var fileDescriptor_40684369d1fc796f = []byte{
	// 522 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xcd, 0x6e, 0xd4, 0x30,
	0x10, 0xde, 0x50, 0xba, 0xa8, 0xe6, 0x47, 0xec, 0xb0, 0x42, 0x10, 0xd1, 0xa4, 0x18, 0xa9, 0xda,
	0x56, 0x55, 0x4c, 0xcb, 0x8d, 0x1b, 0x01, 0xc1, 0x01, 0x2a, 0x41, 0x8e, 0x5c, 0x2a, 0x6f, 0x6b,
	0x85, 0x95, 0x9a, 0x38, 0x8d, 0xbd, 0x88, 0xa8, 0xea, 0x85, 0x27, 0x40, 0x82, 0x1b, 0x4f, 0xc0,
	0x91, 0xb7, 0xe8, 0xb1, 0x12, 0x17, 0x4e, 0x11, 0xda, 0xe5, 0x09, 0xf6, 0x09, 0x50, 0xec, 0x29,
	0x0d, 0x24, 0xad, 0xf6, 0x94, 0x68, 0xbe, 0xcf, 0xdf, 0x7c, 0xf3, 0x8d, 0x4d, 0x7a, 0x7b, 0x22,
	0x95, 0x09, 0x3b, 0x18, 0x8b, 0xbc, 0x08, 0xb2, 0x5c, 0x6a, 0x09, 0x24, 0x16, 0xb9, 0x0e, 0x4c,
	0xdd, 0xed, 0xc7, 0x32, 0x96, 0xa6, 0xcc, 0xaa, 0x3f, 0xcb, 0x70, 0xef, 0xc5, 0x52, 0xc6, 0xfb,
	0x82, 0xf1, 0x6c, 0xc4, 0x78, 0x9a, 0x4a, 0xcd, 0xf5, 0x48, 0xa6, 0x0a, 0xd1, 0xf5, 0x5d, 0xa9,
	0x12, 0xa9, 0xd8, 0x90, 0x2b, 0x61, 0x85, 0xd9, 0xfb, 0xcd, 0xa1, 0xd0, 0x7c, 0x93, 0x65, 0x3c,
	0x1e, 0xa5, 0x86, 0x8c, 0xdc, 0x65, 0xdb, 0x9e, 0x8f, 0xf5, 0x3b, 0x99, 0x8f, 0x74, 0xb1, 0x2d,
	0x34, 0xdf, 0xe3, 0x9a, 0x23, 0x0c, 0x16, 0xce, 0x78, 0xce, 0x13, 0x94, 0xa7, 0x7d, 0x02, 0x6f,
	0x2a, 0xd1, 0xd7, 0xa6, 0x18, 0x89, 0x83, 0xb1, 0x50, 0x9a, 0xbe, 0x20, 0xb7, 0xfe, 0xa9, 0xaa,
	0x4c, 0xa6, 0x4a, 0xc0, 0x43, 0xd2, 0xb5, 0x87, 0xef, 0x38, 0x2b, 0xce, 0xe0, 0xea, 0x16, 0x04,
	0x67, 0xc3, 0x05, 0x96, 0x1b, 0x5e, 0x3e, 0x2e, 0xfd, 0x4e, 0x84, 0x3c, 0xfa, 0x8a, 0x50, 0x23,
	0xf4, 0xac, 0xa2, 0x3c, 0xf9, 0xdf, 0x17, 0xb6, 0x83, 0x55, 0xb2, 0x68, 0x34, 0x8c, 0xec, 0x52,
	0x78, 0x73, 0x56, 0xfa, 0xd7, 0x0a, 0x9e, 0xec, 0x3f, 0xa6, 0xa6, 0x4c, 0x23, 0x0b, 0xd3, 0xaf,
	0x0e, 0x79, 0x70, 0xa1, 0x1c, 0xfa, 0xd4, 0x04, 0xfe, 0x66, 0xb0, 0x93, 0x20, 0x8a, 0x9e, 0x69,
	0xdd, 0x73, 0xbb, 0x4e, 0x78, 0xbf, 0x9a, 0x61, 0x56, 0xfa, 0x77, 0xad, 0x89, 0xa6, 0x16, 0x8d,
	0x7a, 0x8d, 0x90, 0xe9, 0x36, 0x59, 0x3e, 0x33, 0xa7, 0x9e, 0xe7, 0x32, 0x79, 0x9a, 0x0b, 0xae,
	0x65, 0x7e, 0x3a, 0xe6, 0x06, 0xb9, 0xb2, 0x6b, 0x2b, 0x38, 0x28, 0xcc, 0x4a, 0xff, 0x86, 0xed,
	0x81, 0x00, 0x8d, 0x4e, 0x29, 0xf4, 0x25, 0xf1, 0xce, 0x93, 0xc3, 0x31, 0xd7, 0x48, 0xd7, 0x8c,
	0x51, 0xad, 0x63, 0x61, 0xb0, 0x14, 0xf6, 0x66, 0xa5, 0x7f, 0xbd, 0x96, 0x9b, 0xa2, 0x11, 0x12,
	0xb6, 0xbe, 0x2f, 0x90, 0x45, 0xa3, 0x06, 0x82, 0x74, 0xed, 0xa6, 0xc0, 0xab, 0x27, 0xd1, 0xbc,
	0x04, 0xae, 0x7f, 0x2e, 0x6e, 0xfb, 0x53, 0xf7, 0xe3, 0x8f, 0xdf, 0x9f, 0x2f, 0xf5, 0x01, 0x58,
	0x45, 0xb4, 0x9d, 0xf0, 0x76, 0xc1, 0x37, 0x87, 0xdc, 0x6e, 0x4f, 0x17, 0x82, 0x86, 0xee, 0x85,
	0xb7, 0xc3, 0x65, 0x73, 0xf3, 0xd1, 0x57, 0x60, 0x7c, 0x0d, 0x60, 0xb5, 0xee, 0xeb, 0xd0, 0x7c,
	0x8f, 0x58, 0x73, 0x99, 0xf0, 0xc5, 0x21, 0xbd, 0x46, 0xca, 0xb0, 0xd6, 0xde, 0xb6, 0x65, 0xb1,
	0xee, 0xfa, 0x3c, 0x54, 0x34, 0x37, 0x30, 0xe6, 0x28, 0xac, 0xd4, 0xcd, 0x0d, 0x8b, 0x1d, 0x5c,
	0x3b, 0x3b, 0xc4, 0x9f, 0xa3, 0x70, 0xe3, 0x78, 0xe2, 0x39, 0x27, 0x13, 0xcf, 0xf9, 0x35, 0xf1,
	0x9c, 0x4f, 0x53, 0xaf, 0x73, 0x32, 0xf5, 0x3a, 0x3f, 0xa7, 0x5e, 0xe7, 0x2d, 0x54, 0x47, 0xd9,
	0x07, 0x66, 0xdf, 0xb3, 0x2e, 0x32, 0xa1, 0x86, 0x5d, 0xf3, 0x9e, 0x1f, 0xfd, 0x09, 0x00, 0x00,
	0xff, 0xff, 0xa0, 0xed, 0x87, 0x1a, 0x83, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Params returns the total set of minting parameters.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	DenomAuthorityMetadata(ctx context.Context, in *QueryDenomAuthorityMetadataRequest, opts ...grpc.CallOption) (*QueryDenomAuthorityMetadataResponse, error)
	DenomsFromCreator(ctx context.Context, in *QueryDenomsFromCreatorRequest, opts ...grpc.CallOption) (*QueryDenomsFromCreatorResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/gert.denom.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DenomAuthorityMetadata(ctx context.Context, in *QueryDenomAuthorityMetadataRequest, opts ...grpc.CallOption) (*QueryDenomAuthorityMetadataResponse, error) {
	out := new(QueryDenomAuthorityMetadataResponse)
	err := c.cc.Invoke(ctx, "/gert.denom.Query/DenomAuthorityMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DenomsFromCreator(ctx context.Context, in *QueryDenomsFromCreatorRequest, opts ...grpc.CallOption) (*QueryDenomsFromCreatorResponse, error) {
	out := new(QueryDenomsFromCreatorResponse)
	err := c.cc.Invoke(ctx, "/gert.denom.Query/DenomsFromCreator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Params returns the total set of minting parameters.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	DenomAuthorityMetadata(context.Context, *QueryDenomAuthorityMetadataRequest) (*QueryDenomAuthorityMetadataResponse, error)
	DenomsFromCreator(context.Context, *QueryDenomsFromCreatorRequest) (*QueryDenomsFromCreatorResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) DenomAuthorityMetadata(ctx context.Context, req *QueryDenomAuthorityMetadataRequest) (*QueryDenomAuthorityMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DenomAuthorityMetadata not implemented")
}
func (*UnimplementedQueryServer) DenomsFromCreator(ctx context.Context, req *QueryDenomsFromCreatorRequest) (*QueryDenomsFromCreatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DenomsFromCreator not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gert.denom.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DenomAuthorityMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDenomAuthorityMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DenomAuthorityMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gert.denom.Query/DenomAuthorityMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DenomAuthorityMetadata(ctx, req.(*QueryDenomAuthorityMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DenomsFromCreator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDenomsFromCreatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DenomsFromCreator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gert.denom.Query/DenomsFromCreator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DenomsFromCreator(ctx, req.(*QueryDenomsFromCreatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gert.denom.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "DenomAuthorityMetadata",
			Handler:    _Query_DenomAuthorityMetadata_Handler,
		},
		{
			MethodName: "DenomsFromCreator",
			Handler:    _Query_DenomsFromCreator_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "denom/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryDenomAuthorityMetadataRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDenomAuthorityMetadataRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDenomAuthorityMetadataRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryDenomAuthorityMetadataResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDenomAuthorityMetadataResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDenomAuthorityMetadataResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.AuthorityMetadata.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryDenomsFromCreatorRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDenomsFromCreatorRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDenomsFromCreatorRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryDenomsFromCreatorResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDenomsFromCreatorResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDenomsFromCreatorResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Denoms) > 0 {
		for iNdEx := len(m.Denoms) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Denoms[iNdEx])
			copy(dAtA[i:], m.Denoms[iNdEx])
			i = encodeVarintQuery(dAtA, i, uint64(len(m.Denoms[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryDenomAuthorityMetadataRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryDenomAuthorityMetadataResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.AuthorityMetadata.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryDenomsFromCreatorRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryDenomsFromCreatorResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Denoms) > 0 {
		for _, s := range m.Denoms {
			l = len(s)
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryDenomAuthorityMetadataRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryDenomAuthorityMetadataRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDenomAuthorityMetadataRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryDenomAuthorityMetadataResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryDenomAuthorityMetadataResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDenomAuthorityMetadataResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthorityMetadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AuthorityMetadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryDenomsFromCreatorRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryDenomsFromCreatorRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDenomsFromCreatorRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryDenomsFromCreatorResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryDenomsFromCreatorResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDenomsFromCreatorResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denoms", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denoms = append(m.Denoms, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
