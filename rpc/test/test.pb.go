// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: test.proto

/*
	Package test is a generated protocol buffer package.

	It is generated from these files:
		test.proto

	It has these top-level messages:
		Empty
		SimpleTypes
		ComplexTypes
*/
package test

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

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

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptorTest, []int{0} }

type SimpleTypes struct {
	Number int32  `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	Str    string `protobuf:"bytes,2,opt,name=str,proto3" json:"str,omitempty"`
	Logic  bool   `protobuf:"varint,3,opt,name=logic,proto3" json:"logic,omitempty"`
}

func (m *SimpleTypes) Reset()                    { *m = SimpleTypes{} }
func (m *SimpleTypes) String() string            { return proto.CompactTextString(m) }
func (*SimpleTypes) ProtoMessage()               {}
func (*SimpleTypes) Descriptor() ([]byte, []int) { return fileDescriptorTest, []int{1} }

func (m *SimpleTypes) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *SimpleTypes) GetStr() string {
	if m != nil {
		return m.Str
	}
	return ""
}

func (m *SimpleTypes) GetLogic() bool {
	if m != nil {
		return m.Logic
	}
	return false
}

type ComplexTypes struct {
	Array []int32          `protobuf:"varint,2,rep,packed,name=array" json:"array,omitempty"`
	Dict  map[int32]string `protobuf:"bytes,3,rep,name=dict" json:"dict,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *ComplexTypes) Reset()                    { *m = ComplexTypes{} }
func (m *ComplexTypes) String() string            { return proto.CompactTextString(m) }
func (*ComplexTypes) ProtoMessage()               {}
func (*ComplexTypes) Descriptor() ([]byte, []int) { return fileDescriptorTest, []int{2} }

func (m *ComplexTypes) GetArray() []int32 {
	if m != nil {
		return m.Array
	}
	return nil
}

func (m *ComplexTypes) GetDict() map[int32]string {
	if m != nil {
		return m.Dict
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "test.Empty")
	proto.RegisterType((*SimpleTypes)(nil), "test.SimpleTypes")
	proto.RegisterType((*ComplexTypes)(nil), "test.ComplexTypes")
}
func (this *Empty) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Empty)
	if !ok {
		that2, ok := that.(Empty)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	return true
}
func (this *SimpleTypes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SimpleTypes)
	if !ok {
		that2, ok := that.(SimpleTypes)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Number != that1.Number {
		return false
	}
	if this.Str != that1.Str {
		return false
	}
	if this.Logic != that1.Logic {
		return false
	}
	return true
}
func (this *ComplexTypes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ComplexTypes)
	if !ok {
		that2, ok := that.(ComplexTypes)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Array) != len(that1.Array) {
		return false
	}
	for i := range this.Array {
		if this.Array[i] != that1.Array[i] {
			return false
		}
	}
	if len(this.Dict) != len(that1.Dict) {
		return false
	}
	for i := range this.Dict {
		if this.Dict[i] != that1.Dict[i] {
			return false
		}
	}
	return true
}
func (m *Empty) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Empty) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *SimpleTypes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SimpleTypes) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Number != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintTest(dAtA, i, uint64(m.Number))
	}
	if len(m.Str) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintTest(dAtA, i, uint64(len(m.Str)))
		i += copy(dAtA[i:], m.Str)
	}
	if m.Logic {
		dAtA[i] = 0x18
		i++
		if m.Logic {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *ComplexTypes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ComplexTypes) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Array) > 0 {
		dAtA2 := make([]byte, len(m.Array)*10)
		var j1 int
		for _, num1 := range m.Array {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		dAtA[i] = 0x12
		i++
		i = encodeVarintTest(dAtA, i, uint64(j1))
		i += copy(dAtA[i:], dAtA2[:j1])
	}
	if len(m.Dict) > 0 {
		for k, _ := range m.Dict {
			dAtA[i] = 0x1a
			i++
			v := m.Dict[k]
			mapSize := 1 + sovTest(uint64(k)) + 1 + len(v) + sovTest(uint64(len(v)))
			i = encodeVarintTest(dAtA, i, uint64(mapSize))
			dAtA[i] = 0x8
			i++
			i = encodeVarintTest(dAtA, i, uint64(k))
			dAtA[i] = 0x12
			i++
			i = encodeVarintTest(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	return i, nil
}

func encodeVarintTest(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Empty) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *SimpleTypes) Size() (n int) {
	var l int
	_ = l
	if m.Number != 0 {
		n += 1 + sovTest(uint64(m.Number))
	}
	l = len(m.Str)
	if l > 0 {
		n += 1 + l + sovTest(uint64(l))
	}
	if m.Logic {
		n += 2
	}
	return n
}

func (m *ComplexTypes) Size() (n int) {
	var l int
	_ = l
	if len(m.Array) > 0 {
		l = 0
		for _, e := range m.Array {
			l += sovTest(uint64(e))
		}
		n += 1 + sovTest(uint64(l)) + l
	}
	if len(m.Dict) > 0 {
		for k, v := range m.Dict {
			_ = k
			_ = v
			mapEntrySize := 1 + sovTest(uint64(k)) + 1 + len(v) + sovTest(uint64(len(v)))
			n += mapEntrySize + 1 + sovTest(uint64(mapEntrySize))
		}
	}
	return n
}

func sovTest(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTest(x uint64) (n int) {
	return sovTest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Empty) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTest
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
			return fmt.Errorf("proto: Empty: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Empty: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTest
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
func (m *SimpleTypes) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTest
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
			return fmt.Errorf("proto: SimpleTypes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SimpleTypes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Number", wireType)
			}
			m.Number = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Number |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Str", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTest
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Str = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Logic", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Logic = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipTest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTest
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
func (m *ComplexTypes) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTest
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
			return fmt.Errorf("proto: ComplexTypes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ComplexTypes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType == 0 {
				var v int32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTest
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (int32(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Array = append(m.Array, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTest
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
					return ErrInvalidLengthTest
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v int32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTest
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (int32(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Array = append(m.Array, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Array", wireType)
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dict", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTest
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Dict == nil {
				m.Dict = make(map[int32]string)
			}
			var mapkey int32
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTest
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
				if fieldNum == 1 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTest
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapkey |= (int32(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTest
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthTest
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipTest(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthTest
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Dict[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTest
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
func skipTest(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTest
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
					return 0, ErrIntOverflowTest
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
					return 0, ErrIntOverflowTest
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
				return 0, ErrInvalidLengthTest
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTest
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
				next, err := skipTest(dAtA[start:])
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
	ErrInvalidLengthTest = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTest   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("test.proto", fileDescriptorTest) }

var fileDescriptorTest = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xc1, 0x4e, 0xbb, 0x40,
	0x10, 0xc6, 0xb3, 0xa5, 0xf4, 0xff, 0xef, 0xe0, 0xa1, 0x4e, 0x1a, 0x43, 0x88, 0x21, 0x84, 0x13,
	0x17, 0xa9, 0xa9, 0x31, 0x1a, 0x8f, 0xd6, 0xc6, 0x93, 0x17, 0xec, 0x0b, 0x14, 0x5c, 0x91, 0x58,
	0x58, 0xb2, 0x2c, 0xc6, 0x7d, 0x0a, 0x4f, 0xbe, 0x83, 0x8f, 0xe2, 0xd1, 0x47, 0x30, 0x3c, 0x89,
	0xd9, 0x5d, 0x62, 0x48, 0xea, 0xed, 0xfb, 0xcd, 0x30, 0x33, 0xdf, 0xc7, 0x02, 0x08, 0xda, 0x88,
	0xb8, 0xe6, 0x4c, 0x30, 0x1c, 0x2b, 0xed, 0x9d, 0xe4, 0x85, 0x78, 0x6a, 0xd3, 0x38, 0x63, 0xe5,
	0x22, 0x67, 0x39, 0x5b, 0xe8, 0x66, 0xda, 0x3e, 0x6a, 0xd2, 0xa0, 0x95, 0x19, 0x0a, 0xff, 0x81,
	0xbd, 0x2e, 0x6b, 0x21, 0xc3, 0x3b, 0x70, 0xee, 0x8b, 0xb2, 0xde, 0xd1, 0x8d, 0xac, 0x69, 0x83,
	0x47, 0x30, 0xa9, 0xda, 0x32, 0xa5, 0xdc, 0x25, 0x01, 0x89, 0xec, 0xa4, 0x27, 0x9c, 0x81, 0xd5,
	0x08, 0xee, 0x8e, 0x02, 0x12, 0x4d, 0x13, 0x25, 0x71, 0x0e, 0xf6, 0x8e, 0xe5, 0x45, 0xe6, 0x5a,
	0x01, 0x89, 0xfe, 0x27, 0x06, 0xc2, 0x37, 0x02, 0x07, 0x2b, 0xa6, 0xf6, 0xbd, 0x9a, 0x85, 0x73,
	0xb0, 0xb7, 0x9c, 0x6f, 0xa5, 0x3b, 0x0a, 0xac, 0xc8, 0x4e, 0x0c, 0xe0, 0x29, 0x8c, 0x1f, 0x8a,
	0x4c, 0xb8, 0x56, 0x60, 0x45, 0xce, 0xf2, 0x38, 0xd6, 0x71, 0x86, 0x73, 0xf1, 0x4d, 0x91, 0x89,
	0x75, 0x25, 0xb8, 0x4c, 0xf4, 0x97, 0xde, 0x05, 0x4c, 0x7f, 0x4b, 0xca, 0xcd, 0x33, 0x95, 0xbd,
	0x45, 0x25, 0xd5, 0x99, 0x97, 0xed, 0xae, 0xa5, 0xbd, 0x43, 0x03, 0x57, 0xa3, 0x4b, 0xb2, 0x7c,
	0x27, 0x30, 0xde, 0xd0, 0x46, 0xe0, 0x12, 0x60, 0xc5, 0x6a, 0x69, 0xd2, 0xe2, 0xa1, 0xb9, 0x39,
	0xc8, 0xee, 0xed, 0x97, 0x30, 0x84, 0xc9, 0x2d, 0xad, 0xd6, 0x9c, 0xa3, 0x63, 0x9a, 0xfa, 0xa7,
	0x79, 0x43, 0xc0, 0x73, 0x70, 0xd4, 0xde, 0xde, 0x3d, 0xe2, 0x7e, 0x18, 0xef, 0x8f, 0xda, 0xf5,
	0xec, 0xa3, 0xf3, 0xc9, 0x67, 0xe7, 0x93, 0xaf, 0xce, 0x27, 0xdf, 0x9d, 0x4f, 0xd2, 0x89, 0x7e,
	0x9a, 0xb3, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x05, 0x0c, 0x89, 0x89, 0xdd, 0x01, 0x00, 0x00,
}
