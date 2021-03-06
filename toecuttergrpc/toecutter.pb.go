// Code generated by protoc-gen-go. DO NOT EDIT.
// source: toecutter.proto

package toecuttergrpc

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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_9669e087ad5603ae, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

// RespHashes ...
type RespHashes struct {
	Message              uint64   `protobuf:"varint,1,opt,name=message,proto3" json:"message,omitempty"`
	Status               string   `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RespHashes) Reset()         { *m = RespHashes{} }
func (m *RespHashes) String() string { return proto.CompactTextString(m) }
func (*RespHashes) ProtoMessage()    {}
func (*RespHashes) Descriptor() ([]byte, []int) {
	return fileDescriptor_9669e087ad5603ae, []int{1}
}

func (m *RespHashes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RespHashes.Unmarshal(m, b)
}
func (m *RespHashes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RespHashes.Marshal(b, m, deterministic)
}
func (m *RespHashes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RespHashes.Merge(m, src)
}
func (m *RespHashes) XXX_Size() int {
	return xxx_messageInfo_RespHashes.Size(m)
}
func (m *RespHashes) XXX_DiscardUnknown() {
	xxx_messageInfo_RespHashes.DiscardUnknown(m)
}

var xxx_messageInfo_RespHashes proto.InternalMessageInfo

func (m *RespHashes) GetMessage() uint64 {
	if m != nil {
		return m.Message
	}
	return 0
}

func (m *RespHashes) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type RespReward struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Status               string   `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RespReward) Reset()         { *m = RespReward{} }
func (m *RespReward) String() string { return proto.CompactTextString(m) }
func (*RespReward) ProtoMessage()    {}
func (*RespReward) Descriptor() ([]byte, []int) {
	return fileDescriptor_9669e087ad5603ae, []int{2}
}

func (m *RespReward) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RespReward.Unmarshal(m, b)
}
func (m *RespReward) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RespReward.Marshal(b, m, deterministic)
}
func (m *RespReward) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RespReward.Merge(m, src)
}
func (m *RespReward) XXX_Size() int {
	return xxx_messageInfo_RespReward.Size(m)
}
func (m *RespReward) XXX_DiscardUnknown() {
	xxx_messageInfo_RespReward.DiscardUnknown(m)
}

var xxx_messageInfo_RespReward proto.InternalMessageInfo

func (m *RespReward) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *RespReward) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type RespStats struct {
	Hasherate            float32  `protobuf:"fixed32,1,opt,name=hasherate,proto3" json:"hasherate,omitempty"`
	Hashes               uint64   `protobuf:"varint,2,opt,name=hashes,proto3" json:"hashes,omitempty"`
	Reward               string   `protobuf:"bytes,3,opt,name=reward,proto3" json:"reward,omitempty"`
	Rewardpending        string   `protobuf:"bytes,4,opt,name=rewardpending,proto3" json:"rewardpending,omitempty"`
	Status               string   `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RespStats) Reset()         { *m = RespStats{} }
func (m *RespStats) String() string { return proto.CompactTextString(m) }
func (*RespStats) ProtoMessage()    {}
func (*RespStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_9669e087ad5603ae, []int{3}
}

func (m *RespStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RespStats.Unmarshal(m, b)
}
func (m *RespStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RespStats.Marshal(b, m, deterministic)
}
func (m *RespStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RespStats.Merge(m, src)
}
func (m *RespStats) XXX_Size() int {
	return xxx_messageInfo_RespStats.Size(m)
}
func (m *RespStats) XXX_DiscardUnknown() {
	xxx_messageInfo_RespStats.DiscardUnknown(m)
}

var xxx_messageInfo_RespStats proto.InternalMessageInfo

func (m *RespStats) GetHasherate() float32 {
	if m != nil {
		return m.Hasherate
	}
	return 0
}

func (m *RespStats) GetHashes() uint64 {
	if m != nil {
		return m.Hashes
	}
	return 0
}

func (m *RespStats) GetReward() string {
	if m != nil {
		return m.Reward
	}
	return ""
}

func (m *RespStats) GetRewardpending() string {
	if m != nil {
		return m.Rewardpending
	}
	return ""
}

func (m *RespStats) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type Site struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Sitekey              string   `protobuf:"bytes,2,opt,name=sitekey,proto3" json:"sitekey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Site) Reset()         { *m = Site{} }
func (m *Site) String() string { return proto.CompactTextString(m) }
func (*Site) ProtoMessage()    {}
func (*Site) Descriptor() ([]byte, []int) {
	return fileDescriptor_9669e087ad5603ae, []int{4}
}

func (m *Site) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Site.Unmarshal(m, b)
}
func (m *Site) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Site.Marshal(b, m, deterministic)
}
func (m *Site) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Site.Merge(m, src)
}
func (m *Site) XXX_Size() int {
	return xxx_messageInfo_Site.Size(m)
}
func (m *Site) XXX_DiscardUnknown() {
	xxx_messageInfo_Site.DiscardUnknown(m)
}

var xxx_messageInfo_Site proto.InternalMessageInfo

func (m *Site) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Site) GetSitekey() string {
	if m != nil {
		return m.Sitekey
	}
	return ""
}

type RespSite struct {
	Sites                []*Site  `protobuf:"bytes,1,rep,name=sites,proto3" json:"sites,omitempty"`
	Status               string   `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RespSite) Reset()         { *m = RespSite{} }
func (m *RespSite) String() string { return proto.CompactTextString(m) }
func (*RespSite) ProtoMessage()    {}
func (*RespSite) Descriptor() ([]byte, []int) {
	return fileDescriptor_9669e087ad5603ae, []int{5}
}

func (m *RespSite) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RespSite.Unmarshal(m, b)
}
func (m *RespSite) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RespSite.Marshal(b, m, deterministic)
}
func (m *RespSite) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RespSite.Merge(m, src)
}
func (m *RespSite) XXX_Size() int {
	return xxx_messageInfo_RespSite.Size(m)
}
func (m *RespSite) XXX_DiscardUnknown() {
	xxx_messageInfo_RespSite.DiscardUnknown(m)
}

var xxx_messageInfo_RespSite proto.InternalMessageInfo

func (m *RespSite) GetSites() []*Site {
	if m != nil {
		return m.Sites
	}
	return nil
}

func (m *RespSite) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type Payment struct {
	Amount               string   `protobuf:"bytes,1,opt,name=amount,proto3" json:"amount,omitempty"`
	Fee                  string   `protobuf:"bytes,2,opt,name=fee,proto3" json:"fee,omitempty"`
	Status               string   `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	Walletaddress        string   `protobuf:"bytes,4,opt,name=walletaddress,proto3" json:"walletaddress,omitempty"`
	Txid                 string   `protobuf:"bytes,5,opt,name=txid,proto3" json:"txid,omitempty"`
	Timestamp            uint64   `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Payment) Reset()         { *m = Payment{} }
func (m *Payment) String() string { return proto.CompactTextString(m) }
func (*Payment) ProtoMessage()    {}
func (*Payment) Descriptor() ([]byte, []int) {
	return fileDescriptor_9669e087ad5603ae, []int{6}
}

func (m *Payment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Payment.Unmarshal(m, b)
}
func (m *Payment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Payment.Marshal(b, m, deterministic)
}
func (m *Payment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Payment.Merge(m, src)
}
func (m *Payment) XXX_Size() int {
	return xxx_messageInfo_Payment.Size(m)
}
func (m *Payment) XXX_DiscardUnknown() {
	xxx_messageInfo_Payment.DiscardUnknown(m)
}

var xxx_messageInfo_Payment proto.InternalMessageInfo

func (m *Payment) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func (m *Payment) GetFee() string {
	if m != nil {
		return m.Fee
	}
	return ""
}

func (m *Payment) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Payment) GetWalletaddress() string {
	if m != nil {
		return m.Walletaddress
	}
	return ""
}

func (m *Payment) GetTxid() string {
	if m != nil {
		return m.Txid
	}
	return ""
}

func (m *Payment) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type RespPayments struct {
	Payments             []*Payment `protobuf:"bytes,1,rep,name=payments,proto3" json:"payments,omitempty"`
	Status               string     `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *RespPayments) Reset()         { *m = RespPayments{} }
func (m *RespPayments) String() string { return proto.CompactTextString(m) }
func (*RespPayments) ProtoMessage()    {}
func (*RespPayments) Descriptor() ([]byte, []int) {
	return fileDescriptor_9669e087ad5603ae, []int{7}
}

func (m *RespPayments) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RespPayments.Unmarshal(m, b)
}
func (m *RespPayments) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RespPayments.Marshal(b, m, deterministic)
}
func (m *RespPayments) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RespPayments.Merge(m, src)
}
func (m *RespPayments) XXX_Size() int {
	return xxx_messageInfo_RespPayments.Size(m)
}
func (m *RespPayments) XXX_DiscardUnknown() {
	xxx_messageInfo_RespPayments.DiscardUnknown(m)
}

var xxx_messageInfo_RespPayments proto.InternalMessageInfo

func (m *RespPayments) GetPayments() []*Payment {
	if m != nil {
		return m.Payments
	}
	return nil
}

func (m *RespPayments) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type RespSiteStats struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Hashrate             float32  `protobuf:"fixed32,2,opt,name=hashrate,proto3" json:"hashrate,omitempty"`
	Hashes               uint64   `protobuf:"varint,3,opt,name=hashes,proto3" json:"hashes,omitempty"`
	Reward               string   `protobuf:"bytes,4,opt,name=reward,proto3" json:"reward,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RespSiteStats) Reset()         { *m = RespSiteStats{} }
func (m *RespSiteStats) String() string { return proto.CompactTextString(m) }
func (*RespSiteStats) ProtoMessage()    {}
func (*RespSiteStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_9669e087ad5603ae, []int{8}
}

func (m *RespSiteStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RespSiteStats.Unmarshal(m, b)
}
func (m *RespSiteStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RespSiteStats.Marshal(b, m, deterministic)
}
func (m *RespSiteStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RespSiteStats.Merge(m, src)
}
func (m *RespSiteStats) XXX_Size() int {
	return xxx_messageInfo_RespSiteStats.Size(m)
}
func (m *RespSiteStats) XXX_DiscardUnknown() {
	xxx_messageInfo_RespSiteStats.DiscardUnknown(m)
}

var xxx_messageInfo_RespSiteStats proto.InternalMessageInfo

func (m *RespSiteStats) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RespSiteStats) GetHashrate() float32 {
	if m != nil {
		return m.Hashrate
	}
	return 0
}

func (m *RespSiteStats) GetHashes() uint64 {
	if m != nil {
		return m.Hashes
	}
	return 0
}

func (m *RespSiteStats) GetReward() string {
	if m != nil {
		return m.Reward
	}
	return ""
}

type RespPayout1MHash struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Status               string   `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RespPayout1MHash) Reset()         { *m = RespPayout1MHash{} }
func (m *RespPayout1MHash) String() string { return proto.CompactTextString(m) }
func (*RespPayout1MHash) ProtoMessage()    {}
func (*RespPayout1MHash) Descriptor() ([]byte, []int) {
	return fileDescriptor_9669e087ad5603ae, []int{9}
}

func (m *RespPayout1MHash) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RespPayout1MHash.Unmarshal(m, b)
}
func (m *RespPayout1MHash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RespPayout1MHash.Marshal(b, m, deterministic)
}
func (m *RespPayout1MHash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RespPayout1MHash.Merge(m, src)
}
func (m *RespPayout1MHash) XXX_Size() int {
	return xxx_messageInfo_RespPayout1MHash.Size(m)
}
func (m *RespPayout1MHash) XXX_DiscardUnknown() {
	xxx_messageInfo_RespPayout1MHash.DiscardUnknown(m)
}

var xxx_messageInfo_RespPayout1MHash proto.InternalMessageInfo

func (m *RespPayout1MHash) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func (m *RespPayout1MHash) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*Empty)(nil), "toecuttergrpc.Empty")
	proto.RegisterType((*RespHashes)(nil), "toecuttergrpc.RespHashes")
	proto.RegisterType((*RespReward)(nil), "toecuttergrpc.RespReward")
	proto.RegisterType((*RespStats)(nil), "toecuttergrpc.RespStats")
	proto.RegisterType((*Site)(nil), "toecuttergrpc.Site")
	proto.RegisterType((*RespSite)(nil), "toecuttergrpc.RespSite")
	proto.RegisterType((*Payment)(nil), "toecuttergrpc.Payment")
	proto.RegisterType((*RespPayments)(nil), "toecuttergrpc.RespPayments")
	proto.RegisterType((*RespSiteStats)(nil), "toecuttergrpc.RespSiteStats")
	proto.RegisterType((*RespPayout1MHash)(nil), "toecuttergrpc.RespPayout1MHash")
}

func init() {
	proto.RegisterFile("toecutter.proto", fileDescriptor_9669e087ad5603ae)
}

var fileDescriptor_9669e087ad5603ae = []byte{
	// 566 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0xb5, 0x63, 0xe7, 0x75, 0x9b, 0xa8, 0xd5, 0x80, 0x8a, 0x29, 0x95, 0x88, 0x46, 0x08, 0x85,
	0x8d, 0x25, 0x02, 0x1b, 0x84, 0x54, 0xd1, 0x20, 0x1e, 0x9b, 0x4a, 0x95, 0xbb, 0x63, 0x37, 0xb1,
	0x87, 0xc4, 0x6a, 0xfc, 0x90, 0xe7, 0x5a, 0x25, 0x5f, 0xc2, 0x07, 0xf0, 0x55, 0xfc, 0x0d, 0x9a,
	0x87, 0x9d, 0x38, 0xb2, 0x45, 0x77, 0x73, 0x6f, 0xee, 0x39, 0xb9, 0xe7, 0xcc, 0x19, 0xc3, 0x29,
	0x66, 0x3c, 0x2c, 0x11, 0x79, 0xe1, 0xe7, 0x45, 0x86, 0x19, 0x99, 0xd6, 0x8d, 0x75, 0x91, 0x87,
	0x74, 0x08, 0xfd, 0x2f, 0x49, 0x8e, 0x3b, 0x7a, 0x05, 0x10, 0x70, 0x91, 0x7f, 0x67, 0x62, 0xc3,
	0x05, 0xf1, 0x60, 0x98, 0x70, 0x21, 0xd8, 0x9a, 0x7b, 0xf6, 0xcc, 0x9e, 0xbb, 0x41, 0x55, 0x92,
	0x73, 0x18, 0x08, 0x64, 0x58, 0x0a, 0xaf, 0x37, 0xb3, 0xe7, 0xe3, 0xc0, 0x54, 0x15, 0x3e, 0xe0,
	0x0f, 0xac, 0x88, 0x8e, 0xf1, 0xe3, 0xff, 0xe3, 0x7f, 0xdb, 0x30, 0x96, 0x04, 0x77, 0xc8, 0x50,
	0x90, 0x4b, 0x18, 0x6f, 0xe4, 0x26, 0x05, 0x43, 0xcd, 0xd0, 0x0b, 0xf6, 0x0d, 0xc9, 0xa1, 0x0a,
	0xcd, 0xe1, 0x06, 0xa6, 0x92, 0xfd, 0x42, 0xfd, 0xbf, 0xe7, 0x68, 0x6e, 0x5d, 0x91, 0x57, 0x30,
	0xd5, 0xa7, 0x9c, 0xa7, 0x51, 0x9c, 0xae, 0x3d, 0x57, 0xfd, 0xdc, 0x6c, 0x1e, 0x6c, 0xd6, 0x6f,
	0x6c, 0xf6, 0x1e, 0xdc, 0xbb, 0x18, 0x39, 0x21, 0xe0, 0xa6, 0x2c, 0xa9, 0x04, 0xa9, 0xb3, 0xd4,
	0x29, 0x62, 0xe4, 0xf7, 0x7c, 0x67, 0xe4, 0x54, 0x25, 0xbd, 0x81, 0x91, 0x92, 0x23, 0x91, 0x6f,
	0xa0, 0x2f, 0xdb, 0xc2, 0xb3, 0x67, 0xce, 0xfc, 0x64, 0xf1, 0xc4, 0x6f, 0xdc, 0x81, 0x2f, 0x67,
	0x02, 0x3d, 0xd1, 0x69, 0xcf, 0x1f, 0x1b, 0x86, 0xb7, 0x6c, 0x97, 0xf0, 0x14, 0xe5, 0x0c, 0x4b,
	0xb2, 0x32, 0x45, 0xb3, 0x8a, 0xa9, 0xc8, 0x19, 0x38, 0x3f, 0x39, 0x37, 0x40, 0x79, 0x3c, 0x60,
	0x73, 0x0e, 0xd9, 0xa4, 0x21, 0x0f, 0x6c, 0xbb, 0xe5, 0xc8, 0xa2, 0xa8, 0xe0, 0x42, 0x54, 0x86,
	0x34, 0x9a, 0x52, 0x30, 0xfe, 0x8a, 0x23, 0x63, 0x87, 0x3a, 0xcb, 0x8b, 0xc1, 0x38, 0xe1, 0x02,
	0x59, 0x92, 0x7b, 0x03, 0xe5, 0xfe, 0xbe, 0x41, 0x7f, 0xc0, 0x44, 0x8a, 0x36, 0x8b, 0x0a, 0xb2,
	0x80, 0x51, 0x6e, 0xce, 0x46, 0xfb, 0xf9, 0x91, 0x76, 0x33, 0x1a, 0xd4, 0x73, 0x9d, 0x0e, 0x64,
	0x30, 0xad, 0x0c, 0xd5, 0x19, 0x69, 0xbb, 0x8f, 0x0b, 0x18, 0xc9, 0x2c, 0xa8, 0xd8, 0xf4, 0x54,
	0x6c, 0xea, 0xfa, 0x20, 0x35, 0x4e, 0x47, 0x6a, 0xdc, 0xc3, 0xd4, 0xd0, 0x25, 0x9c, 0x19, 0x31,
	0x59, 0x89, 0x6f, 0x6f, 0xe4, 0xcb, 0xd0, 0xb3, 0xa2, 0xdc, 0xd6, 0xd6, 0xeb, 0xaa, 0x6b, 0xe9,
	0xc5, 0x5f, 0x07, 0x86, 0x9f, 0xb3, 0x38, 0x8d, 0x93, 0x9c, 0x7c, 0x84, 0x81, 0x79, 0x5d, 0x4f,
	0x8f, 0x4c, 0x50, 0x2f, 0xf0, 0xe2, 0xf9, 0x51, 0x77, 0xff, 0x1c, 0xa9, 0x25, 0xc1, 0xe6, 0x69,
	0x3d, 0x1e, 0xac, 0x01, 0xd4, 0x22, 0x9f, 0x60, 0x72, 0x1d, 0x86, 0x32, 0x23, 0xda, 0xb9, 0x76,
	0x0a, 0xaf, 0x85, 0x42, 0xcd, 0x53, 0x8b, 0x5c, 0xc1, 0x49, 0xc5, 0x20, 0x03, 0xdd, 0x4e, 0xf0,
	0xac, 0x8d, 0x20, 0x46, 0x4e, 0x2d, 0xf2, 0x15, 0x4e, 0x0d, 0xbe, 0xce, 0x46, 0x3b, 0xc7, 0x8b,
	0x16, 0x8e, 0x0a, 0x42, 0x2d, 0x72, 0x0d, 0xe3, 0x7d, 0x00, 0xda, 0x19, 0x2e, 0x3b, 0xb6, 0xa8,
	0xa4, 0x7c, 0x83, 0x49, 0xe3, 0x4a, 0xdb, 0x59, 0x5e, 0xb6, 0xef, 0x51, 0xc3, 0xa8, 0xb5, 0xfc,
	0x00, 0xaf, 0xc3, 0x2c, 0xf1, 0xd7, 0x31, 0x6e, 0xca, 0x95, 0xbf, 0xd9, 0xe5, 0x59, 0xc4, 0x90,
	0xad, 0x58, 0x7a, 0xef, 0x6f, 0xb3, 0x90, 0x6d, 0x43, 0x16, 0x6e, 0xb8, 0x04, 0x2f, 0x9b, 0xdf,
	0xdc, 0x5b, 0x7b, 0x35, 0x50, 0xdf, 0xe2, 0x77, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xbf, 0x62,
	0xf6, 0x32, 0x9e, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CoinimpClient is the client API for Coinimp service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CoinimpClient interface {
	// receipt coinimp hashes
	Hashes(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*RespHashes, error)
	// receipt coinimp reward
	Reward(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*RespReward, error)
	// receipt coinimp account stats
	AccountStats(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*RespStats, error)
	// receipt coinimp account sit
	AccountSite(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*RespSite, error)
	// receipt coinimp account payment
	AccountPayments(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*RespPayments, error)
	// receipt coinimp site stats
	SiteStats(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*RespSiteStats, error)
	// receipt coinimp payout 1Mhash
	Payout1MHash(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*RespPayout1MHash, error)
}

type coinimpClient struct {
	cc grpc.ClientConnInterface
}

func NewCoinimpClient(cc grpc.ClientConnInterface) CoinimpClient {
	return &coinimpClient{cc}
}

func (c *coinimpClient) Hashes(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*RespHashes, error) {
	out := new(RespHashes)
	err := c.cc.Invoke(ctx, "/toecuttergrpc.Coinimp/Hashes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coinimpClient) Reward(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*RespReward, error) {
	out := new(RespReward)
	err := c.cc.Invoke(ctx, "/toecuttergrpc.Coinimp/Reward", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coinimpClient) AccountStats(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*RespStats, error) {
	out := new(RespStats)
	err := c.cc.Invoke(ctx, "/toecuttergrpc.Coinimp/AccountStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coinimpClient) AccountSite(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*RespSite, error) {
	out := new(RespSite)
	err := c.cc.Invoke(ctx, "/toecuttergrpc.Coinimp/AccountSite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coinimpClient) AccountPayments(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*RespPayments, error) {
	out := new(RespPayments)
	err := c.cc.Invoke(ctx, "/toecuttergrpc.Coinimp/AccountPayments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coinimpClient) SiteStats(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*RespSiteStats, error) {
	out := new(RespSiteStats)
	err := c.cc.Invoke(ctx, "/toecuttergrpc.Coinimp/SiteStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coinimpClient) Payout1MHash(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*RespPayout1MHash, error) {
	out := new(RespPayout1MHash)
	err := c.cc.Invoke(ctx, "/toecuttergrpc.Coinimp/Payout1MHash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoinimpServer is the server API for Coinimp service.
type CoinimpServer interface {
	// receipt coinimp hashes
	Hashes(context.Context, *Empty) (*RespHashes, error)
	// receipt coinimp reward
	Reward(context.Context, *Empty) (*RespReward, error)
	// receipt coinimp account stats
	AccountStats(context.Context, *Empty) (*RespStats, error)
	// receipt coinimp account sit
	AccountSite(context.Context, *Empty) (*RespSite, error)
	// receipt coinimp account payment
	AccountPayments(context.Context, *Empty) (*RespPayments, error)
	// receipt coinimp site stats
	SiteStats(context.Context, *Empty) (*RespSiteStats, error)
	// receipt coinimp payout 1Mhash
	Payout1MHash(context.Context, *Empty) (*RespPayout1MHash, error)
}

// UnimplementedCoinimpServer can be embedded to have forward compatible implementations.
type UnimplementedCoinimpServer struct {
}

func (*UnimplementedCoinimpServer) Hashes(ctx context.Context, req *Empty) (*RespHashes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hashes not implemented")
}
func (*UnimplementedCoinimpServer) Reward(ctx context.Context, req *Empty) (*RespReward, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reward not implemented")
}
func (*UnimplementedCoinimpServer) AccountStats(ctx context.Context, req *Empty) (*RespStats, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccountStats not implemented")
}
func (*UnimplementedCoinimpServer) AccountSite(ctx context.Context, req *Empty) (*RespSite, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccountSite not implemented")
}
func (*UnimplementedCoinimpServer) AccountPayments(ctx context.Context, req *Empty) (*RespPayments, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccountPayments not implemented")
}
func (*UnimplementedCoinimpServer) SiteStats(ctx context.Context, req *Empty) (*RespSiteStats, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SiteStats not implemented")
}
func (*UnimplementedCoinimpServer) Payout1MHash(ctx context.Context, req *Empty) (*RespPayout1MHash, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Payout1MHash not implemented")
}

func RegisterCoinimpServer(s *grpc.Server, srv CoinimpServer) {
	s.RegisterService(&_Coinimp_serviceDesc, srv)
}

func _Coinimp_Hashes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoinimpServer).Hashes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/toecuttergrpc.Coinimp/Hashes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoinimpServer).Hashes(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Coinimp_Reward_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoinimpServer).Reward(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/toecuttergrpc.Coinimp/Reward",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoinimpServer).Reward(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Coinimp_AccountStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoinimpServer).AccountStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/toecuttergrpc.Coinimp/AccountStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoinimpServer).AccountStats(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Coinimp_AccountSite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoinimpServer).AccountSite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/toecuttergrpc.Coinimp/AccountSite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoinimpServer).AccountSite(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Coinimp_AccountPayments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoinimpServer).AccountPayments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/toecuttergrpc.Coinimp/AccountPayments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoinimpServer).AccountPayments(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Coinimp_SiteStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoinimpServer).SiteStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/toecuttergrpc.Coinimp/SiteStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoinimpServer).SiteStats(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Coinimp_Payout1MHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoinimpServer).Payout1MHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/toecuttergrpc.Coinimp/Payout1MHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoinimpServer).Payout1MHash(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Coinimp_serviceDesc = grpc.ServiceDesc{
	ServiceName: "toecuttergrpc.Coinimp",
	HandlerType: (*CoinimpServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hashes",
			Handler:    _Coinimp_Hashes_Handler,
		},
		{
			MethodName: "Reward",
			Handler:    _Coinimp_Reward_Handler,
		},
		{
			MethodName: "AccountStats",
			Handler:    _Coinimp_AccountStats_Handler,
		},
		{
			MethodName: "AccountSite",
			Handler:    _Coinimp_AccountSite_Handler,
		},
		{
			MethodName: "AccountPayments",
			Handler:    _Coinimp_AccountPayments_Handler,
		},
		{
			MethodName: "SiteStats",
			Handler:    _Coinimp_SiteStats_Handler,
		},
		{
			MethodName: "Payout1MHash",
			Handler:    _Coinimp_Payout1MHash_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "toecutter.proto",
}
