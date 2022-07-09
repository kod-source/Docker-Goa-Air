//go:generate goagen bootstrap -d github.com/kod-source/docker-goa-mysql/design

package main

import (
	"github.com/kod-source/docker-goa-mysql/app"
	goa "github.com/shogo82148/goa-v1"
	"github.com/shogo82148/goa-v1/middleware"
)

func main() {
	// Create service
	service := goa.New("adder")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "operands" controller
	c := NewTestController(service)
	app.MountTestController(service, c)

	// Start service
	if err := service.ListenAndServe(":3000"); err != nil {
		service.LogError("startup", "err", err)
	}
}
