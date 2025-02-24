// Code generated by protoc-gen-go. DO NOT EDIT.
// source: orc8r/protos/digest.proto

package protos

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

// DigestTree contains the full set of digest information for a particular network.
// DigestTree is similar to a depth=2 Merkle tree.
type DigestTree struct {
	// root_digest is the amalgum of all leaf digests.
	RootDigest *Digest `protobuf:"bytes,1,opt,name=root_digest,json=rootDigest,proto3" json:"root_digest,omitempty"`
	// leaf_digests contains per-object digests, along with the object IDs, sorted by ID.
	LeafDigests          []*LeafDigest `protobuf:"bytes,2,rep,name=leaf_digests,json=leafDigests,proto3" json:"leaf_digests,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *DigestTree) Reset()         { *m = DigestTree{} }
func (m *DigestTree) String() string { return proto.CompactTextString(m) }
func (*DigestTree) ProtoMessage()    {}
func (*DigestTree) Descriptor() ([]byte, []int) {
	return fileDescriptor_01bfdb0cdb6f10f9, []int{0}
}

func (m *DigestTree) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DigestTree.Unmarshal(m, b)
}
func (m *DigestTree) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DigestTree.Marshal(b, m, deterministic)
}
func (m *DigestTree) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DigestTree.Merge(m, src)
}
func (m *DigestTree) XXX_Size() int {
	return xxx_messageInfo_DigestTree.Size(m)
}
func (m *DigestTree) XXX_DiscardUnknown() {
	xxx_messageInfo_DigestTree.DiscardUnknown(m)
}

var xxx_messageInfo_DigestTree proto.InternalMessageInfo

func (m *DigestTree) GetRootDigest() *Digest {
	if m != nil {
		return m.RootDigest
	}
	return nil
}

func (m *DigestTree) GetLeafDigests() []*LeafDigest {
	if m != nil {
		return m.LeafDigests
	}
	return nil
}

// Digest contains the digest (hash) of some object.
type Digest struct {
	// md5_base64_digest is a base64-encoded MD5 digest.
	Md5Base64Digest      string   `protobuf:"bytes,1,opt,name=md5_base64_digest,json=md5Base64Digest,proto3" json:"md5_base64_digest,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Digest) Reset()         { *m = Digest{} }
func (m *Digest) String() string { return proto.CompactTextString(m) }
func (*Digest) ProtoMessage()    {}
func (*Digest) Descriptor() ([]byte, []int) {
	return fileDescriptor_01bfdb0cdb6f10f9, []int{1}
}

func (m *Digest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Digest.Unmarshal(m, b)
}
func (m *Digest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Digest.Marshal(b, m, deterministic)
}
func (m *Digest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Digest.Merge(m, src)
}
func (m *Digest) XXX_Size() int {
	return xxx_messageInfo_Digest.Size(m)
}
func (m *Digest) XXX_DiscardUnknown() {
	xxx_messageInfo_Digest.DiscardUnknown(m)
}

var xxx_messageInfo_Digest proto.InternalMessageInfo

func (m *Digest) GetMd5Base64Digest() string {
	if m != nil {
		return m.Md5Base64Digest
	}
	return ""
}

type LeafDigest struct {
	// id is the network-wide unique identifier of the object.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// digest is the deterministic digest of the object.
	Digest               *Digest  `protobuf:"bytes,2,opt,name=digest,proto3" json:"digest,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LeafDigest) Reset()         { *m = LeafDigest{} }
func (m *LeafDigest) String() string { return proto.CompactTextString(m) }
func (*LeafDigest) ProtoMessage()    {}
func (*LeafDigest) Descriptor() ([]byte, []int) {
	return fileDescriptor_01bfdb0cdb6f10f9, []int{2}
}

func (m *LeafDigest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LeafDigest.Unmarshal(m, b)
}
func (m *LeafDigest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LeafDigest.Marshal(b, m, deterministic)
}
func (m *LeafDigest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeafDigest.Merge(m, src)
}
func (m *LeafDigest) XXX_Size() int {
	return xxx_messageInfo_LeafDigest.Size(m)
}
func (m *LeafDigest) XXX_DiscardUnknown() {
	xxx_messageInfo_LeafDigest.DiscardUnknown(m)
}

var xxx_messageInfo_LeafDigest proto.InternalMessageInfo

func (m *LeafDigest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *LeafDigest) GetDigest() *Digest {
	if m != nil {
		return m.Digest
	}
	return nil
}

// LeafDigests is used to encapsulate a list of leaf digests exclusively for
// serialization en masse.
// NOTE: In a proto message used by gRPC endpoints (e.g. DigestTree), the leaf
// digests should still be represented directly as a list for simplicity.
type LeafDigests struct {
	Digests              []*LeafDigest `protobuf:"bytes,1,rep,name=digests,proto3" json:"digests,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *LeafDigests) Reset()         { *m = LeafDigests{} }
func (m *LeafDigests) String() string { return proto.CompactTextString(m) }
func (*LeafDigests) ProtoMessage()    {}
func (*LeafDigests) Descriptor() ([]byte, []int) {
	return fileDescriptor_01bfdb0cdb6f10f9, []int{3}
}

func (m *LeafDigests) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LeafDigests.Unmarshal(m, b)
}
func (m *LeafDigests) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LeafDigests.Marshal(b, m, deterministic)
}
func (m *LeafDigests) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeafDigests.Merge(m, src)
}
func (m *LeafDigests) XXX_Size() int {
	return xxx_messageInfo_LeafDigests.Size(m)
}
func (m *LeafDigests) XXX_DiscardUnknown() {
	xxx_messageInfo_LeafDigests.DiscardUnknown(m)
}

var xxx_messageInfo_LeafDigests proto.InternalMessageInfo

func (m *LeafDigests) GetDigests() []*LeafDigest {
	if m != nil {
		return m.Digests
	}
	return nil
}

func init() {
	proto.RegisterType((*DigestTree)(nil), "magma.orc8r.DigestTree")
	proto.RegisterType((*Digest)(nil), "magma.orc8r.Digest")
	proto.RegisterType((*LeafDigest)(nil), "magma.orc8r.LeafDigest")
	proto.RegisterType((*LeafDigests)(nil), "magma.orc8r.LeafDigests")
}

func init() { proto.RegisterFile("orc8r/protos/digest.proto", fileDescriptor_01bfdb0cdb6f10f9) }

var fileDescriptor_01bfdb0cdb6f10f9 = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcc, 0x2f, 0x4a, 0xb6,
	0x28, 0xd2, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x2f, 0xd6, 0x4f, 0xc9, 0x4c, 0x4f, 0x2d, 0x2e, 0xd1,
	0x03, 0xf3, 0x84, 0xb8, 0x73, 0x13, 0xd3, 0x73, 0x13, 0xf5, 0xc0, 0x0a, 0x94, 0xea, 0xb8, 0xb8,
	0x5c, 0xc0, 0x92, 0x21, 0x45, 0xa9, 0xa9, 0x42, 0x26, 0x5c, 0xdc, 0x45, 0xf9, 0xf9, 0x25, 0xf1,
	0x10, 0xf5, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0xc2, 0x7a, 0x48, 0x1a, 0xf4, 0x20, 0xaa,
	0x83, 0xb8, 0x40, 0xea, 0x20, 0x6c, 0x21, 0x2b, 0x2e, 0x9e, 0x9c, 0xd4, 0xc4, 0x34, 0xa8, 0xae,
	0x62, 0x09, 0x26, 0x05, 0x66, 0x0d, 0x6e, 0x23, 0x71, 0x14, 0x6d, 0x3e, 0xa9, 0x89, 0x69, 0x50,
	0xad, 0xdc, 0x39, 0x70, 0x76, 0xb1, 0x92, 0x09, 0x17, 0x1b, 0xd4, 0x14, 0x2d, 0x2e, 0xc1, 0xdc,
	0x14, 0xd3, 0xf8, 0xa4, 0xc4, 0xe2, 0x54, 0x33, 0x13, 0x64, 0x17, 0x70, 0x06, 0xf1, 0xe7, 0xa6,
	0x98, 0x3a, 0x81, 0xc5, 0x21, 0x6a, 0x95, 0x3c, 0xb9, 0xb8, 0x10, 0x06, 0x0a, 0xf1, 0x71, 0x31,
	0x65, 0xa6, 0x40, 0x95, 0x32, 0x65, 0xa6, 0x08, 0x69, 0x73, 0xb1, 0x41, 0xb5, 0x33, 0xe1, 0xf6,
	0x00, 0x54, 0x89, 0x92, 0x03, 0x17, 0x37, 0xc2, 0xa8, 0x62, 0x21, 0x43, 0x2e, 0x76, 0x98, 0x37,
	0x18, 0xf1, 0x7b, 0x03, 0xa6, 0xce, 0x49, 0x3a, 0x4a, 0x12, 0xac, 0x44, 0x1f, 0x12, 0xe4, 0x39,
	0x99, 0x49, 0xfa, 0xe9, 0xf9, 0xd0, 0x90, 0x4f, 0x62, 0x03, 0xd3, 0xc6, 0x80, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x4d, 0x64, 0x34, 0x12, 0x90, 0x01, 0x00, 0x00,
}
