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
		Response(OK, URLMedia)
		Response(NotFound)
	})
})

var URLMedia = MediaType("application/vnd.url+json", func() {
	Description("A tenant account")
	Attributes(func() {
		Attribute("id", Integer, "ID of account", func() {
			Example(1)
		})
		Attribute("amount", Integer, "API href of amount", func() {
			Example(300)
		})

		Required("id")
	})

	View("default", func() {
		Attribute("id")
		Attribute("amount")
	})
})
