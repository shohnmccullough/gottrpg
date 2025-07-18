package goback

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

type Middleware func(HandlerFunc) HandlerFunc

type StandardRouter struct {
	mux         *http.ServeMux
	middlewares []Middleware
}

func NewStandardRouter() *StandardRouter {
	return &StandardRouter{
		mux:         http.NewServeMux(),
		middlewares: []Middleware{},
	}
}

func (sr *StandardRouter) Use(middleware Middleware) {
	sr.middlewares = append(sr.middlewares, middleware)
}

func (sr *StandardRouter) Handle(pattern string, handler HandlerFunc) {
	finalHandler := sr.applyMiddleware(handler)
	sr.mux.HandleFunc(pattern, finalHandler)
}

func (sr *StandardRouter) GET(pattern string, handler HandlerFunc) {
	sr.Handle(pattern, sr.methodFilter("GET", handler))
}

func (sr *StandardRouter) POST(pattern string, handler HandlerFunc) {
	sr.Handle(pattern, sr.methodFilter("POST", handler))
}

func (sr *StandardRouter) PUT(pattern string, handler HandlerFunc) {
	sr.Handle(pattern, sr.methodFilter("PUT", handler))
}

func (sr *StandardRouter) DELETE(pattern string, handler HandlerFunc) {
	sr.Handle(pattern, sr.methodFilter("DELETE", handler))
}

func (sr *StandardRouter) PATCH(pattern string, handler HandlerFunc) {
	sr.Handle(pattern, sr.methodFilter("PATCH", handler))
}

func (sr *StandardRouter) HEAD(pattern string, handler HandlerFunc) {
	sr.Handle(pattern, sr.methodFilter("HEAD", handler))
}

func (sr *StandardRouter) OPTIONS(pattern string, handler HandlerFunc) {
	sr.Handle(pattern, sr.methodFilter("OPTIONS", handler))
}

func (sr *StandardRouter) Group(prefix string) *RouteGroup {
	return &RouteGroup{
		router: sr,
		prefix: strings.TrimSuffix(prefix, "/"),
	}
}

func (sr *StandardRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sr.mux.ServeHTTP(w, r)
}

func (sr *StandardRouter) methodFilter(method string, handler HandlerFunc) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			JSONError(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		handler(w, r)
	}
}

func (sr *StandardRouter) applyMiddleware(handler HandlerFunc) HandlerFunc {
	for i := len(sr.middlewares) - 1; i >= 0; i-- {
		handler = sr.middlewares[i](handler)
	}
	return handler
}

type RouteGroup struct {
	router *StandardRouter
	prefix string
}

func (rg *RouteGroup) GET(pattern string, handler HandlerFunc) {
	rg.router.GET(rg.prefix+pattern, handler)
}

func (rg *RouteGroup) POST(pattern string, handler HandlerFunc) {
	rg.router.POST(rg.prefix+pattern, handler)
}

func (rg *RouteGroup) PUT(pattern string, handler HandlerFunc) {
	rg.router.PUT(rg.prefix+pattern, handler)
}

func (rg *RouteGroup) DELETE(pattern string, handler HandlerFunc) {
	rg.router.DELETE(rg.prefix+pattern, handler)
}

func (rg *RouteGroup) PATCH(pattern string, handler HandlerFunc) {
	rg.router.PATCH(rg.prefix+pattern, handler)
}

func (rg *RouteGroup) HEAD(pattern string, handler HandlerFunc) {
	rg.router.HEAD(rg.prefix+pattern, handler)
}

func (rg *RouteGroup) Use(middleware Middleware) {
	rg.router.Use(middleware)
}

func JSONResponse(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Error encoding JSON response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func JSONError(w http.ResponseWriter, message string, status int) {
	errorResponse := map[string]string{"error": message}
	JSONResponse(w, errorResponse, status)
}

func LoggingMiddleware(next HandlerFunc) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.Method, r.URL.Path, r.RemoteAddr)
		next(w, r)
	}
}

func CORSMiddleware(next HandlerFunc) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		next(w, r)
	}
}

func RecoveryMiddleware(next HandlerFunc) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic recovered: %v", err)
				JSONError(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		next(w, r)
	}
}