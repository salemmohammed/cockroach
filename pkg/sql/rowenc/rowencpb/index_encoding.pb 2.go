// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sql/rowenc/rowencpb/index_encoding.proto

package rowencpb

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

// Wrapper for the bytes of the value of an index that also contains bit for
// whether or not the value was deleted. A wrapper of arbitrary index values
// with the additional delete bit was chosen over a separate index encoding type
// because there would have to be a separate type for each encoding that we
// already have for indexes. The alternative would get harder to maintain if we
// added more index encodings in the future.
type IndexValueWrapper struct {
	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	// If deleted is true, value will always be nil.
	Deleted bool `protobuf:"varint,2,opt,name=deleted,proto3" json:"deleted,omitempty"`
}

func (m *IndexValueWrapper) Reset()         { *m = IndexValueWrapper{} }
func (m *IndexValueWrapper) String() string { return proto.CompactTextString(m) }
func (*IndexValueWrapper) ProtoMessage()    {}
func (*IndexValueWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_75ef4c67be01dfc7, []int{0}
}
func (m *IndexValueWrapper) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IndexValueWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *IndexValueWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexValueWrapper.Merge(m, src)
}
func (m *IndexValueWrapper) XXX_Size() int {
	return m.Size()
}
func (m *IndexValueWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexValueWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_IndexValueWrapper proto.InternalMessageInfo

func init() {
	proto.RegisterType((*IndexValueWrapper)(nil), "cockroach.rowenc.IndexValueWrapper")
}

func init() {
	proto.RegisterFile("sql/rowenc/rowencpb/index_encoding.proto", fileDescriptor_75ef4c67be01dfc7)
}

var fileDescriptor_75ef4c67be01dfc7 = []byte{
	// 174 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x28, 0x2e, 0xcc, 0xd1,
	0x2f, 0xca, 0x2f, 0x4f, 0xcd, 0x4b, 0x86, 0x52, 0x05, 0x49, 0xfa, 0x99, 0x79, 0x29, 0xa9, 0x15,
	0xf1, 0xa9, 0x79, 0xc9, 0xf9, 0x29, 0x99, 0x79, 0xe9, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42,
	0x02, 0xc9, 0xf9, 0xc9, 0xd9, 0x45, 0xf9, 0x89, 0xc9, 0x19, 0x7a, 0x10, 0x85, 0x4a, 0xce, 0x5c,
	0x82, 0x9e, 0x20, 0x95, 0x61, 0x89, 0x39, 0xa5, 0xa9, 0xe1, 0x45, 0x89, 0x05, 0x05, 0xa9, 0x45,
	0x42, 0x22, 0x5c, 0xac, 0x65, 0x20, 0xbe, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x4f, 0x10, 0x84, 0x23,
	0x24, 0xc1, 0xc5, 0x9e, 0x92, 0x9a, 0x93, 0x5a, 0x92, 0x9a, 0x22, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1,
	0x11, 0x04, 0xe3, 0x3a, 0x69, 0x9d, 0x78, 0x28, 0xc7, 0x70, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47,
	0x72, 0x8c, 0x37, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85,
	0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0x71, 0xc0, 0xdc, 0x95, 0xc4, 0x06, 0x76, 0x89,
	0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xd6, 0x9d, 0xa5, 0x04, 0xb5, 0x00, 0x00, 0x00,
}

func (m *IndexValueWrapper) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IndexValueWrapper) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IndexValueWrapper) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Deleted {
		i--
		if m.Deleted {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintIndexEncoding(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintIndexEncoding(dAtA []byte, offset int, v uint64) int {
	offset -= sovIndexEncoding(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *IndexValueWrapper) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovIndexEncoding(uint64(l))
	}
	if m.Deleted {
		n += 2
	}
	return n
}

func sovIndexEncoding(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIndexEncoding(x uint64) (n int) {
	return sovIndexEncoding(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *IndexValueWrapper) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIndexEncoding
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
			return fmt.Errorf("proto: IndexValueWrapper: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IndexValueWrapper: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIndexEncoding
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
				return ErrInvalidLengthIndexEncoding
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIndexEncoding
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value[:0], dAtA[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deleted", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIndexEncoding
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Deleted = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipIndexEncoding(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIndexEncoding
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
func skipIndexEncoding(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIndexEncoding
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
					return 0, ErrIntOverflowIndexEncoding
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
					return 0, ErrIntOverflowIndexEncoding
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
				return 0, ErrInvalidLengthIndexEncoding
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIndexEncoding
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIndexEncoding
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIndexEncoding        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIndexEncoding          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIndexEncoding = fmt.Errorf("proto: unexpected end of group")
)

