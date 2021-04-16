// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: news/news.proto

package news

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/micro/v3/service/api"
	client "github.com/micro/micro/v3/service/client"
	server "github.com/micro/micro/v3/service/server"
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
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for News service

func NewNewsEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for News service

type NewsService interface {
	News(ctx context.Context, in *EmptyMsg, opts ...client.CallOption) (News_NewsService, error)
	NewsList(ctx context.Context, in *EmptyMsg, opts ...client.CallOption) (*NewsResponse, error)
}

type newsService struct {
	c    client.Client
	name string
}

func NewNewsService(name string, c client.Client) NewsService {
	return &newsService{
		c:    c,
		name: name,
	}
}

func (c *newsService) News(ctx context.Context, in *EmptyMsg, opts ...client.CallOption) (News_NewsService, error) {
	req := c.c.NewRequest(c.name, "News.News", &EmptyMsg{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &newsServiceNews{stream}, nil
}

type News_NewsService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*NewsItem, error)
}

type newsServiceNews struct {
	stream client.Stream
}

func (x *newsServiceNews) Close() error {
	return x.stream.Close()
}

func (x *newsServiceNews) Context() context.Context {
	return x.stream.Context()
}

func (x *newsServiceNews) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *newsServiceNews) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *newsServiceNews) Recv() (*NewsItem, error) {
	m := new(NewsItem)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *newsService) NewsList(ctx context.Context, in *EmptyMsg, opts ...client.CallOption) (*NewsResponse, error) {
	req := c.c.NewRequest(c.name, "News.NewsList", in)
	out := new(NewsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for News service

type NewsHandler interface {
	News(context.Context, *EmptyMsg, News_NewsStream) error
	NewsList(context.Context, *EmptyMsg, *NewsResponse) error
}

func RegisterNewsHandler(s server.Server, hdlr NewsHandler, opts ...server.HandlerOption) error {
	type news interface {
		News(ctx context.Context, stream server.Stream) error
		NewsList(ctx context.Context, in *EmptyMsg, out *NewsResponse) error
	}
	type News struct {
		news
	}
	h := &newsHandler{hdlr}
	return s.Handle(s.NewHandler(&News{h}, opts...))
}

type newsHandler struct {
	NewsHandler
}

func (h *newsHandler) News(ctx context.Context, stream server.Stream) error {
	m := new(EmptyMsg)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.NewsHandler.News(ctx, m, &newsNewsStream{stream})
}

type News_NewsStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*NewsItem) error
}

type newsNewsStream struct {
	stream server.Stream
}

func (x *newsNewsStream) Close() error {
	return x.stream.Close()
}

func (x *newsNewsStream) Context() context.Context {
	return x.stream.Context()
}

func (x *newsNewsStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *newsNewsStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *newsNewsStream) Send(m *NewsItem) error {
	return x.stream.Send(m)
}

func (h *newsHandler) NewsList(ctx context.Context, in *EmptyMsg, out *NewsResponse) error {
	return h.NewsHandler.NewsList(ctx, in, out)
}
