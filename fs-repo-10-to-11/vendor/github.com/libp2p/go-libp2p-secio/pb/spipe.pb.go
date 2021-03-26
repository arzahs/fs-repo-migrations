// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: spipe.proto

package spipe_pb

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

type Propose struct {
	Rand      []byte `protobuf:"bytes,1,opt,name=rand" json:"rand"`
	Pubkey    []byte `protobuf:"bytes,2,opt,name=pubkey" json:"pubkey"`
	Exchanges string `protobuf:"bytes,3,opt,name=exchanges" json:"exchanges"`
	Ciphers   string `protobuf:"bytes,4,opt,name=ciphers" json:"ciphers"`
	Hashes    string `protobuf:"bytes,5,opt,name=hashes" json:"hashes"`
}

func (m *Propose) Reset()         { *m = Propose{} }
func (m *Propose) String() string { return proto.CompactTextString(m) }
func (*Propose) ProtoMessage()    {}
func (*Propose) Descriptor() ([]byte, []int) {
	return fileDescriptor_c474ec75f0379e64, []int{0}
}
func (m *Propose) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Propose) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Propose.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Propose) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Propose.Merge(m, src)
}
func (m *Propose) XXX_Size() int {
	return m.Size()
}
func (m *Propose) XXX_DiscardUnknown() {
	xxx_messageInfo_Propose.DiscardUnknown(m)
}

var xxx_messageInfo_Propose proto.InternalMessageInfo

func (m *Propose) GetRand() []byte {
	if m != nil {
		return m.Rand
	}
	return nil
}

func (m *Propose) GetPubkey() []byte {
	if m != nil {
		return m.Pubkey
	}
	return nil
}

func (m *Propose) GetExchanges() string {
	if m != nil {
		return m.Exchanges
	}
	return ""
}

func (m *Propose) GetCiphers() string {
	if m != nil {
		return m.Ciphers
	}
	return ""
}

func (m *Propose) GetHashes() string {
	if m != nil {
		return m.Hashes
	}
	return ""
}

type Exchange struct {
	Epubkey   []byte `protobuf:"bytes,1,opt,name=epubkey" json:"epubkey"`
	Signature []byte `protobuf:"bytes,2,opt,name=signature" json:"signature"`
}

func (m *Exchange) Reset()         { *m = Exchange{} }
func (m *Exchange) String() string { return proto.CompactTextString(m) }
func (*Exchange) ProtoMessage()    {}
func (*Exchange) Descriptor() ([]byte, []int) {
	return fileDescriptor_c474ec75f0379e64, []int{1}
}
func (m *Exchange) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Exchange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Exchange.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Exchange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Exchange.Merge(m, src)
}
func (m *Exchange) XXX_Size() int {
	return m.Size()
}
func (m *Exchange) XXX_DiscardUnknown() {
	xxx_messageInfo_Exchange.DiscardUnknown(m)
}

var xxx_messageInfo_Exchange proto.InternalMessageInfo

func (m *Exchange) GetEpubkey() []byte {
	if m != nil {
		return m.Epubkey
	}
	return nil
}

func (m *Exchange) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func init() {
	proto.RegisterType((*Propose)(nil), "spipe.pb.Propose")
	proto.RegisterType((*Exchange)(nil), "spipe.pb.Exchange")
}

func init() { proto.RegisterFile("spipe.proto", fileDescriptor_c474ec75f0379e64) }

var fileDescriptor_c474ec75f0379e64 = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x2e, 0xc8, 0x2c,
	0x48, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0x72, 0x92, 0x94, 0x16, 0x33, 0x72,
	0xb1, 0x07, 0x14, 0xe5, 0x17, 0xe4, 0x17, 0xa7, 0x0a, 0x49, 0x70, 0xb1, 0x14, 0x25, 0xe6, 0xa5,
	0x48, 0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x38, 0xb1, 0x9c, 0xb8, 0x27, 0xcf, 0x10, 0x04, 0x16, 0x11,
	0x92, 0xe1, 0x62, 0x2b, 0x28, 0x4d, 0xca, 0x4e, 0xad, 0x94, 0x60, 0x42, 0x92, 0x83, 0x8a, 0x09,
	0x29, 0x71, 0x71, 0xa6, 0x56, 0x24, 0x67, 0x24, 0xe6, 0xa5, 0xa7, 0x16, 0x4b, 0x30, 0x2b, 0x30,
	0x6a, 0x70, 0x42, 0x15, 0x20, 0x84, 0x85, 0xe4, 0xb8, 0xd8, 0x93, 0x33, 0x0b, 0x32, 0x52, 0x8b,
	0x8a, 0x25, 0x58, 0x90, 0x54, 0xc0, 0x04, 0x41, 0x36, 0x64, 0x24, 0x16, 0x67, 0xa4, 0x16, 0x4b,
	0xb0, 0x22, 0x49, 0x43, 0xc5, 0x94, 0xfc, 0xb8, 0x38, 0x5c, 0xa1, 0x46, 0x81, 0x4c, 0x4a, 0x85,
	0x3a, 0x06, 0xd9, 0xa1, 0x30, 0x41, 0x90, 0x6b, 0x8a, 0x33, 0xd3, 0xf3, 0x12, 0x4b, 0x4a, 0x8b,
	0x52, 0x51, 0x9c, 0x8b, 0x10, 0x76, 0x92, 0x38, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6,
	0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39,
	0x06, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7c, 0x0d, 0xc4, 0xcb, 0x27, 0x01, 0x00, 0x00,
}

func (m *Propose) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Propose) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Propose) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	i -= len(m.Hashes)
	copy(dAtA[i:], m.Hashes)
	i = encodeVarintSpipe(dAtA, i, uint64(len(m.Hashes)))
	i--
	dAtA[i] = 0x2a
	i -= len(m.Ciphers)
	copy(dAtA[i:], m.Ciphers)
	i = encodeVarintSpipe(dAtA, i, uint64(len(m.Ciphers)))
	i--
	dAtA[i] = 0x22
	i -= len(m.Exchanges)
	copy(dAtA[i:], m.Exchanges)
	i = encodeVarintSpipe(dAtA, i, uint64(len(m.Exchanges)))
	i--
	dAtA[i] = 0x1a
	if m.Pubkey != nil {
		i -= len(m.Pubkey)
		copy(dAtA[i:], m.Pubkey)
		i = encodeVarintSpipe(dAtA, i, uint64(len(m.Pubkey)))
		i--
		dAtA[i] = 0x12
	}
	if m.Rand != nil {
		i -= len(m.Rand)
		copy(dAtA[i:], m.Rand)
		i = encodeVarintSpipe(dAtA, i, uint64(len(m.Rand)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Exchange) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Exchange) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Exchange) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Signature != nil {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintSpipe(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x12
	}
	if m.Epubkey != nil {
		i -= len(m.Epubkey)
		copy(dAtA[i:], m.Epubkey)
		i = encodeVarintSpipe(dAtA, i, uint64(len(m.Epubkey)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSpipe(dAtA []byte, offset int, v uint64) int {
	offset -= sovSpipe(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Propose) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Rand != nil {
		l = len(m.Rand)
		n += 1 + l + sovSpipe(uint64(l))
	}
	if m.Pubkey != nil {
		l = len(m.Pubkey)
		n += 1 + l + sovSpipe(uint64(l))
	}
	l = len(m.Exchanges)
	n += 1 + l + sovSpipe(uint64(l))
	l = len(m.Ciphers)
	n += 1 + l + sovSpipe(uint64(l))
	l = len(m.Hashes)
	n += 1 + l + sovSpipe(uint64(l))
	return n
}

func (m *Exchange) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Epubkey != nil {
		l = len(m.Epubkey)
		n += 1 + l + sovSpipe(uint64(l))
	}
	if m.Signature != nil {
		l = len(m.Signature)
		n += 1 + l + sovSpipe(uint64(l))
	}
	return n
}

func sovSpipe(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSpipe(x uint64) (n int) {
	return sovSpipe(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Propose) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSpipe
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
			return fmt.Errorf("proto: Propose: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Propose: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rand", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpipe
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
				return ErrInvalidLengthSpipe
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSpipe
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Rand = append(m.Rand[:0], dAtA[iNdEx:postIndex]...)
			if m.Rand == nil {
				m.Rand = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pubkey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpipe
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
				return ErrInvalidLengthSpipe
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSpipe
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pubkey = append(m.Pubkey[:0], dAtA[iNdEx:postIndex]...)
			if m.Pubkey == nil {
				m.Pubkey = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Exchanges", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpipe
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
				return ErrInvalidLengthSpipe
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSpipe
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Exchanges = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ciphers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpipe
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
				return ErrInvalidLengthSpipe
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSpipe
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ciphers = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hashes", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpipe
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
				return ErrInvalidLengthSpipe
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSpipe
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hashes = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSpipe(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSpipe
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSpipe
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
func (m *Exchange) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSpipe
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
			return fmt.Errorf("proto: Exchange: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Exchange: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Epubkey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpipe
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
				return ErrInvalidLengthSpipe
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSpipe
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Epubkey = append(m.Epubkey[:0], dAtA[iNdEx:postIndex]...)
			if m.Epubkey == nil {
				m.Epubkey = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpipe
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
				return ErrInvalidLengthSpipe
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSpipe
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = append(m.Signature[:0], dAtA[iNdEx:postIndex]...)
			if m.Signature == nil {
				m.Signature = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSpipe(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSpipe
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSpipe
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
func skipSpipe(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSpipe
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
					return 0, ErrIntOverflowSpipe
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
					return 0, ErrIntOverflowSpipe
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
				return 0, ErrInvalidLengthSpipe
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSpipe
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSpipe
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSpipe        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSpipe          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSpipe = fmt.Errorf("proto: unexpected end of group")
)