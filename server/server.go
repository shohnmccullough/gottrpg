package goback

import (
	"fmt"
	"net/http"
	"time"
)

// would be simpler as a const
var INTERNAL_SERVER_ERROR []byte = []byte("Internal Server Error")

/*
10 seconds
*/
func DefaultTimeout() time.Duration {
	return 10 * time.Second
}

/*
8 kilobytes
*/
func DefaultMaxHeaderBytes() int {
	return 8 * 1000
}

/*
Default HTTP response writer
*/
func WriteResponse(w http.ResponseWriter, status int, body []byte) {
	w.WriteHeader(status)
	w.Write(body)
}

/*
Struct to wrap around *http.Server
*/
type Server struct {
	Serve *http.Server
}

/*
Create a new HTTP server to listen and serve requests
*/
func InitServer() *Server {
	return &Server{
		Serve: &http.Server{
			Addr:           fmt.Sprintf("%s:%d", "127.0.0.1", 8080),
			Handler:        http.DefaultServeMux,
			ReadTimeout:    DefaultTimeout(),
			WriteTimeout:   DefaultTimeout(),
			MaxHeaderBytes: DefaultMaxHeaderBytes(),
		},
	}
}

/*
Assign a host and port to listen on
*/
func (s *Server) Address(host string, port int) *Server {
	s.Serve.Addr = fmt.Sprintf("%s:%d", host, port)
	return s
}

/*
Assign a custom multiplexer to the Server
*/
func (s *Server) Handler(mux *http.ServeMux) *Server {
	s.Serve.Handler = mux
	return s
}

/*
Start listening and serving requests using the constructed *http.Server object
*/
func (s *Server) ListenAndServe() error {
	return s.Serve.ListenAndServe()
}

/*
Assign a custom max header bytes value
*/
func (s *Server) MaxHeaderBytes(kb int) *Server {
	s.Serve.MaxHeaderBytes = kb * 1000
	return s
}

/*
Assign a custom read timeout value
*/
func (s *Server) ReadTimeout(seconds time.Duration) *Server {
	s.Serve.ReadTimeout = seconds * time.Second
	return s
}

/*
Assign a custom write timeout value
*/
func (s *Server) WriteTimeout(seconds time.Duration) *Server {
	s.Serve.WriteTimeout = seconds * time.Second
	return s
}
