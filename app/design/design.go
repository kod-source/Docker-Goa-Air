package design

import (
	. "github.com/shogo82148/goa-v1/design"
	. "github.com/shogo82148/goa-v1/design/apidsl"
)

var _ = API("test_build", func() {
	Title("The adder API")
	Description("A teaser for goa")
	Host("localhost:3000")
	Scheme("http")
})

var _ = Resource("test", func() {
	Action("add", func() {
		Routing(GET("add/:left/:right"))
		Description("add returns the sum of the left and right parameters in the response body")
		Params(func() {
			Param("left", Integer, "Left operand")
			Param("right", Integer, "Right operand")
		})
		Response(OK, "text/plain")
		Response(NotFound)
	})
})
