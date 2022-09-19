package api

import (
	"context"
	"database/sql"
	"fmt"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Service struct {
	fx.In

	DB     *sql.DB
	Logger *zap.SugaredLogger
}

func (s *Service) CreateHelloWorld(_ context.Context, request *CreateHelloWorldRequest) (*HelloWorldResponse, error) {
	_, err := s.DB.Exec("INSERT INTO test_service.hello_world (name, last_name) VALUES ($1, $2)", request.Name, request.LastName)
	if err != nil {
		s.Logger.Errorf("Error while inserting hello world: %s", err)
		return nil, err
	}

	message := fmt.Sprintf("Hello, %s %s", request.Name, request.LastName)
	return &HelloWorldResponse{message}, nil
}

func (s *Service) GetHelloWorld(_ context.Context, request *GetHelloWorldRequest) (*HelloWorldResponse, error) {
	var id int
	var name, lastName string
	err := s.DB.QueryRow("SELECT * FROM test_service.hello_world WHERE name=$1", request.Name).Scan(&id, &name, &lastName)
	if err != nil {
		s.Logger.Errorf("Error while getting hello world: %s", err)
		return nil, err
	}

	message := fmt.Sprintf("Hello, %s %s", name, lastName)
	return &HelloWorldResponse{message}, nil
}
