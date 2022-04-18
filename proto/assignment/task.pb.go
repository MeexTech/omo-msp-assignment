// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/assignment/task.proto

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

type RecordInfo struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Created              int64    `protobuf:"varint,2,opt,name=created,proto3" json:"created,omitempty"`
	Creator              string   `protobuf:"bytes,3,opt,name=creator,proto3" json:"creator,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Status               uint32   `protobuf:"varint,5,opt,name=status,proto3" json:"status,omitempty"`
	Remark               string   `protobuf:"bytes,6,opt,name=remark,proto3" json:"remark,omitempty"`
	Executor             string   `protobuf:"bytes,7,opt,name=executor,proto3" json:"executor,omitempty"`
	Tags                 []string `protobuf:"bytes,8,rep,name=tags,proto3" json:"tags,omitempty"`
	Assets               []string `protobuf:"bytes,9,rep,name=assets,proto3" json:"assets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecordInfo) Reset()         { *m = RecordInfo{} }
func (m *RecordInfo) String() string { return proto.CompactTextString(m) }
func (*RecordInfo) ProtoMessage()    {}
func (*RecordInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb1a60d5b88542d3, []int{0}
}

func (m *RecordInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecordInfo.Unmarshal(m, b)
}
func (m *RecordInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecordInfo.Marshal(b, m, deterministic)
}
func (m *RecordInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecordInfo.Merge(m, src)
}
func (m *RecordInfo) XXX_Size() int {
	return xxx_messageInfo_RecordInfo.Size(m)
}
func (m *RecordInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RecordInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RecordInfo proto.InternalMessageInfo

func (m *RecordInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *RecordInfo) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *RecordInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *RecordInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RecordInfo) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *RecordInfo) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *RecordInfo) GetExecutor() string {
	if m != nil {
		return m.Executor
	}
	return ""
}

func (m *RecordInfo) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *RecordInfo) GetAssets() []string {
	if m != nil {
		return m.Assets
	}
	return nil
}

type TaskInfo struct {
	Uid                  string        `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Id                   uint64        `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Created              int64         `protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"`
	Updated              int64         `protobuf:"varint,4,opt,name=updated,proto3" json:"updated,omitempty"`
	Creator              string        `protobuf:"bytes,5,opt,name=creator,proto3" json:"creator,omitempty"`
	Operator             string        `protobuf:"bytes,6,opt,name=operator,proto3" json:"operator,omitempty"`
	Target               string        `protobuf:"bytes,7,opt,name=target,proto3" json:"target,omitempty"`
	Name                 string        `protobuf:"bytes,8,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string        `protobuf:"bytes,9,opt,name=remark,proto3" json:"remark,omitempty"`
	Type                 uint32        `protobuf:"varint,10,opt,name=type,proto3" json:"type,omitempty"`
	Status               uint32        `protobuf:"varint,11,opt,name=status,proto3" json:"status,omitempty"`
	Owner                string        `protobuf:"bytes,12,opt,name=owner,proto3" json:"owner,omitempty"`
	Way                  string        `protobuf:"bytes,13,opt,name=way,proto3" json:"way,omitempty"`
	Duration             *DateInfo     `protobuf:"bytes,14,opt,name=duration,proto3" json:"duration,omitempty"`
	Executors            []string      `protobuf:"bytes,15,rep,name=executors,proto3" json:"executors,omitempty"`
	Pretasks             []string      `protobuf:"bytes,16,rep,name=pretasks,proto3" json:"pretasks,omitempty"`
	Tags                 []string      `protobuf:"bytes,17,rep,name=tags,proto3" json:"tags,omitempty"`
	Regions              []string      `protobuf:"bytes,18,rep,name=regions,proto3" json:"regions,omitempty"`
	Assets               []string      `protobuf:"bytes,19,rep,name=assets,proto3" json:"assets,omitempty"`
	Records              []*RecordInfo `protobuf:"bytes,20,rep,name=records,proto3" json:"records,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *TaskInfo) Reset()         { *m = TaskInfo{} }
func (m *TaskInfo) String() string { return proto.CompactTextString(m) }
func (*TaskInfo) ProtoMessage()    {}
func (*TaskInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb1a60d5b88542d3, []int{1}
}

func (m *TaskInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskInfo.Unmarshal(m, b)
}
func (m *TaskInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskInfo.Marshal(b, m, deterministic)
}
func (m *TaskInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskInfo.Merge(m, src)
}
func (m *TaskInfo) XXX_Size() int {
	return xxx_messageInfo_TaskInfo.Size(m)
}
func (m *TaskInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TaskInfo proto.InternalMessageInfo

func (m *TaskInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *TaskInfo) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *TaskInfo) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *TaskInfo) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *TaskInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *TaskInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *TaskInfo) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *TaskInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TaskInfo) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *TaskInfo) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *TaskInfo) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *TaskInfo) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *TaskInfo) GetWay() string {
	if m != nil {
		return m.Way
	}
	return ""
}

func (m *TaskInfo) GetDuration() *DateInfo {
	if m != nil {
		return m.Duration
	}
	return nil
}

func (m *TaskInfo) GetExecutors() []string {
	if m != nil {
		return m.Executors
	}
	return nil
}

func (m *TaskInfo) GetPretasks() []string {
	if m != nil {
		return m.Pretasks
	}
	return nil
}

func (m *TaskInfo) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *TaskInfo) GetRegions() []string {
	if m != nil {
		return m.Regions
	}
	return nil
}

func (m *TaskInfo) GetAssets() []string {
	if m != nil {
		return m.Assets
	}
	return nil
}

func (m *TaskInfo) GetRecords() []*RecordInfo {
	if m != nil {
		return m.Records
	}
	return nil
}

type ReqTaskAdd struct {
	Name                 string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type                 int32     `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	Remark               string    `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	Owner                string    `protobuf:"bytes,4,opt,name=owner,proto3" json:"owner,omitempty"`
	Target               string    `protobuf:"bytes,5,opt,name=target,proto3" json:"target,omitempty"`
	Operator             string    `protobuf:"bytes,6,opt,name=operator,proto3" json:"operator,omitempty"`
	Way                  string    `protobuf:"bytes,7,opt,name=way,proto3" json:"way,omitempty"`
	Duration             *DateInfo `protobuf:"bytes,8,opt,name=duration,proto3" json:"duration,omitempty"`
	Regions              []string  `protobuf:"bytes,9,rep,name=regions,proto3" json:"regions,omitempty"`
	Pretasks             []string  `protobuf:"bytes,10,rep,name=pretasks,proto3" json:"pretasks,omitempty"`
	Assets               []string  `protobuf:"bytes,11,rep,name=assets,proto3" json:"assets,omitempty"`
	Tags                 []string  `protobuf:"bytes,12,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ReqTaskAdd) Reset()         { *m = ReqTaskAdd{} }
func (m *ReqTaskAdd) String() string { return proto.CompactTextString(m) }
func (*ReqTaskAdd) ProtoMessage()    {}
func (*ReqTaskAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb1a60d5b88542d3, []int{2}
}

func (m *ReqTaskAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqTaskAdd.Unmarshal(m, b)
}
func (m *ReqTaskAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqTaskAdd.Marshal(b, m, deterministic)
}
func (m *ReqTaskAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqTaskAdd.Merge(m, src)
}
func (m *ReqTaskAdd) XXX_Size() int {
	return xxx_messageInfo_ReqTaskAdd.Size(m)
}
func (m *ReqTaskAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqTaskAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ReqTaskAdd proto.InternalMessageInfo

func (m *ReqTaskAdd) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqTaskAdd) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ReqTaskAdd) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqTaskAdd) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ReqTaskAdd) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *ReqTaskAdd) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqTaskAdd) GetWay() string {
	if m != nil {
		return m.Way
	}
	return ""
}

func (m *ReqTaskAdd) GetDuration() *DateInfo {
	if m != nil {
		return m.Duration
	}
	return nil
}

func (m *ReqTaskAdd) GetRegions() []string {
	if m != nil {
		return m.Regions
	}
	return nil
}

func (m *ReqTaskAdd) GetPretasks() []string {
	if m != nil {
		return m.Pretasks
	}
	return nil
}

func (m *ReqTaskAdd) GetAssets() []string {
	if m != nil {
		return m.Assets
	}
	return nil
}

func (m *ReqTaskAdd) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type ReplyTaskOne struct {
	Info                 *TaskInfo    `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Status               *ReplyStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyTaskOne) Reset()         { *m = ReplyTaskOne{} }
func (m *ReplyTaskOne) String() string { return proto.CompactTextString(m) }
func (*ReplyTaskOne) ProtoMessage()    {}
func (*ReplyTaskOne) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb1a60d5b88542d3, []int{3}
}

func (m *ReplyTaskOne) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyTaskOne.Unmarshal(m, b)
}
func (m *ReplyTaskOne) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyTaskOne.Marshal(b, m, deterministic)
}
func (m *ReplyTaskOne) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyTaskOne.Merge(m, src)
}
func (m *ReplyTaskOne) XXX_Size() int {
	return xxx_messageInfo_ReplyTaskOne.Size(m)
}
func (m *ReplyTaskOne) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyTaskOne.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyTaskOne proto.InternalMessageInfo

func (m *ReplyTaskOne) GetInfo() *TaskInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *ReplyTaskOne) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type ReplyTaskList struct {
	Total                uint32       `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	PageMax              uint32       `protobuf:"varint,2,opt,name=pageMax,proto3" json:"pageMax,omitempty"`
	PageNow              uint32       `protobuf:"varint,3,opt,name=pageNow,proto3" json:"pageNow,omitempty"`
	List                 []*TaskInfo  `protobuf:"bytes,4,rep,name=list,proto3" json:"list,omitempty"`
	Status               *ReplyStatus `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyTaskList) Reset()         { *m = ReplyTaskList{} }
func (m *ReplyTaskList) String() string { return proto.CompactTextString(m) }
func (*ReplyTaskList) ProtoMessage()    {}
func (*ReplyTaskList) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb1a60d5b88542d3, []int{4}
}

func (m *ReplyTaskList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyTaskList.Unmarshal(m, b)
}
func (m *ReplyTaskList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyTaskList.Marshal(b, m, deterministic)
}
func (m *ReplyTaskList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyTaskList.Merge(m, src)
}
func (m *ReplyTaskList) XXX_Size() int {
	return xxx_messageInfo_ReplyTaskList.Size(m)
}
func (m *ReplyTaskList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyTaskList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyTaskList proto.InternalMessageInfo

func (m *ReplyTaskList) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ReplyTaskList) GetPageMax() uint32 {
	if m != nil {
		return m.PageMax
	}
	return 0
}

func (m *ReplyTaskList) GetPageNow() uint32 {
	if m != nil {
		return m.PageNow
	}
	return 0
}

func (m *ReplyTaskList) GetList() []*TaskInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *ReplyTaskList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type ReqTaskUpdate struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string   `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	Operator             string   `protobuf:"bytes,4,opt,name=operator,proto3" json:"operator,omitempty"`
	Assets               []string `protobuf:"bytes,5,rep,name=assets,proto3" json:"assets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqTaskUpdate) Reset()         { *m = ReqTaskUpdate{} }
func (m *ReqTaskUpdate) String() string { return proto.CompactTextString(m) }
func (*ReqTaskUpdate) ProtoMessage()    {}
func (*ReqTaskUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb1a60d5b88542d3, []int{5}
}

func (m *ReqTaskUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqTaskUpdate.Unmarshal(m, b)
}
func (m *ReqTaskUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqTaskUpdate.Marshal(b, m, deterministic)
}
func (m *ReqTaskUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqTaskUpdate.Merge(m, src)
}
func (m *ReqTaskUpdate) XXX_Size() int {
	return xxx_messageInfo_ReqTaskUpdate.Size(m)
}
func (m *ReqTaskUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqTaskUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_ReqTaskUpdate proto.InternalMessageInfo

func (m *ReqTaskUpdate) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqTaskUpdate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqTaskUpdate) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqTaskUpdate) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqTaskUpdate) GetAssets() []string {
	if m != nil {
		return m.Assets
	}
	return nil
}

type ReqTaskRecord struct {
	Task                 string   `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	Creator              string   `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string   `protobuf:"bytes,4,opt,name=remark,proto3" json:"remark,omitempty"`
	Executor             string   `protobuf:"bytes,5,opt,name=executor,proto3" json:"executor,omitempty"`
	Status               uint32   `protobuf:"varint,6,opt,name=status,proto3" json:"status,omitempty"`
	Tags                 []string `protobuf:"bytes,7,rep,name=tags,proto3" json:"tags,omitempty"`
	Assets               []string `protobuf:"bytes,8,rep,name=assets,proto3" json:"assets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqTaskRecord) Reset()         { *m = ReqTaskRecord{} }
func (m *ReqTaskRecord) String() string { return proto.CompactTextString(m) }
func (*ReqTaskRecord) ProtoMessage()    {}
func (*ReqTaskRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb1a60d5b88542d3, []int{6}
}

func (m *ReqTaskRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqTaskRecord.Unmarshal(m, b)
}
func (m *ReqTaskRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqTaskRecord.Marshal(b, m, deterministic)
}
func (m *ReqTaskRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqTaskRecord.Merge(m, src)
}
func (m *ReqTaskRecord) XXX_Size() int {
	return xxx_messageInfo_ReqTaskRecord.Size(m)
}
func (m *ReqTaskRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqTaskRecord.DiscardUnknown(m)
}

var xxx_messageInfo_ReqTaskRecord proto.InternalMessageInfo

func (m *ReqTaskRecord) GetTask() string {
	if m != nil {
		return m.Task
	}
	return ""
}

func (m *ReqTaskRecord) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *ReqTaskRecord) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqTaskRecord) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqTaskRecord) GetExecutor() string {
	if m != nil {
		return m.Executor
	}
	return ""
}

func (m *ReqTaskRecord) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *ReqTaskRecord) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *ReqTaskRecord) GetAssets() []string {
	if m != nil {
		return m.Assets
	}
	return nil
}

type ReplyTaskRecords struct {
	Status               *ReplyStatus  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Task                 string        `protobuf:"bytes,2,opt,name=task,proto3" json:"task,omitempty"`
	List                 []*RecordInfo `protobuf:"bytes,3,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ReplyTaskRecords) Reset()         { *m = ReplyTaskRecords{} }
func (m *ReplyTaskRecords) String() string { return proto.CompactTextString(m) }
func (*ReplyTaskRecords) ProtoMessage()    {}
func (*ReplyTaskRecords) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb1a60d5b88542d3, []int{7}
}

func (m *ReplyTaskRecords) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyTaskRecords.Unmarshal(m, b)
}
func (m *ReplyTaskRecords) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyTaskRecords.Marshal(b, m, deterministic)
}
func (m *ReplyTaskRecords) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyTaskRecords.Merge(m, src)
}
func (m *ReplyTaskRecords) XXX_Size() int {
	return xxx_messageInfo_ReplyTaskRecords.Size(m)
}
func (m *ReplyTaskRecords) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyTaskRecords.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyTaskRecords proto.InternalMessageInfo

func (m *ReplyTaskRecords) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyTaskRecords) GetTask() string {
	if m != nil {
		return m.Task
	}
	return ""
}

func (m *ReplyTaskRecords) GetList() []*RecordInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*RecordInfo)(nil), "omo.msp.assignment.RecordInfo")
	proto.RegisterType((*TaskInfo)(nil), "omo.msp.assignment.TaskInfo")
	proto.RegisterType((*ReqTaskAdd)(nil), "omo.msp.assignment.ReqTaskAdd")
	proto.RegisterType((*ReplyTaskOne)(nil), "omo.msp.assignment.ReplyTaskOne")
	proto.RegisterType((*ReplyTaskList)(nil), "omo.msp.assignment.ReplyTaskList")
	proto.RegisterType((*ReqTaskUpdate)(nil), "omo.msp.assignment.ReqTaskUpdate")
	proto.RegisterType((*ReqTaskRecord)(nil), "omo.msp.assignment.ReqTaskRecord")
	proto.RegisterType((*ReplyTaskRecords)(nil), "omo.msp.assignment.ReplyTaskRecords")
}

func init() {
	proto.RegisterFile("proto/assignment/task.proto", fileDescriptor_eb1a60d5b88542d3)
}

var fileDescriptor_eb1a60d5b88542d3 = []byte{
	// 926 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x56, 0x4b, 0x8e, 0x1b, 0x37,
	0x10, 0x9d, 0x56, 0xb7, 0x7e, 0x25, 0x69, 0x3c, 0x61, 0x8c, 0x80, 0x50, 0xec, 0x44, 0x6e, 0x64,
	0xa1, 0x95, 0x1c, 0x4c, 0x16, 0xf1, 0x76, 0x8c, 0xc0, 0x83, 0x00, 0xf6, 0x24, 0x68, 0x39, 0x08,
	0x9c, 0x1d, 0xad, 0xa6, 0x95, 0xc6, 0x48, 0xcd, 0x36, 0x49, 0x59, 0xd6, 0x32, 0x67, 0xc8, 0x0d,
	0x72, 0x82, 0x1c, 0xc2, 0x9b, 0x5c, 0x22, 0x67, 0x31, 0x8a, 0x64, 0xff, 0x46, 0xdf, 0xd9, 0x75,
	0x91, 0xc5, 0xa7, 0xaa, 0xf7, 0x5e, 0x15, 0x04, 0x5f, 0x67, 0x52, 0x68, 0xf1, 0x94, 0x29, 0x95,
	0xcc, 0xd3, 0x25, 0x4f, 0xf5, 0x53, 0xcd, 0xd4, 0xed, 0xc4, 0x9c, 0x12, 0x22, 0x96, 0x62, 0xb2,
	0x54, 0xd9, 0xa4, 0xbc, 0x1e, 0x3e, 0xde, 0x7a, 0x30, 0x13, 0xcb, 0xa5, 0x48, 0xed, 0x93, 0xf0,
	0x7f, 0x0f, 0x20, 0xe2, 0x33, 0x21, 0xe3, 0x9f, 0xd3, 0x77, 0x82, 0x5c, 0x80, 0xbf, 0x4a, 0x62,
	0xea, 0x8d, 0xbc, 0x71, 0x37, 0xc2, 0x4f, 0x42, 0xa1, 0x3d, 0x93, 0x9c, 0x69, 0x1e, 0xd3, 0xc6,
	0xc8, 0x1b, 0xfb, 0x51, 0x1e, 0x16, 0x37, 0x42, 0x52, 0xdf, 0xe4, 0xe7, 0x21, 0x21, 0x10, 0xa4,
	0x6c, 0xc9, 0x69, 0x60, 0x8e, 0xcd, 0x37, 0xf9, 0x0a, 0x5a, 0x4a, 0x33, 0xbd, 0x52, 0xb4, 0x39,
	0xf2, 0xc6, 0x83, 0xc8, 0x45, 0x78, 0x2e, 0xf9, 0x92, 0xc9, 0x5b, 0xda, 0x32, 0xd9, 0x2e, 0x22,
	0x43, 0xe8, 0xf0, 0x8f, 0x7c, 0xb6, 0x42, 0xf8, 0xb6, 0xb9, 0x29, 0x62, 0xc4, 0xd7, 0x6c, 0xae,
	0x68, 0x67, 0xe4, 0x23, 0x3e, 0x7e, 0x23, 0x0e, 0x53, 0x8a, 0x6b, 0x45, 0xbb, 0xe6, 0xd4, 0x45,
	0xe1, 0x3f, 0x01, 0x74, 0x5e, 0x33, 0x75, 0xbb, 0xa7, 0xbd, 0x73, 0x68, 0x24, 0xb6, 0xb3, 0x20,
	0x6a, 0xd4, 0xdb, 0xf5, 0xb7, 0xda, 0x5d, 0x65, 0xb1, 0xb9, 0x09, 0xec, 0x8d, 0x0b, 0xab, 0x44,
	0x34, 0xeb, 0x44, 0x0c, 0xa1, 0x23, 0x32, 0x2e, 0xcd, 0x95, 0x6d, 0xaf, 0x88, 0xb1, 0x60, 0xcd,
	0xe4, 0x9c, 0x6b, 0xd7, 0x9e, 0x8b, 0x0a, 0xf2, 0x3a, 0x75, 0xf2, 0x1c, 0x49, 0xdd, 0x1a, 0x49,
	0x48, 0xc4, 0x26, 0xe3, 0x14, 0x0c, 0xa5, 0xe6, 0xbb, 0x42, 0x74, 0xaf, 0x46, 0xf4, 0x43, 0x68,
	0x8a, 0x75, 0xca, 0x25, 0xed, 0x1b, 0x08, 0x1b, 0x20, 0x23, 0x6b, 0xb6, 0xa1, 0x03, 0xcb, 0xc8,
	0x9a, 0x6d, 0xc8, 0x33, 0xe8, 0xc4, 0x2b, 0xc9, 0x74, 0x22, 0x52, 0x7a, 0x3e, 0xf2, 0xc6, 0xbd,
	0xcb, 0x47, 0x93, 0x6d, 0x5f, 0x4d, 0x7e, 0x62, 0x9a, 0x23, 0xa7, 0x51, 0x91, 0x4d, 0x1e, 0x41,
	0x37, 0x97, 0x48, 0xd1, 0x07, 0x46, 0x85, 0xf2, 0x00, 0xb9, 0xc8, 0x24, 0x47, 0xb7, 0x2a, 0x7a,
	0x61, 0x2e, 0x8b, 0xb8, 0x10, 0xf4, 0x8b, 0x8a, 0xa0, 0x14, 0xda, 0x92, 0xcf, 0x13, 0x91, 0x2a,
	0x4a, 0xcc, 0x71, 0x1e, 0x56, 0xa4, 0xfe, 0xb2, 0x2a, 0x35, 0x79, 0x86, 0x2f, 0xd0, 0xca, 0x8a,
	0x3e, 0x1c, 0xf9, 0xe3, 0xde, 0xe5, 0x37, 0xbb, 0x0a, 0x2f, 0xdd, 0x1e, 0xe5, 0xe9, 0xe1, 0xa7,
	0x06, 0x4e, 0xc1, 0x7b, 0xf4, 0xc9, 0x55, 0x1c, 0x17, 0x12, 0x78, 0x15, 0x09, 0x72, 0xaa, 0xd1,
	0x2a, 0xcd, 0x92, 0x6a, 0x27, 0x8b, 0x5f, 0x93, 0xa5, 0xa0, 0x3a, 0xa8, 0x52, 0x5d, 0x0a, 0xde,
	0xac, 0x09, 0x7e, 0xc8, 0x24, 0x4e, 0x9e, 0xf6, 0x6e, 0x79, 0x3a, 0xf7, 0x92, 0xa7, 0x42, 0x68,
	0xb7, 0x4e, 0x68, 0x55, 0x1a, 0xb8, 0x23, 0x4d, 0x49, 0x76, 0xaf, 0x46, 0x76, 0x2e, 0x59, 0xbf,
	0x94, 0x2c, 0xdc, 0x40, 0x3f, 0xe2, 0xd9, 0x62, 0x83, 0x3c, 0xfe, 0x92, 0x72, 0xf2, 0x3d, 0x04,
	0x49, 0xfa, 0x4e, 0x18, 0x1e, 0xf7, 0xd4, 0x99, 0x8f, 0x66, 0x64, 0x32, 0xc9, 0x8f, 0x85, 0x79,
	0x1b, 0xe6, 0xcd, 0xb7, 0xbb, 0x15, 0xcc, 0x16, 0x9b, 0xa9, 0x49, 0xcb, 0xdd, 0x1d, 0x7e, 0xf2,
	0x60, 0x50, 0xfc, 0xf6, 0xcb, 0x44, 0x69, 0x14, 0x41, 0x0b, 0xcd, 0x16, 0xe6, 0xd7, 0x07, 0x91,
	0x0d, 0x90, 0x84, 0x8c, 0xcd, 0xf9, 0x2b, 0xf6, 0xd1, 0xfc, 0xc2, 0x20, 0xca, 0xc3, 0xfc, 0xe6,
	0x46, 0xac, 0x8d, 0x9a, 0xee, 0xe6, 0x46, 0xac, 0xb1, 0x8d, 0x45, 0xa2, 0x34, 0x0d, 0x8c, 0xa9,
	0x8e, 0xb4, 0x81, 0x99, 0x95, 0x36, 0x9a, 0xf7, 0x6b, 0xe3, 0x2f, 0xd3, 0x86, 0x31, 0xe2, 0x6f,
	0x66, 0xbb, 0xec, 0x58, 0x59, 0xb9, 0x3b, 0x1b, 0x3b, 0x17, 0x84, 0x7f, 0x77, 0x8b, 0x16, 0xde,
	0x0a, 0xb6, 0x17, 0x90, 0x53, 0xb6, 0x59, 0xdb, 0x98, 0xff, 0x95, 0x35, 0xd8, 0x59, 0xb1, 0x5a,
	0xab, 0xdb, 0x7c, 0x1e, 0xf0, 0xbb, 0xba, 0xf4, 0x1a, 0xbb, 0xb7, 0xbf, 0xbf, 0xb3, 0xbe, 0x60,
	0xef, 0x96, 0x6f, 0xde, 0xd9, 0xf2, 0xe5, 0x22, 0x6b, 0xd5, 0x16, 0x59, 0xee, 0xbc, 0xf6, 0xce,
	0xed, 0xdf, 0xa9, 0xf5, 0xf2, 0xb7, 0x07, 0x17, 0x85, 0x2d, 0x6c, 0x37, 0xaa, 0xa2, 0x8e, 0x77,
	0x2f, 0x75, 0x0a, 0x1e, 0x1a, 0x15, 0x1e, 0x2e, 0x9d, 0x39, 0xfc, 0x93, 0x36, 0x8e, 0xc9, 0xbd,
	0xfc, 0xb7, 0x0d, 0x3d, 0x2c, 0x68, 0xca, 0xe5, 0x87, 0x64, 0xc6, 0xc9, 0x4b, 0x68, 0x5d, 0xc5,
	0x31, 0x4e, 0xcc, 0x9e, 0xf7, 0xf9, 0x66, 0x1a, 0x8e, 0xf6, 0x96, 0xea, 0x66, 0x2e, 0x3c, 0x23,
	0xaf, 0xa0, 0x75, 0xcd, 0x35, 0xa2, 0xed, 0x69, 0xec, 0xfd, 0x8a, 0x2b, 0x8d, 0xe5, 0x9c, 0x08,
	0xd7, 0x8d, 0xf8, 0x52, 0x7c, 0xe0, 0x27, 0x21, 0x3e, 0xde, 0x8b, 0x88, 0xd7, 0xe1, 0x19, 0xb9,
	0x81, 0xd6, 0x94, 0x33, 0x39, 0xfb, 0xf3, 0x38, 0xd6, 0x93, 0x83, 0xd5, 0xe1, 0x90, 0x87, 0x67,
	0xe4, 0x0d, 0x3c, 0xb8, 0xe6, 0x1a, 0x83, 0xe7, 0x9b, 0x17, 0xc9, 0x42, 0x73, 0x49, 0x9e, 0x1c,
	0x00, 0xb6, 0x29, 0xa7, 0x41, 0xff, 0x0e, 0xfd, 0x6b, 0xae, 0xd1, 0x03, 0x89, 0xd2, 0xc9, 0xec,
	0x14, 0xdc, 0xf0, 0xa0, 0x95, 0x0c, 0x4c, 0x78, 0x46, 0x5e, 0xc3, 0xb9, 0x9d, 0xee, 0x93, 0x4a,
	0xb6, 0xa9, 0xc7, 0x99, 0xfd, 0x15, 0xc0, 0xa1, 0x32, 0xc5, 0xf7, 0x22, 0x96, 0xab, 0xe5, 0x38,
	0xe2, 0x14, 0xfa, 0x36, 0xd5, 0xce, 0x01, 0x09, 0x0f, 0x2a, 0xa6, 0x5f, 0x2c, 0xd8, 0xfc, 0x38,
	0xe8, 0x1b, 0xe8, 0x5f, 0x65, 0x19, 0x4f, 0x63, 0xb7, 0x5c, 0x0e, 0x15, 0x6a, 0x53, 0x86, 0xdf,
	0x1d, 0x54, 0xcb, 0x8d, 0xb5, 0x81, 0x3e, 0x9f, 0xae, 0xde, 0x6a, 0xc9, 0x66, 0xda, 0x81, 0x1f,
	0xf5, 0xd8, 0x89, 0xd0, 0xcf, 0xc9, 0x1f, 0x17, 0x77, 0xff, 0x48, 0xbf, 0x6d, 0x99, 0x93, 0x1f,
	0x3e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x24, 0x9a, 0x2e, 0x1d, 0x94, 0x0b, 0x00, 0x00,
}
