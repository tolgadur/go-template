package http

import (
	httptransport "github.com/go-kit/kit/transport/http"
)

func (s *Server) RegisterRoutes() error {
	s.Logger.Info("Registering routes")
	s.Router.Handle("/{name}", httptransport.NewServer(
		s.Endpoints.HelloWorld,
		s.Codec.decodeHelloWorld2Request,
		s.Codec.encode,
	)).Methods("GET")

	s.Router.Handle("/hello-world", httptransport.NewServer(
		s.Endpoints.HelloWorld,
		s.Codec.decodeHelloWorldRequest,
		s.Codec.encode,
	)).Methods("GET")

	return nil
}
