// Code generated by protoc-gen-gogo.
// source: Scaler.proto
// DO NOT EDIT!

package CoreML

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// *
// A scaling operation.
//
// This function has the following formula:
//
// .. math::
//     f(x) = scaleValue \cdot (x + shiftValue)
//
// If the ``scaleValue`` is not given, the default value 1 is used.
// If the ``shiftValue`` is not given, the default value 0 is used.
//
// If ``scaleValue`` and ``shiftValue`` are each a single value
// and the input is an array, then the scale and shift are applied
// to each element of the array.
//
// If the input is an integer, then it is converted to a double to
// perform the scaling operation. If the output type is an integer,
// then it is cast to an integer. If that cast is lossy, then an
// error is generated.
type Scaler struct {
	ShiftValue []float64 `protobuf:"fixed64,1,rep,packed,name=shiftValue" json:"shiftValue,omitempty"`
	ScaleValue []float64 `protobuf:"fixed64,2,rep,packed,name=scaleValue" json:"scaleValue,omitempty"`
}

func (m *Scaler) Reset()                    { *m = Scaler{} }
func (m *Scaler) String() string            { return proto.CompactTextString(m) }
func (*Scaler) ProtoMessage()               {}
func (*Scaler) Descriptor() ([]byte, []int) { return fileDescriptorScaler, []int{0} }

func (m *Scaler) GetShiftValue() []float64 {
	if m != nil {
		return m.ShiftValue
	}
	return nil
}

func (m *Scaler) GetScaleValue() []float64 {
	if m != nil {
		return m.ScaleValue
	}
	return nil
}

func init() {
	proto.RegisterType((*Scaler)(nil), "CoreML.Scaler")
}
func (m *Scaler) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Scaler) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ShiftValue) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintScaler(dAtA, i, uint64(len(m.ShiftValue)*8))
		for _, num := range m.ShiftValue {
			f1 := math.Float64bits(float64(num))
			dAtA[i] = uint8(f1)
			i++
			dAtA[i] = uint8(f1 >> 8)
			i++
			dAtA[i] = uint8(f1 >> 16)
			i++
			dAtA[i] = uint8(f1 >> 24)
			i++
			dAtA[i] = uint8(f1 >> 32)
			i++
			dAtA[i] = uint8(f1 >> 40)
			i++
			dAtA[i] = uint8(f1 >> 48)
			i++
			dAtA[i] = uint8(f1 >> 56)
			i++
		}
	}
	if len(m.ScaleValue) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintScaler(dAtA, i, uint64(len(m.ScaleValue)*8))
		for _, num := range m.ScaleValue {
			f2 := math.Float64bits(float64(num))
			dAtA[i] = uint8(f2)
			i++
			dAtA[i] = uint8(f2 >> 8)
			i++
			dAtA[i] = uint8(f2 >> 16)
			i++
			dAtA[i] = uint8(f2 >> 24)
			i++
			dAtA[i] = uint8(f2 >> 32)
			i++
			dAtA[i] = uint8(f2 >> 40)
			i++
			dAtA[i] = uint8(f2 >> 48)
			i++
			dAtA[i] = uint8(f2 >> 56)
			i++
		}
	}
	return i, nil
}

func encodeFixed64Scaler(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Scaler(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintScaler(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Scaler) Size() (n int) {
	var l int
	_ = l
	if len(m.ShiftValue) > 0 {
		n += 1 + sovScaler(uint64(len(m.ShiftValue)*8)) + len(m.ShiftValue)*8
	}
	if len(m.ScaleValue) > 0 {
		n += 1 + sovScaler(uint64(len(m.ScaleValue)*8)) + len(m.ScaleValue)*8
	}
	return n
}

func sovScaler(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozScaler(x uint64) (n int) {
	return sovScaler(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Scaler) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowScaler
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
			return fmt.Errorf("proto: Scaler: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Scaler: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 1 {
				var v uint64
				if (iNdEx + 8) > l {
					return io.ErrUnexpectedEOF
				}
				iNdEx += 8
				v = uint64(dAtA[iNdEx-8])
				v |= uint64(dAtA[iNdEx-7]) << 8
				v |= uint64(dAtA[iNdEx-6]) << 16
				v |= uint64(dAtA[iNdEx-5]) << 24
				v |= uint64(dAtA[iNdEx-4]) << 32
				v |= uint64(dAtA[iNdEx-3]) << 40
				v |= uint64(dAtA[iNdEx-2]) << 48
				v |= uint64(dAtA[iNdEx-1]) << 56
				v2 := float64(math.Float64frombits(v))
				m.ShiftValue = append(m.ShiftValue, v2)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowScaler
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthScaler
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v uint64
					if (iNdEx + 8) > l {
						return io.ErrUnexpectedEOF
					}
					iNdEx += 8
					v = uint64(dAtA[iNdEx-8])
					v |= uint64(dAtA[iNdEx-7]) << 8
					v |= uint64(dAtA[iNdEx-6]) << 16
					v |= uint64(dAtA[iNdEx-5]) << 24
					v |= uint64(dAtA[iNdEx-4]) << 32
					v |= uint64(dAtA[iNdEx-3]) << 40
					v |= uint64(dAtA[iNdEx-2]) << 48
					v |= uint64(dAtA[iNdEx-1]) << 56
					v2 := float64(math.Float64frombits(v))
					m.ShiftValue = append(m.ShiftValue, v2)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field ShiftValue", wireType)
			}
		case 2:
			if wireType == 1 {
				var v uint64
				if (iNdEx + 8) > l {
					return io.ErrUnexpectedEOF
				}
				iNdEx += 8
				v = uint64(dAtA[iNdEx-8])
				v |= uint64(dAtA[iNdEx-7]) << 8
				v |= uint64(dAtA[iNdEx-6]) << 16
				v |= uint64(dAtA[iNdEx-5]) << 24
				v |= uint64(dAtA[iNdEx-4]) << 32
				v |= uint64(dAtA[iNdEx-3]) << 40
				v |= uint64(dAtA[iNdEx-2]) << 48
				v |= uint64(dAtA[iNdEx-1]) << 56
				v2 := float64(math.Float64frombits(v))
				m.ScaleValue = append(m.ScaleValue, v2)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowScaler
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthScaler
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v uint64
					if (iNdEx + 8) > l {
						return io.ErrUnexpectedEOF
					}
					iNdEx += 8
					v = uint64(dAtA[iNdEx-8])
					v |= uint64(dAtA[iNdEx-7]) << 8
					v |= uint64(dAtA[iNdEx-6]) << 16
					v |= uint64(dAtA[iNdEx-5]) << 24
					v |= uint64(dAtA[iNdEx-4]) << 32
					v |= uint64(dAtA[iNdEx-3]) << 40
					v |= uint64(dAtA[iNdEx-2]) << 48
					v |= uint64(dAtA[iNdEx-1]) << 56
					v2 := float64(math.Float64frombits(v))
					m.ScaleValue = append(m.ScaleValue, v2)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field ScaleValue", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipScaler(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthScaler
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
func skipScaler(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowScaler
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
					return 0, ErrIntOverflowScaler
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
					return 0, ErrIntOverflowScaler
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
				return 0, ErrInvalidLengthScaler
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowScaler
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
				next, err := skipScaler(dAtA[start:])
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
	ErrInvalidLengthScaler = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowScaler   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("Scaler.proto", fileDescriptorScaler) }

var fileDescriptorScaler = []byte{
	// 112 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x09, 0x4e, 0x4e, 0xcc,
	0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x73, 0xce, 0x2f, 0x4a, 0xf5, 0xf5,
	0x51, 0xf2, 0xe0, 0x62, 0x83, 0x88, 0x0b, 0xc9, 0x71, 0x71, 0x15, 0x67, 0x64, 0xa6, 0x95, 0x84,
	0x25, 0xe6, 0x94, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6b, 0x30, 0x06, 0x21, 0x89, 0x80, 0xe5, 0x41,
	0x2a, 0x21, 0xf2, 0x4c, 0x50, 0x79, 0xb8, 0x88, 0x93, 0xd0, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e,
	0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0x83, 0x07, 0x73, 0x12, 0x1b, 0xd8,
	0x32, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa5, 0xfc, 0x7c, 0xa4, 0x7c, 0x00, 0x00, 0x00,
}