package kbrouter

import (
	"fmt"
	"net/http"
)

type Router struct {
	routes map[string]map[string]func(req *KBRequest, res *KBResponse)
}

type HealthzResponse struct {
	Test string `json:"test"`
}

// Create a new kbrouter
func NewRouter() *Router {
	router := &Router{
		routes: make(map[string]map[string]func(req *KBRequest, res *KBResponse)),
	}
	router.AddRoute("GET", "/healthz", func(req *KBRequest, res *KBResponse) {
		// Implementation for creating a new product
		fmt.Println("got request to /healthz")
		res.SendString("OKAY\n")
	})
	return router
}

func (r *Router) ServeHTTP(w http.ResponseWriter, httpReq *http.Request) {
	if handlers, ok := r.routes[httpReq.URL.Path]; ok {
		if handler, methodExists := handlers[httpReq.Method]; methodExists {
			req := &KBRequest{
				httpReq: httpReq,
				Host:    httpReq.URL.Host,
				Path:    httpReq.URL.Path,
			}
			res := &KBResponse{
				writer: w,
			}
			handler(req, res)
			return
		}
	}
	http.NotFound(w, httpReq)
}
func (r *Router) Listen(port int, cb func()) error {
	addr := fmt.Sprintf("%s%d", ":", port)
	server := http.Server{
		Addr:    addr,
		Handler: r,
	}
	cb()
	return server.ListenAndServe()
}

func (r *Router) AddRoute(method, path string, handler func(req *KBRequest, res *KBResponse)) {
	if r.routes[path] == nil {
		r.routes[path] = make(map[string]func(req *KBRequest, res *KBResponse))
	}
	r.routes[path][method] = handler
}
