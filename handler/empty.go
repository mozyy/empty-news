package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"
	"google.golang.org/protobuf/types/known/emptypb"

	"empty/crawler"
	empty "empty/proto"
)

type Empty struct{}

// News is a single request handler called via client.Call or the generated client code
func (e *Empty) News(ctx context.Context, req *emptypb.Empty, stream empty.Empty_NewsStream) error {
	log.Info("Received Empty.News request")
	crawler.News(func(list []*empty.NewsItem, err error) {
		for _, item := range list {
			stream.Send(item)
		}
		stream.Close()
	})
	return nil
}

// NewsList is a single request handler called via client.Call or the generated client code
func (e *Empty) NewsList(ctx context.Context, req *emptypb.Empty, rsp *empty.NewsResponse) error {
	log.Info("Received Empty.NewsList request")
	crawler.News(func(list []*empty.NewsItem, err error) {
		rsp.List = list
	})
	return nil
}

// Call is a single request handler called via client.Call or the generated client code
func (e *Empty) Call(ctx context.Context, req *empty.Request, rsp *empty.Response) error {
	log.Info("Received Empty.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Empty) Stream(ctx context.Context, req *empty.StreamingRequest, stream empty.Empty_StreamStream) error {
	log.Infof("Received Empty.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&empty.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Empty) PingPong(ctx context.Context, stream empty.Empty_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&empty.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
