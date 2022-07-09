package main

import (
	"strconv"

	"github.com/kod-source/Docker-Goa-Air/app"
	goa "github.com/shogo82148/goa-v1"
)

// URLController implements the url resource.
type URLController struct {
	*goa.Controller
}

// NewURLController creates a url controller.
func NewURLController(service *goa.Service) *URLController {
	return &URLController{Controller: service.NewController("URLController")}
}

// URLAdd runs the url_add action.
func (c *URLController) URLAdd(ctx *app.URLAddURLContext) error {
	sum := ctx.Left + ctx.Middle + ctx.Right
	return ctx.OK([]byte(strconv.Itoa(sum)))
}
