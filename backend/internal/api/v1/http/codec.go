package http

import (
	"context"
	"encoding/json"
	"errors"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/tolgadur/email-project/backend/internal/api"
	"go.uber.org/fx"
	"net/http"
)

var (
	errorMalformedRequest = errors.New("request is malformed. Either request or body is nil")
)

type Codec struct {
	fx.In
}

func (c *Codec) decodeHelloWorldRequest(_ context.Context, r *http.Request) (interface{}, error) {
	if r == nil || r.Body == nil {
		return nil, errorMalformedRequest
	}
	var request api.HelloWorldRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	if request.Name == "" {
		return nil, errorMalformedRequest
	}
	return &request, nil
}

func (c *Codec) decodeHelloWorld2Request(_ context.Context, r *http.Request) (interface{}, error) {
	if r == nil || r.Body == nil {
		return nil, errorMalformedRequest
	}
	vars := mux.Vars(r)
	return &api.HelloWorldRequest{vars["name"]}, nil
}

func (c *Codec) encode(ctx context.Context, w http.ResponseWriter, resp interface{}) error {
	return httptransport.EncodeJSONResponse(ctx, w, resp)
}
