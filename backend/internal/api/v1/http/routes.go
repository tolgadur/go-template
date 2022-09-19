package http

import (
	httptransport "github.com/go-kit/kit/transport/http"
)

func (s *Server) RegisterRoutes() error {
	s.Logger.Info("Registering routes")
	s.Router.Handle("/hello-world/{name}", httptransport.NewServer(
		s.Endpoints.GetHelloWorld,
		s.Codec.decodeGetHelloWorldRequest,
		s.Codec.encode,
	)).Methods("GET")

	s.Router.Handle("/hello-world", httptransport.NewServer(
		s.Endpoints.CreateHelloWorld,
		s.Codec.decodeCreateHelloWorldRequest,
		s.Codec.encode,
	)).Methods("POST")

	return nil
}
