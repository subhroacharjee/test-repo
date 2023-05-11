package routes

import "net/http"

type (
	Handler func(res http.ResponseWriter, req *http.Request)
	Route   map[string]Handler
)

type Routes struct {
	routes Route
}

func (r *Routes) Initalize() {
	r.routes = make(Route)
}

func (r *Routes) SetRoute(uri string, handler Handler) {
	r.routes[uri] = handler
}

func (r *Routes) GetRoutes() Route {
	return r.routes
}
