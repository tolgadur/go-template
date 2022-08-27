package v1

import (
	"context"
	"errors"
	"github.com/tolgadur/email-project/backend/internal/api"
	"go.uber.org/fx"
)

type Endpoints struct {
	fx.In

	Service api.Service
}

func (s *Endpoints) HelloWorld(ctx context.Context, req interface{}) (interface{}, error) {
	request, ok := req.(*api.HelloWorldRequest)
	if !ok {
		return nil, errors.New("new internal server error")
	}
	return s.Service.HelloWorld(ctx, request)
}
