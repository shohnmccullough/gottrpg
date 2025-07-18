package goback

import (
	"strings"
)

/*
Capture incoming HTTP requests and run them through Router for handling
*/
type Router struct{}

/**/
func (r *Router) Use(path string, routes ...*Route) {
	for _, route := range routes {
		if !strings.EqualFold(route.Pattern, "/") {
			route.Pattern = strings.Join([]string{path, route.Pattern}, "/")
		}

		// TODO: temporary - figure out how to handle this better (how to allow a handler overwrite?)
		route.Handler = route.DefaultHandler
		route.RegisterRoute()
	}
}
