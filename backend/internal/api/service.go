package api

import (
	"context"
	"database/sql"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Service struct {
	fx.In

	DB     *sql.DB
	Logger *zap.SugaredLogger
}

func (s *Service) CreateHelloWorld(_ context.Context, request *HelloWorldRequest) (*HelloWorldResponse, error) {
	_, err := s.DB.Exec("INSERT INTO hello_world (name) VALUES ($1)", request.Name)
	if err != nil {
		s.Logger.Errorf("Error while inserting hello world: %s", err)
		return nil, err
	}
	return &HelloWorldResponse{"Hello, " + request.Name}, nil
}

func (s *Service) GetHelloWorld(_ context.Context, request *HelloWorldRequest) (*HelloWorldResponse, error) {
	return &HelloWorldResponse{"Hello, " + request.Name}, nil
}
