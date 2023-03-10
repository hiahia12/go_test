// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package microuser

import (
	"context"

	"go_test"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Request  = __microuser.Request
	Response = __microuser.Response

	Microuser interface {
		GetUser(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	}

	defaultMicrouser struct {
		cli zrpc.Client
	}
)

func NewMicrouser(cli zrpc.Client) Microuser {
	return &defaultMicrouser{
		cli: cli,
	}
}

func (m *defaultMicrouser) GetUser(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := __microuser.NewMicrouserClient(m.cli.Conn())
	return client.GetUser(ctx, in, opts...)
}
