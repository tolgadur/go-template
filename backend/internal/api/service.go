package api

import (
	"context"
	"go.uber.org/fx"
)

type Service struct {
	fx.In
}

func (s *Service) HelloWorld(_ context.Context, request *HelloWorldRequest) (*HelloWorldResponse, error) {
	return &HelloWorldResponse{"Hello, " + request.Name}, nil
}
