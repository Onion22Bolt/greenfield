// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: greenfield/payment/events.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
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

// EventCreatePaymentAccount is emitted on MsgCreatePaymentAccount
type EventCreatePaymentAccount struct {
	// address of the payment account
	Addr string `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	// owner address of the payment account
	Owner string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	// index of the payment account of the owner
	Index uint64 `protobuf:"varint,3,opt,name=index,proto3" json:"index,omitempty"`
}

func (m *EventCreatePaymentAccount) Reset()         { *m = EventCreatePaymentAccount{} }
func (m *EventCreatePaymentAccount) String() string { return proto.CompactTextString(m) }
func (*EventCreatePaymentAccount) ProtoMessage()    {}
func (*EventCreatePaymentAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_befcc80e27bc8df9, []int{0}
}
func (m *EventCreatePaymentAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventCreatePaymentAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventCreatePaymentAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventCreatePaymentAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventCreatePaymentAccount.Merge(m, src)
}
func (m *EventCreatePaymentAccount) XXX_Size() int {
	return m.Size()
}
func (m *EventCreatePaymentAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_EventCreatePaymentAccount.DiscardUnknown(m)
}

var xxx_messageInfo_EventCreatePaymentAccount proto.InternalMessageInfo

func (m *EventCreatePaymentAccount) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *EventCreatePaymentAccount) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *EventCreatePaymentAccount) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

// EventForceSettle may be emitted on all Msgs and EndBlocker when a payment account's
// balance or net outflow rate is changed
type EventForceSettle struct {
	// address of the payment account
	Addr string `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	// left balance of the payment account after force settlement
	// if the balance is positive, it will go to the governance stream account
	// if the balance is negative, it's the debt of the system, which will be paid by the governance stream account
	SettledBalance github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=settled_balance,json=settledBalance,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"settled_balance"`
}

func (m *EventForceSettle) Reset()         { *m = EventForceSettle{} }
func (m *EventForceSettle) String() string { return proto.CompactTextString(m) }
func (*EventForceSettle) ProtoMessage()    {}
func (*EventForceSettle) Descriptor() ([]byte, []int) {
	return fileDescriptor_befcc80e27bc8df9, []int{1}
}
func (m *EventForceSettle) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventForceSettle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventForceSettle.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventForceSettle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventForceSettle.Merge(m, src)
}
func (m *EventForceSettle) XXX_Size() int {
	return m.Size()
}
func (m *EventForceSettle) XXX_DiscardUnknown() {
	xxx_messageInfo_EventForceSettle.DiscardUnknown(m)
}

var xxx_messageInfo_EventForceSettle proto.InternalMessageInfo

func (m *EventForceSettle) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func init() {
	proto.RegisterType((*EventCreatePaymentAccount)(nil), "bnbchain.greenfield.payment.EventCreatePaymentAccount")
	proto.RegisterType((*EventForceSettle)(nil), "bnbchain.greenfield.payment.EventForceSettle")
}

func init() { proto.RegisterFile("greenfield/payment/events.proto", fileDescriptor_befcc80e27bc8df9) }

var fileDescriptor_befcc80e27bc8df9 = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x31, 0x4f, 0x02, 0x31,
	0x14, 0xc7, 0xaf, 0x0a, 0x26, 0x76, 0x50, 0x73, 0x61, 0x38, 0x30, 0x39, 0x08, 0x83, 0x61, 0xf0,
	0xee, 0x06, 0x57, 0x17, 0x30, 0x9a, 0xb0, 0x19, 0xd8, 0x5c, 0xc8, 0xb5, 0x7d, 0x1e, 0x17, 0xa1,
	0x25, 0x6d, 0x51, 0xf8, 0x14, 0xfa, 0x2d, 0xfc, 0x02, 0x7c, 0x08, 0x46, 0xc2, 0x64, 0x1c, 0x88,
	0xe1, 0xbe, 0x88, 0xa1, 0x6d, 0x94, 0xcd, 0x38, 0xdd, 0xbd, 0x97, 0xdf, 0xfb, 0xbd, 0x97, 0xfe,
	0x71, 0x3d, 0x93, 0x00, 0xfc, 0x31, 0x87, 0x11, 0x4b, 0x26, 0xe9, 0x7c, 0x0c, 0x5c, 0x27, 0xf0,
	0x0c, 0x5c, 0xab, 0x78, 0x22, 0x85, 0x16, 0xfe, 0x39, 0xe1, 0x84, 0x0e, 0xd3, 0x9c, 0xc7, 0xbf,
	0x64, 0xec, 0xc8, 0x5a, 0x95, 0x0a, 0x35, 0x16, 0x6a, 0x60, 0xd0, 0xc4, 0x16, 0x76, 0xae, 0x56,
	0xc9, 0x44, 0x26, 0x6c, 0x7f, 0xf7, 0x67, 0xbb, 0xcd, 0x57, 0x84, 0xab, 0xb7, 0x3b, 0xfd, 0x8d,
	0x84, 0x54, 0xc3, 0xbd, 0xf5, 0xb4, 0x29, 0x15, 0x53, 0xae, 0xfd, 0x4b, 0x5c, 0x4a, 0x19, 0x93,
	0x01, 0x6a, 0xa0, 0xd6, 0x71, 0x27, 0x58, 0x2f, 0xa2, 0x8a, 0x73, 0xb6, 0x19, 0x93, 0xa0, 0x54,
	0x5f, 0xcb, 0x9c, 0x67, 0x3d, 0x43, 0xf9, 0x31, 0x2e, 0x8b, 0x17, 0x0e, 0x32, 0x38, 0xf8, 0x03,
	0xb7, 0x98, 0x5f, 0xc1, 0xe5, 0x9c, 0x33, 0x98, 0x05, 0x87, 0x0d, 0xd4, 0x2a, 0xf5, 0x6c, 0xd1,
	0x7c, 0x47, 0xf8, 0xcc, 0x5c, 0x74, 0x27, 0x24, 0x85, 0x3e, 0x68, 0x3d, 0x82, 0x7f, 0x1e, 0x02,
	0xf8, 0x54, 0x99, 0x39, 0x36, 0x20, 0xe9, 0x28, 0xe5, 0x14, 0xdc, 0x49, 0xd7, 0xcb, 0x4d, 0xdd,
	0xfb, 0xdc, 0xd4, 0x2f, 0xb2, 0x5c, 0x0f, 0xa7, 0x24, 0xa6, 0x62, 0xec, 0x1e, 0xc9, 0x7d, 0x22,
	0xc5, 0x9e, 0x12, 0x3d, 0x9f, 0x80, 0x8a, 0xbb, 0x5c, 0xaf, 0x17, 0x11, 0x76, 0x6b, 0xba, 0x5c,
	0xf7, 0x4e, 0x9c, 0xb4, 0x63, 0x9d, 0x9d, 0xee, 0x72, 0x1b, 0xa2, 0xd5, 0x36, 0x44, 0x5f, 0xdb,
	0x10, 0xbd, 0x15, 0xa1, 0xb7, 0x2a, 0x42, 0xef, 0xa3, 0x08, 0xbd, 0x87, 0x64, 0xcf, 0x4f, 0x38,
	0x89, 0x4c, 0x5e, 0xc9, 0x5e, 0xb2, 0xb3, 0x9f, 0x6c, 0xcd, 0x32, 0x72, 0x64, 0xd2, 0xb8, 0xfa,
	0x0e, 0x00, 0x00, 0xff, 0xff, 0xaa, 0xbc, 0x87, 0xa3, 0xfe, 0x01, 0x00, 0x00,
}

func (m *EventCreatePaymentAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventCreatePaymentAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventCreatePaymentAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Index != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Addr) > 0 {
		i -= len(m.Addr)
		copy(dAtA[i:], m.Addr)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Addr)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventForceSettle) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventForceSettle) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventForceSettle) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.SettledBalance.Size()
		i -= size
		if _, err := m.SettledBalance.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEvents(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Addr) > 0 {
		i -= len(m.Addr)
		copy(dAtA[i:], m.Addr)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Addr)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventCreatePaymentAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Addr)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.Index != 0 {
		n += 1 + sovEvents(uint64(m.Index))
	}
	return n
}

func (m *EventForceSettle) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Addr)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = m.SettledBalance.Size()
	n += 1 + l + sovEvents(uint64(l))
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventCreatePaymentAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventCreatePaymentAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventCreatePaymentAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Addr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventForceSettle) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventForceSettle: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventForceSettle: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Addr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SettledBalance", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SettledBalance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)
