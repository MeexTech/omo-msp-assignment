// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/assignment/meeting.proto

package assignment

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

type MeetingInfo struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Id                   uint64   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Created              int64    `protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"`
	Updated              int64    `protobuf:"varint,4,opt,name=updated,proto3" json:"updated,omitempty"`
	Creator              string   `protobuf:"bytes,5,opt,name=creator,proto3" json:"creator,omitempty"`
	Operator             string   `protobuf:"bytes,6,opt,name=operator,proto3" json:"operator,omitempty"`
	Status               uint32   `protobuf:"varint,7,opt,name=status,proto3" json:"status,omitempty"`
	Type                 uint32   `protobuf:"varint,8,opt,name=type,proto3" json:"type,omitempty"`
	Stopped              int64    `protobuf:"varint,9,opt,name=stopped,proto3" json:"stopped,omitempty"`
	Started              int64    `protobuf:"varint,10,opt,name=started,proto3" json:"started,omitempty"`
	Owner                string   `protobuf:"bytes,11,opt,name=owner,proto3" json:"owner,omitempty"`
	Group                string   `protobuf:"bytes,12,opt,name=group,proto3" json:"group,omitempty"`
	Location             string   `protobuf:"bytes,13,opt,name=location,proto3" json:"location,omitempty"`
	Duration             uint32   `protobuf:"varint,14,opt,name=duration,proto3" json:"duration,omitempty"`
	Appointed            string   `protobuf:"bytes,15,opt,name=appointed,proto3" json:"appointed,omitempty"`
	Name                 string   `protobuf:"bytes,16,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string   `protobuf:"bytes,17,opt,name=remark,proto3" json:"remark,omitempty"`
	Signs                []string `protobuf:"bytes,20,rep,name=signs,proto3" json:"signs,omitempty"`
	Submits              []string `protobuf:"bytes,21,rep,name=submits,proto3" json:"submits,omitempty"`
	Notifies             []string `protobuf:"bytes,22,rep,name=notifies,proto3" json:"notifies,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MeetingInfo) Reset()         { *m = MeetingInfo{} }
func (m *MeetingInfo) String() string { return proto.CompactTextString(m) }
func (*MeetingInfo) ProtoMessage()    {}
func (*MeetingInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_42ac28e9a2c56839, []int{0}
}

func (m *MeetingInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeetingInfo.Unmarshal(m, b)
}
func (m *MeetingInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeetingInfo.Marshal(b, m, deterministic)
}
func (m *MeetingInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeetingInfo.Merge(m, src)
}
func (m *MeetingInfo) XXX_Size() int {
	return xxx_messageInfo_MeetingInfo.Size(m)
}
func (m *MeetingInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MeetingInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MeetingInfo proto.InternalMessageInfo

func (m *MeetingInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *MeetingInfo) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *MeetingInfo) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *MeetingInfo) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *MeetingInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MeetingInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *MeetingInfo) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *MeetingInfo) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *MeetingInfo) GetStopped() int64 {
	if m != nil {
		return m.Stopped
	}
	return 0
}

func (m *MeetingInfo) GetStarted() int64 {
	if m != nil {
		return m.Started
	}
	return 0
}

func (m *MeetingInfo) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *MeetingInfo) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *MeetingInfo) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *MeetingInfo) GetDuration() uint32 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *MeetingInfo) GetAppointed() string {
	if m != nil {
		return m.Appointed
	}
	return ""
}

func (m *MeetingInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MeetingInfo) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *MeetingInfo) GetSigns() []string {
	if m != nil {
		return m.Signs
	}
	return nil
}

func (m *MeetingInfo) GetSubmits() []string {
	if m != nil {
		return m.Submits
	}
	return nil
}

func (m *MeetingInfo) GetNotifies() []string {
	if m != nil {
		return m.Notifies
	}
	return nil
}

type ReqMeetingAdd struct {
	Type                 uint32   `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Owner                string   `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Group                string   `protobuf:"bytes,3,opt,name=group,proto3" json:"group,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string   `protobuf:"bytes,5,opt,name=remark,proto3" json:"remark,omitempty"`
	Appointed            string   `protobuf:"bytes,6,opt,name=appointed,proto3" json:"appointed,omitempty"`
	Location             string   `protobuf:"bytes,7,opt,name=location,proto3" json:"location,omitempty"`
	Operator             string   `protobuf:"bytes,8,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqMeetingAdd) Reset()         { *m = ReqMeetingAdd{} }
func (m *ReqMeetingAdd) String() string { return proto.CompactTextString(m) }
func (*ReqMeetingAdd) ProtoMessage()    {}
func (*ReqMeetingAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_42ac28e9a2c56839, []int{1}
}

func (m *ReqMeetingAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqMeetingAdd.Unmarshal(m, b)
}
func (m *ReqMeetingAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqMeetingAdd.Marshal(b, m, deterministic)
}
func (m *ReqMeetingAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqMeetingAdd.Merge(m, src)
}
func (m *ReqMeetingAdd) XXX_Size() int {
	return xxx_messageInfo_ReqMeetingAdd.Size(m)
}
func (m *ReqMeetingAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqMeetingAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ReqMeetingAdd proto.InternalMessageInfo

func (m *ReqMeetingAdd) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ReqMeetingAdd) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ReqMeetingAdd) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *ReqMeetingAdd) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqMeetingAdd) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqMeetingAdd) GetAppointed() string {
	if m != nil {
		return m.Appointed
	}
	return ""
}

func (m *ReqMeetingAdd) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *ReqMeetingAdd) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

type ReqMeetingUpdate struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string   `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	Operator             string   `protobuf:"bytes,4,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqMeetingUpdate) Reset()         { *m = ReqMeetingUpdate{} }
func (m *ReqMeetingUpdate) String() string { return proto.CompactTextString(m) }
func (*ReqMeetingUpdate) ProtoMessage()    {}
func (*ReqMeetingUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_42ac28e9a2c56839, []int{2}
}

func (m *ReqMeetingUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqMeetingUpdate.Unmarshal(m, b)
}
func (m *ReqMeetingUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqMeetingUpdate.Marshal(b, m, deterministic)
}
func (m *ReqMeetingUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqMeetingUpdate.Merge(m, src)
}
func (m *ReqMeetingUpdate) XXX_Size() int {
	return xxx_messageInfo_ReqMeetingUpdate.Size(m)
}
func (m *ReqMeetingUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqMeetingUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_ReqMeetingUpdate proto.InternalMessageInfo

func (m *ReqMeetingUpdate) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqMeetingUpdate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqMeetingUpdate) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqMeetingUpdate) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

type ReqMeetingSign struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Member               string   `protobuf:"bytes,2,opt,name=member,proto3" json:"member,omitempty"`
	Location             string   `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	Operator             string   `protobuf:"bytes,4,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqMeetingSign) Reset()         { *m = ReqMeetingSign{} }
func (m *ReqMeetingSign) String() string { return proto.CompactTextString(m) }
func (*ReqMeetingSign) ProtoMessage()    {}
func (*ReqMeetingSign) Descriptor() ([]byte, []int) {
	return fileDescriptor_42ac28e9a2c56839, []int{3}
}

func (m *ReqMeetingSign) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqMeetingSign.Unmarshal(m, b)
}
func (m *ReqMeetingSign) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqMeetingSign.Marshal(b, m, deterministic)
}
func (m *ReqMeetingSign) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqMeetingSign.Merge(m, src)
}
func (m *ReqMeetingSign) XXX_Size() int {
	return xxx_messageInfo_ReqMeetingSign.Size(m)
}
func (m *ReqMeetingSign) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqMeetingSign.DiscardUnknown(m)
}

var xxx_messageInfo_ReqMeetingSign proto.InternalMessageInfo

func (m *ReqMeetingSign) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqMeetingSign) GetMember() string {
	if m != nil {
		return m.Member
	}
	return ""
}

func (m *ReqMeetingSign) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *ReqMeetingSign) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

type ReplyMeetingOne struct {
	Info                 *MeetingInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Status               *ReplyStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyMeetingOne) Reset()         { *m = ReplyMeetingOne{} }
func (m *ReplyMeetingOne) String() string { return proto.CompactTextString(m) }
func (*ReplyMeetingOne) ProtoMessage()    {}
func (*ReplyMeetingOne) Descriptor() ([]byte, []int) {
	return fileDescriptor_42ac28e9a2c56839, []int{4}
}

func (m *ReplyMeetingOne) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyMeetingOne.Unmarshal(m, b)
}
func (m *ReplyMeetingOne) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyMeetingOne.Marshal(b, m, deterministic)
}
func (m *ReplyMeetingOne) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyMeetingOne.Merge(m, src)
}
func (m *ReplyMeetingOne) XXX_Size() int {
	return xxx_messageInfo_ReplyMeetingOne.Size(m)
}
func (m *ReplyMeetingOne) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyMeetingOne.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyMeetingOne proto.InternalMessageInfo

func (m *ReplyMeetingOne) GetInfo() *MeetingInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *ReplyMeetingOne) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type ReplyMeetingList struct {
	Total                uint32         `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	PageMax              uint32         `protobuf:"varint,2,opt,name=pageMax,proto3" json:"pageMax,omitempty"`
	PageNow              uint32         `protobuf:"varint,3,opt,name=pageNow,proto3" json:"pageNow,omitempty"`
	List                 []*MeetingInfo `protobuf:"bytes,4,rep,name=list,proto3" json:"list,omitempty"`
	Status               *ReplyStatus   `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ReplyMeetingList) Reset()         { *m = ReplyMeetingList{} }
func (m *ReplyMeetingList) String() string { return proto.CompactTextString(m) }
func (*ReplyMeetingList) ProtoMessage()    {}
func (*ReplyMeetingList) Descriptor() ([]byte, []int) {
	return fileDescriptor_42ac28e9a2c56839, []int{5}
}

func (m *ReplyMeetingList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyMeetingList.Unmarshal(m, b)
}
func (m *ReplyMeetingList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyMeetingList.Marshal(b, m, deterministic)
}
func (m *ReplyMeetingList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyMeetingList.Merge(m, src)
}
func (m *ReplyMeetingList) XXX_Size() int {
	return xxx_messageInfo_ReplyMeetingList.Size(m)
}
func (m *ReplyMeetingList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyMeetingList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyMeetingList proto.InternalMessageInfo

func (m *ReplyMeetingList) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ReplyMeetingList) GetPageMax() uint32 {
	if m != nil {
		return m.PageMax
	}
	return 0
}

func (m *ReplyMeetingList) GetPageNow() uint32 {
	if m != nil {
		return m.PageNow
	}
	return 0
}

func (m *ReplyMeetingList) GetList() []*MeetingInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *ReplyMeetingList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*MeetingInfo)(nil), "omo.msp.assignment.MeetingInfo")
	proto.RegisterType((*ReqMeetingAdd)(nil), "omo.msp.assignment.ReqMeetingAdd")
	proto.RegisterType((*ReqMeetingUpdate)(nil), "omo.msp.assignment.ReqMeetingUpdate")
	proto.RegisterType((*ReqMeetingSign)(nil), "omo.msp.assignment.ReqMeetingSign")
	proto.RegisterType((*ReplyMeetingOne)(nil), "omo.msp.assignment.ReplyMeetingOne")
	proto.RegisterType((*ReplyMeetingList)(nil), "omo.msp.assignment.ReplyMeetingList")
}

func init() {
	proto.RegisterFile("proto/assignment/meeting.proto", fileDescriptor_42ac28e9a2c56839)
}

var fileDescriptor_42ac28e9a2c56839 = []byte{
	// 739 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x56, 0x41, 0x4f, 0xdb, 0x4c,
	0x10, 0x8d, 0x63, 0xc7, 0x90, 0x85, 0x84, 0x7c, 0x2b, 0x3e, 0xb4, 0x8a, 0x4a, 0x9b, 0xba, 0x1c,
	0x72, 0x0a, 0x12, 0x1c, 0x7a, 0x86, 0x03, 0xa8, 0x6a, 0xa1, 0x95, 0x69, 0x55, 0xa9, 0xea, 0xc5,
	0xc4, 0x4b, 0xb4, 0x6a, 0xec, 0x35, 0xde, 0x35, 0x94, 0x53, 0xcf, 0xfd, 0x6b, 0xbd, 0xf5, 0xb7,
	0xf4, 0x0f, 0x54, 0x33, 0x6b, 0x3b, 0x4e, 0x48, 0x82, 0x7b, 0xf3, 0xdb, 0xb7, 0x3b, 0xf3, 0xde,
	0xdb, 0xd9, 0x28, 0xe4, 0x79, 0x92, 0x4a, 0x2d, 0x0f, 0x03, 0xa5, 0xc4, 0x24, 0x8e, 0x78, 0xac,
	0x0f, 0x23, 0xce, 0xb5, 0x88, 0x27, 0x23, 0x24, 0x28, 0x95, 0x91, 0x1c, 0x45, 0x2a, 0x19, 0xcd,
	0x76, 0xf4, 0xf7, 0x1f, 0x9d, 0x19, 0xcb, 0x28, 0x92, 0xb1, 0x39, 0xe2, 0xfd, 0xb1, 0xc9, 0xd6,
	0x85, 0x29, 0xf2, 0x26, 0xbe, 0x91, 0xb4, 0x47, 0xec, 0x4c, 0x84, 0xcc, 0x1a, 0x58, 0xc3, 0xb6,
	0x0f, 0x9f, 0xb4, 0x4b, 0x9a, 0x22, 0x64, 0xcd, 0x81, 0x35, 0x74, 0xfc, 0xa6, 0x08, 0x29, 0x23,
	0x1b, 0xe3, 0x94, 0x07, 0x9a, 0x87, 0xcc, 0x1e, 0x58, 0x43, 0xdb, 0x2f, 0x20, 0x30, 0x59, 0x12,
	0x22, 0xe3, 0x18, 0x26, 0x87, 0xe5, 0x19, 0x99, 0xb2, 0x16, 0x56, 0x2e, 0x20, 0xed, 0x93, 0x4d,
	0x99, 0xf0, 0x14, 0x29, 0x17, 0xa9, 0x12, 0xd3, 0x3d, 0xe2, 0x2a, 0x1d, 0xe8, 0x4c, 0xb1, 0x8d,
	0x81, 0x35, 0xec, 0xf8, 0x39, 0xa2, 0x94, 0x38, 0xfa, 0x21, 0xe1, 0x6c, 0x13, 0x57, 0xf1, 0x1b,
	0x3a, 0x28, 0x2d, 0x93, 0x84, 0x87, 0xac, 0x6d, 0x7a, 0xe7, 0xd0, 0x30, 0x41, 0x0a, 0xaa, 0x48,
	0xc1, 0x20, 0xa4, 0xbb, 0xa4, 0x25, 0xef, 0x63, 0x9e, 0xb2, 0x2d, 0x6c, 0x6c, 0x00, 0xac, 0x4e,
	0x52, 0x99, 0x25, 0x6c, 0xdb, 0xac, 0x22, 0x00, 0x9d, 0x53, 0x39, 0x0e, 0xb4, 0x90, 0x31, 0xeb,
	0x18, 0x9d, 0x05, 0x06, 0x2e, 0xcc, 0x52, 0xc3, 0x75, 0x51, 0x53, 0x89, 0xe9, 0x33, 0xd2, 0x0e,
	0x92, 0x44, 0x8a, 0x18, 0xfa, 0xef, 0xe0, 0xc1, 0xd9, 0x02, 0x38, 0x89, 0x83, 0x88, 0xb3, 0x1e,
	0x12, 0xf8, 0x0d, 0xae, 0x53, 0x1e, 0x05, 0xe9, 0x37, 0xf6, 0x1f, 0xae, 0xe6, 0x08, 0x74, 0xc1,
	0x15, 0x2a, 0xb6, 0x3b, 0xb0, 0x41, 0x17, 0x02, 0x74, 0x97, 0x5d, 0x47, 0x42, 0x2b, 0xf6, 0x3f,
	0xae, 0x17, 0x10, 0x54, 0xc5, 0x52, 0x8b, 0x1b, 0xc1, 0x15, 0xdb, 0x43, 0xaa, 0xc4, 0xde, 0x6f,
	0x8b, 0x74, 0x7c, 0x7e, 0x9b, 0x5f, 0xfc, 0x49, 0x18, 0x96, 0x99, 0x5a, 0x95, 0x4c, 0xcb, 0x7c,
	0x9a, 0x4b, 0xf3, 0xb1, 0xab, 0xf9, 0x14, 0x4e, 0x9c, 0xa5, 0x4e, 0x5a, 0x73, 0x4e, 0xe6, 0x32,
	0x71, 0x17, 0x33, 0xa9, 0x26, 0xbd, 0xf1, 0x38, 0xe9, 0x72, 0x5a, 0x36, 0xe7, 0xa7, 0xc5, 0x9b,
	0x92, 0xde, 0xcc, 0xd2, 0x27, 0x1c, 0xbc, 0x25, 0xd3, 0x5c, 0xe8, 0x6c, 0x2e, 0xd5, 0x69, 0xcf,
	0xe9, 0xac, 0x76, 0x73, 0x16, 0xba, 0xa5, 0xa4, 0x3b, 0xeb, 0x76, 0x25, 0x26, 0xf1, 0x92, 0x5e,
	0x7b, 0xc4, 0x8d, 0x78, 0x74, 0x5d, 0x06, 0x98, 0xa3, 0x39, 0x87, 0xf6, 0x1a, 0x87, 0x8b, 0x3d,
	0x7f, 0x90, 0x1d, 0x9f, 0x27, 0xd3, 0x87, 0xbc, 0xeb, 0xfb, 0x98, 0xd3, 0x63, 0xe2, 0x88, 0xf8,
	0x46, 0x62, 0xd7, 0xad, 0xa3, 0x17, 0xa3, 0xc7, 0x3f, 0x00, 0xa3, 0xca, 0xeb, 0xf6, 0x71, 0x33,
	0x7d, 0x5d, 0xbe, 0xab, 0xe6, 0xea, 0x63, 0xd8, 0xe9, 0x0a, 0xb7, 0x15, 0x0f, 0xcf, 0xfb, 0x65,
	0x41, 0xc6, 0x33, 0x05, 0xef, 0x84, 0xd2, 0x30, 0x0f, 0x5a, 0xea, 0x60, 0x9a, 0x8f, 0x8e, 0x01,
	0x30, 0x97, 0x49, 0x30, 0xe1, 0x17, 0xc1, 0x77, 0x6c, 0xd2, 0xf1, 0x0b, 0x58, 0x30, 0x97, 0xf2,
	0x1e, 0xcd, 0xe7, 0xcc, 0xa5, 0xbc, 0x07, 0x33, 0x53, 0xa1, 0x34, 0x73, 0x06, 0x76, 0x2d, 0x33,
	0xb0, 0xb9, 0x62, 0xa6, 0xf5, 0x4f, 0x66, 0x8e, 0x7e, 0xba, 0xa4, 0x5b, 0xdc, 0x1f, 0x4f, 0xef,
	0xc4, 0x98, 0x53, 0x9f, 0xb8, 0x27, 0x61, 0x08, 0xb9, 0xbe, 0x5c, 0x5e, 0xa5, 0xf2, 0x62, 0xfa,
	0xaf, 0x56, 0x36, 0x9a, 0xdd, 0x8f, 0xd7, 0xa0, 0x1f, 0x88, 0x7b, 0xce, 0x35, 0xd4, 0x5c, 0xa1,
	0xec, 0x36, 0xe3, 0x4a, 0x83, 0xa1, 0xba, 0x15, 0x2f, 0x48, 0xdb, 0xe7, 0x91, 0xbc, 0xe3, 0xb5,
	0x8a, 0xee, 0xaf, 0x2c, 0x0a, 0xb4, 0xd7, 0xa0, 0x5f, 0xc9, 0xce, 0x39, 0xd7, 0x70, 0x95, 0xa7,
	0x0f, 0x67, 0x62, 0xaa, 0x79, 0xba, 0xd2, 0x3d, 0x14, 0x35, 0x5b, 0xfa, 0x07, 0x4f, 0x69, 0x85,
	0x82, 0x5e, 0x83, 0x7e, 0x26, 0xdb, 0xe7, 0x5c, 0x43, 0xf4, 0x42, 0x69, 0x31, 0xae, 0x53, 0xda,
	0x5b, 0x7b, 0x83, 0x58, 0xc6, 0x6b, 0xd0, 0x8f, 0xa4, 0x6b, 0x1e, 0x79, 0x2d, 0xd5, 0x66, 0xeb,
	0xd3, 0x61, 0x5c, 0x91, 0x6d, 0xb3, 0xd5, 0x0c, 0x0b, 0xf5, 0xd6, 0xc6, 0xab, 0xcf, 0xa6, 0xc1,
	0xa4, 0x4e, 0x51, 0x92, 0x4b, 0x0d, 0x14, 0xa7, 0x07, 0xeb, 0x47, 0xab, 0xae, 0xd2, 0xb7, 0xc4,
	0xc1, 0x9f, 0x1d, 0x6f, 0x7d, 0x39, 0xd8, 0xf3, 0x64, 0xb1, 0x53, 0xfa, 0xa5, 0xb7, 0xf8, 0x37,
	0xe1, 0xda, 0xc5, 0x95, 0xe3, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe4, 0x1e, 0x68, 0x7d, 0x75,
	0x08, 0x00, 0x00,
}