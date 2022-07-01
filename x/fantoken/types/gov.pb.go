// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bitsong/fantoken/v1beta1/gov.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
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

type UpdateFeesProposal struct {
	Title       string     `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string     `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	IssueFee    types.Coin `protobuf:"bytes,3,opt,name=issue_fee,json=issueFee,proto3" json:"issue_fee" yaml:"issue_fee"`
	MintFee     types.Coin `protobuf:"bytes,4,opt,name=mint_fee,json=mintFee,proto3" json:"mint_fee" yaml:"mint_fee"`
	BurnFee     types.Coin `protobuf:"bytes,5,opt,name=burn_fee,json=burnFee,proto3" json:"burn_fee" yaml:"burn_fee"`
}

func (m *UpdateFeesProposal) Reset()      { *m = UpdateFeesProposal{} }
func (*UpdateFeesProposal) ProtoMessage() {}
func (*UpdateFeesProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_1525a26433a8d1c3, []int{0}
}
func (m *UpdateFeesProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpdateFeesProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpdateFeesProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpdateFeesProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateFeesProposal.Merge(m, src)
}
func (m *UpdateFeesProposal) XXX_Size() int {
	return m.Size()
}
func (m *UpdateFeesProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateFeesProposal.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateFeesProposal proto.InternalMessageInfo

type UpdateFeesProposalWithDeposit struct {
	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	IssueFee    string `protobuf:"bytes,3,opt,name=issue_fee,json=issueFee,proto3" json:"issue_fee,omitempty"`
	MintFee     string `protobuf:"bytes,4,opt,name=mint_fee,json=mintFee,proto3" json:"mint_fee,omitempty"`
	BurnFee     string `protobuf:"bytes,5,opt,name=burn_fee,json=burnFee,proto3" json:"burn_fee,omitempty"`
	Deposit     string `protobuf:"bytes,7,opt,name=deposit,proto3" json:"deposit,omitempty"`
}

func (m *UpdateFeesProposalWithDeposit) Reset()         { *m = UpdateFeesProposalWithDeposit{} }
func (m *UpdateFeesProposalWithDeposit) String() string { return proto.CompactTextString(m) }
func (*UpdateFeesProposalWithDeposit) ProtoMessage()    {}
func (*UpdateFeesProposalWithDeposit) Descriptor() ([]byte, []int) {
	return fileDescriptor_1525a26433a8d1c3, []int{1}
}
func (m *UpdateFeesProposalWithDeposit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpdateFeesProposalWithDeposit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpdateFeesProposalWithDeposit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpdateFeesProposalWithDeposit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateFeesProposalWithDeposit.Merge(m, src)
}
func (m *UpdateFeesProposalWithDeposit) XXX_Size() int {
	return m.Size()
}
func (m *UpdateFeesProposalWithDeposit) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateFeesProposalWithDeposit.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateFeesProposalWithDeposit proto.InternalMessageInfo

func init() {
	proto.RegisterType((*UpdateFeesProposal)(nil), "bitsong.fantoken.v1beta1.UpdateFeesProposal")
	proto.RegisterType((*UpdateFeesProposalWithDeposit)(nil), "bitsong.fantoken.v1beta1.UpdateFeesProposalWithDeposit")
}

func init() {
	proto.RegisterFile("bitsong/fantoken/v1beta1/gov.proto", fileDescriptor_1525a26433a8d1c3)
}

var fileDescriptor_1525a26433a8d1c3 = []byte{
	// 409 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x4d, 0xab, 0xd3, 0x40,
	0x14, 0xcd, 0x3c, 0xdf, 0x33, 0xcd, 0xbc, 0x85, 0x12, 0x1e, 0x98, 0x3e, 0x71, 0x52, 0xb2, 0x7a,
	0x1b, 0x13, 0xaa, 0xe0, 0xa2, 0xcb, 0x2a, 0xdd, 0x09, 0x25, 0x28, 0x82, 0x1b, 0x49, 0xd2, 0x9b,
	0x74, 0x30, 0x99, 0x1b, 0x32, 0xd3, 0x62, 0xff, 0x85, 0xcb, 0x2e, 0xfb, 0x6b, 0xa4, 0xb8, 0xea,
	0xd2, 0x55, 0xd1, 0x76, 0xe3, 0xda, 0x5f, 0x20, 0xf9, 0x6a, 0x4b, 0x5d, 0x88, 0xee, 0xce, 0xcc,
	0xbd, 0xe7, 0x70, 0xee, 0xe1, 0x50, 0x27, 0xe4, 0x4a, 0xa2, 0x48, 0xbc, 0x38, 0x10, 0x0a, 0x3f,
	0x82, 0xf0, 0xe6, 0xfd, 0x10, 0x54, 0xd0, 0xf7, 0x12, 0x9c, 0xbb, 0x79, 0x81, 0x0a, 0x4d, 0xab,
	0xd9, 0x71, 0xdb, 0x1d, 0xb7, 0xd9, 0xb9, 0x65, 0x11, 0xca, 0x0c, 0xa5, 0x17, 0x06, 0x12, 0x0e,
	0xc4, 0x08, 0xb9, 0xa8, 0x99, 0xb7, 0x37, 0x09, 0x26, 0x58, 0x41, 0xaf, 0x44, 0xf5, 0xaf, 0xf3,
	0xe5, 0x82, 0x9a, 0x6f, 0xf3, 0x49, 0xa0, 0x60, 0x04, 0x20, 0xc7, 0x05, 0xe6, 0x28, 0x83, 0xd4,
	0xbc, 0xa1, 0x57, 0x8a, 0xab, 0x14, 0x2c, 0xd2, 0x23, 0x77, 0x86, 0x5f, 0x3f, 0xcc, 0x1e, 0xbd,
	0x9e, 0x80, 0x8c, 0x0a, 0x9e, 0x2b, 0x8e, 0xc2, 0xba, 0xa8, 0x66, 0xa7, 0x5f, 0xe6, 0x98, 0x1a,
	0x5c, 0xca, 0x19, 0x7c, 0x88, 0x01, 0xac, 0x7b, 0x3d, 0x72, 0x77, 0xfd, 0xac, 0xeb, 0xd6, 0xc6,
	0xdc, 0xd2, 0x58, 0xeb, 0xd6, 0x7d, 0x89, 0x5c, 0x0c, 0xad, 0xf5, 0xd6, 0xd6, 0x7e, 0x6d, 0xed,
	0x87, 0x8b, 0x20, 0x4b, 0x07, 0xce, 0x81, 0xe9, 0xf8, 0x9d, 0x0a, 0x8f, 0x00, 0xcc, 0xd7, 0xb4,
	0x93, 0x71, 0xa1, 0x2a, 0xc1, 0xcb, 0xbf, 0x09, 0x3e, 0x6a, 0x04, 0x1f, 0xd4, 0x82, 0x2d, 0xd1,
	0xf1, 0xf5, 0x12, 0x36, 0x72, 0xe1, 0xac, 0x10, 0x95, 0xdc, 0xd5, 0x3f, 0xca, 0xb5, 0x44, 0xc7,
	0xd7, 0x4b, 0x38, 0x02, 0x18, 0x74, 0x96, 0x2b, 0x5b, 0xfb, 0xb9, 0xb2, 0x89, 0xf3, 0x95, 0xd0,
	0x27, 0x7f, 0x06, 0xf9, 0x8e, 0xab, 0xe9, 0x2b, 0xc8, 0x51, 0x72, 0xf5, 0xdf, 0x99, 0x3e, 0x3e,
	0xcf, 0xd4, 0x38, 0x89, 0xa7, 0x7b, 0x16, 0x8f, 0x71, 0x3c, 0xb5, 0x7b, 0x76, 0xaa, 0x71, 0xb0,
	0x6d, 0x5a, 0x54, 0x9f, 0xd4, 0xae, 0x2c, 0xbd, 0x9e, 0x34, 0xcf, 0xc1, 0xe5, 0x72, 0x65, 0x93,
	0xe1, 0x9b, 0xf5, 0x0f, 0xa6, 0xad, 0x77, 0x8c, 0x6c, 0x76, 0x8c, 0x7c, 0xdf, 0x31, 0xf2, 0x79,
	0xcf, 0xb4, 0xcd, 0x9e, 0x69, 0xdf, 0xf6, 0x4c, 0x7b, 0xff, 0x22, 0xe1, 0x6a, 0x3a, 0x0b, 0xdd,
	0x08, 0x33, 0xaf, 0xa9, 0x23, 0xc6, 0x31, 0x8f, 0x78, 0x90, 0x7a, 0x09, 0x3e, 0x6d, 0x5b, 0xfc,
	0xe9, 0xd8, 0x63, 0xb5, 0xc8, 0x41, 0x86, 0xf7, 0xab, 0xca, 0x3d, 0xff, 0x1d, 0x00, 0x00, 0xff,
	0xff, 0x1d, 0xf5, 0x21, 0x4f, 0xe8, 0x02, 0x00, 0x00,
}

func (this *UpdateFeesProposal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpdateFeesProposal)
	if !ok {
		that2, ok := that.(UpdateFeesProposal)
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
	if this.Title != that1.Title {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if !this.IssueFee.Equal(&that1.IssueFee) {
		return false
	}
	if !this.MintFee.Equal(&that1.MintFee) {
		return false
	}
	if !this.BurnFee.Equal(&that1.BurnFee) {
		return false
	}
	return true
}
func (m *UpdateFeesProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdateFeesProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UpdateFeesProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.BurnFee.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGov(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size, err := m.MintFee.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGov(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size, err := m.IssueFee.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGov(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UpdateFeesProposalWithDeposit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdateFeesProposalWithDeposit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UpdateFeesProposalWithDeposit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Deposit) > 0 {
		i -= len(m.Deposit)
		copy(dAtA[i:], m.Deposit)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Deposit)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.BurnFee) > 0 {
		i -= len(m.BurnFee)
		copy(dAtA[i:], m.BurnFee)
		i = encodeVarintGov(dAtA, i, uint64(len(m.BurnFee)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.MintFee) > 0 {
		i -= len(m.MintFee)
		copy(dAtA[i:], m.MintFee)
		i = encodeVarintGov(dAtA, i, uint64(len(m.MintFee)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.IssueFee) > 0 {
		i -= len(m.IssueFee)
		copy(dAtA[i:], m.IssueFee)
		i = encodeVarintGov(dAtA, i, uint64(len(m.IssueFee)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGov(dAtA []byte, offset int, v uint64) int {
	offset -= sovGov(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *UpdateFeesProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = m.IssueFee.Size()
	n += 1 + l + sovGov(uint64(l))
	l = m.MintFee.Size()
	n += 1 + l + sovGov(uint64(l))
	l = m.BurnFee.Size()
	n += 1 + l + sovGov(uint64(l))
	return n
}

func (m *UpdateFeesProposalWithDeposit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.IssueFee)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.MintFee)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.BurnFee)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.Deposit)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	return n
}

func sovGov(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGov(x uint64) (n int) {
	return sovGov(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UpdateFeesProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
			return fmt.Errorf("proto: UpdateFeesProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdateFeesProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IssueFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.IssueFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MintFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BurnFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BurnFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGov
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
func (m *UpdateFeesProposalWithDeposit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
			return fmt.Errorf("proto: UpdateFeesProposalWithDeposit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdateFeesProposalWithDeposit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IssueFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IssueFee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MintFee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BurnFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BurnFee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deposit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Deposit = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGov
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
func skipGov(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGov
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
					return 0, ErrIntOverflowGov
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
					return 0, ErrIntOverflowGov
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
				return 0, ErrInvalidLengthGov
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGov
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGov
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGov        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGov          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGov = fmt.Errorf("proto: unexpected end of group")
)