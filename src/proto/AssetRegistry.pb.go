// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/AssetRegistry.proto

package assetregistry

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

type NoAssetParams struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NoAssetParams) Reset()         { *m = NoAssetParams{} }
func (m *NoAssetParams) String() string { return proto.CompactTextString(m) }
func (*NoAssetParams) ProtoMessage()    {}
func (*NoAssetParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4521fb1970e84c3, []int{0}
}

func (m *NoAssetParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NoAssetParams.Unmarshal(m, b)
}
func (m *NoAssetParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NoAssetParams.Marshal(b, m, deterministic)
}
func (m *NoAssetParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NoAssetParams.Merge(m, src)
}
func (m *NoAssetParams) XXX_Size() int {
	return xxx_messageInfo_NoAssetParams.Size(m)
}
func (m *NoAssetParams) XXX_DiscardUnknown() {
	xxx_messageInfo_NoAssetParams.DiscardUnknown(m)
}

var xxx_messageInfo_NoAssetParams proto.InternalMessageInfo

type AssetList struct {
	Assets               []*Asset `protobuf:"bytes,1,rep,name=assets,proto3" json:"assets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AssetList) Reset()         { *m = AssetList{} }
func (m *AssetList) String() string { return proto.CompactTextString(m) }
func (*AssetList) ProtoMessage()    {}
func (*AssetList) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4521fb1970e84c3, []int{1}
}

func (m *AssetList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssetList.Unmarshal(m, b)
}
func (m *AssetList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssetList.Marshal(b, m, deterministic)
}
func (m *AssetList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssetList.Merge(m, src)
}
func (m *AssetList) XXX_Size() int {
	return xxx_messageInfo_AssetList.Size(m)
}
func (m *AssetList) XXX_DiscardUnknown() {
	xxx_messageInfo_AssetList.DiscardUnknown(m)
}

var xxx_messageInfo_AssetList proto.InternalMessageInfo

func (m *AssetList) GetAssets() []*Asset {
	if m != nil {
		return m.Assets
	}
	return nil
}

type Asset struct {
	AssetId              string   `protobuf:"bytes,1,opt,name=assetId,proto3" json:"assetId,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Protocol             string   `protobuf:"bytes,4,opt,name=protocol,proto3" json:"protocol,omitempty"`
	SecondaryProtocol    string   `protobuf:"bytes,5,opt,name=secondary_protocol,json=secondaryProtocol,proto3" json:"secondary_protocol,omitempty"`
	PrimaryEndpoint      string   `protobuf:"bytes,6,opt,name=primary_endpoint,json=primaryEndpoint,proto3" json:"primary_endpoint,omitempty"`
	SecondaryEndpoint    string   `protobuf:"bytes,7,opt,name=secondary_endpoint,json=secondaryEndpoint,proto3" json:"secondary_endpoint,omitempty"`
	Timeout              string   `protobuf:"bytes,8,opt,name=timeout,proto3" json:"timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Asset) Reset()         { *m = Asset{} }
func (m *Asset) String() string { return proto.CompactTextString(m) }
func (*Asset) ProtoMessage()    {}
func (*Asset) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4521fb1970e84c3, []int{2}
}

func (m *Asset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Asset.Unmarshal(m, b)
}
func (m *Asset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Asset.Marshal(b, m, deterministic)
}
func (m *Asset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Asset.Merge(m, src)
}
func (m *Asset) XXX_Size() int {
	return xxx_messageInfo_Asset.Size(m)
}
func (m *Asset) XXX_DiscardUnknown() {
	xxx_messageInfo_Asset.DiscardUnknown(m)
}

var xxx_messageInfo_Asset proto.InternalMessageInfo

func (m *Asset) GetAssetId() string {
	if m != nil {
		return m.AssetId
	}
	return ""
}

func (m *Asset) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Asset) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Asset) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *Asset) GetSecondaryProtocol() string {
	if m != nil {
		return m.SecondaryProtocol
	}
	return ""
}

func (m *Asset) GetPrimaryEndpoint() string {
	if m != nil {
		return m.PrimaryEndpoint
	}
	return ""
}

func (m *Asset) GetSecondaryEndpoint() string {
	if m != nil {
		return m.SecondaryEndpoint
	}
	return ""
}

func (m *Asset) GetTimeout() string {
	if m != nil {
		return m.Timeout
	}
	return ""
}

func init() {
	proto.RegisterType((*NoAssetParams)(nil), "NoAssetParams")
	proto.RegisterType((*AssetList)(nil), "AssetList")
	proto.RegisterType((*Asset)(nil), "Asset")
}

func init() { proto.RegisterFile("proto/AssetRegistry.proto", fileDescriptor_f4521fb1970e84c3) }

var fileDescriptor_f4521fb1970e84c3 = []byte{
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0xdd, 0xbf, 0x6e, 0x7b, 0xe7, 0x9c, 0xe6, 0x14, 0x77, 0x90, 0xd2, 0x8b, 0x1b, 0x6a,
	0x0a, 0xf3, 0xec, 0x61, 0x82, 0x88, 0x20, 0x32, 0x76, 0xf4, 0x22, 0x59, 0x1b, 0x66, 0xb4, 0x49,
	0x4a, 0x92, 0x1d, 0xfa, 0x45, 0xfc, 0xbc, 0xb2, 0x77, 0x5d, 0xb0, 0xde, 0xfa, 0xfc, 0xde, 0x1f,
	0x0f, 0xcd, 0x03, 0x97, 0xa5, 0x35, 0xde, 0xa4, 0x4b, 0xe7, 0x84, 0x5f, 0x8b, 0xad, 0x74, 0xde,
	0x56, 0x0c, 0x59, 0x32, 0x81, 0xf1, 0x9b, 0xc1, 0xc3, 0x8a, 0x5b, 0xae, 0x5c, 0x72, 0x03, 0x43,
	0x8c, 0xaf, 0xd2, 0x79, 0x72, 0x05, 0x11, 0xdf, 0x07, 0x47, 0x5b, 0x71, 0x67, 0x36, 0x5a, 0x44,
	0xec, 0xd0, 0x51, 0xd3, 0xe4, 0xa7, 0x0d, 0x3d, 0x24, 0x84, 0x42, 0x1f, 0xd9, 0x4b, 0x4e, 0x5b,
	0x71, 0x6b, 0x36, 0x5c, 0x1f, 0x23, 0x21, 0xd0, 0xd5, 0x5c, 0x09, 0xda, 0x46, 0x8c, 0xdf, 0x24,
	0x86, 0x51, 0x2e, 0x5c, 0x66, 0x65, 0xe9, 0xa5, 0xd1, 0xb4, 0x83, 0xa7, 0xbf, 0x88, 0x4c, 0x61,
	0x80, 0x3f, 0x98, 0x99, 0x82, 0x76, 0xf1, 0x1c, 0x32, 0xb9, 0x03, 0xe2, 0x44, 0x66, 0x74, 0xce,
	0x6d, 0xf5, 0x11, 0xac, 0x1e, 0x5a, 0x17, 0xe1, 0xb2, 0x3a, 0xea, 0x73, 0x38, 0x2f, 0xad, 0x54,
	0x7b, 0x59, 0xe8, 0xbc, 0x34, 0x52, 0x7b, 0x1a, 0xa1, 0x3c, 0xa9, 0xf9, 0x53, 0x8d, 0x9b, 0xcd,
	0x41, 0xee, 0xff, 0x6b, 0x0e, 0x3a, 0x85, 0xbe, 0x97, 0x4a, 0x98, 0x9d, 0xa7, 0x83, 0xc3, 0xa3,
	0xeb, 0xb8, 0x78, 0x80, 0x71, 0x63, 0x6d, 0x72, 0x0b, 0xa7, 0xcf, 0xc2, 0x2f, 0x8b, 0x02, 0xb1,
	0x23, 0x67, 0xac, 0x31, 0xfb, 0x14, 0x58, 0x58, 0x3d, 0x39, 0x79, 0x9c, 0xbf, 0x5f, 0xb3, 0x74,
	0x2b, 0xfd, 0xe7, 0x6e, 0xc3, 0x32, 0xa3, 0xd2, 0xef, 0x2f, 0x5e, 0x71, 0xed, 0x95, 0xd0, 0x46,
	0xa7, 0x38, 0xac, 0xad, 0x8b, 0x37, 0x11, 0x0e, 0x70, 0xff, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x8b,
	0x3a, 0x3b, 0xc8, 0xe4, 0x01, 0x00, 0x00,
}
