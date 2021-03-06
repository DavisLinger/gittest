// Code generated by protoc-gen-go. DO NOT EDIT.
// source: interactivefb.proto

package interactivefb

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

type InteractiveFb struct {
	PlatId               string   `protobuf:"bytes,1,opt,name=plat_id,json=platId,proto3" json:"plat_id,omitempty"`
	MediaId              string   `protobuf:"bytes,2,opt,name=media_id,json=mediaId,proto3" json:"media_id,omitempty"`
	AdSlotId             string   `protobuf:"bytes,3,opt,name=ad_slot_id,json=adSlotId,proto3" json:"ad_slot_id,omitempty"`
	TemplateId           string   `protobuf:"bytes,4,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
	PlanId               string   `protobuf:"bytes,5,opt,name=plan_id,json=planId,proto3" json:"plan_id,omitempty"`
	AwardId              string   `protobuf:"bytes,6,opt,name=award_id,json=awardId,proto3" json:"award_id,omitempty"`
	AdvertiseId          string   `protobuf:"bytes,7,opt,name=advertise_id,json=advertiseId,proto3" json:"advertise_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InteractiveFb) Reset()         { *m = InteractiveFb{} }
func (m *InteractiveFb) String() string { return proto.CompactTextString(m) }
func (*InteractiveFb) ProtoMessage()    {}
func (*InteractiveFb) Descriptor() ([]byte, []int) {
	return fileDescriptor_62a40373f31fa3c1, []int{0}
}

func (m *InteractiveFb) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InteractiveFb.Unmarshal(m, b)
}
func (m *InteractiveFb) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InteractiveFb.Marshal(b, m, deterministic)
}
func (m *InteractiveFb) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InteractiveFb.Merge(m, src)
}
func (m *InteractiveFb) XXX_Size() int {
	return xxx_messageInfo_InteractiveFb.Size(m)
}
func (m *InteractiveFb) XXX_DiscardUnknown() {
	xxx_messageInfo_InteractiveFb.DiscardUnknown(m)
}

var xxx_messageInfo_InteractiveFb proto.InternalMessageInfo

func (m *InteractiveFb) GetPlatId() string {
	if m != nil {
		return m.PlatId
	}
	return ""
}

func (m *InteractiveFb) GetMediaId() string {
	if m != nil {
		return m.MediaId
	}
	return ""
}

func (m *InteractiveFb) GetAdSlotId() string {
	if m != nil {
		return m.AdSlotId
	}
	return ""
}

func (m *InteractiveFb) GetTemplateId() string {
	if m != nil {
		return m.TemplateId
	}
	return ""
}

func (m *InteractiveFb) GetPlanId() string {
	if m != nil {
		return m.PlanId
	}
	return ""
}

func (m *InteractiveFb) GetAwardId() string {
	if m != nil {
		return m.AwardId
	}
	return ""
}

func (m *InteractiveFb) GetAdvertiseId() string {
	if m != nil {
		return m.AdvertiseId
	}
	return ""
}

func init() {
	proto.RegisterType((*InteractiveFb)(nil), "interactive_fb")
}

func init() { proto.RegisterFile("interactivefb.proto", fileDescriptor_62a40373f31fa3c1) }

var fileDescriptor_62a40373f31fa3c1 = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0x3d, 0x0a, 0xc2, 0x40,
	0x10, 0x85, 0x89, 0x3f, 0x49, 0x9c, 0x88, 0xc5, 0x5a, 0xa8, 0x20, 0xf8, 0x53, 0x59, 0xd9, 0x78,
	0x8a, 0x6d, 0xf5, 0x00, 0x61, 0xe2, 0x6c, 0x60, 0x21, 0xc9, 0x86, 0xcd, 0x12, 0xef, 0xea, 0x69,
	0x64, 0x26, 0x1a, 0x2c, 0xdf, 0xfb, 0x98, 0xf7, 0x31, 0xb0, 0xb6, 0x4d, 0x30, 0x1e, 0x9f, 0xc1,
	0xf6, 0xa6, 0x2c, 0xae, 0xad, 0x77, 0xc1, 0x9d, 0xdf, 0x11, 0xac, 0xfe, 0xfa, 0xbc, 0x2c, 0xd4,
	0x06, 0x92, 0xb6, 0xc2, 0x90, 0x5b, 0xda, 0x46, 0xc7, 0xe8, 0xb2, 0xb8, 0xc7, 0x1c, 0x35, 0xa9,
	0x1d, 0xa4, 0xb5, 0x21, 0x8b, 0x4c, 0x26, 0x42, 0x12, 0xc9, 0x9a, 0xd4, 0x1e, 0x00, 0x29, 0xef,
	0x2a, 0x27, 0x67, 0x53, 0x81, 0x29, 0xd2, 0xa3, 0x72, 0x7c, 0x78, 0x80, 0x2c, 0x98, 0x9a, 0x57,
	0x0c, 0xe3, 0x99, 0x60, 0xf8, 0x55, 0x9a, 0xbe, 0xca, 0x86, 0xe1, 0x7c, 0x54, 0x36, 0x83, 0x12,
	0x5f, 0xe8, 0x89, 0x49, 0x3c, 0x28, 0x25, 0x6b, 0x52, 0x27, 0x58, 0x22, 0xf5, 0xc6, 0x07, 0xdb,
	0xc9, 0x6a, 0x22, 0x38, 0x1b, 0x3b, 0x4d, 0x45, 0x2c, 0x3f, 0xde, 0x3e, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x87, 0x6a, 0x2c, 0x64, 0xfa, 0x00, 0x00, 0x00,
}
