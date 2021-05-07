// Code generated by jrpc. DO NOT EDIT.

package routes

import (
	context "context"

	taxi "github.com/jakewright/home-automation/libraries/go/taxi"
	def "github.com/jakewright/home-automation/services/dummy/def"
)

// taxiRouter is an interface implemented by taxi.Router
type taxiRouter interface {
	HandleFunc(method, path string, handler func(context.Context, taxi.Decoder) (interface{}, error))
}

type handler interface {
	Log(ctx context.Context, body *def.LogRequest) (*def.LogResponse, error)
	Panic(ctx context.Context, body *def.PanicRequest) (*def.PanicResponse, error)
}

// Register adds the service's routes to the router
func Register(r taxiRouter, h handler) {
	r.HandleFunc("POST", "/log", func(ctx context.Context, decode taxi.Decoder) (interface{}, error) {
		body := &def.LogRequest{}
		if err := decode(body); err != nil {
			return nil, err
		}

		if err := body.Validate(); err != nil {
			return nil, err
		}

		return h.Log(ctx, body)
	})

	r.HandleFunc("POST", "/panic", func(ctx context.Context, decode taxi.Decoder) (interface{}, error) {
		body := &def.PanicRequest{}
		if err := decode(body); err != nil {
			return nil, err
		}

		if err := body.Validate(); err != nil {
			return nil, err
		}

		return h.Panic(ctx, body)
	})

}