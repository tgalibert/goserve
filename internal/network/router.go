package network

import (
	"fmt"
	"goserve/internal/constant"
)

type HandleFunc func(req *Request, res *Response)
type Routes map[string]map[RequestMethod]HandleFunc

type Router struct {
	Routes Routes
}

func NewRouter() *Router {
	return &Router{
		Routes: make(map[string]map[RequestMethod]HandleFunc, 0),
	}
}

func (r *Router) GET(route string, handler HandleFunc) {
	r.addHandler(GET, route, handler)
}

func (r *Router) POST(route string, handler HandleFunc) {
	r.addHandler(POST, route, handler)
}

func (r *Router) PUT(route string, handler HandleFunc) {
	r.addHandler(PUT, route, handler)
}

func (r *Router) DELETE(route string, handler HandleFunc) {
	r.addHandler(DELETE, route, handler)
}

func (r *Router) PATCH(route string, handler HandleFunc) {
	r.addHandler(PATCH, route, handler)
}

func (r *Router) addHandler(method RequestMethod, route string, handler HandleFunc) {
	_, exist := r.Routes[route]
	if !exist {
		r.Routes[route] = make(map[RequestMethod]HandleFunc)
	}
	r.Routes[route][method] = handler
}

func (r *Router) Serve(req *Request, res *Response) {
	sub, exist := r.Routes[req.Path]
	if !exist {
		res.StatusCode = constant.StatusNotFound
		res.SetBody([]byte("404 Not Found"))
		return
	}
	handler, methodExist := sub[req.Method]
	if !methodExist {
		res.StatusCode = constant.StatusMethodNotAllowed
		res.SetBody([]byte("405 Method Not Allowed"))
		allowed := ""
		for key := range sub {
			if allowed == "" {
				allowed = fmt.Sprintf("%s", key)
			} else {
				allowed += ", " + string(key)
			}
		}
		res.SetHeader("Allow", allowed)
		return
	}
	handler(req, res)
}