// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/batch/codec.proto

/*
Package batch is a generated protocol buffer package.

It is generated from these files:
	x/batch/codec.proto

It has these top-level messages:
	ByteArrayList
*/
package batch

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ByteArrayList struct {
	Elements [][]byte `protobuf:"bytes,1,rep,name=elements" json:"elements,omitempty"`
}

func (m *ByteArrayList) Reset()                    { *m = ByteArrayList{} }
func (m *ByteArrayList) String() string            { return proto.CompactTextString(m) }
func (*ByteArrayList) ProtoMessage()               {}
func (*ByteArrayList) Descriptor() ([]byte, []int) { return fileDescriptorCodec, []int{0} }

func (m *ByteArrayList) GetElements() [][]byte {
	if m != nil {
		return m.Elements
	}
	return nil
}

func init() {
	proto.RegisterType((*ByteArrayList)(nil), "batch.ByteArrayList")
}
func (m *ByteArrayList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ByteArrayList) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Elements) > 0 {
		for _, b := range m.Elements {
			dAtA[i] = 0xa
			i++
			i = encodeVarintCodec(dAtA, i, uint64(len(b)))
			i += copy(dAtA[i:], b)
		}
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
func (m *ByteArrayList) Size() (n int) {
	var l int
	_ = l
	if len(m.Elements) > 0 {
		for _, b := range m.Elements {
			l = len(b)
			n += 1 + l + sovCodec(uint64(l))
		}
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
func (m *ByteArrayList) Unmarshal(dAtA []byte) error {
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
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ByteArrayList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ByteArrayList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Elements", wireType)
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
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Elements = append(m.Elements, make([]byte, postIndex-iNdEx))
			copy(m.Elements[len(m.Elements)-1], dAtA[iNdEx:postIndex])
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
			iNdEx += length
			if length < 0 {
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

func init() { proto.RegisterFile("x/batch/codec.proto", fileDescriptorCodec) }

var fileDescriptorCodec = []byte{
	// 111 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xae, 0xd0, 0x4f, 0x4a,
	0x2c, 0x49, 0xce, 0xd0, 0x4f, 0xce, 0x4f, 0x49, 0x4d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x05, 0x0b, 0x29, 0x69, 0x73, 0xf1, 0x3a, 0x55, 0x96, 0xa4, 0x3a, 0x16, 0x15, 0x25, 0x56,
	0xfa, 0x64, 0x16, 0x97, 0x08, 0x49, 0x71, 0x71, 0xa4, 0xe6, 0xa4, 0xe6, 0xa6, 0xe6, 0x95, 0x14,
	0x4b, 0x30, 0x2a, 0x30, 0x6b, 0xf0, 0x04, 0xc1, 0xf9, 0x4e, 0x02, 0x27, 0x1e, 0xc9, 0x31, 0x5e,
	0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x49, 0x6c, 0x60, 0xc3,
	0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf6, 0x1e, 0x39, 0x05, 0x63, 0x00, 0x00, 0x00,
}