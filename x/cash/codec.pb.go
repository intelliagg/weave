// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/cash/codec.proto

package cash

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_iov_one_weave "github.com/iov-one/weave"
	weave "github.com/iov-one/weave"
	coin "github.com/iov-one/weave/coin"
	io "io"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Set may contain Coin of many different currencies.
// It handles adding and subtracting sets of currencies.
type Set struct {
	Metadata *weave.Metadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Coins    []*coin.Coin    `protobuf:"bytes,2,rep,name=coins,proto3" json:"coins,omitempty"`
}

func (m *Set) Reset()         { *m = Set{} }
func (m *Set) String() string { return proto.CompactTextString(m) }
func (*Set) ProtoMessage()    {}
func (*Set) Descriptor() ([]byte, []int) {
	return fileDescriptor_7149e4b58e322390, []int{0}
}
func (m *Set) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Set) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Set.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Set) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Set.Merge(m, src)
}
func (m *Set) XXX_Size() int {
	return m.Size()
}
func (m *Set) XXX_DiscardUnknown() {
	xxx_messageInfo_Set.DiscardUnknown(m)
}

var xxx_messageInfo_Set proto.InternalMessageInfo

func (m *Set) GetMetadata() *weave.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Set) GetCoins() []*coin.Coin {
	if m != nil {
		return m.Coins
	}
	return nil
}

// SendMsg is a request to move these coins from the given
// source to the given destination address.
// memo is an optional human-readable message
// ref is optional binary data, that can refer to another
// eg. tx hash
type SendMsg struct {
	Metadata *weave.Metadata                  `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Src      github_com_iov_one_weave.Address `protobuf:"bytes,2,opt,name=src,proto3,casttype=github.com/iov-one/weave.Address" json:"src,omitempty"`
	Dest     github_com_iov_one_weave.Address `protobuf:"bytes,3,opt,name=dest,proto3,casttype=github.com/iov-one/weave.Address" json:"dest,omitempty"`
	Amount   *coin.Coin                       `protobuf:"bytes,4,opt,name=amount,proto3" json:"amount,omitempty"`
	// max length 128 character
	Memo string `protobuf:"bytes,5,opt,name=memo,proto3" json:"memo,omitempty"`
	// max length 64 bytes
	Ref []byte `protobuf:"bytes,6,opt,name=ref,proto3" json:"ref,omitempty"`
}

func (m *SendMsg) Reset()         { *m = SendMsg{} }
func (m *SendMsg) String() string { return proto.CompactTextString(m) }
func (*SendMsg) ProtoMessage()    {}
func (*SendMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_7149e4b58e322390, []int{1}
}
func (m *SendMsg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SendMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SendMsg.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SendMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendMsg.Merge(m, src)
}
func (m *SendMsg) XXX_Size() int {
	return m.Size()
}
func (m *SendMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_SendMsg.DiscardUnknown(m)
}

var xxx_messageInfo_SendMsg proto.InternalMessageInfo

func (m *SendMsg) GetMetadata() *weave.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *SendMsg) GetSrc() github_com_iov_one_weave.Address {
	if m != nil {
		return m.Src
	}
	return nil
}

func (m *SendMsg) GetDest() github_com_iov_one_weave.Address {
	if m != nil {
		return m.Dest
	}
	return nil
}

func (m *SendMsg) GetAmount() *coin.Coin {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *SendMsg) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

func (m *SendMsg) GetRef() []byte {
	if m != nil {
		return m.Ref
	}
	return nil
}

// FeeInfo records who pays what fees to have this
// message processed
type FeeInfo struct {
	Metadata *weave.Metadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Payer    []byte          `protobuf:"bytes,2,opt,name=payer,proto3" json:"payer,omitempty"`
	Fees     *coin.Coin      `protobuf:"bytes,3,opt,name=fees,proto3" json:"fees,omitempty"`
}

func (m *FeeInfo) Reset()         { *m = FeeInfo{} }
func (m *FeeInfo) String() string { return proto.CompactTextString(m) }
func (*FeeInfo) ProtoMessage()    {}
func (*FeeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_7149e4b58e322390, []int{2}
}
func (m *FeeInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FeeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FeeInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FeeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeeInfo.Merge(m, src)
}
func (m *FeeInfo) XXX_Size() int {
	return m.Size()
}
func (m *FeeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_FeeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_FeeInfo proto.InternalMessageInfo

func (m *FeeInfo) GetMetadata() *weave.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *FeeInfo) GetPayer() []byte {
	if m != nil {
		return m.Payer
	}
	return nil
}

func (m *FeeInfo) GetFees() *coin.Coin {
	if m != nil {
		return m.Fees
	}
	return nil
}

type Configuration struct {
	// TODO: add schema uint32 here
	Owner            github_com_iov_one_weave.Address `protobuf:"bytes,2,opt,name=owner,proto3,casttype=github.com/iov-one/weave.Address" json:"owner,omitempty"`
	CollectorAddress github_com_iov_one_weave.Address `protobuf:"bytes,3,opt,name=collector_address,json=collectorAddress,proto3,casttype=github.com/iov-one/weave.Address" json:"collector_address,omitempty"`
	MinimalFee       coin.Coin                        `protobuf:"bytes,4,opt,name=minimal_fee,json=minimalFee,proto3" json:"minimal_fee"`
}

func (m *Configuration) Reset()         { *m = Configuration{} }
func (m *Configuration) String() string { return proto.CompactTextString(m) }
func (*Configuration) ProtoMessage()    {}
func (*Configuration) Descriptor() ([]byte, []int) {
	return fileDescriptor_7149e4b58e322390, []int{3}
}
func (m *Configuration) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Configuration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Configuration.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Configuration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Configuration.Merge(m, src)
}
func (m *Configuration) XXX_Size() int {
	return m.Size()
}
func (m *Configuration) XXX_DiscardUnknown() {
	xxx_messageInfo_Configuration.DiscardUnknown(m)
}

var xxx_messageInfo_Configuration proto.InternalMessageInfo

func (m *Configuration) GetOwner() github_com_iov_one_weave.Address {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *Configuration) GetCollectorAddress() github_com_iov_one_weave.Address {
	if m != nil {
		return m.CollectorAddress
	}
	return nil
}

func (m *Configuration) GetMinimalFee() coin.Coin {
	if m != nil {
		return m.MinimalFee
	}
	return coin.Coin{}
}

type ConfigurationMsg struct {
	// TODO: add schema uint32 here
	Patch *Configuration `protobuf:"bytes,2,opt,name=patch,proto3" json:"patch,omitempty"`
}

func (m *ConfigurationMsg) Reset()         { *m = ConfigurationMsg{} }
func (m *ConfigurationMsg) String() string { return proto.CompactTextString(m) }
func (*ConfigurationMsg) ProtoMessage()    {}
func (*ConfigurationMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_7149e4b58e322390, []int{4}
}
func (m *ConfigurationMsg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConfigurationMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConfigurationMsg.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConfigurationMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigurationMsg.Merge(m, src)
}
func (m *ConfigurationMsg) XXX_Size() int {
	return m.Size()
}
func (m *ConfigurationMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigurationMsg.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigurationMsg proto.InternalMessageInfo

func (m *ConfigurationMsg) GetPatch() *Configuration {
	if m != nil {
		return m.Patch
	}
	return nil
}

func init() {
	proto.RegisterType((*Set)(nil), "cash.Set")
	proto.RegisterType((*SendMsg)(nil), "cash.SendMsg")
	proto.RegisterType((*FeeInfo)(nil), "cash.FeeInfo")
	proto.RegisterType((*Configuration)(nil), "cash.Configuration")
	proto.RegisterType((*ConfigurationMsg)(nil), "cash.ConfigurationMsg")
}

func init() { proto.RegisterFile("x/cash/codec.proto", fileDescriptor_7149e4b58e322390) }

var fileDescriptor_7149e4b58e322390 = []byte{
	// 446 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xcd, 0xd6, 0x4e, 0x0a, 0x13, 0x10, 0x61, 0xe1, 0x60, 0xe5, 0xe0, 0x5a, 0x56, 0x0f, 0xa9,
	0x50, 0xd7, 0x22, 0x48, 0x08, 0x21, 0x71, 0x20, 0x95, 0x2a, 0x71, 0xe8, 0x01, 0x97, 0x7b, 0xb5,
	0xb1, 0xc7, 0x8e, 0xa5, 0x78, 0xa7, 0xb2, 0x37, 0x2d, 0xfc, 0x0b, 0x7e, 0x56, 0x8f, 0xe5, 0xc6,
	0xa9, 0x42, 0xc9, 0x3f, 0xe0, 0xc8, 0x09, 0xed, 0x3a, 0x44, 0x89, 0x10, 0xa8, 0xbe, 0xcd, 0xc7,
	0x7b, 0xcf, 0xfb, 0x9e, 0xc6, 0xc0, 0x3f, 0x47, 0x89, 0xac, 0x67, 0x51, 0x42, 0x29, 0x26, 0xe2,
	0xb2, 0x22, 0x4d, 0xdc, 0x35, 0x93, 0xe1, 0x71, 0x5e, 0xe8, 0xd9, 0x62, 0x2a, 0x12, 0x2a, 0xa3,
	0x9c, 0x72, 0x8a, 0xec, 0x72, 0xba, 0xc8, 0x6c, 0x67, 0x1b, 0x5b, 0x35, 0xa4, 0xe1, 0xe1, 0x16,
	0xbc, 0xa0, 0xab, 0x63, 0x52, 0x18, 0x5d, 0xa3, 0xbc, 0xc2, 0x6d, 0xe9, 0xe1, 0xd1, 0x7f, 0x50,
	0x85, 0xda, 0x86, 0x86, 0x9f, 0xc0, 0x39, 0x47, 0xcd, 0x5f, 0xc0, 0x83, 0x12, 0xb5, 0x4c, 0xa5,
	0x96, 0x1e, 0x0b, 0xd8, 0xa8, 0x3f, 0x7e, 0x22, 0x2c, 0x43, 0x9c, 0xad, 0xc7, 0xf1, 0x06, 0xc0,
	0x03, 0xe8, 0x1a, 0x9d, 0xda, 0xdb, 0x0b, 0x9c, 0x51, 0x7f, 0x0c, 0xc2, 0x74, 0xe2, 0x84, 0x0a,
	0x15, 0x37, 0x8b, 0xf0, 0x27, 0x83, 0xfd, 0x73, 0x54, 0xe9, 0x59, 0x9d, 0xb7, 0x93, 0x7e, 0x0d,
	0x4e, 0x5d, 0x25, 0xde, 0x5e, 0xc0, 0x46, 0x8f, 0x26, 0x87, 0xbf, 0xee, 0x0e, 0x82, 0x7f, 0x59,
	0x11, 0xef, 0xd3, 0xb4, 0xc2, 0xba, 0x8e, 0x0d, 0x81, 0xbf, 0x01, 0x37, 0xc5, 0x5a, 0x7b, 0x4e,
	0x0b, 0xa2, 0x65, 0xf0, 0x10, 0x7a, 0xb2, 0xa4, 0x85, 0xd2, 0x9e, 0x6b, 0x1f, 0xb7, 0xed, 0x66,
	0xbd, 0xe1, 0x1c, 0xdc, 0x12, 0x4b, 0xf2, 0xba, 0x01, 0x1b, 0x3d, 0x8c, 0x6d, 0xcd, 0x07, 0xe0,
	0x54, 0x98, 0x79, 0x3d, 0xf3, 0xc1, 0xd8, 0x94, 0xe1, 0x1c, 0xf6, 0x4f, 0x11, 0x3f, 0xa8, 0x8c,
	0xda, 0x79, 0x7e, 0x0e, 0xdd, 0x4b, 0xf9, 0x05, 0xab, 0xc6, 0x75, 0xdc, 0x34, 0xdc, 0x07, 0x37,
	0x43, 0xac, 0xad, 0xa3, 0xdd, 0x57, 0xd9, 0x79, 0xf8, 0x8d, 0xc1, 0xe3, 0x13, 0x52, 0x59, 0x91,
	0x2f, 0x2a, 0xa9, 0x0b, 0x52, 0xfc, 0x2d, 0x74, 0xe9, 0x5a, 0xfd, 0xd1, 0xb9, 0x67, 0x08, 0x0d,
	0x85, 0x7f, 0x84, 0xa7, 0x09, 0xcd, 0xe7, 0x98, 0x68, 0xaa, 0x2e, 0x64, 0xb3, 0x6b, 0x15, 0xe6,
	0x60, 0x43, 0x5f, 0x4f, 0xf8, 0x4b, 0xe8, 0x97, 0x85, 0x2a, 0x4a, 0x39, 0xbf, 0xc8, 0x10, 0xff,
	0x4e, 0x77, 0xe2, 0xde, 0xdc, 0x1d, 0x74, 0x62, 0x58, 0x83, 0x4e, 0x11, 0xc3, 0x77, 0x30, 0xd8,
	0xb1, 0x64, 0xce, 0xe7, 0xc8, 0xa4, 0xa3, 0x93, 0x99, 0x75, 0xd5, 0x1f, 0x3f, 0x13, 0xe6, 0xb7,
	0x11, 0x3b, 0xb0, 0xb8, 0x41, 0x4c, 0xbc, 0x9b, 0xa5, 0xcf, 0x6e, 0x97, 0x3e, 0xfb, 0xb1, 0xf4,
	0xd9, 0xd7, 0x95, 0xdf, 0xb9, 0x5d, 0xf9, 0x9d, 0xef, 0x2b, 0xbf, 0x33, 0xed, 0xd9, 0x63, 0x7f,
	0xf5, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x0f, 0xd5, 0x15, 0x8c, 0x88, 0x03, 0x00, 0x00,
}

func (m *Set) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Set) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Metadata != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.Metadata.Size()))
		n1, err := m.Metadata.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.Coins) > 0 {
		for _, msg := range m.Coins {
			dAtA[i] = 0x12
			i++
			i = encodeVarintCodec(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *SendMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SendMsg) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Metadata != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.Metadata.Size()))
		n2, err := m.Metadata.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if len(m.Src) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Src)))
		i += copy(dAtA[i:], m.Src)
	}
	if len(m.Dest) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Dest)))
		i += copy(dAtA[i:], m.Dest)
	}
	if m.Amount != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.Amount.Size()))
		n3, err := m.Amount.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if len(m.Memo) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Memo)))
		i += copy(dAtA[i:], m.Memo)
	}
	if len(m.Ref) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Ref)))
		i += copy(dAtA[i:], m.Ref)
	}
	return i, nil
}

func (m *FeeInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FeeInfo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Metadata != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.Metadata.Size()))
		n4, err := m.Metadata.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if len(m.Payer) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Payer)))
		i += copy(dAtA[i:], m.Payer)
	}
	if m.Fees != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.Fees.Size()))
		n5, err := m.Fees.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	return i, nil
}

func (m *Configuration) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Configuration) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Owner) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Owner)))
		i += copy(dAtA[i:], m.Owner)
	}
	if len(m.CollectorAddress) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintCodec(dAtA, i, uint64(len(m.CollectorAddress)))
		i += copy(dAtA[i:], m.CollectorAddress)
	}
	dAtA[i] = 0x22
	i++
	i = encodeVarintCodec(dAtA, i, uint64(m.MinimalFee.Size()))
	n6, err := m.MinimalFee.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n6
	return i, nil
}

func (m *ConfigurationMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConfigurationMsg) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Patch != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCodec(dAtA, i, uint64(m.Patch.Size()))
		n7, err := m.Patch.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n7
	}
	return i, nil
}

func encodeVarintCodec(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Set) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovCodec(uint64(l))
	}
	if len(m.Coins) > 0 {
		for _, e := range m.Coins {
			l = e.Size()
			n += 1 + l + sovCodec(uint64(l))
		}
	}
	return n
}

func (m *SendMsg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovCodec(uint64(l))
	}
	l = len(m.Src)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	l = len(m.Dest)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	if m.Amount != nil {
		l = m.Amount.Size()
		n += 1 + l + sovCodec(uint64(l))
	}
	l = len(m.Memo)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	l = len(m.Ref)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	return n
}

func (m *FeeInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovCodec(uint64(l))
	}
	l = len(m.Payer)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	if m.Fees != nil {
		l = m.Fees.Size()
		n += 1 + l + sovCodec(uint64(l))
	}
	return n
}

func (m *Configuration) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	l = len(m.CollectorAddress)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	l = m.MinimalFee.Size()
	n += 1 + l + sovCodec(uint64(l))
	return n
}

func (m *ConfigurationMsg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Patch != nil {
		l = m.Patch.Size()
		n += 1 + l + sovCodec(uint64(l))
	}
	return n
}

func sovCodec(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCodec(x uint64) (n int) {
	return sovCodec(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Set) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCodec
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
			return fmt.Errorf("proto: Set: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Set: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &weave.Metadata{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coins = append(m.Coins, &coin.Coin{})
			if err := m.Coins[len(m.Coins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCodec
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCodec
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
func (m *SendMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCodec
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
			return fmt.Errorf("proto: SendMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SendMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &weave.Metadata{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Src", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Src = append(m.Src[:0], dAtA[iNdEx:postIndex]...)
			if m.Src == nil {
				m.Src = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dest", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Dest = append(m.Dest[:0], dAtA[iNdEx:postIndex]...)
			if m.Dest == nil {
				m.Dest = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Amount == nil {
				m.Amount = &coin.Coin{}
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Memo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Memo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ref", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ref = append(m.Ref[:0], dAtA[iNdEx:postIndex]...)
			if m.Ref == nil {
				m.Ref = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCodec
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCodec
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
func (m *FeeInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCodec
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
			return fmt.Errorf("proto: FeeInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FeeInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &weave.Metadata{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payer", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payer = append(m.Payer[:0], dAtA[iNdEx:postIndex]...)
			if m.Payer == nil {
				m.Payer = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fees", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Fees == nil {
				m.Fees = &coin.Coin{}
			}
			if err := m.Fees.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCodec
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCodec
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
func (m *Configuration) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCodec
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
			return fmt.Errorf("proto: Configuration: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Configuration: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = append(m.Owner[:0], dAtA[iNdEx:postIndex]...)
			if m.Owner == nil {
				m.Owner = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollectorAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CollectorAddress = append(m.CollectorAddress[:0], dAtA[iNdEx:postIndex]...)
			if m.CollectorAddress == nil {
				m.CollectorAddress = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinimalFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinimalFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCodec
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCodec
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
func (m *ConfigurationMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCodec
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
			return fmt.Errorf("proto: ConfigurationMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConfigurationMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Patch", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Patch == nil {
				m.Patch = &Configuration{}
			}
			if err := m.Patch.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCodec
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCodec
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
func skipCodec(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCodec
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
					return 0, ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCodec
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
				return 0, ErrInvalidLengthCodec
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthCodec
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCodec
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipCodec(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthCodec
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthCodec = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCodec   = fmt.Errorf("proto: integer overflow")
)
