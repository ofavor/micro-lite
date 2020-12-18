package toto

import (
	"context"

	"github.com/ofavor/micro-lite/server"

	"github.com/ofavor/micro-lite/client"
)

type TotoService interface {
	Multiply(ctx context.Context, req *Request, opts ...client.CallOption) (*Response, error)
}

func NewTotoService(c client.Client) TotoService {
	return &totoService{
		c: c,
	}
}

type totoService struct {
	c client.Client
}

func (s *totoService) Multiply(ctx context.Context, req *Request, opts ...client.CallOption) (*Response, error) {
	in := client.NewRequest("simple.server", "Toto.Multiply", req)
	out := new(Response)
	if err := s.c.Call(ctx, in, out); err != nil {
		return nil, err
	}
	return out, nil
}

type TotoHandler interface {
	Multiply(ctx context.Context, req *Request, rsp *Response) error
}

func RegisterTotoHandler(s server.Server, h TotoHandler) {
	hdr := server.NewHandler("Toto", h)
	s.Handle(hdr)
}
