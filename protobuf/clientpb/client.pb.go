// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: protobuf/clientpb/client.proto

/*
	Package clientpb is a generated protocol buffer package.

	It is generated from these files:
		protobuf/clientpb/client.proto

	It has these top-level messages:
		ExampleFile
		ExampleIndex
*/
package clientpb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/timestamp"
import _ "github.com/gogo/protobuf/gogoproto"

import time "time"

import strings "strings"
import reflect "reflect"

import github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ExampleFile struct {
	Name      string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Created   time.Time `protobuf:"bytes,2,opt,name=created,stdtime" json:"created"`
	TotalSize int64     `protobuf:"varint,3,opt,name=total_size,json=totalSize,proto3" json:"total_size,omitempty"`
	Positions []int32   `protobuf:"varint,4,rep,packed,name=positions" json:"positions,omitempty"`
}

func (m *ExampleFile) Reset()                    { *m = ExampleFile{} }
func (*ExampleFile) ProtoMessage()               {}
func (*ExampleFile) Descriptor() ([]byte, []int) { return fileDescriptorClient, []int{0} }

func (m *ExampleFile) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ExampleFile) GetCreated() time.Time {
	if m != nil {
		return m.Created
	}
	return time.Time{}
}

func (m *ExampleFile) GetTotalSize() int64 {
	if m != nil {
		return m.TotalSize
	}
	return 0
}

func (m *ExampleFile) GetPositions() []int32 {
	if m != nil {
		return m.Positions
	}
	return nil
}

type ExampleIndex struct {
	TotalExamples int64         `protobuf:"varint,1,opt,name=total_examples,json=totalExamples,proto3" json:"total_examples,omitempty"`
	TotalSize     int64         `protobuf:"varint,2,opt,name=total_size,json=totalSize,proto3" json:"total_size,omitempty"`
	Files         []ExampleFile `protobuf:"bytes,3,rep,name=files" json:"files"`
}

func (m *ExampleIndex) Reset()                    { *m = ExampleIndex{} }
func (*ExampleIndex) ProtoMessage()               {}
func (*ExampleIndex) Descriptor() ([]byte, []int) { return fileDescriptorClient, []int{1} }

func (m *ExampleIndex) GetTotalExamples() int64 {
	if m != nil {
		return m.TotalExamples
	}
	return 0
}

func (m *ExampleIndex) GetTotalSize() int64 {
	if m != nil {
		return m.TotalSize
	}
	return 0
}

func (m *ExampleIndex) GetFiles() []ExampleFile {
	if m != nil {
		return m.Files
	}
	return nil
}

func init() {
	proto.RegisterType((*ExampleFile)(nil), "clientpb.ExampleFile")
	proto.RegisterType((*ExampleIndex)(nil), "clientpb.ExampleIndex")
}
func (this *ExampleFile) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*ExampleFile)
	if !ok {
		that2, ok := that.(ExampleFile)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if !this.Created.Equal(that1.Created) {
		return false
	}
	if this.TotalSize != that1.TotalSize {
		return false
	}
	if len(this.Positions) != len(that1.Positions) {
		return false
	}
	for i := range this.Positions {
		if this.Positions[i] != that1.Positions[i] {
			return false
		}
	}
	return true
}
func (this *ExampleIndex) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*ExampleIndex)
	if !ok {
		that2, ok := that.(ExampleIndex)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.TotalExamples != that1.TotalExamples {
		return false
	}
	if this.TotalSize != that1.TotalSize {
		return false
	}
	if len(this.Files) != len(that1.Files) {
		return false
	}
	for i := range this.Files {
		if !this.Files[i].Equal(&that1.Files[i]) {
			return false
		}
	}
	return true
}
func (this *ExampleFile) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&clientpb.ExampleFile{")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "Created: "+fmt.Sprintf("%#v", this.Created)+",\n")
	s = append(s, "TotalSize: "+fmt.Sprintf("%#v", this.TotalSize)+",\n")
	s = append(s, "Positions: "+fmt.Sprintf("%#v", this.Positions)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ExampleIndex) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&clientpb.ExampleIndex{")
	s = append(s, "TotalExamples: "+fmt.Sprintf("%#v", this.TotalExamples)+",\n")
	s = append(s, "TotalSize: "+fmt.Sprintf("%#v", this.TotalSize)+",\n")
	if this.Files != nil {
		vs := make([]*ExampleFile, len(this.Files))
		for i := range vs {
			vs[i] = &this.Files[i]
		}
		s = append(s, "Files: "+fmt.Sprintf("%#v", vs)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringClient(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *ExampleFile) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExampleFile) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintClient(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	dAtA[i] = 0x12
	i++
	i = encodeVarintClient(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdTime(m.Created)))
	n1, err := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Created, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if m.TotalSize != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintClient(dAtA, i, uint64(m.TotalSize))
	}
	if len(m.Positions) > 0 {
		dAtA3 := make([]byte, len(m.Positions)*10)
		var j2 int
		for _, num1 := range m.Positions {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA3[j2] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j2++
			}
			dAtA3[j2] = uint8(num)
			j2++
		}
		dAtA[i] = 0x22
		i++
		i = encodeVarintClient(dAtA, i, uint64(j2))
		i += copy(dAtA[i:], dAtA3[:j2])
	}
	return i, nil
}

func (m *ExampleIndex) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExampleIndex) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.TotalExamples != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintClient(dAtA, i, uint64(m.TotalExamples))
	}
	if m.TotalSize != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintClient(dAtA, i, uint64(m.TotalSize))
	}
	if len(m.Files) > 0 {
		for _, msg := range m.Files {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintClient(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeFixed64Client(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32Client(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintClient(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ExampleFile) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovClient(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Created)
	n += 1 + l + sovClient(uint64(l))
	if m.TotalSize != 0 {
		n += 1 + sovClient(uint64(m.TotalSize))
	}
	if len(m.Positions) > 0 {
		l = 0
		for _, e := range m.Positions {
			l += sovClient(uint64(e))
		}
		n += 1 + sovClient(uint64(l)) + l
	}
	return n
}

func (m *ExampleIndex) Size() (n int) {
	var l int
	_ = l
	if m.TotalExamples != 0 {
		n += 1 + sovClient(uint64(m.TotalExamples))
	}
	if m.TotalSize != 0 {
		n += 1 + sovClient(uint64(m.TotalSize))
	}
	if len(m.Files) > 0 {
		for _, e := range m.Files {
			l = e.Size()
			n += 1 + l + sovClient(uint64(l))
		}
	}
	return n
}

func sovClient(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozClient(x uint64) (n int) {
	return sovClient(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ExampleFile) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ExampleFile{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Created:` + strings.Replace(strings.Replace(this.Created.String(), "Timestamp", "google_protobuf.Timestamp", 1), `&`, ``, 1) + `,`,
		`TotalSize:` + fmt.Sprintf("%v", this.TotalSize) + `,`,
		`Positions:` + fmt.Sprintf("%v", this.Positions) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ExampleIndex) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ExampleIndex{`,
		`TotalExamples:` + fmt.Sprintf("%v", this.TotalExamples) + `,`,
		`TotalSize:` + fmt.Sprintf("%v", this.TotalSize) + `,`,
		`Files:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Files), "ExampleFile", "ExampleFile", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringClient(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ExampleFile) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClient
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
			return fmt.Errorf("proto: ExampleFile: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExampleFile: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClient
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
				return ErrInvalidLengthClient
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Created", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClient
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
				return ErrInvalidLengthClient
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Created, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalSize", wireType)
			}
			m.TotalSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClient
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalSize |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType == 0 {
				var v int32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClient
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
				m.Positions = append(m.Positions, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClient
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
					return ErrInvalidLengthClient
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v int32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowClient
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
					m.Positions = append(m.Positions, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Positions", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipClient(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthClient
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
func (m *ExampleIndex) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClient
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
			return fmt.Errorf("proto: ExampleIndex: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExampleIndex: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalExamples", wireType)
			}
			m.TotalExamples = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClient
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalExamples |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalSize", wireType)
			}
			m.TotalSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClient
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalSize |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Files", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClient
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
				return ErrInvalidLengthClient
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Files = append(m.Files, ExampleFile{})
			if err := m.Files[len(m.Files)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClient(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthClient
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
func skipClient(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowClient
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
					return 0, ErrIntOverflowClient
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
					return 0, ErrIntOverflowClient
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
				return 0, ErrInvalidLengthClient
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowClient
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
				next, err := skipClient(dAtA[start:])
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
	ErrInvalidLengthClient = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowClient   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("protobuf/clientpb/client.proto", fileDescriptorClient) }

var fileDescriptorClient = []byte{
	// 338 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x50, 0x31, 0x4b, 0xfb, 0x40,
	0x1c, 0xcd, 0xaf, 0x69, 0xff, 0xff, 0xf6, 0xaa, 0x0e, 0x07, 0x42, 0x28, 0x78, 0x0d, 0x05, 0x21,
	0x83, 0x5e, 0xb1, 0xee, 0x0e, 0x05, 0x05, 0xd7, 0xe8, 0x2e, 0x49, 0x7b, 0x8d, 0x07, 0x49, 0x2e,
	0x34, 0x57, 0x28, 0x9d, 0xdc, 0x5c, 0xfb, 0x21, 0x1c, 0xfc, 0x28, 0x1d, 0x3b, 0x3a, 0xa9, 0x3d,
	0x17, 0xc7, 0x7e, 0x04, 0xc9, 0x5d, 0x62, 0xa5, 0xdb, 0xbb, 0x77, 0xef, 0xbd, 0xdf, 0xe3, 0x21,
	0x92, 0x4d, 0x85, 0x14, 0xe1, 0x6c, 0xd2, 0x1f, 0xc5, 0x9c, 0xa5, 0x32, 0x0b, 0x4b, 0x40, 0xf5,
	0x07, 0x6e, 0x56, 0x74, 0xa7, 0x1b, 0x09, 0x11, 0xc5, 0xac, 0xff, 0x6b, 0x90, 0x3c, 0x61, 0xb9,
	0x0c, 0x92, 0xcc, 0x48, 0x3b, 0xe7, 0x11, 0x97, 0x8f, 0xb3, 0x90, 0x8e, 0x44, 0xd2, 0x8f, 0x44,
	0x24, 0x76, 0xca, 0xe2, 0xa5, 0x1f, 0x1a, 0x19, 0x79, 0xef, 0x05, 0x50, 0xfb, 0x7a, 0x1e, 0x24,
	0x59, 0xcc, 0x6e, 0x78, 0xcc, 0x30, 0x46, 0xf5, 0x34, 0x48, 0x98, 0x03, 0x2e, 0x78, 0x2d, 0x5f,
	0x63, 0x7c, 0x85, 0xfe, 0x8f, 0xa6, 0x2c, 0x90, 0x6c, 0xec, 0xd4, 0x5c, 0xf0, 0xda, 0x83, 0x0e,
	0x35, 0x2d, 0x68, 0x95, 0x4d, 0xef, 0xab, 0x16, 0xc3, 0xe6, 0xea, 0xbd, 0x6b, 0x2d, 0x3f, 0xba,
	0xe0, 0x57, 0x26, 0x7c, 0x82, 0x90, 0x14, 0x32, 0x88, 0x1f, 0x72, 0xbe, 0x60, 0x8e, 0xed, 0x82,
	0x67, 0xfb, 0x2d, 0xcd, 0xdc, 0xf1, 0x05, 0xc3, 0x3d, 0xd4, 0xca, 0x44, 0xce, 0x25, 0x17, 0x69,
	0xee, 0xd4, 0x5d, 0xdb, 0x6b, 0x0c, 0xeb, 0x45, 0x88, 0xbf, 0xa3, 0x7b, 0xcf, 0x80, 0x0e, 0xca,
	0x9a, 0xb7, 0xe9, 0x98, 0xcd, 0xf1, 0x29, 0x3a, 0x32, 0x99, 0xcc, 0xb0, 0xb9, 0x6e, 0x6c, 0xfb,
	0x87, 0x9a, 0x2d, 0xa5, 0xf9, 0xde, 0xe9, 0xda, 0xfe, 0xe9, 0x0b, 0xd4, 0x98, 0xf0, 0xc2, 0x6c,
	0xbb, 0xb6, 0xd7, 0x1e, 0x1c, 0xd3, 0x6a, 0x67, 0xfa, 0x67, 0x93, 0xb2, 0x8d, 0x51, 0x0e, 0xcf,
	0xd6, 0x1b, 0x62, 0xbd, 0x6d, 0x88, 0xb5, 0xdd, 0x10, 0x78, 0x52, 0x04, 0x5e, 0x15, 0x81, 0x95,
	0x22, 0xb0, 0x56, 0x04, 0x3e, 0x15, 0x81, 0x6f, 0x45, 0xac, 0xad, 0x22, 0xb0, 0xfc, 0x22, 0x56,
	0xf8, 0x4f, 0x2f, 0x74, 0xf9, 0x13, 0x00, 0x00, 0xff, 0xff, 0x04, 0x00, 0xc0, 0xa2, 0xe1, 0x01,
	0x00, 0x00,
}
