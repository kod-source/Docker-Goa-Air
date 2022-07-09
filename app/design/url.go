package design

import (
	. "github.com/shogo82148/goa-v1/design"
	. "github.com/shogo82148/goa-v1/design/apidsl"
)

var _ = Resource("url", func() {
	Action("url_add", func() {
		Routing(GET("url/:left/:middle/:right"))
		Description("add returns the sum")
		Params(func() {
			Param("left", Integer, "Left operand")
			Param("middle", Integer, "Middle operand")
			Param("right", Integer, "Right operand")
		})
		Response(OK, "text/plain")
		Response(NotFound)
	})
})
