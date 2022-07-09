// Code generated by goagen v1.5.13, DO NOT EDIT.
//
// API "test_build": Application Controllers
//
// Command:
// $ goagen
// --design=github.com/kod-source/Docker-Goa-Air/design
// --out=/Users/horikoudai/program-practice/docker-goa/app
// --version=v1.5.13

package app

import (
	"context"
	goa "github.com/shogo82148/goa-v1"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// TestController is the controller interface for the Test actions.
type TestController interface {
	goa.Muxer
	Add(*AddTestContext) error
}

// MountTestController "mounts" a Test resource controller on the given service.
func MountTestController(service *goa.Service, ctrl TestController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewAddTestContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Add(rctx)
	}
	service.Mux.Handle("GET", "/add/:left/:right", ctrl.MuxHandler("add", h, nil))
	service.LogInfo("mount", "ctrl", "Test", "action", "Add", "route", "GET /add/:left/:right")
}

// URLController is the controller interface for the URL actions.
type URLController interface {
	goa.Muxer
	URLAdd(*URLAddURLContext) error
}

// MountURLController "mounts" a URL resource controller on the given service.
func MountURLController(service *goa.Service, ctrl URLController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewURLAddURLContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.URLAdd(rctx)
	}
	service.Mux.Handle("GET", "/url/:left/:middle/:right", ctrl.MuxHandler("url_add", h, nil))
	service.LogInfo("mount", "ctrl", "URL", "action", "URLAdd", "route", "GET /url/:left/:middle/:right")
}
