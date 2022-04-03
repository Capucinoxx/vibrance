package router

import (
	"fmt"
	"net/http"
)

type Router interface {
	Append(routes ...Route) Router
	AddSubRouter(rr Router)
	Consume()
}

type router struct {
	prefix      string
	routes      map[string]map[Method]Route
	middlewares Middlewares
	printRoute  func(pattern string, methods ...string)
	isConsummed bool
}

func printMethod(pattern string, methods ...string) {
	for i := 0; i < len(methods); i++ {
		fmt.Printf("[%-6s] %s\n", methods[i], pattern)
	}
}

func emptyMiddlewares(middlewares ...Middleware) bool {
	if len(middlewares) == 0 {
		return false
	}

	for _, m := range middlewares {
		if m != nil {
			return false
		}
	}

	return true
}

func Consumer(prefix string, routes Routes, middlewares ...Middleware) Router {
	if emptyMiddlewares(middlewares...) {
		middlewares = []Middleware{}
	}

	middlewares = append(middlewares, nil)
	copy(middlewares[1:], middlewares)
	middlewares[0] = Logger

	r := &router{
		prefix,
		make(map[string]map[Method]Route),
		middlewares,
		printMethod,
		false,
	}

	r.makeRoutes(routes...)

	return r
}

func (r *router) Append(routes ...Route) Router {
	r.makeRoutes(routes...)
	return r
}

func (r *router) AddSubRouter(rr Router) {
	switch v := rr.(type) {
	case *router:
		r.mergeRoutes(rr)
		v.isConsummed = true
	}
}

func (r *router) Consume() {
	if r.isConsummed {
		panic("error")
	}
	r.isConsummed = true

	for pattern, methods := range r.routes {
		func(pattern string, methods map[Method]Route) {
			handler := CORS(func(w http.ResponseWriter, rq *http.Request) {
				if next, ok := methods[Method(rq.Method)]; ok {
					if len(r.middlewares) == 0 {
						next.HandlerFunc(w, rq)
						return
					}

					wrapped := next.HandlerFunc
					for i := len(r.middlewares) - 1; i >= 0; i-- {
						wrapped = r.middlewares[i](wrapped)
					}
					wrapped(w, rq)
				}
			})

			http.HandleFunc(pattern, handler)

			for method := range methods {
				fmt.Printf("[%-7s] %s\n", method, pattern)
			}

		}(pattern, methods)
	}
}

func (r *router) makeRoutes(routes ...Route) {
	for i := 0; i < len(routes); i++ {
		routes[i].Pattern = r.prefix + routes[i].Pattern
		if _, ok := r.routes[routes[i].Pattern]; !ok {
			r.routes[routes[i].Pattern] = make(map[Method]Route)
		}
		r.routes[routes[i].Pattern][routes[i].Method] = routes[i]
	}
}

func (r *router) mergeRoutes(rr Router) {
	v := rr.(*router)

	for pattern, methods := range v.routes {
		pattern = r.prefix + pattern

		if _, ok := r.routes[pattern]; !ok {
			r.routes[pattern] = make(map[Method]Route)
		}

		for method, route := range methods {
			r.routes[pattern][method] = route
		}
	}
}
