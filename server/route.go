package goback

import (
	"encoding/json"
	"net/http"
)

/**/
type Route struct {
	Children []Route
	Handler  func(w http.ResponseWriter, req *http.Request)
	Methods  Method
	Pattern  string
	Table    string
}

/**/
func (r *Route) DefaultHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if fn, ok := r.Methods.methods[req.Method]; !ok {
		// attempt to inform client that the method does not exist, otherwise 500
		if b, err := json.Marshal(r.DefaultMethodHandler()); err != nil {
			WriteResponse(w, http.StatusInternalServerError, INTERNAL_SERVER_ERROR)
		} else {
			// teapot, haha
			WriteResponse(w, http.StatusTeapot, b)
		}
	} else {
		if res, err := json.Marshal(fn(req)); err != nil {
			WriteResponse(w, http.StatusInternalServerError, INTERNAL_SERVER_ERROR)
		} else {
			WriteResponse(w, http.StatusOK, res)
		}
	}
}

/**/
func (r *Route) DefaultMethodHandler() Response {
	// TODO: decide upon a default Reponse struct
	return Response{
		Message: "method does not exist",
	}
}

/*
Override DefaultMethodHandler() for incoming Delete requests
*/
func (r *Route) Delete(handler func(req *http.Request) interface{}) *Route {
	r.Methods.SetMethod(0, handler)
	return r
}

/*
Override DefaultMethodHandler() for incoming Get requests
*/
func (r *Route) Get(handler func(req *http.Request) interface{}) *Route {
	r.Methods.SetMethod(1, handler)
	return r
}

/*
Override DefaultMethodHandler() for incoming Head requests
*/
func (r *Route) Head(handler func(req *http.Request) interface{}) *Route {
	r.Methods.SetMethod(2, handler)
	return r
}

/*
Override DefaultMethodHandler() for incoming Patch requests
*/
func (r *Route) Patch(handler func(req *http.Request) interface{}) *Route {
	r.Methods.SetMethod(3, handler)
	return r
}

/*
Override DefaultMethodHandler() for incoming Post requests
*/
func (r *Route) Post(handler func(req *http.Request) interface{}) *Route {
	r.Methods.SetMethod(4, handler)
	return r
}

/*
Override DefaultMethodHandler() for incoming Put requests
*/
func (r *Route) Put(handler func(req *http.Request) interface{}) *Route {
	r.Methods.SetMethod(5, handler)
	return r
}

/*
Create a Route struct that contains a pattern and handler
*/
func CreateRoute(pattern string) *Route {
	var method Method
	return &Route{
		Pattern: pattern,
		Methods: *method.Init(),
	}
}

/*
Register a Route struct to http.HandleFunc()
*/
func (r *Route) RegisterRoute() {
	http.HandleFunc(r.Pattern, r.Handler)
}
