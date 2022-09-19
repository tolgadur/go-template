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

func (s *Endpoints) CreateHelloWorld(ctx context.Context, req interface{}) (interface{}, error) {
	request, ok := req.(*api.CreateHelloWorldRequest)
	if !ok {
		return nil, errors.New("new internal server error")
	}
	return s.Service.CreateHelloWorld(ctx, request)
}

func (s *Endpoints) GetHelloWorld(ctx context.Context, req interface{}) (interface{}, error) {
	request, ok := req.(*api.GetHelloWorldRequest)
	if !ok {
		return nil, errors.New("new internal server error")
	}
	return s.Service.GetHelloWorld(ctx, request)
}
