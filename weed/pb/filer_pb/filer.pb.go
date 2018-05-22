// Code generated by protoc-gen-go.
// source: filer.proto
// DO NOT EDIT!

/*
Package filer_pb is a generated protocol buffer package.

It is generated from these files:
	filer.proto

It has these top-level messages:
	LookupDirectoryEntryRequest
	LookupDirectoryEntryResponse
	ListEntriesRequest
	ListEntriesResponse
	Entry
	FuseAttributes
	GetFileAttributesRequest
	GetFileAttributesResponse
	GetFileContentRequest
	GetFileContentResponse
	DeleteEntryRequest
	DeleteEntryResponse
*/
package filer_pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type LookupDirectoryEntryRequest struct {
	Directory string `protobuf:"bytes,1,opt,name=directory" json:"directory,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *LookupDirectoryEntryRequest) Reset()                    { *m = LookupDirectoryEntryRequest{} }
func (m *LookupDirectoryEntryRequest) String() string            { return proto.CompactTextString(m) }
func (*LookupDirectoryEntryRequest) ProtoMessage()               {}
func (*LookupDirectoryEntryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LookupDirectoryEntryRequest) GetDirectory() string {
	if m != nil {
		return m.Directory
	}
	return ""
}

func (m *LookupDirectoryEntryRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type LookupDirectoryEntryResponse struct {
	Entry *Entry `protobuf:"bytes,1,opt,name=entry" json:"entry,omitempty"`
}

func (m *LookupDirectoryEntryResponse) Reset()                    { *m = LookupDirectoryEntryResponse{} }
func (m *LookupDirectoryEntryResponse) String() string            { return proto.CompactTextString(m) }
func (*LookupDirectoryEntryResponse) ProtoMessage()               {}
func (*LookupDirectoryEntryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LookupDirectoryEntryResponse) GetEntry() *Entry {
	if m != nil {
		return m.Entry
	}
	return nil
}

type ListEntriesRequest struct {
	Directory string `protobuf:"bytes,1,opt,name=directory" json:"directory,omitempty"`
}

func (m *ListEntriesRequest) Reset()                    { *m = ListEntriesRequest{} }
func (m *ListEntriesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListEntriesRequest) ProtoMessage()               {}
func (*ListEntriesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ListEntriesRequest) GetDirectory() string {
	if m != nil {
		return m.Directory
	}
	return ""
}

type ListEntriesResponse struct {
	Entries []*Entry `protobuf:"bytes,1,rep,name=entries" json:"entries,omitempty"`
}

func (m *ListEntriesResponse) Reset()                    { *m = ListEntriesResponse{} }
func (m *ListEntriesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListEntriesResponse) ProtoMessage()               {}
func (*ListEntriesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ListEntriesResponse) GetEntries() []*Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type Entry struct {
	Name        string          `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	IsDirectory bool            `protobuf:"varint,2,opt,name=is_directory,json=isDirectory" json:"is_directory,omitempty"`
	FileId      string          `protobuf:"bytes,3,opt,name=file_id,json=fileId" json:"file_id,omitempty"`
	Attributes  *FuseAttributes `protobuf:"bytes,4,opt,name=attributes" json:"attributes,omitempty"`
}

func (m *Entry) Reset()                    { *m = Entry{} }
func (m *Entry) String() string            { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()               {}
func (*Entry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Entry) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Entry) GetIsDirectory() bool {
	if m != nil {
		return m.IsDirectory
	}
	return false
}

func (m *Entry) GetFileId() string {
	if m != nil {
		return m.FileId
	}
	return ""
}

func (m *Entry) GetAttributes() *FuseAttributes {
	if m != nil {
		return m.Attributes
	}
	return nil
}

type FuseAttributes struct {
	FileSize uint64 `protobuf:"varint,1,opt,name=file_size,json=fileSize" json:"file_size,omitempty"`
	Mtime    int64  `protobuf:"varint,2,opt,name=mtime" json:"mtime,omitempty"`
	FileMode uint32 `protobuf:"varint,3,opt,name=file_mode,json=fileMode" json:"file_mode,omitempty"`
	Uid      uint32 `protobuf:"varint,4,opt,name=uid" json:"uid,omitempty"`
	Gid      uint32 `protobuf:"varint,5,opt,name=gid" json:"gid,omitempty"`
}

func (m *FuseAttributes) Reset()                    { *m = FuseAttributes{} }
func (m *FuseAttributes) String() string            { return proto.CompactTextString(m) }
func (*FuseAttributes) ProtoMessage()               {}
func (*FuseAttributes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *FuseAttributes) GetFileSize() uint64 {
	if m != nil {
		return m.FileSize
	}
	return 0
}

func (m *FuseAttributes) GetMtime() int64 {
	if m != nil {
		return m.Mtime
	}
	return 0
}

func (m *FuseAttributes) GetFileMode() uint32 {
	if m != nil {
		return m.FileMode
	}
	return 0
}

func (m *FuseAttributes) GetUid() uint32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *FuseAttributes) GetGid() uint32 {
	if m != nil {
		return m.Gid
	}
	return 0
}

type GetFileAttributesRequest struct {
	Name      string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	ParentDir string `protobuf:"bytes,2,opt,name=parent_dir,json=parentDir" json:"parent_dir,omitempty"`
	FileId    string `protobuf:"bytes,3,opt,name=file_id,json=fileId" json:"file_id,omitempty"`
}

func (m *GetFileAttributesRequest) Reset()                    { *m = GetFileAttributesRequest{} }
func (m *GetFileAttributesRequest) String() string            { return proto.CompactTextString(m) }
func (*GetFileAttributesRequest) ProtoMessage()               {}
func (*GetFileAttributesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *GetFileAttributesRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetFileAttributesRequest) GetParentDir() string {
	if m != nil {
		return m.ParentDir
	}
	return ""
}

func (m *GetFileAttributesRequest) GetFileId() string {
	if m != nil {
		return m.FileId
	}
	return ""
}

type GetFileAttributesResponse struct {
	Attributes *FuseAttributes `protobuf:"bytes,1,opt,name=attributes" json:"attributes,omitempty"`
}

func (m *GetFileAttributesResponse) Reset()                    { *m = GetFileAttributesResponse{} }
func (m *GetFileAttributesResponse) String() string            { return proto.CompactTextString(m) }
func (*GetFileAttributesResponse) ProtoMessage()               {}
func (*GetFileAttributesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *GetFileAttributesResponse) GetAttributes() *FuseAttributes {
	if m != nil {
		return m.Attributes
	}
	return nil
}

type GetFileContentRequest struct {
	FileId string `protobuf:"bytes,1,opt,name=file_id,json=fileId" json:"file_id,omitempty"`
}

func (m *GetFileContentRequest) Reset()                    { *m = GetFileContentRequest{} }
func (m *GetFileContentRequest) String() string            { return proto.CompactTextString(m) }
func (*GetFileContentRequest) ProtoMessage()               {}
func (*GetFileContentRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *GetFileContentRequest) GetFileId() string {
	if m != nil {
		return m.FileId
	}
	return ""
}

type GetFileContentResponse struct {
	Content []byte `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (m *GetFileContentResponse) Reset()                    { *m = GetFileContentResponse{} }
func (m *GetFileContentResponse) String() string            { return proto.CompactTextString(m) }
func (*GetFileContentResponse) ProtoMessage()               {}
func (*GetFileContentResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *GetFileContentResponse) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type DeleteEntryRequest struct {
	Directory   string `protobuf:"bytes,1,opt,name=directory" json:"directory,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	IsDirectory bool   `protobuf:"varint,3,opt,name=is_directory,json=isDirectory" json:"is_directory,omitempty"`
}

func (m *DeleteEntryRequest) Reset()                    { *m = DeleteEntryRequest{} }
func (m *DeleteEntryRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteEntryRequest) ProtoMessage()               {}
func (*DeleteEntryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *DeleteEntryRequest) GetDirectory() string {
	if m != nil {
		return m.Directory
	}
	return ""
}

func (m *DeleteEntryRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DeleteEntryRequest) GetIsDirectory() bool {
	if m != nil {
		return m.IsDirectory
	}
	return false
}

type DeleteEntryResponse struct {
}

func (m *DeleteEntryResponse) Reset()                    { *m = DeleteEntryResponse{} }
func (m *DeleteEntryResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteEntryResponse) ProtoMessage()               {}
func (*DeleteEntryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func init() {
	proto.RegisterType((*LookupDirectoryEntryRequest)(nil), "filer_pb.LookupDirectoryEntryRequest")
	proto.RegisterType((*LookupDirectoryEntryResponse)(nil), "filer_pb.LookupDirectoryEntryResponse")
	proto.RegisterType((*ListEntriesRequest)(nil), "filer_pb.ListEntriesRequest")
	proto.RegisterType((*ListEntriesResponse)(nil), "filer_pb.ListEntriesResponse")
	proto.RegisterType((*Entry)(nil), "filer_pb.Entry")
	proto.RegisterType((*FuseAttributes)(nil), "filer_pb.FuseAttributes")
	proto.RegisterType((*GetFileAttributesRequest)(nil), "filer_pb.GetFileAttributesRequest")
	proto.RegisterType((*GetFileAttributesResponse)(nil), "filer_pb.GetFileAttributesResponse")
	proto.RegisterType((*GetFileContentRequest)(nil), "filer_pb.GetFileContentRequest")
	proto.RegisterType((*GetFileContentResponse)(nil), "filer_pb.GetFileContentResponse")
	proto.RegisterType((*DeleteEntryRequest)(nil), "filer_pb.DeleteEntryRequest")
	proto.RegisterType((*DeleteEntryResponse)(nil), "filer_pb.DeleteEntryResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SeaweedFiler service

type SeaweedFilerClient interface {
	LookupDirectoryEntry(ctx context.Context, in *LookupDirectoryEntryRequest, opts ...grpc.CallOption) (*LookupDirectoryEntryResponse, error)
	ListEntries(ctx context.Context, in *ListEntriesRequest, opts ...grpc.CallOption) (*ListEntriesResponse, error)
	GetFileAttributes(ctx context.Context, in *GetFileAttributesRequest, opts ...grpc.CallOption) (*GetFileAttributesResponse, error)
	GetFileContent(ctx context.Context, in *GetFileContentRequest, opts ...grpc.CallOption) (*GetFileContentResponse, error)
	DeleteEntry(ctx context.Context, in *DeleteEntryRequest, opts ...grpc.CallOption) (*DeleteEntryResponse, error)
}

type seaweedFilerClient struct {
	cc *grpc.ClientConn
}

func NewSeaweedFilerClient(cc *grpc.ClientConn) SeaweedFilerClient {
	return &seaweedFilerClient{cc}
}

func (c *seaweedFilerClient) LookupDirectoryEntry(ctx context.Context, in *LookupDirectoryEntryRequest, opts ...grpc.CallOption) (*LookupDirectoryEntryResponse, error) {
	out := new(LookupDirectoryEntryResponse)
	err := grpc.Invoke(ctx, "/filer_pb.SeaweedFiler/LookupDirectoryEntry", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seaweedFilerClient) ListEntries(ctx context.Context, in *ListEntriesRequest, opts ...grpc.CallOption) (*ListEntriesResponse, error) {
	out := new(ListEntriesResponse)
	err := grpc.Invoke(ctx, "/filer_pb.SeaweedFiler/ListEntries", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seaweedFilerClient) GetFileAttributes(ctx context.Context, in *GetFileAttributesRequest, opts ...grpc.CallOption) (*GetFileAttributesResponse, error) {
	out := new(GetFileAttributesResponse)
	err := grpc.Invoke(ctx, "/filer_pb.SeaweedFiler/GetFileAttributes", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seaweedFilerClient) GetFileContent(ctx context.Context, in *GetFileContentRequest, opts ...grpc.CallOption) (*GetFileContentResponse, error) {
	out := new(GetFileContentResponse)
	err := grpc.Invoke(ctx, "/filer_pb.SeaweedFiler/GetFileContent", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seaweedFilerClient) DeleteEntry(ctx context.Context, in *DeleteEntryRequest, opts ...grpc.CallOption) (*DeleteEntryResponse, error) {
	out := new(DeleteEntryResponse)
	err := grpc.Invoke(ctx, "/filer_pb.SeaweedFiler/DeleteEntry", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SeaweedFiler service

type SeaweedFilerServer interface {
	LookupDirectoryEntry(context.Context, *LookupDirectoryEntryRequest) (*LookupDirectoryEntryResponse, error)
	ListEntries(context.Context, *ListEntriesRequest) (*ListEntriesResponse, error)
	GetFileAttributes(context.Context, *GetFileAttributesRequest) (*GetFileAttributesResponse, error)
	GetFileContent(context.Context, *GetFileContentRequest) (*GetFileContentResponse, error)
	DeleteEntry(context.Context, *DeleteEntryRequest) (*DeleteEntryResponse, error)
}

func RegisterSeaweedFilerServer(s *grpc.Server, srv SeaweedFilerServer) {
	s.RegisterService(&_SeaweedFiler_serviceDesc, srv)
}

func _SeaweedFiler_LookupDirectoryEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupDirectoryEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeaweedFilerServer).LookupDirectoryEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filer_pb.SeaweedFiler/LookupDirectoryEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeaweedFilerServer).LookupDirectoryEntry(ctx, req.(*LookupDirectoryEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SeaweedFiler_ListEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeaweedFilerServer).ListEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filer_pb.SeaweedFiler/ListEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeaweedFilerServer).ListEntries(ctx, req.(*ListEntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SeaweedFiler_GetFileAttributes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFileAttributesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeaweedFilerServer).GetFileAttributes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filer_pb.SeaweedFiler/GetFileAttributes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeaweedFilerServer).GetFileAttributes(ctx, req.(*GetFileAttributesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SeaweedFiler_GetFileContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFileContentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeaweedFilerServer).GetFileContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filer_pb.SeaweedFiler/GetFileContent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeaweedFilerServer).GetFileContent(ctx, req.(*GetFileContentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SeaweedFiler_DeleteEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeaweedFilerServer).DeleteEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filer_pb.SeaweedFiler/DeleteEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeaweedFilerServer).DeleteEntry(ctx, req.(*DeleteEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SeaweedFiler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "filer_pb.SeaweedFiler",
	HandlerType: (*SeaweedFilerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LookupDirectoryEntry",
			Handler:    _SeaweedFiler_LookupDirectoryEntry_Handler,
		},
		{
			MethodName: "ListEntries",
			Handler:    _SeaweedFiler_ListEntries_Handler,
		},
		{
			MethodName: "GetFileAttributes",
			Handler:    _SeaweedFiler_GetFileAttributes_Handler,
		},
		{
			MethodName: "GetFileContent",
			Handler:    _SeaweedFiler_GetFileContent_Handler,
		},
		{
			MethodName: "DeleteEntry",
			Handler:    _SeaweedFiler_DeleteEntry_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "filer.proto",
}

func init() { proto.RegisterFile("filer.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 532 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xad, 0x71, 0xd2, 0x34, 0x93, 0xb4, 0xc0, 0xb4, 0x05, 0x93, 0xa6, 0x22, 0x2c, 0x2a, 0x82,
	0x4b, 0x84, 0xc2, 0x85, 0x23, 0x88, 0xb4, 0x08, 0x29, 0x08, 0xc9, 0x55, 0xaf, 0x44, 0x49, 0x3d,
	0x8d, 0x56, 0x24, 0x76, 0xf0, 0xae, 0x85, 0xda, 0x33, 0x7f, 0x80, 0xbf, 0xc5, 0xaf, 0x42, 0xbb,
	0xeb, 0x8f, 0x35, 0x76, 0x0a, 0x88, 0x9b, 0xf7, 0xcd, 0xbe, 0x99, 0x37, 0x6f, 0x66, 0x0d, 0x9d,
	0x2b, 0xbe, 0xa4, 0x78, 0xb8, 0x8e, 0x23, 0x19, 0xe1, 0x8e, 0x3e, 0x4c, 0xd7, 0x73, 0xf6, 0x09,
	0x8e, 0x26, 0x51, 0xf4, 0x25, 0x59, 0x8f, 0x79, 0x4c, 0x97, 0x32, 0x8a, 0xaf, 0x4f, 0x43, 0x19,
	0x5f, 0xfb, 0xf4, 0x35, 0x21, 0x21, 0xb1, 0x0f, 0xed, 0x20, 0x0b, 0x78, 0xce, 0xc0, 0x79, 0xde,
	0xf6, 0x0b, 0x00, 0x11, 0x1a, 0xe1, 0x6c, 0x45, 0xde, 0x1d, 0x1d, 0xd0, 0xdf, 0xec, 0x14, 0xfa,
	0xf5, 0x09, 0xc5, 0x3a, 0x0a, 0x05, 0xe1, 0x09, 0x34, 0x49, 0x01, 0x3a, 0x5b, 0x67, 0x74, 0x77,
	0x98, 0x49, 0x19, 0x9a, 0x7b, 0x26, 0xca, 0x46, 0x80, 0x13, 0x2e, 0xa4, 0xc2, 0x38, 0x89, 0xbf,
	0x92, 0xc3, 0xde, 0xc0, 0x7e, 0x89, 0x93, 0x56, 0x7c, 0x01, 0x2d, 0x32, 0x90, 0xe7, 0x0c, 0xdc,
	0xba, 0x9a, 0x59, 0x9c, 0xfd, 0x70, 0xa0, 0xa9, 0xa1, 0xbc, 0x35, 0xa7, 0x68, 0x0d, 0x9f, 0x40,
	0x97, 0x8b, 0x69, 0x21, 0x40, 0xb5, 0xbd, 0xe3, 0x77, 0xb8, 0xc8, 0x5b, 0xc5, 0x87, 0xd0, 0x52,
	0xb9, 0xa7, 0x3c, 0xf0, 0x5c, 0xcd, 0xdc, 0x56, 0xc7, 0x0f, 0x01, 0xbe, 0x06, 0x98, 0x49, 0x19,
	0xf3, 0x79, 0x22, 0x49, 0x78, 0x0d, 0xdd, 0xbb, 0x57, 0xe8, 0x38, 0x4b, 0x04, 0xbd, 0xcd, 0xe3,
	0xbe, 0x75, 0x97, 0x7d, 0x77, 0x60, 0xaf, 0x1c, 0xc6, 0x23, 0x68, 0xeb, 0x2a, 0x82, 0xdf, 0x18,
	0x85, 0x0d, 0x5f, 0x4f, 0xf4, 0x9c, 0xdf, 0x10, 0x1e, 0x40, 0x73, 0x25, 0x79, 0x3a, 0x15, 0xd7,
	0x37, 0x87, 0x9c, 0xb2, 0x8a, 0x02, 0xd2, 0xd2, 0x76, 0x0d, 0xe5, 0x63, 0x14, 0x10, 0xde, 0x03,
	0x37, 0xe1, 0x81, 0x56, 0xb5, 0xeb, 0xab, 0x4f, 0x85, 0x2c, 0x78, 0xe0, 0x35, 0x0d, 0xb2, 0xe0,
	0x01, 0xbb, 0x02, 0xef, 0x3d, 0xc9, 0x33, 0xbe, 0xb4, 0x75, 0xa6, 0x63, 0xa9, 0x33, 0xeb, 0x18,
	0x60, 0x3d, 0x8b, 0x29, 0x94, 0xca, 0xb0, 0x74, 0x43, 0xda, 0x06, 0x19, 0xf3, 0x78, 0xa3, 0x51,
	0xec, 0x02, 0x1e, 0xd5, 0xd4, 0x49, 0x47, 0x59, 0x76, 0xd1, 0xf9, 0x07, 0x17, 0x5f, 0xc2, 0x61,
	0x9a, 0xf6, 0x5d, 0x14, 0x4a, 0x0a, 0x65, 0xa6, 0xdd, 0x12, 0xe2, 0x94, 0x84, 0x8c, 0xe0, 0xc1,
	0xef, 0x8c, 0x54, 0x85, 0x07, 0xad, 0x4b, 0x03, 0x69, 0x4a, 0xd7, 0xcf, 0x8e, 0x8c, 0x03, 0x8e,
	0x69, 0x49, 0x92, 0xfe, 0xef, 0x11, 0x55, 0x36, 0xcd, 0xad, 0x6c, 0x1a, 0x3b, 0x84, 0xfd, 0x52,
	0x29, 0xa3, 0x6d, 0xf4, 0xd3, 0x85, 0xee, 0x39, 0xcd, 0xbe, 0x11, 0x05, 0x4a, 0x7a, 0x8c, 0x0b,
	0x38, 0xa8, 0x7b, 0x8f, 0x78, 0x52, 0xd8, 0x76, 0xcb, 0x0f, 0xa0, 0xf7, 0xec, 0x4f, 0xd7, 0x4c,
	0x5d, 0xb6, 0x85, 0x13, 0xe8, 0x58, 0xaf, 0x0f, 0xfb, 0x16, 0xb1, 0xf2, 0x90, 0x7b, 0xc7, 0x1b,
	0xa2, 0x79, 0xb6, 0xcf, 0x70, 0xbf, 0xb2, 0x06, 0xc8, 0x0a, 0xd6, 0xa6, 0x5d, 0xec, 0x3d, 0xbd,
	0xf5, 0x4e, 0x9e, 0xff, 0x02, 0xf6, 0xca, 0xd3, 0xc5, 0xc7, 0x15, 0x62, 0x79, 0x53, 0x7a, 0x83,
	0xcd, 0x17, 0x6c, 0x13, 0xac, 0xa9, 0xd8, 0x26, 0x54, 0xf7, 0xc2, 0x36, 0xa1, 0x66, 0x94, 0x6c,
	0x6b, 0xbe, 0xad, 0xff, 0xd6, 0xaf, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0x6a, 0x83, 0x8a, 0x32,
	0xbc, 0x05, 0x00, 0x00,
}