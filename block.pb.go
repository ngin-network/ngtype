// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: block.proto

package ngtype

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Block struct {
	Height               uint64       `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Timestamp            int64        `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	TrieHash             []byte       `protobuf:"bytes,3,opt,name=trie_hash,json=trieHash,proto3" json:"trie_hash,omitempty"`
	PrevBlockHash        []byte       `protobuf:"bytes,4,opt,name=prev_block_hash,json=prevBlockHash,proto3" json:"prev_block_hash,omitempty"`
	PrevVaultHash        []byte       `protobuf:"bytes,5,opt,name=prev_vault_hash,json=prevVaultHash,proto3" json:"prev_vault_hash,omitempty"`
	Beneficiary          []byte       `protobuf:"bytes,6,opt,name=beneficiary,proto3" json:"beneficiary,omitempty"`
	Difficulty           []byte       `protobuf:"bytes,7,opt,name=difficulty,proto3" json:"difficulty,omitempty"`
	Nonce                []byte       `protobuf:"bytes,8,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Operations           []*Operation `protobuf:"bytes,9,rep,name=operations,proto3" json:"operations,omitempty"`
	Hash                 []byte       `protobuf:"bytes,10,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e550b1f5926e92d, []int{0}
}
func (m *Block) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Block.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(m, src)
}
func (m *Block) XXX_Size() int {
	return m.Size()
}
func (m *Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Block proto.InternalMessageInfo

func (m *Block) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Block) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Block) GetTrieHash() []byte {
	if m != nil {
		return m.TrieHash
	}
	return nil
}

func (m *Block) GetPrevBlockHash() []byte {
	if m != nil {
		return m.PrevBlockHash
	}
	return nil
}

func (m *Block) GetPrevVaultHash() []byte {
	if m != nil {
		return m.PrevVaultHash
	}
	return nil
}

func (m *Block) GetBeneficiary() []byte {
	if m != nil {
		return m.Beneficiary
	}
	return nil
}

func (m *Block) GetDifficulty() []byte {
	if m != nil {
		return m.Difficulty
	}
	return nil
}

func (m *Block) GetNonce() []byte {
	if m != nil {
		return m.Nonce
	}
	return nil
}

func (m *Block) GetOperations() []*Operation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *Block) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func init() {
	proto.RegisterType((*Block)(nil), "ngtype.Block")
}

func init() { proto.RegisterFile("block.proto", fileDescriptor_8e550b1f5926e92d) }

var fileDescriptor_8e550b1f5926e92d = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xcf, 0x4e, 0x83, 0x40,
	0x10, 0xc6, 0xdd, 0xf2, 0xc7, 0x32, 0x68, 0xaa, 0x13, 0x63, 0x36, 0x6a, 0x08, 0xf1, 0x60, 0x38,
	0x91, 0xa8, 0x6f, 0xd0, 0x93, 0x37, 0x13, 0x0e, 0x5e, 0x9b, 0x05, 0x97, 0xb2, 0x91, 0xb2, 0x04,
	0xb6, 0x4d, 0x78, 0x13, 0x1f, 0x49, 0x6f, 0x3e, 0x82, 0xc1, 0x17, 0x31, 0x0c, 0x6a, 0xf1, 0xb6,
	0xf3, 0xfb, 0x7e, 0xd9, 0x99, 0x7c, 0xe0, 0xa7, 0xa5, 0xce, 0x5e, 0xe2, 0xba, 0xd1, 0x46, 0xa3,
	0x5b, 0xad, 0x4d, 0x57, 0xcb, 0x8b, 0x85, 0xae, 0x65, 0x23, 0x8c, 0xd2, 0xd5, 0x18, 0x5c, 0xbf,
	0xcf, 0xc0, 0x59, 0x0e, 0x22, 0x9e, 0x83, 0x5b, 0x48, 0xb5, 0x2e, 0x0c, 0x67, 0x21, 0x8b, 0xec,
	0xe4, 0x67, 0xc2, 0x2b, 0xf0, 0x8c, 0xda, 0xc8, 0xd6, 0x88, 0x4d, 0xcd, 0x67, 0x21, 0x8b, 0xac,
	0x64, 0x0f, 0xf0, 0x12, 0x3c, 0xd3, 0x28, 0xb9, 0x2a, 0x44, 0x5b, 0x70, 0x2b, 0x64, 0xd1, 0x51,
	0x32, 0x1f, 0xc0, 0x83, 0x68, 0x0b, 0xbc, 0x81, 0x45, 0xdd, 0xc8, 0xdd, 0x8a, 0x2e, 0x19, 0x15,
	0x9b, 0x94, 0xe3, 0x01, 0xd3, 0xda, 0x7f, 0xde, 0x4e, 0x6c, 0x4b, 0x33, 0x7a, 0xce, 0xde, 0x7b,
	0x1a, 0x28, 0x79, 0x21, 0xf8, 0xa9, 0xac, 0x64, 0xae, 0x32, 0x25, 0x9a, 0x8e, 0xbb, 0xe4, 0x4c,
	0x11, 0x06, 0x00, 0xcf, 0x2a, 0xcf, 0x55, 0xb6, 0x2d, 0x4d, 0xc7, 0x0f, 0x49, 0x98, 0x10, 0x3c,
	0x03, 0xa7, 0xd2, 0x55, 0x26, 0xf9, 0x9c, 0xa2, 0x71, 0xc0, 0x5b, 0x80, 0xbf, 0x5e, 0x5a, 0xee,
	0x85, 0x56, 0xe4, 0xdf, 0x9d, 0xc6, 0x63, 0x65, 0xf1, 0xe3, 0x6f, 0x92, 0x4c, 0x24, 0x44, 0xb0,
	0xe9, 0x4e, 0xa0, 0x7f, 0xe8, 0xbd, 0x3c, 0x79, 0xeb, 0x03, 0xf6, 0xd1, 0x07, 0xec, 0xb3, 0x0f,
	0xd8, 0xeb, 0x57, 0x70, 0x90, 0xba, 0x54, 0xf2, 0xfd, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0c,
	0x27, 0x07, 0xcb, 0x8c, 0x01, 0x00, 0x00,
}

func (m *Block) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Block) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Block) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintBlock(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.Operations) > 0 {
		for iNdEx := len(m.Operations) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Operations[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBlock(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	if len(m.Nonce) > 0 {
		i -= len(m.Nonce)
		copy(dAtA[i:], m.Nonce)
		i = encodeVarintBlock(dAtA, i, uint64(len(m.Nonce)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Difficulty) > 0 {
		i -= len(m.Difficulty)
		copy(dAtA[i:], m.Difficulty)
		i = encodeVarintBlock(dAtA, i, uint64(len(m.Difficulty)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Beneficiary) > 0 {
		i -= len(m.Beneficiary)
		copy(dAtA[i:], m.Beneficiary)
		i = encodeVarintBlock(dAtA, i, uint64(len(m.Beneficiary)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.PrevVaultHash) > 0 {
		i -= len(m.PrevVaultHash)
		copy(dAtA[i:], m.PrevVaultHash)
		i = encodeVarintBlock(dAtA, i, uint64(len(m.PrevVaultHash)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.PrevBlockHash) > 0 {
		i -= len(m.PrevBlockHash)
		copy(dAtA[i:], m.PrevBlockHash)
		i = encodeVarintBlock(dAtA, i, uint64(len(m.PrevBlockHash)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.TrieHash) > 0 {
		i -= len(m.TrieHash)
		copy(dAtA[i:], m.TrieHash)
		i = encodeVarintBlock(dAtA, i, uint64(len(m.TrieHash)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Timestamp != 0 {
		i = encodeVarintBlock(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x10
	}
	if m.Height != 0 {
		i = encodeVarintBlock(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintBlock(dAtA []byte, offset int, v uint64) int {
	offset -= sovBlock(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Block) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovBlock(uint64(m.Height))
	}
	if m.Timestamp != 0 {
		n += 1 + sovBlock(uint64(m.Timestamp))
	}
	l = len(m.TrieHash)
	if l > 0 {
		n += 1 + l + sovBlock(uint64(l))
	}
	l = len(m.PrevBlockHash)
	if l > 0 {
		n += 1 + l + sovBlock(uint64(l))
	}
	l = len(m.PrevVaultHash)
	if l > 0 {
		n += 1 + l + sovBlock(uint64(l))
	}
	l = len(m.Beneficiary)
	if l > 0 {
		n += 1 + l + sovBlock(uint64(l))
	}
	l = len(m.Difficulty)
	if l > 0 {
		n += 1 + l + sovBlock(uint64(l))
	}
	l = len(m.Nonce)
	if l > 0 {
		n += 1 + l + sovBlock(uint64(l))
	}
	if len(m.Operations) > 0 {
		for _, e := range m.Operations {
			l = e.Size()
			n += 1 + l + sovBlock(uint64(l))
		}
	}
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovBlock(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovBlock(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBlock(x uint64) (n int) {
	return sovBlock(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Block) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlock
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
			return fmt.Errorf("proto: Block: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Block: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
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
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TrieHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
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
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TrieHash = append(m.TrieHash[:0], dAtA[iNdEx:postIndex]...)
			if m.TrieHash == nil {
				m.TrieHash = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrevBlockHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
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
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PrevBlockHash = append(m.PrevBlockHash[:0], dAtA[iNdEx:postIndex]...)
			if m.PrevBlockHash == nil {
				m.PrevBlockHash = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrevVaultHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
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
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PrevVaultHash = append(m.PrevVaultHash[:0], dAtA[iNdEx:postIndex]...)
			if m.PrevVaultHash == nil {
				m.PrevVaultHash = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Beneficiary", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
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
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Beneficiary = append(m.Beneficiary[:0], dAtA[iNdEx:postIndex]...)
			if m.Beneficiary == nil {
				m.Beneficiary = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Difficulty", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
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
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Difficulty = append(m.Difficulty[:0], dAtA[iNdEx:postIndex]...)
			if m.Difficulty == nil {
				m.Difficulty = []byte{}
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
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
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nonce = append(m.Nonce[:0], dAtA[iNdEx:postIndex]...)
			if m.Nonce == nil {
				m.Nonce = []byte{}
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Operations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
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
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Operations = append(m.Operations, &Operation{})
			if err := m.Operations[len(m.Operations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
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
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = append(m.Hash[:0], dAtA[iNdEx:postIndex]...)
			if m.Hash == nil {
				m.Hash = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBlock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBlock
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBlock
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
func skipBlock(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBlock
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
					return 0, ErrIntOverflowBlock
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
					return 0, ErrIntOverflowBlock
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
				return 0, ErrInvalidLengthBlock
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBlock
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBlock
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBlock        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBlock          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBlock = fmt.Errorf("proto: unexpected end of group")
)
