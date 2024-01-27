// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/assignment/question.proto

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

// Client API for QuestionService service

type QuestionService interface {
	AddOne(ctx context.Context, in *ReqQuestionAdd, opts ...client.CallOption) (*ReplyQuestionOne, error)
	GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyQuestionOne, error)
	RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error)
	GetListByFilter(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyQuestionList, error)
	GetStatistic(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyStatistic, error)
	UpdateBase(ctx context.Context, in *ReqQuestionUpdate, opts ...client.CallOption) (*ReplyInfo, error)
	UpdateByFilter(ctx context.Context, in *RequestUpdate, opts ...client.CallOption) (*ReplyInfo, error)
}

type questionService struct {
	c    client.Client
	name string
}

func NewQuestionService(name string, c client.Client) QuestionService {
	return &questionService{
		c:    c,
		name: name,
	}
}

func (c *questionService) AddOne(ctx context.Context, in *ReqQuestionAdd, opts ...client.CallOption) (*ReplyQuestionOne, error) {
	req := c.c.NewRequest(c.name, "QuestionService.AddOne", in)
	out := new(ReplyQuestionOne)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionService) GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyQuestionOne, error) {
	req := c.c.NewRequest(c.name, "QuestionService.GetOne", in)
	out := new(ReplyQuestionOne)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionService) RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "QuestionService.RemoveOne", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionService) GetListByFilter(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyQuestionList, error) {
	req := c.c.NewRequest(c.name, "QuestionService.GetListByFilter", in)
	out := new(ReplyQuestionList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionService) GetStatistic(ctx context.Context, in *RequestFilter, opts ...client.CallOption) (*ReplyStatistic, error) {
	req := c.c.NewRequest(c.name, "QuestionService.GetStatistic", in)
	out := new(ReplyStatistic)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionService) UpdateBase(ctx context.Context, in *ReqQuestionUpdate, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "QuestionService.UpdateBase", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionService) UpdateByFilter(ctx context.Context, in *RequestUpdate, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "QuestionService.UpdateByFilter", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for QuestionService service

type QuestionServiceHandler interface {
	AddOne(context.Context, *ReqQuestionAdd, *ReplyQuestionOne) error
	GetOne(context.Context, *RequestInfo, *ReplyQuestionOne) error
	RemoveOne(context.Context, *RequestInfo, *ReplyInfo) error
	GetListByFilter(context.Context, *RequestFilter, *ReplyQuestionList) error
	GetStatistic(context.Context, *RequestFilter, *ReplyStatistic) error
	UpdateBase(context.Context, *ReqQuestionUpdate, *ReplyInfo) error
	UpdateByFilter(context.Context, *RequestUpdate, *ReplyInfo) error
}

func RegisterQuestionServiceHandler(s server.Server, hdlr QuestionServiceHandler, opts ...server.HandlerOption) error {
	type questionService interface {
		AddOne(ctx context.Context, in *ReqQuestionAdd, out *ReplyQuestionOne) error
		GetOne(ctx context.Context, in *RequestInfo, out *ReplyQuestionOne) error
		RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error
		GetListByFilter(ctx context.Context, in *RequestFilter, out *ReplyQuestionList) error
		GetStatistic(ctx context.Context, in *RequestFilter, out *ReplyStatistic) error
		UpdateBase(ctx context.Context, in *ReqQuestionUpdate, out *ReplyInfo) error
		UpdateByFilter(ctx context.Context, in *RequestUpdate, out *ReplyInfo) error
	}
	type QuestionService struct {
		questionService
	}
	h := &questionServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&QuestionService{h}, opts...))
}

type questionServiceHandler struct {
	QuestionServiceHandler
}

func (h *questionServiceHandler) AddOne(ctx context.Context, in *ReqQuestionAdd, out *ReplyQuestionOne) error {
	return h.QuestionServiceHandler.AddOne(ctx, in, out)
}

func (h *questionServiceHandler) GetOne(ctx context.Context, in *RequestInfo, out *ReplyQuestionOne) error {
	return h.QuestionServiceHandler.GetOne(ctx, in, out)
}

func (h *questionServiceHandler) RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error {
	return h.QuestionServiceHandler.RemoveOne(ctx, in, out)
}

func (h *questionServiceHandler) GetListByFilter(ctx context.Context, in *RequestFilter, out *ReplyQuestionList) error {
	return h.QuestionServiceHandler.GetListByFilter(ctx, in, out)
}

func (h *questionServiceHandler) GetStatistic(ctx context.Context, in *RequestFilter, out *ReplyStatistic) error {
	return h.QuestionServiceHandler.GetStatistic(ctx, in, out)
}

func (h *questionServiceHandler) UpdateBase(ctx context.Context, in *ReqQuestionUpdate, out *ReplyInfo) error {
	return h.QuestionServiceHandler.UpdateBase(ctx, in, out)
}

func (h *questionServiceHandler) UpdateByFilter(ctx context.Context, in *RequestUpdate, out *ReplyInfo) error {
	return h.QuestionServiceHandler.UpdateByFilter(ctx, in, out)
}