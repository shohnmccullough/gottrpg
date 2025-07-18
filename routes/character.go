package routes

import (
	"net/http"

	server "example.com/gottrpg/v2/server"
)

// func Character() *server.Route {
// 	return server.CreateRoute("character").
// 		Get(func(req *http.Request) interface{} {
// 			return ""
// 		})
// }

var Character *server.Route = server.CreateRoute("character").
	Get(func(req *http.Request) interface{} {
		return ""
	})
