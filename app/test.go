package main

import (
	"strconv"

	"github.com/kod-source/docker-goa-mysql/app"
	goa "github.com/shogo82148/goa-v1"
)

// TestController implements the test resource.
type TestController struct {
	*goa.Controller
}

// NewTestController creates a test controller.
func NewTestController(service *goa.Service) *TestController {
	return &TestController{Controller: service.NewController("TestController")}
}

// Add runs the add action.
func (c *TestController) Add(ctx *app.AddTestContext) error {
	amount := ctx.Left + ctx.Right
	return ctx.OK([]byte(strconv.Itoa(amount)))
	// TestController_Add: end_implement
}
