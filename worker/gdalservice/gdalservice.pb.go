// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gdalservice.proto

/*
Package gdalservice is a generated protocol buffer package.

It is generated from these files:
	gdalservice.proto

It has these top-level messages:
	GeoRPCGranule
	Raster
	TimeSeries
	Overview
	GeoMetaData
	GeoFile
	WorkerInfo
	WorkerMetrics
	Result
*/
package gdalservice

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

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

type GeoRPCGranule struct {
	Operation        string    `protobuf:"bytes,1,opt,name=operation" json:"operation,omitempty"`
	Path             string    `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	Geometry         string    `protobuf:"bytes,3,opt,name=geometry" json:"geometry,omitempty"`
	Bands            []int32   `protobuf:"varint,4,rep,packed,name=bands" json:"bands,omitempty"`
	Height           float32   `protobuf:"fixed32,5,opt,name=height" json:"height,omitempty"`
	Width            float32   `protobuf:"fixed32,6,opt,name=width" json:"width,omitempty"`
	SrcSRS           string    `protobuf:"bytes,7,opt,name=srcSRS" json:"srcSRS,omitempty"`
	SrcGeot          []float64 `protobuf:"fixed64,8,rep,packed,name=srcGeot" json:"srcGeot,omitempty"`
	DstSRS           string    `protobuf:"bytes,9,opt,name=dstSRS" json:"dstSRS,omitempty"`
	DstGeot          []float64 `protobuf:"fixed64,10,rep,packed,name=dstGeot" json:"dstGeot,omitempty"`
	BandStrides      int32     `protobuf:"varint,11,opt,name=bandStrides" json:"bandStrides,omitempty"`
	GeoLocOpts       []string  `protobuf:"bytes,12,rep,name=geoLocOpts" json:"geoLocOpts,omitempty"`
	DrillDecileCount int32     `protobuf:"varint,13,opt,name=drillDecileCount" json:"drillDecileCount,omitempty"`
	ClipUpper        float32   `protobuf:"fixed32,14,opt,name=clipUpper" json:"clipUpper,omitempty"`
	ClipLower        float32   `protobuf:"fixed32,15,opt,name=clipLower" json:"clipLower,omitempty"`
	SRSCf            int32     `protobuf:"varint,16,opt,name=sRSCf" json:"sRSCf,omitempty"`
	PixelCount       int32     `protobuf:"varint,17,opt,name=pixelCount" json:"pixelCount,omitempty"`
	VRT              string    `protobuf:"bytes,18,opt,name=vRT" json:"vRT,omitempty"`
}

func (m *GeoRPCGranule) Reset()                    { *m = GeoRPCGranule{} }
func (m *GeoRPCGranule) String() string            { return proto.CompactTextString(m) }
func (*GeoRPCGranule) ProtoMessage()               {}
func (*GeoRPCGranule) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GeoRPCGranule) GetOperation() string {
	if m != nil {
		return m.Operation
	}
	return ""
}

func (m *GeoRPCGranule) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *GeoRPCGranule) GetGeometry() string {
	if m != nil {
		return m.Geometry
	}
	return ""
}

func (m *GeoRPCGranule) GetBands() []int32 {
	if m != nil {
		return m.Bands
	}
	return nil
}

func (m *GeoRPCGranule) GetHeight() float32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *GeoRPCGranule) GetWidth() float32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *GeoRPCGranule) GetSrcSRS() string {
	if m != nil {
		return m.SrcSRS
	}
	return ""
}

func (m *GeoRPCGranule) GetSrcGeot() []float64 {
	if m != nil {
		return m.SrcGeot
	}
	return nil
}

func (m *GeoRPCGranule) GetDstSRS() string {
	if m != nil {
		return m.DstSRS
	}
	return ""
}

func (m *GeoRPCGranule) GetDstGeot() []float64 {
	if m != nil {
		return m.DstGeot
	}
	return nil
}

func (m *GeoRPCGranule) GetBandStrides() int32 {
	if m != nil {
		return m.BandStrides
	}
	return 0
}

func (m *GeoRPCGranule) GetGeoLocOpts() []string {
	if m != nil {
		return m.GeoLocOpts
	}
	return nil
}

func (m *GeoRPCGranule) GetDrillDecileCount() int32 {
	if m != nil {
		return m.DrillDecileCount
	}
	return 0
}

func (m *GeoRPCGranule) GetClipUpper() float32 {
	if m != nil {
		return m.ClipUpper
	}
	return 0
}

func (m *GeoRPCGranule) GetClipLower() float32 {
	if m != nil {
		return m.ClipLower
	}
	return 0
}

func (m *GeoRPCGranule) GetSRSCf() int32 {
	if m != nil {
		return m.SRSCf
	}
	return 0
}

func (m *GeoRPCGranule) GetPixelCount() int32 {
	if m != nil {
		return m.PixelCount
	}
	return 0
}

func (m *GeoRPCGranule) GetVRT() string {
	if m != nil {
		return m.VRT
	}
	return ""
}

type Raster struct {
	Data       []byte  `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	NoData     float64 `protobuf:"fixed64,2,opt,name=noData" json:"noData,omitempty"`
	RasterType string  `protobuf:"bytes,3,opt,name=rasterType" json:"rasterType,omitempty"`
	Bbox       []int32 `protobuf:"varint,4,rep,packed,name=bbox" json:"bbox,omitempty"`
}

func (m *Raster) Reset()                    { *m = Raster{} }
func (m *Raster) String() string            { return proto.CompactTextString(m) }
func (*Raster) ProtoMessage()               {}
func (*Raster) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Raster) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Raster) GetNoData() float64 {
	if m != nil {
		return m.NoData
	}
	return 0
}

func (m *Raster) GetRasterType() string {
	if m != nil {
		return m.RasterType
	}
	return ""
}

func (m *Raster) GetBbox() []int32 {
	if m != nil {
		return m.Bbox
	}
	return nil
}

type TimeSeries struct {
	Value float64 `protobuf:"fixed64,1,opt,name=value" json:"value,omitempty"`
	Count int32   `protobuf:"varint,2,opt,name=count" json:"count,omitempty"`
}

func (m *TimeSeries) Reset()                    { *m = TimeSeries{} }
func (m *TimeSeries) String() string            { return proto.CompactTextString(m) }
func (*TimeSeries) ProtoMessage()               {}
func (*TimeSeries) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *TimeSeries) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *TimeSeries) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type Overview struct {
	XSize int32 `protobuf:"varint,1,opt,name=xSize" json:"xSize,omitempty"`
	YSize int32 `protobuf:"varint,2,opt,name=ySize" json:"ySize,omitempty"`
}

func (m *Overview) Reset()                    { *m = Overview{} }
func (m *Overview) String() string            { return proto.CompactTextString(m) }
func (*Overview) ProtoMessage()               {}
func (*Overview) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Overview) GetXSize() int32 {
	if m != nil {
		return m.XSize
	}
	return 0
}

func (m *Overview) GetYSize() int32 {
	if m != nil {
		return m.YSize
	}
	return 0
}

type GeoMetaData struct {
	DatasetName  string                       `protobuf:"bytes,1,opt,name=datasetName" json:"datasetName,omitempty"`
	NameSpace    string                       `protobuf:"bytes,2,opt,name=nameSpace" json:"nameSpace,omitempty"`
	Type         string                       `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	RasterCount  int32                        `protobuf:"varint,4,opt,name=rasterCount" json:"rasterCount,omitempty"`
	TimeStamps   []*google_protobuf.Timestamp `protobuf:"bytes,5,rep,name=timeStamps" json:"timeStamps,omitempty"`
	Height       []float64                    `protobuf:"fixed64,6,rep,packed,name=height" json:"height,omitempty"`
	Overviews    []*Overview                  `protobuf:"bytes,7,rep,name=overviews" json:"overviews,omitempty"`
	XSize        int32                        `protobuf:"varint,8,opt,name=xSize" json:"xSize,omitempty"`
	YSize        int32                        `protobuf:"varint,9,opt,name=ySize" json:"ySize,omitempty"`
	GeoTransform []float64                    `protobuf:"fixed64,10,rep,packed,name=geoTransform" json:"geoTransform,omitempty"`
	Polygon      string                       `protobuf:"bytes,11,opt,name=polygon" json:"polygon,omitempty"`
	ProjWKT      string                       `protobuf:"bytes,12,opt,name=projWKT" json:"projWKT,omitempty"`
	Proj4        string                       `protobuf:"bytes,13,opt,name=proj4" json:"proj4,omitempty"`
}

func (m *GeoMetaData) Reset()                    { *m = GeoMetaData{} }
func (m *GeoMetaData) String() string            { return proto.CompactTextString(m) }
func (*GeoMetaData) ProtoMessage()               {}
func (*GeoMetaData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GeoMetaData) GetDatasetName() string {
	if m != nil {
		return m.DatasetName
	}
	return ""
}

func (m *GeoMetaData) GetNameSpace() string {
	if m != nil {
		return m.NameSpace
	}
	return ""
}

func (m *GeoMetaData) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *GeoMetaData) GetRasterCount() int32 {
	if m != nil {
		return m.RasterCount
	}
	return 0
}

func (m *GeoMetaData) GetTimeStamps() []*google_protobuf.Timestamp {
	if m != nil {
		return m.TimeStamps
	}
	return nil
}

func (m *GeoMetaData) GetHeight() []float64 {
	if m != nil {
		return m.Height
	}
	return nil
}

func (m *GeoMetaData) GetOverviews() []*Overview {
	if m != nil {
		return m.Overviews
	}
	return nil
}

func (m *GeoMetaData) GetXSize() int32 {
	if m != nil {
		return m.XSize
	}
	return 0
}

func (m *GeoMetaData) GetYSize() int32 {
	if m != nil {
		return m.YSize
	}
	return 0
}

func (m *GeoMetaData) GetGeoTransform() []float64 {
	if m != nil {
		return m.GeoTransform
	}
	return nil
}

func (m *GeoMetaData) GetPolygon() string {
	if m != nil {
		return m.Polygon
	}
	return ""
}

func (m *GeoMetaData) GetProjWKT() string {
	if m != nil {
		return m.ProjWKT
	}
	return ""
}

func (m *GeoMetaData) GetProj4() string {
	if m != nil {
		return m.Proj4
	}
	return ""
}

type GeoFile struct {
	FileName string         `protobuf:"bytes,1,opt,name=fileName" json:"fileName,omitempty"`
	Driver   string         `protobuf:"bytes,2,opt,name=driver" json:"driver,omitempty"`
	DataSets []*GeoMetaData `protobuf:"bytes,3,rep,name=dataSets" json:"dataSets,omitempty"`
}

func (m *GeoFile) Reset()                    { *m = GeoFile{} }
func (m *GeoFile) String() string            { return proto.CompactTextString(m) }
func (*GeoFile) ProtoMessage()               {}
func (*GeoFile) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GeoFile) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

func (m *GeoFile) GetDriver() string {
	if m != nil {
		return m.Driver
	}
	return ""
}

func (m *GeoFile) GetDataSets() []*GeoMetaData {
	if m != nil {
		return m.DataSets
	}
	return nil
}

type WorkerInfo struct {
	PoolSize int32 `protobuf:"varint,1,opt,name=poolSize" json:"poolSize,omitempty"`
}

func (m *WorkerInfo) Reset()                    { *m = WorkerInfo{} }
func (m *WorkerInfo) String() string            { return proto.CompactTextString(m) }
func (*WorkerInfo) ProtoMessage()               {}
func (*WorkerInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *WorkerInfo) GetPoolSize() int32 {
	if m != nil {
		return m.PoolSize
	}
	return 0
}

type WorkerMetrics struct {
	BytesRead     int64 `protobuf:"varint,1,opt,name=bytesRead" json:"bytesRead,omitempty"`
	UserTime      int64 `protobuf:"varint,2,opt,name=userTime" json:"userTime,omitempty"`
	SysTime       int64 `protobuf:"varint,3,opt,name=sysTime" json:"sysTime,omitempty"`
	WallTimeStart int64 `protobuf:"varint,4,opt,name=wallTimeStart" json:"wallTimeStart,omitempty"`
	WallTimeEnd   int64 `protobuf:"varint,5,opt,name=wallTimeEnd" json:"wallTimeEnd,omitempty"`
}

func (m *WorkerMetrics) Reset()                    { *m = WorkerMetrics{} }
func (m *WorkerMetrics) String() string            { return proto.CompactTextString(m) }
func (*WorkerMetrics) ProtoMessage()               {}
func (*WorkerMetrics) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *WorkerMetrics) GetBytesRead() int64 {
	if m != nil {
		return m.BytesRead
	}
	return 0
}

func (m *WorkerMetrics) GetUserTime() int64 {
	if m != nil {
		return m.UserTime
	}
	return 0
}

func (m *WorkerMetrics) GetSysTime() int64 {
	if m != nil {
		return m.SysTime
	}
	return 0
}

func (m *WorkerMetrics) GetWallTimeStart() int64 {
	if m != nil {
		return m.WallTimeStart
	}
	return 0
}

func (m *WorkerMetrics) GetWallTimeEnd() int64 {
	if m != nil {
		return m.WallTimeEnd
	}
	return 0
}

type Result struct {
	TimeSeries []*TimeSeries  `protobuf:"bytes,1,rep,name=timeSeries" json:"timeSeries,omitempty"`
	Raster     *Raster        `protobuf:"bytes,2,opt,name=raster" json:"raster,omitempty"`
	Info       *GeoFile       `protobuf:"bytes,3,opt,name=info" json:"info,omitempty"`
	Error      string         `protobuf:"bytes,4,opt,name=error" json:"error,omitempty"`
	Shape      []int32        `protobuf:"varint,5,rep,packed,name=shape" json:"shape,omitempty"`
	WorkerInfo *WorkerInfo    `protobuf:"bytes,6,opt,name=workerInfo" json:"workerInfo,omitempty"`
	Metrics    *WorkerMetrics `protobuf:"bytes,7,opt,name=metrics" json:"metrics,omitempty"`
}

func (m *Result) Reset()                    { *m = Result{} }
func (m *Result) String() string            { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()               {}
func (*Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Result) GetTimeSeries() []*TimeSeries {
	if m != nil {
		return m.TimeSeries
	}
	return nil
}

func (m *Result) GetRaster() *Raster {
	if m != nil {
		return m.Raster
	}
	return nil
}

func (m *Result) GetInfo() *GeoFile {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *Result) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *Result) GetShape() []int32 {
	if m != nil {
		return m.Shape
	}
	return nil
}

func (m *Result) GetWorkerInfo() *WorkerInfo {
	if m != nil {
		return m.WorkerInfo
	}
	return nil
}

func (m *Result) GetMetrics() *WorkerMetrics {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func init() {
	proto.RegisterType((*GeoRPCGranule)(nil), "gdalservice.GeoRPCGranule")
	proto.RegisterType((*Raster)(nil), "gdalservice.Raster")
	proto.RegisterType((*TimeSeries)(nil), "gdalservice.TimeSeries")
	proto.RegisterType((*Overview)(nil), "gdalservice.Overview")
	proto.RegisterType((*GeoMetaData)(nil), "gdalservice.GeoMetaData")
	proto.RegisterType((*GeoFile)(nil), "gdalservice.GeoFile")
	proto.RegisterType((*WorkerInfo)(nil), "gdalservice.WorkerInfo")
	proto.RegisterType((*WorkerMetrics)(nil), "gdalservice.WorkerMetrics")
	proto.RegisterType((*Result)(nil), "gdalservice.Result")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for GDAL service

type GDALClient interface {
	Process(ctx context.Context, in *GeoRPCGranule, opts ...grpc.CallOption) (*Result, error)
}

type gDALClient struct {
	cc *grpc.ClientConn
}

func NewGDALClient(cc *grpc.ClientConn) GDALClient {
	return &gDALClient{cc}
}

func (c *gDALClient) Process(ctx context.Context, in *GeoRPCGranule, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := grpc.Invoke(ctx, "/gdalservice.GDAL/Process", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GDAL service

type GDALServer interface {
	Process(context.Context, *GeoRPCGranule) (*Result, error)
}

func RegisterGDALServer(s *grpc.Server, srv GDALServer) {
	s.RegisterService(&_GDAL_serviceDesc, srv)
}

func _GDAL_Process_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GeoRPCGranule)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GDALServer).Process(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gdalservice.GDAL/Process",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GDALServer).Process(ctx, req.(*GeoRPCGranule))
	}
	return interceptor(ctx, in, info, handler)
}

var _GDAL_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gdalservice.GDAL",
	HandlerType: (*GDALServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Process",
			Handler:    _GDAL_Process_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gdalservice.proto",
}

func init() { proto.RegisterFile("gdalservice.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 911 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x55, 0x4d, 0x8f, 0xe3, 0x44,
	0x10, 0x55, 0xc6, 0xf9, 0xac, 0xcc, 0xc0, 0x6c, 0xb3, 0x40, 0x2b, 0x42, 0x10, 0x59, 0x1c, 0x22,
	0x90, 0xb2, 0xd2, 0xec, 0x08, 0xd0, 0xde, 0x60, 0x06, 0x22, 0xc4, 0x2c, 0xbb, 0xea, 0x04, 0xed,
	0xd9, 0xb1, 0x2b, 0x89, 0xc1, 0x71, 0x5b, 0xdd, 0x9d, 0x64, 0xc2, 0x0f, 0xe2, 0xc2, 0x91, 0x1f,
	0xc7, 0x15, 0x55, 0xb5, 0x1d, 0x3b, 0xb3, 0xdc, 0xfa, 0xbd, 0xaa, 0x6a, 0x57, 0xbd, 0xaa, 0x6a,
	0xc3, 0xb3, 0x75, 0x12, 0x65, 0x16, 0xcd, 0x3e, 0x8d, 0x71, 0x5a, 0x18, 0xed, 0xb4, 0x18, 0x36,
	0xa8, 0xd1, 0x17, 0x6b, 0xad, 0xd7, 0x19, 0xbe, 0x60, 0xd3, 0x72, 0xb7, 0x7a, 0xe1, 0xd2, 0x2d,
	0x5a, 0x17, 0x6d, 0x0b, 0xef, 0x1d, 0xfe, 0x1b, 0xc0, 0xd5, 0x0c, 0xb5, 0x7a, 0x7b, 0x37, 0x33,
	0x51, 0xbe, 0xcb, 0x50, 0x7c, 0x06, 0x03, 0x5d, 0xa0, 0x89, 0x5c, 0xaa, 0x73, 0xd9, 0x1a, 0xb7,
	0x26, 0x03, 0x55, 0x13, 0x42, 0x40, 0xbb, 0x88, 0xdc, 0x46, 0x5e, 0xb0, 0x81, 0xcf, 0x62, 0x04,
	0xfd, 0x35, 0xea, 0x2d, 0x3a, 0x73, 0x94, 0x01, 0xf3, 0x27, 0x2c, 0x9e, 0x43, 0x67, 0x19, 0xe5,
	0x89, 0x95, 0xed, 0x71, 0x30, 0xe9, 0x28, 0x0f, 0xc4, 0x27, 0xd0, 0xdd, 0x60, 0xba, 0xde, 0x38,
	0xd9, 0x19, 0xb7, 0x26, 0x17, 0xaa, 0x44, 0xe4, 0x7d, 0x48, 0x13, 0xb7, 0x91, 0x5d, 0xa6, 0x3d,
	0x20, 0x6f, 0x6b, 0xe2, 0xb9, 0x9a, 0xcb, 0x1e, 0xdf, 0x5e, 0x22, 0x21, 0xa1, 0x67, 0x4d, 0x3c,
	0x43, 0xed, 0x64, 0x7f, 0x1c, 0x4c, 0x5a, 0xaa, 0x82, 0x14, 0x91, 0x58, 0x47, 0x11, 0x03, 0x1f,
	0xe1, 0x11, 0x45, 0x24, 0xd6, 0x71, 0x04, 0xf8, 0x88, 0x12, 0x8a, 0x31, 0x0c, 0x29, 0xb5, 0xb9,
	0x33, 0x69, 0x82, 0x56, 0x0e, 0xc7, 0xad, 0x49, 0x47, 0x35, 0x29, 0xf1, 0x39, 0xc0, 0x1a, 0xf5,
	0x83, 0x8e, 0xdf, 0x14, 0xce, 0xca, 0xcb, 0x71, 0x30, 0x19, 0xa8, 0x06, 0x23, 0xbe, 0x82, 0xeb,
	0xc4, 0xa4, 0x59, 0x76, 0x8f, 0x71, 0x9a, 0xe1, 0x9d, 0xde, 0xe5, 0x4e, 0x5e, 0xf1, 0x35, 0xef,
	0xf1, 0xa4, 0x71, 0x9c, 0xa5, 0xc5, 0x6f, 0x45, 0x81, 0x46, 0x7e, 0xc0, 0xb5, 0xd6, 0x44, 0x65,
	0x7d, 0xd0, 0x07, 0x34, 0xf2, 0xc3, 0xda, 0xca, 0x04, 0x69, 0x64, 0xd5, 0xfc, 0x6e, 0x25, 0xaf,
	0xf9, 0x72, 0x0f, 0x28, 0xbb, 0x22, 0x7d, 0xc4, 0xcc, 0x7f, 0xf7, 0x19, 0x9b, 0x1a, 0x8c, 0xb8,
	0x86, 0x60, 0xaf, 0x16, 0x52, 0xb0, 0x1c, 0x74, 0x0c, 0x37, 0xd0, 0x55, 0x91, 0x75, 0x68, 0xa8,
	0xa7, 0x49, 0xe4, 0x22, 0x6e, 0xf6, 0xa5, 0xe2, 0x33, 0x29, 0x98, 0xeb, 0x7b, 0x62, 0xa9, 0xd3,
	0x2d, 0x55, 0x22, 0xfa, 0x8e, 0xe1, 0xa8, 0xc5, 0xb1, 0xc0, 0xb2, 0xdb, 0x0d, 0x86, 0xee, 0x5a,
	0x2e, 0xf5, 0x63, 0xd9, 0x6e, 0x3e, 0x87, 0xdf, 0x01, 0x2c, 0xd2, 0x2d, 0xce, 0xd1, 0xa4, 0x68,
	0x29, 0xff, 0x7d, 0x94, 0xed, 0x90, 0x3f, 0xd7, 0x52, 0x1e, 0x10, 0x1b, 0x73, 0xea, 0x17, 0xbe,
	0x2a, 0x06, 0xe1, 0x37, 0xd0, 0x7f, 0xb3, 0xa7, 0x51, 0xc6, 0x03, 0x79, 0x3c, 0xce, 0xd3, 0x3f,
	0x7d, 0x5c, 0x47, 0x79, 0x40, 0xec, 0x91, 0xd9, 0x32, 0x8e, 0x41, 0xf8, 0x57, 0x00, 0xc3, 0x19,
	0xea, 0xd7, 0xe8, 0x22, 0xce, 0x7a, 0x0c, 0x43, 0xaa, 0xca, 0xa2, 0xfb, 0x35, 0xda, 0x62, 0x39,
	0xd5, 0x4d, 0x8a, 0x34, 0xcf, 0xa3, 0x2d, 0xce, 0x8b, 0x28, 0xc6, 0x72, 0xb8, 0x6b, 0x82, 0xaa,
	0x72, 0x75, 0xbd, 0x7c, 0xa6, 0x3b, 0x7d, 0xdd, 0x5e, 0xf2, 0xb6, 0x9f, 0x98, 0x06, 0x25, 0x5e,
	0x01, 0xd0, 0xba, 0xcd, 0x69, 0xdd, 0xac, 0xec, 0x8c, 0x83, 0xc9, 0xf0, 0x66, 0x34, 0xf5, 0x1b,
	0x39, 0xad, 0x36, 0x72, 0xba, 0xa8, 0x36, 0x52, 0x35, 0xbc, 0x1b, 0x1b, 0xd2, 0xe5, 0x41, 0xad,
	0x36, 0xe4, 0x25, 0x0c, 0x74, 0xa9, 0x88, 0x95, 0x3d, 0xbe, 0xf2, 0xe3, 0x69, 0xf3, 0x11, 0xa8,
	0xf4, 0x52, 0xb5, 0x5f, 0x2d, 0x5d, 0xff, 0x7f, 0xa5, 0x1b, 0x34, 0xa4, 0x13, 0x21, 0x5c, 0xae,
	0x51, 0x2f, 0x4c, 0x94, 0xdb, 0x95, 0x36, 0xdb, 0x72, 0x4f, 0xce, 0x38, 0x5a, 0xa3, 0x42, 0x67,
	0xc7, 0xb5, 0xce, 0x79, 0x51, 0x06, 0xaa, 0x82, 0x6c, 0x31, 0xfa, 0xf7, 0x77, 0xbf, 0x2c, 0xe4,
	0x65, 0x69, 0xf1, 0x90, 0xbe, 0x46, 0xc7, 0x5b, 0xde, 0x89, 0x81, 0xf2, 0x20, 0xb4, 0xd0, 0x9b,
	0xa1, 0xfe, 0x29, 0xcd, 0x90, 0x5e, 0x91, 0x55, 0x9a, 0x61, 0xa3, 0x41, 0x27, 0xcc, 0xfb, 0x6c,
	0xd2, 0x3d, 0x9a, 0xb2, 0x35, 0x25, 0x12, 0xb7, 0xd0, 0xa7, 0x26, 0xce, 0xd1, 0x59, 0x19, 0xb0,
	0x18, 0xf2, 0x4c, 0x8c, 0xc6, 0x0c, 0xa8, 0x93, 0x67, 0x38, 0x01, 0x78, 0xa7, 0xcd, 0x1f, 0x68,
	0x7e, 0xce, 0x57, 0x9a, 0xbe, 0x5b, 0x68, 0x9d, 0x35, 0x46, 0xeb, 0x84, 0xc3, 0xbf, 0x5b, 0x70,
	0xe5, 0x5d, 0x5f, 0xa3, 0x33, 0x69, 0x6c, 0x69, 0x4e, 0x96, 0x47, 0x87, 0x56, 0x61, 0x94, 0xb0,
	0x7b, 0xa0, 0x6a, 0x82, 0xee, 0xda, 0x59, 0x34, 0xd4, 0x52, 0xce, 0x34, 0x50, 0x27, 0xcc, 0xaf,
	0xd5, 0xd1, 0xb2, 0x29, 0x60, 0x53, 0x05, 0xc5, 0x97, 0x70, 0x75, 0x88, 0xb2, 0x6c, 0xe1, 0xbb,
	0x6f, 0xfc, 0x2c, 0x05, 0xea, 0x9c, 0xa4, 0x79, 0xab, 0x88, 0x1f, 0xf3, 0x84, 0x1f, 0xce, 0x40,
	0x35, 0xa9, 0xf0, 0x9f, 0x0b, 0xe8, 0x2a, 0xb4, 0xbb, 0xcc, 0x89, 0x6f, 0xcb, 0xd1, 0xe3, 0x95,
	0x93, 0x2d, 0x96, 0xe6, 0xd3, 0x33, 0x69, 0xea, 0x8d, 0x54, 0x0d, 0x57, 0xf1, 0x35, 0x74, 0xfd,
	0x08, 0x73, 0xfe, 0xc3, 0x9b, 0x8f, 0xce, 0x82, 0xfc, 0x83, 0xa1, 0x4a, 0x17, 0x31, 0x81, 0x76,
	0x9a, 0xaf, 0x34, 0xd7, 0x33, 0xbc, 0x79, 0xfe, 0x54, 0x7a, 0x6a, 0xab, 0x62, 0x0f, 0xea, 0x3e,
	0x1a, 0xa3, 0x0d, 0x97, 0x36, 0x50, 0x1e, 0xf0, 0x53, 0xb6, 0x89, 0x0a, 0xe4, 0xdd, 0xa0, 0xa7,
	0x8c, 0x00, 0xe5, 0x7e, 0x38, 0xb5, 0x87, 0xff, 0x04, 0x4f, 0x73, 0xaf, 0xbb, 0xa7, 0x1a, 0xae,
	0xe2, 0x16, 0x7a, 0x5b, 0xdf, 0x26, 0xfe, 0x51, 0xf0, 0xb2, 0xbd, 0x17, 0x55, 0x36, 0x52, 0x55,
	0xae, 0x37, 0x3f, 0x40, 0x7b, 0x76, 0xff, 0xfd, 0x83, 0x78, 0x05, 0xbd, 0xb7, 0x46, 0xc7, 0x68,
	0xad, 0x18, 0x3d, 0xad, 0xa4, 0xfe, 0x3d, 0x8e, 0x9e, 0x08, 0xc2, 0x72, 0x2f, 0xbb, 0xbc, 0xcd,
	0x2f, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x5d, 0x26, 0x4b, 0xec, 0x8f, 0x07, 0x00, 0x00,
}
