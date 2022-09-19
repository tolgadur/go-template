package http

import (
	"context"
	"encoding/json"
	"errors"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/tolgadur/email-project/backend/internal/api"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net/http"
)

var (
	errorMalformedRequest = errors.New("request is malformed. Either request or body is nil")
)

type Codec struct {
	fx.In

	Logger *zap.SugaredLogger
}

func (c *Codec) decodeCreateHelloWorldRequest(_ context.Context, r *http.Request) (interface{}, error) {
	c.Logger.Info("decodeHelloWorldRequest")
	if r == nil || r.Body == http.NoBody || r.Body == nil {
		c.Logger.Error("Error malformed request")
		return nil, errorMalformedRequest
	}
	var request api.HelloWorldRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	if request.Name == "" {
		c.Logger.Error("Error malformed request: name is missing")
		return nil, errorMalformedRequest
	}
	return &request, nil
}

func (c *Codec) decodeGetHelloWorldRequest(_ context.Context, r *http.Request) (interface{}, error) {
	c.Logger.Info("decodeHelloWorldRequest")
	if r == nil || r.URL == nil {
		c.Logger.Error("Error malformed request")
		return nil, errorMalformedRequest
	}
	vars := mux.Vars(r)
	return &api.HelloWorldRequest{vars["name"]}, nil
}

func (c *Codec) encode(ctx context.Context, w http.ResponseWriter, resp interface{}) error {
	c.Logger.Info("encoding json response")
	return httptransport.EncodeJSONResponse(ctx, w, resp)
}
