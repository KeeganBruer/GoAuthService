package kbrouter

import (
	"fmt"
	"net/http"
	"strings"
)

type KBRouteHandler = func(req *KBRequest, res *KBResponse)

type Router struct {
	middlewares []KBRouteHandler
	routes      map[string]map[string]KBRouteHandler
	subRouters  map[string]*Router
}

type HealthzResponse struct {
	Test string `json:"test"`
}

// Create a new kbrouter
func NewRouter() *Router {
	middleware := &[]KBRouteHandler{}
	router := &Router{
		middlewares: *middleware,
		routes:      make(map[string]map[string]KBRouteHandler),
		subRouters:  make(map[string]*Router),
	}

	return router
}

func (r *Router) ServeHTTP(w http.ResponseWriter, httpReq *http.Request) {
	r.HandleServe(w, httpReq, "")
}
func (r *Router) HandleServe(w http.ResponseWriter, httpReq *http.Request, basepath string) {
	CurrPath := strings.Replace(httpReq.URL.Path, basepath, "", 1)
	//Empty route should map to the / route
	if CurrPath == "" {
		CurrPath = "/"
	}
	req := &KBRequest{
		httpReq:  httpReq,
		Host:     httpReq.URL.Host,
		CurrPath: CurrPath,
		Path:     httpReq.URL.Path,
	}
	splitPath := strings.Split(req.CurrPath, "/")
	res := NewKBResponse(w)

	if httpReq.Method == "OPTIONS" {
		res.SetHeader("Allow", "*")
		res.SetHeader("Access-Control-Allow-Credentials", "true")
		res.SetHeader("Access-Control-Allow-Origin", "*")
		res.SetHeader("Vary", "Origin")
		res.SetHeader("Access-Control-Allow-Headers", "*")
		res.SendString("OKAY")
		return
	}

	//Run middlewares
	for i := range r.middlewares {
		middleware := r.middlewares[i]
		middleware(req, res)
		if !res.IsOpen {
			return
		}
	}

	//Handle this router's routes
	if handlers, ok := r.routes[req.CurrPath]; ok {
		if handler, methodExists := handlers[httpReq.Method]; methodExists {
			handler(req, res)
			return
		}
	}

	subPath := "/"
	if len(splitPath) > 1 {
		subPath = fmt.Sprintf("/%s", splitPath[1])
	}
	//Check for sub routers
	if r.subRouters[subPath] != nil {
		fullPath := fmt.Sprintf("%s%s", basepath, subPath)
		r.subRouters[subPath].HandleServe(w, httpReq, fullPath)
		return
	}

	http.NotFound(w, httpReq)
}

// Set up a http server using the router to handle serving routes
func (r *Router) Listen(port int, cb func(port int)) error {
	addr := fmt.Sprintf("%s%d", ":", port)
	server := http.Server{
		Addr:    addr,
		Handler: r,
	}
	cb(port)
	return server.ListenAndServe()
}

func (r *Router) AddMiddleware(handlers ...KBRouteHandler) {
	r.middlewares = append(r.middlewares, handlers...)
}

// Add a route to the router
func (r *Router) AddRoute(method, path string, handlers ...KBRouteHandler) {
	if r.routes[path] == nil {
		r.routes[path] = make(map[string]KBRouteHandler)
	}
	if len(handlers) > 1 {
		r.routes[path][method] = func(req *KBRequest, res *KBResponse) {
			for i := range handlers {
				handlers[i](req, res)
			}
		}
	} else if len(handlers) == 1 {
		r.routes[path][method] = handlers[0]
	}
}
func (r *Router) AddSubRouter(basePath string, router *Router) {
	r.subRouters[basePath] = router
}

func (r *Router) AddHealthRoute(route string) {
	r.AddRoute("GET", route, func(req *KBRequest, res *KBResponse) {
		res.SendString("OKAY\n")
	})
}
