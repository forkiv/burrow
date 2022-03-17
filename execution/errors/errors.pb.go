// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: errors.proto

package errors

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Exception struct {
	CodeNumber           uint32   `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Exception            string   `protobuf:"bytes,2,opt,name=Exception,proto3" json:"Exception,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Exception) Reset()      { *m = Exception{} }
func (*Exception) ProtoMessage() {}
func (*Exception) Descriptor() ([]byte, []int) {
	return fileDescriptor_24fe73c7f0ddb19c, []int{0}
}
func (m *Exception) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Exception) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *Exception) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Exception.Merge(m, src)
}
func (m *Exception) XXX_Size() int {
	return m.Size()
}
func (m *Exception) XXX_DiscardUnknown() {
	xxx_messageInfo_Exception.DiscardUnknown(m)
}

var xxx_messageInfo_Exception proto.InternalMessageInfo

func (m *Exception) GetCodeNumber() uint32 {
	if m != nil {
		return m.CodeNumber
	}
	return 0
}

func (m *Exception) GetException() string {
	if m != nil {
		return m.Exception
	}
	return ""
}

func (*Exception) XXX_MessageName() string {
	return "errors.Exception"
}
func init() {
	proto.RegisterType((*Exception)(nil), "errors.Exception")
	golang_proto.RegisterType((*Exception)(nil), "errors.Exception")
}

func init() { proto.RegisterFile("errors.proto", fileDescriptor_24fe73c7f0ddb19c) }
func init() { golang_proto.RegisterFile("errors.proto", fileDescriptor_24fe73c7f0ddb19c) }

var fileDescriptor_24fe73c7f0ddb19c = []byte{
	// 193 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0x2d, 0x2a, 0xca,
	0x2f, 0x2a, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0xa4, 0x44, 0xd2, 0xf3,
	0xd3, 0xf3, 0xc1, 0x42, 0xfa, 0x20, 0x16, 0x44, 0x56, 0x29, 0x98, 0x8b, 0xd3, 0xb5, 0x22, 0x39,
	0xb5, 0xa0, 0x24, 0x33, 0x3f, 0x4f, 0x48, 0x89, 0x8b, 0xc5, 0x39, 0x3f, 0x25, 0x55, 0x82, 0x51,
	0x81, 0x51, 0x83, 0xd7, 0x89, 0xef, 0xd1, 0x3d, 0x79, 0x2e, 0x10, 0xdf, 0xaf, 0x34, 0x37, 0x29,
	0xb5, 0x28, 0x08, 0x2c, 0x27, 0x24, 0x83, 0xa4, 0x41, 0x82, 0x49, 0x81, 0x51, 0x83, 0x33, 0x08,
	0x21, 0x60, 0xc5, 0x32, 0x63, 0x81, 0x3c, 0x83, 0x93, 0xc7, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e,
	0xc9, 0x31, 0xde, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x81, 0xc7, 0x72, 0x8c, 0x27,
	0x1e, 0xcb, 0x31, 0x46, 0xe9, 0xa5, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea,
	0x67, 0x54, 0x16, 0xa4, 0x16, 0xe5, 0xa4, 0xa6, 0xa4, 0xa7, 0x16, 0xe9, 0x27, 0x95, 0x16, 0x15,
	0xe5, 0x97, 0xeb, 0xa7, 0x56, 0xa4, 0x26, 0x97, 0x82, 0x0c, 0xd2, 0x87, 0x38, 0x3a, 0x89, 0x0d,
	0xec, 0x4a, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe0, 0x6d, 0xcb, 0x28, 0xd3, 0x00, 0x00,
	0x00,
}

func (m *Exception) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Exception) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Exception) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Exception) > 0 {
		i -= len(m.Exception)
		copy(dAtA[i:], m.Exception)
		i = encodeVarintErrors(dAtA, i, uint64(len(m.Exception)))
		i--
		dAtA[i] = 0x12
	}
	if m.CodeNumber != 0 {
		i = encodeVarintErrors(dAtA, i, uint64(m.CodeNumber))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintErrors(dAtA []byte, offset int, v uint64) int {
	offset -= sovErrors(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Exception) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CodeNumber != 0 {
		n += 1 + sovErrors(uint64(m.CodeNumber))
	}
	l = len(m.Exception)
	if l > 0 {
		n += 1 + l + sovErrors(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovErrors(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozErrors(x uint64) (n int) {
	return sovErrors(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Exception) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowErrors
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
			return fmt.Errorf("proto: Exception: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Exception: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CodeNumber", wireType)
			}
			m.CodeNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrors
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CodeNumber |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Exception", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowErrors
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
				return ErrInvalidLengthErrors
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthErrors
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Exception = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipErrors(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthErrors
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipErrors(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowErrors
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
					return 0, ErrIntOverflowErrors
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
					return 0, ErrIntOverflowErrors
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
				return 0, ErrInvalidLengthErrors
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupErrors
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthErrors
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthErrors        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowErrors          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupErrors = fmt.Errorf("proto: unexpected end of group")
)
