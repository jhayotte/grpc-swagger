// Code generated by protoc-gen-go.
// source: point.proto
// DO NOT EDIT!

/*
Package pointpb is a generated protocol buffer package.

It is generated from these files:
	point.proto

It has these top-level messages:
	Location
	Point
	Request
	Response
*/
package pointpb

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

type Location struct {
	Points []*Point `protobuf:"bytes,1,rep,name=Points" json:"Points,omitempty"`
}

func (m *Location) Reset()                    { *m = Location{} }
func (m *Location) String() string            { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()               {}
func (*Location) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Location) GetPoints() []*Point {
	if m != nil {
		return m.Points
	}
	return nil
}

type Point struct {
	AssetID  string  `protobuf:"bytes,1,opt,name=assetID" json:"assetID,omitempty"`
	Lat      float64 `protobuf:"fixed64,2,opt,name=Lat" json:"Lat,omitempty"`
	Lng      float64 `protobuf:"fixed64,3,opt,name=Lng" json:"Lng,omitempty"`
	Datetime string  `protobuf:"bytes,4,opt,name=Datetime" json:"Datetime,omitempty"`
}

func (m *Point) Reset()                    { *m = Point{} }
func (m *Point) String() string            { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()               {}
func (*Point) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Request struct {
	AssetID string `protobuf:"bytes,1,opt,name=assetID" json:"assetID,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Response struct {
	Point *Point `protobuf:"bytes,1,opt,name=point" json:"point,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Response) GetPoint() *Point {
	if m != nil {
		return m.Point
	}
	return nil
}

func init() {
	proto.RegisterType((*Location)(nil), "pointpb.Location")
	proto.RegisterType((*Point)(nil), "pointpb.Point")
	proto.RegisterType((*Request)(nil), "pointpb.Request")
	proto.RegisterType((*Response)(nil), "pointpb.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Trackersvc service

type TrackersvcClient interface {
	Point(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	PointStream(ctx context.Context, in *Request, opts ...grpc.CallOption) (Trackersvc_PointStreamClient, error)
}

type trackersvcClient struct {
	cc *grpc.ClientConn
}

func NewTrackersvcClient(cc *grpc.ClientConn) TrackersvcClient {
	return &trackersvcClient{cc}
}

func (c *trackersvcClient) Point(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/pointpb.Trackersvc/Point", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trackersvcClient) PointStream(ctx context.Context, in *Request, opts ...grpc.CallOption) (Trackersvc_PointStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Trackersvc_serviceDesc.Streams[0], c.cc, "/pointpb.Trackersvc/PointStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &trackersvcPointStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Trackersvc_PointStreamClient interface {
	Recv() (*Response, error)
	grpc.ClientStream
}

type trackersvcPointStreamClient struct {
	grpc.ClientStream
}

func (x *trackersvcPointStreamClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Trackersvc service

type TrackersvcServer interface {
	Point(context.Context, *Request) (*Response, error)
	PointStream(*Request, Trackersvc_PointStreamServer) error
}

func RegisterTrackersvcServer(s *grpc.Server, srv TrackersvcServer) {
	s.RegisterService(&_Trackersvc_serviceDesc, srv)
}

func _Trackersvc_Point_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrackersvcServer).Point(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pointpb.Trackersvc/Point",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrackersvcServer).Point(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Trackersvc_PointStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TrackersvcServer).PointStream(m, &trackersvcPointStreamServer{stream})
}

type Trackersvc_PointStreamServer interface {
	Send(*Response) error
	grpc.ServerStream
}

type trackersvcPointStreamServer struct {
	grpc.ServerStream
}

func (x *trackersvcPointStreamServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

var _Trackersvc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pointpb.Trackersvc",
	HandlerType: (*TrackersvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Point",
			Handler:    _Trackersvc_Point_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PointStream",
			Handler:       _Trackersvc_PointStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("point.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x90, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0x89, 0x6b, 0xbb, 0xeb, 0x2c, 0x48, 0xcd, 0x29, 0xf4, 0xb4, 0x44, 0x91, 0x3d, 0xc8,
	0x52, 0x56, 0x7f, 0x42, 0x2f, 0x42, 0x0f, 0x12, 0xfd, 0x03, 0xe9, 0x32, 0x94, 0x45, 0x9a, 0xc4,
	0x64, 0xf4, 0xf7, 0x4b, 0xc7, 0xec, 0x22, 0x08, 0xd2, 0xdb, 0xbc, 0x6f, 0xde, 0xf0, 0x5e, 0x02,
	0x75, 0xf0, 0xa3, 0xa3, 0x2e, 0x44, 0x4f, 0x5e, 0x96, 0x2c, 0xc2, 0x5e, 0xf7, 0x50, 0xed, 0xfc,
	0x60, 0x69, 0xf4, 0x4e, 0xde, 0xc3, 0xf2, 0xe5, 0x84, 0x93, 0x12, 0x4d, 0xd1, 0xd6, 0xfd, 0x75,
	0x97, 0x5d, 0x1d, 0x63, 0x93, 0xb7, 0xda, 0xc2, 0x82, 0x27, 0xa9, 0xa0, 0xb4, 0x29, 0x21, 0x3d,
	0x6f, 0x95, 0x68, 0x44, 0x7b, 0x65, 0x26, 0x29, 0x57, 0x50, 0xec, 0x2c, 0xa9, 0x8b, 0x46, 0xb4,
	0xc2, 0x9c, 0x46, 0x26, 0xee, 0xa0, 0x8a, 0x4c, 0xdc, 0x41, 0xae, 0xa1, 0xda, 0x5a, 0x42, 0x1a,
	0x8f, 0xa8, 0x2e, 0xf9, 0x7c, 0xd6, 0xfa, 0x16, 0x4a, 0x83, 0x1f, 0x9f, 0x98, 0xfe, 0x09, 0xd1,
	0x1b, 0xa8, 0x0c, 0xa6, 0xe0, 0x5d, 0x42, 0x79, 0x07, 0x0b, 0x2e, 0xcb, 0x9e, 0xbf, 0xd5, 0x7f,
	0x96, 0x7d, 0x00, 0x78, 0x8b, 0x76, 0x78, 0xc7, 0x98, 0xbe, 0x06, 0xf9, 0x30, 0xbd, 0x63, 0x35,
	0xbb, 0x73, 0xe8, 0xfa, 0xe6, 0x17, 0xc9, 0x09, 0x4f, 0x50, 0xb3, 0xfb, 0x95, 0x22, 0xda, 0xe3,
	0x59, 0x37, 0x1b, 0xb1, 0x5f, 0xf2, 0x7f, 0x3f, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x87, 0x59,
	0x65, 0xd5, 0x7e, 0x01, 0x00, 0x00,
}
