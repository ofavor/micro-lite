// Code generated by protoc-gen-micro. DO NOT EDIT.

package toto

import (
  "context"
  "github.com/ofavor/micro-lite/server"
  "github.com/ofavor/micro-lite/client"
)

type TotoService interface {
  Multiply(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type totoService struct {
  serviceName string
  c client.Client
}

func NewTotoService(name string, c client.Client) TotoService {
  return &totoService {
    serviceName: name,
    c: c,
  }
}

func (s *totoService)Multiply(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
  req := client.NewRequest(s.serviceName, "Toto.Multiply", in)
  rsp := new(Response)
  err := s.c.Call(ctx, req, rsp, opts...)
  if err != nil {
    return nil, err
  }
  return rsp, nil
}

type TotoHandler interface {
  Multiply(ctx context.Context, in *Request, out *Response) error
}

func RegisterTotoHandler(s server.Server, h TotoHandler) {
  hdr := server.NewHandler("Toto", h)
  s.Handle(hdr)
}

