// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: greenfield/challenge/types.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// VoteResult defines the result attestation for a challenge.
type VoteResult int32

const (
	// The challenge failed.
	CHALLENGE_FAILED VoteResult = 0
	// The challenge succeed.
	CHALLENGE_SUCCEED VoteResult = 1
)

var VoteResult_name = map[int32]string{
	0: "CHALLENGE_FAILED",
	1: "CHALLENGE_SUCCEED",
}

var VoteResult_value = map[string]int32{
	"CHALLENGE_FAILED":  0,
	"CHALLENGE_SUCCEED": 1,
}

func (x VoteResult) String() string {
	return proto.EnumName(VoteResult_name, int32(x))
}

func (VoteResult) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9c297f0764c47d40, []int{0}
}

// Slash records the storage provider slashes, which will be pruned periodically.
type Slash struct {
	// The slashed storage provider.
	SpOperatorAddress []byte `protobuf:"bytes,1,opt,name=sp_operator_address,json=spOperatorAddress,proto3" json:"sp_operator_address,omitempty"`
	// The slashed object info.
	ObjectId Uint `protobuf:"bytes,2,opt,name=object_id,json=objectId,proto3,customtype=Uint" json:"object_id"`
	// The height when the slash happened, which is used for prune purpose.
	Height uint64 `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
}

func (m *Slash) Reset()         { *m = Slash{} }
func (m *Slash) String() string { return proto.CompactTextString(m) }
func (*Slash) ProtoMessage()    {}
func (*Slash) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c297f0764c47d40, []int{0}
}
func (m *Slash) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Slash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Slash.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Slash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Slash.Merge(m, src)
}
func (m *Slash) XXX_Size() int {
	return m.Size()
}
func (m *Slash) XXX_DiscardUnknown() {
	xxx_messageInfo_Slash.DiscardUnknown(m)
}

var xxx_messageInfo_Slash proto.InternalMessageInfo

func (m *Slash) GetSpOperatorAddress() []byte {
	if m != nil {
		return m.SpOperatorAddress
	}
	return nil
}

func (m *Slash) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

// Challenge records the challenge which are not expired yet.
type Challenge struct {
	// The id of the challenge.
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// The height at which the challenge will be expired.
	ExpiredHeight uint64 `protobuf:"varint,2,opt,name=expired_height,json=expiredHeight,proto3" json:"expired_height,omitempty"`
}

func (m *Challenge) Reset()         { *m = Challenge{} }
func (m *Challenge) String() string { return proto.CompactTextString(m) }
func (*Challenge) ProtoMessage()    {}
func (*Challenge) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c297f0764c47d40, []int{1}
}
func (m *Challenge) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Challenge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Challenge.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Challenge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Challenge.Merge(m, src)
}
func (m *Challenge) XXX_Size() int {
	return m.Size()
}
func (m *Challenge) XXX_DiscardUnknown() {
	xxx_messageInfo_Challenge.DiscardUnknown(m)
}

var xxx_messageInfo_Challenge proto.InternalMessageInfo

func (m *Challenge) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Challenge) GetExpiredHeight() uint64 {
	if m != nil {
		return m.ExpiredHeight
	}
	return 0
}

// AttestedChallengeIds stored fixed number of the latest attested challenge ids.
// To use the storage more efficiently, a circular queue will be constructed using these fields.
type AttestedChallengeIds struct {
	// The fixed number of challenge ids to save.
	Size_ uint64 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	// The latest attested challenge ids.
	Ids []uint64 `protobuf:"varint,2,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	// The cursor to retrieve data from the ids field.
	Cursor int64 `protobuf:"varint,3,opt,name=cursor,proto3" json:"cursor,omitempty"`
}

func (m *AttestedChallengeIds) Reset()         { *m = AttestedChallengeIds{} }
func (m *AttestedChallengeIds) String() string { return proto.CompactTextString(m) }
func (*AttestedChallengeIds) ProtoMessage()    {}
func (*AttestedChallengeIds) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c297f0764c47d40, []int{2}
}
func (m *AttestedChallengeIds) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AttestedChallengeIds) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AttestedChallengeIds.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AttestedChallengeIds) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttestedChallengeIds.Merge(m, src)
}
func (m *AttestedChallengeIds) XXX_Size() int {
	return m.Size()
}
func (m *AttestedChallengeIds) XXX_DiscardUnknown() {
	xxx_messageInfo_AttestedChallengeIds.DiscardUnknown(m)
}

var xxx_messageInfo_AttestedChallengeIds proto.InternalMessageInfo

func (m *AttestedChallengeIds) GetSize_() uint64 {
	if m != nil {
		return m.Size_
	}
	return 0
}

func (m *AttestedChallengeIds) GetIds() []uint64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *AttestedChallengeIds) GetCursor() int64 {
	if m != nil {
		return m.Cursor
	}
	return 0
}

func init() {
	proto.RegisterEnum("bnbchain.greenfield.challenge.VoteResult", VoteResult_name, VoteResult_value)
	proto.RegisterType((*Slash)(nil), "bnbchain.greenfield.challenge.Slash")
	proto.RegisterType((*Challenge)(nil), "bnbchain.greenfield.challenge.Challenge")
	proto.RegisterType((*AttestedChallengeIds)(nil), "bnbchain.greenfield.challenge.AttestedChallengeIds")
}

func init() { proto.RegisterFile("greenfield/challenge/types.proto", fileDescriptor_9c297f0764c47d40) }

var fileDescriptor_9c297f0764c47d40 = []byte{
	// 423 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x52, 0xcd, 0x6a, 0xd5, 0x40,
	0x14, 0xce, 0x24, 0xb1, 0x78, 0x47, 0x2d, 0xe9, 0x78, 0x95, 0x78, 0xc5, 0x34, 0x14, 0x84, 0x20,
	0x34, 0x41, 0xdc, 0xb8, 0x93, 0xdc, 0x34, 0xda, 0x8b, 0x17, 0x85, 0xd4, 0xba, 0x70, 0x13, 0x92,
	0xcc, 0x31, 0x19, 0x49, 0x33, 0x61, 0x66, 0x2e, 0x54, 0x9f, 0x40, 0x77, 0xbe, 0x83, 0xaf, 0xe0,
	0x43, 0x74, 0x59, 0x5c, 0x89, 0x8b, 0x22, 0xf7, 0xbe, 0x88, 0xe4, 0xc7, 0x56, 0xdc, 0x9d, 0xef,
	0x3b, 0xdf, 0xf9, 0xe6, 0x1b, 0xce, 0xc1, 0x6e, 0x29, 0x00, 0x9a, 0xf7, 0x0c, 0x6a, 0x1a, 0x14,
	0x55, 0x56, 0xd7, 0xd0, 0x94, 0x10, 0xa8, 0x8f, 0x2d, 0x48, 0xbf, 0x15, 0x5c, 0x71, 0xf2, 0x20,
	0x6f, 0xf2, 0xa2, 0xca, 0x58, 0xe3, 0x5f, 0x49, 0xfd, 0x4b, 0xe9, 0xec, 0x5e, 0xc1, 0xe5, 0x09,
	0x97, 0x69, 0x2f, 0x0e, 0x06, 0x30, 0x4c, 0xce, 0xa6, 0x25, 0x2f, 0xf9, 0xc0, 0x77, 0xd5, 0xc0,
	0xee, 0x7d, 0x41, 0xf8, 0xda, 0x51, 0x9d, 0xc9, 0x8a, 0xf8, 0xf8, 0xb6, 0x6c, 0x53, 0xde, 0x82,
	0xc8, 0x14, 0x17, 0x69, 0x46, 0xa9, 0x00, 0x29, 0x6d, 0xe4, 0x22, 0xef, 0x66, 0xb2, 0x23, 0xdb,
	0xd7, 0x63, 0x27, 0x1c, 0x1a, 0xe4, 0x29, 0x9e, 0xf0, 0xfc, 0x03, 0x14, 0x2a, 0x65, 0xd4, 0xd6,
	0x5d, 0xe4, 0x4d, 0xe6, 0xf7, 0xcf, 0x2e, 0x76, 0xb5, 0x5f, 0x17, 0xbb, 0xe6, 0x31, 0x6b, 0xd4,
	0x8f, 0xef, 0xfb, 0x37, 0xc6, 0x00, 0x1d, 0x4c, 0xae, 0x0f, 0xea, 0x05, 0x25, 0x77, 0xf1, 0x56,
	0x05, 0xac, 0xac, 0x94, 0x6d, 0xb8, 0xc8, 0x33, 0x93, 0x11, 0xed, 0xcd, 0xf1, 0x24, 0xfa, 0xfb,
	0x13, 0xb2, 0x8d, 0x75, 0x46, 0xfb, 0xd7, 0xcd, 0x44, 0x67, 0x94, 0x3c, 0xc4, 0xdb, 0x70, 0xda,
	0x32, 0x01, 0x34, 0x1d, 0x87, 0xf5, 0xbe, 0x77, 0x6b, 0x64, 0x0f, 0x07, 0x8f, 0x37, 0x78, 0x1a,
	0x2a, 0x05, 0x52, 0x01, 0xbd, 0xf4, 0x5a, 0x50, 0x49, 0x08, 0x36, 0x25, 0xfb, 0x04, 0xa3, 0x61,
	0x5f, 0x13, 0x0b, 0x1b, 0x8c, 0x4a, 0x5b, 0x77, 0x0d, 0xcf, 0x4c, 0xba, 0xb2, 0x4b, 0x56, 0xac,
	0x84, 0xe4, 0xa2, 0x4f, 0x66, 0x24, 0x23, 0x7a, 0xf4, 0x0c, 0xe3, 0xb7, 0x5c, 0x41, 0x02, 0x72,
	0x55, 0x2b, 0x32, 0xc5, 0x56, 0x74, 0x18, 0x2e, 0x97, 0xf1, 0xab, 0x17, 0x71, 0xfa, 0x3c, 0x5c,
	0x2c, 0xe3, 0x03, 0x4b, 0x23, 0x77, 0xf0, 0xce, 0x15, 0x7b, 0x74, 0x1c, 0x45, 0x71, 0x7c, 0x60,
	0xa1, 0x99, 0xf9, 0xf9, 0x9b, 0xa3, 0xcd, 0x5f, 0x9e, 0xad, 0x1d, 0x74, 0xbe, 0x76, 0xd0, 0xef,
	0xb5, 0x83, 0xbe, 0x6e, 0x1c, 0xed, 0x7c, 0xe3, 0x68, 0x3f, 0x37, 0x8e, 0xf6, 0xee, 0x71, 0xc9,
	0x54, 0xb5, 0xca, 0xfd, 0x82, 0x9f, 0x04, 0x79, 0x93, 0xef, 0xf7, 0xcb, 0x0d, 0xfe, 0xb9, 0x83,
	0xd3, 0xff, 0x2f, 0x21, 0xdf, 0xea, 0x57, 0xf7, 0xe4, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa3,
	0x26, 0x8b, 0x05, 0x2e, 0x02, 0x00, 0x00,
}

func (m *Slash) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Slash) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Slash) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Height != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x18
	}
	{
		size := m.ObjectId.Size()
		i -= size
		if _, err := m.ObjectId.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.SpOperatorAddress) > 0 {
		i -= len(m.SpOperatorAddress)
		copy(dAtA[i:], m.SpOperatorAddress)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.SpOperatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Challenge) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Challenge) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Challenge) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ExpiredHeight != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.ExpiredHeight))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *AttestedChallengeIds) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AttestedChallengeIds) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AttestedChallengeIds) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Cursor != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Cursor))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Ids) > 0 {
		dAtA2 := make([]byte, len(m.Ids)*10)
		var j1 int
		for _, num := range m.Ids {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintTypes(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x12
	}
	if m.Size_ != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Size_))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Slash) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SpOperatorAddress)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = m.ObjectId.Size()
	n += 1 + l + sovTypes(uint64(l))
	if m.Height != 0 {
		n += 1 + sovTypes(uint64(m.Height))
	}
	return n
}

func (m *Challenge) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovTypes(uint64(m.Id))
	}
	if m.ExpiredHeight != 0 {
		n += 1 + sovTypes(uint64(m.ExpiredHeight))
	}
	return n
}

func (m *AttestedChallengeIds) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Size_ != 0 {
		n += 1 + sovTypes(uint64(m.Size_))
	}
	if len(m.Ids) > 0 {
		l = 0
		for _, e := range m.Ids {
			l += sovTypes(uint64(e))
		}
		n += 1 + sovTypes(uint64(l)) + l
	}
	if m.Cursor != 0 {
		n += 1 + sovTypes(uint64(m.Cursor))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Slash) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Slash: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Slash: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpOperatorAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SpOperatorAddress = append(m.SpOperatorAddress[:0], dAtA[iNdEx:postIndex]...)
			if m.SpOperatorAddress == nil {
				m.SpOperatorAddress = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *Challenge) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Challenge: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Challenge: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpiredHeight", wireType)
			}
			m.ExpiredHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExpiredHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *AttestedChallengeIds) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: AttestedChallengeIds: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AttestedChallengeIds: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Size_", wireType)
			}
			m.Size_ = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Size_ |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTypes
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Ids = append(m.Ids, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTypes
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthTypes
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthTypes
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Ids) == 0 {
					m.Ids = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTypes
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Ids = append(m.Ids, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Ids", wireType)
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cursor", wireType)
			}
			m.Cursor = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Cursor |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
