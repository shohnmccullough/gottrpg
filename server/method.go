package goback

import (
	"fmt"
	"net/http"
)

// type to simplify assigning methods
type RequestMethod int8

const (
	Delete RequestMethod = 0
	Head   RequestMethod = 1
	Get    RequestMethod = 2
	Patch  RequestMethod = 3
	Post   RequestMethod = 4
	Put    RequestMethod = 5
)

/*
Contains a map of functions to handle incoming HTTP requests
  - functions are assigned through Route.SetMethod()
*/
type Method struct {
	methods map[string]func(req *http.Request) interface{}
}

/*
Initialize a new methods map
*/
func (m *Method) Init() *Method {
	m.methods = make(map[string]func(req *http.Request) interface{})
	return m
}

func (m *Method) SetMethod(method RequestMethod, handler func(req *http.Request) interface{}) {
	switch method {
	case Delete:
		m.methods["DELETE"] = handler
	case Head:
		m.methods["HEAD"] = handler
	case Get:
		m.methods["GET"] = handler
	case Patch:
		m.methods["PATCH"] = handler
	case Post:
		m.methods["POST"] = handler
	case Put:
		m.methods["PUT"] = handler
	default:
		fmt.Println("invalid RequestMethod value", method)
	}
}
