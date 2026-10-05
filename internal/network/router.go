package network

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

}

func (r *Router) POST(route string, handler HandleFunc) {

}

func (r *Router) PUT(route string, handler HandleFunc) {

}

func (r *Router) DELETE(route string, handler HandleFunc) {

}

func (r *Router) PATCH(route string, handler HandleFunc) {

}