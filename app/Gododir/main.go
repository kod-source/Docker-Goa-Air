package main

import (
	do "gopkg.in/godo.v2"
)

func tasks(p *do.Project) {

	// p.Task("server", nil, func(c *do.Context) {
	// 	// rebuilds and restarts when a watched file changes
	// 	c.Start(".././docker-goa-mysql", do.M{"$in": "./"})
	// }).Src("*.go", "**/*.go").
	// 	Debounce(3000)

	p.Task("build", do.S{"views", "assets"}, func(c *do.Context) {
		c.Run("GOOS=linux GOARCH=amd64 go build", do.M{"$in": "cmd/server"})
	}).Src("**/*.go")
}

func main() {
	do.Godo(tasks)
}
