//go:generate goagen bootstrap -d github.com/kod-source/Docker-Goa-Air/design

package main

import (
	"github.com/kod-source/Docker-Goa-Air/app"
	goa "github.com/shogo82148/goa-v1"
	"github.com/shogo82148/goa-v1/middleware"
)

func main() {
	// Create service
	service := goa.New("test_build")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "test" controller
	c := NewTestController(service)
	app.MountTestController(service, c)
	u := NewURLController(service)
	app.MountURLController(service, u)

	// Start service
	if err := service.ListenAndServe(":3000"); err != nil {
		service.LogError("startup", "err", err)
	}
}
