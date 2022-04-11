// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/assignment/task.proto

package assignment

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for TaskService service

type TaskService interface {
	AddOne(ctx context.Context, in *ReqTaskAdd, opts ...client.CallOption) (*ReplyTaskOne, error)
	GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyTaskOne, error)
	GetOneByFilter(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyTaskOne, error)
	RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error)
	GetListByFilter(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyTaskList, error)
	GetStatistic(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyStatistic, error)
	UpdateByFilter(ctx context.Context, in *RequestUpdate, opts ...client.CallOption) (*ReplyInfo, error)
	UpdateBase(ctx context.Context, in *ReqTaskUpdate, opts ...client.CallOption) (*ReplyInfo, error)
	UpdateStatus(ctx context.Context, in *RequestIntFlag, opts ...client.CallOption) (*ReplyInfo, error)
	UpdateAgent(ctx context.Context, in *RequestFlag, opts ...client.CallOption) (*ReplyInfo, error)
	AppendRecord(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyList, error)
	SubtractRecord(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyList, error)
}

type taskService struct {
	c    client.Client
	name string
}

func NewTaskService(name string, c client.Client) TaskService {
	return &taskService{
		c:    c,
		name: name,
	}
}

func (c *taskService) AddOne(ctx context.Context, in *ReqTaskAdd, opts ...client.CallOption) (*ReplyTaskOne, error) {
	req := c.c.NewRequest(c.name, "TaskService.AddOne", in)
	out := new(ReplyTaskOne)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskService) GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyTaskOne, error) {
	req := c.c.NewRequest(c.name, "TaskService.GetOne", in)
	out := new(ReplyTaskOne)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskService) GetOneByFilter(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyTaskOne, error) {
	req := c.c.NewRequest(c.name, "TaskService.GetOneByFilter", in)
	out := new(ReplyTaskOne)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskService) RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "TaskService.RemoveOne", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskService) GetListByFilter(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyTaskList, error) {
	req := c.c.NewRequest(c.name, "TaskService.GetListByFilter", in)
	out := new(ReplyTaskList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskService) GetStatistic(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyStatistic, error) {
	req := c.c.NewRequest(c.name, "TaskService.GetStatistic", in)
	out := new(ReplyStatistic)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskService) UpdateByFilter(ctx context.Context, in *RequestUpdate, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "TaskService.UpdateByFilter", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskService) UpdateBase(ctx context.Context, in *ReqTaskUpdate, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "TaskService.UpdateBase", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskService) UpdateStatus(ctx context.Context, in *RequestIntFlag, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "TaskService.UpdateStatus", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskService) UpdateAgent(ctx context.Context, in *RequestFlag, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "TaskService.UpdateAgent", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskService) AppendRecord(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyList, error) {
	req := c.c.NewRequest(c.name, "TaskService.AppendRecord", in)
	out := new(ReplyList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskService) SubtractRecord(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyList, error) {
	req := c.c.NewRequest(c.name, "TaskService.SubtractRecord", in)
	out := new(ReplyList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TaskService service

type TaskServiceHandler interface {
	AddOne(context.Context, *ReqTaskAdd, *ReplyTaskOne) error
	GetOne(context.Context, *RequestInfo, *ReplyTaskOne) error
	GetOneByFilter(context.Context, *RequestInfo, *ReplyTaskOne) error
	RemoveOne(context.Context, *RequestInfo, *ReplyInfo) error
	GetListByFilter(context.Context, *RequestFilter, *ReplyTaskList) error
	GetStatistic(context.Context, *RequestFilter, *ReplyStatistic) error
	UpdateByFilter(context.Context, *RequestUpdate, *ReplyInfo) error
	UpdateBase(context.Context, *ReqTaskUpdate, *ReplyInfo) error
	UpdateStatus(context.Context, *RequestIntFlag, *ReplyInfo) error
	UpdateAgent(context.Context, *RequestFlag, *ReplyInfo) error
	AppendRecord(context.Context, *RequestInfo, *ReplyList) error
	SubtractRecord(context.Context, *RequestInfo, *ReplyList) error
}

func RegisterTaskServiceHandler(s server.Server, hdlr TaskServiceHandler, opts ...server.HandlerOption) error {
	type taskService interface {
		AddOne(ctx context.Context, in *ReqTaskAdd, out *ReplyTaskOne) error
		GetOne(ctx context.Context, in *RequestInfo, out *ReplyTaskOne) error
		GetOneByFilter(ctx context.Context, in *RequestInfo, out *ReplyTaskOne) error
		RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error
		GetListByFilter(ctx context.Context, in *RequestFilter, out *ReplyTaskList) error
		GetStatistic(ctx context.Context, in *RequestFilter, out *ReplyStatistic) error
		UpdateByFilter(ctx context.Context, in *RequestUpdate, out *ReplyInfo) error
		UpdateBase(ctx context.Context, in *ReqTaskUpdate, out *ReplyInfo) error
		UpdateStatus(ctx context.Context, in *RequestIntFlag, out *ReplyInfo) error
		UpdateAgent(ctx context.Context, in *RequestFlag, out *ReplyInfo) error
		AppendRecord(ctx context.Context, in *RequestInfo, out *ReplyList) error
		SubtractRecord(ctx context.Context, in *RequestInfo, out *ReplyList) error
	}
	type TaskService struct {
		taskService
	}
	h := &taskServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&TaskService{h}, opts...))
}

type taskServiceHandler struct {
	TaskServiceHandler
}

func (h *taskServiceHandler) AddOne(ctx context.Context, in *ReqTaskAdd, out *ReplyTaskOne) error {
	return h.TaskServiceHandler.AddOne(ctx, in, out)
}

func (h *taskServiceHandler) GetOne(ctx context.Context, in *RequestInfo, out *ReplyTaskOne) error {
	return h.TaskServiceHandler.GetOne(ctx, in, out)
}

func (h *taskServiceHandler) GetOneByFilter(ctx context.Context, in *RequestInfo, out *ReplyTaskOne) error {
	return h.TaskServiceHandler.GetOneByFilter(ctx, in, out)
}

func (h *taskServiceHandler) RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error {
	return h.TaskServiceHandler.RemoveOne(ctx, in, out)
}

func (h *taskServiceHandler) GetListByFilter(ctx context.Context, in *RequestFilter, out *ReplyTaskList) error {
	return h.TaskServiceHandler.GetListByFilter(ctx, in, out)
}

func (h *taskServiceHandler) GetStatistic(ctx context.Context, in *RequestFilter, out *ReplyStatistic) error {
	return h.TaskServiceHandler.GetStatistic(ctx, in, out)
}

func (h *taskServiceHandler) UpdateByFilter(ctx context.Context, in *RequestUpdate, out *ReplyInfo) error {
	return h.TaskServiceHandler.UpdateByFilter(ctx, in, out)
}

func (h *taskServiceHandler) UpdateBase(ctx context.Context, in *ReqTaskUpdate, out *ReplyInfo) error {
	return h.TaskServiceHandler.UpdateBase(ctx, in, out)
}

func (h *taskServiceHandler) UpdateStatus(ctx context.Context, in *RequestIntFlag, out *ReplyInfo) error {
	return h.TaskServiceHandler.UpdateStatus(ctx, in, out)
}

func (h *taskServiceHandler) UpdateAgent(ctx context.Context, in *RequestFlag, out *ReplyInfo) error {
	return h.TaskServiceHandler.UpdateAgent(ctx, in, out)
}

func (h *taskServiceHandler) AppendRecord(ctx context.Context, in *RequestInfo, out *ReplyList) error {
	return h.TaskServiceHandler.AppendRecord(ctx, in, out)
}

func (h *taskServiceHandler) SubtractRecord(ctx context.Context, in *RequestInfo, out *ReplyList) error {
	return h.TaskServiceHandler.SubtractRecord(ctx, in, out)
}