package router

import (
	"fmt"
	"net/http"
)

// Middleware fonction faisant la passerelle entre la requête client et la logique
// métier de la requête
type Middleware func(http.HandlerFunc) http.HandlerFunc

// Middlewares représente une collection de middleware
type Middlewares []Middleware

// Method est une méthode du protocole
// http ["GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"]
type Method string

// Route est la représentation d'une route http
type Route struct {
	// description de la route
	Name string

	// méthode du protocole http ["GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"]
	Method Method

	// pattern suivant le préfix
	Pattern string

	HandlerFunc func(http.ResponseWriter, *http.Request)
}

// Routes représente une collection de routes
type Routes []Route

// router représentation internet du router.
// À cette struct sera greffé les différentes fonctionnalités du corps de la
// logique métier permettant la gestion de la consommation et mise en place des
// routes du service
type router struct {
	// ensemble des routes du service
	// map[pattern]map[method]Route
	routes map[string]map[Method]Route

	// ensemble des middlewares du service
	middlewares Middlewares

	// état de l'instance si les routes ont été consummées ou non
	isConsummed bool
}

func Router(routes ...Route) *router {
	r := &router{
		make(map[string]map[Method]Route),
		nil,
		false,
	}

	// création des différentes routes pouvant être présente
	r.makeRoutes(routes...)

	return r
}

// makeRoutes fonction interne encapsulant la logique métier concernant l'ajout
// de route dans la représentation
func (r *router) makeRoutes(routes ...Route) {
	for i := 0; i < len(routes); i++ {
		//routes[i].HandlerFunc = Headers(routes[i].HandlerFunc)
		if _, ok := r.routes[routes[i].Pattern]; !ok {
			r.routes[routes[i].Pattern] = make(map[Method]Route)
		}
		r.routes[routes[i].Pattern][routes[i].Method] = routes[i]
	}
}

// AddRoutes ajoute des routes à la liste de routes allant être utilisé
// par le service
func (r *router) AddRoutes(routes ...Route) *router {
	r.makeRoutes(routes...)
	return r
}

// AddMiddlewares ajout d'un ou plusieurs middlewares au router
func (r *router) AddMiddlewares(middlewares ...Middleware) *router {
	r.middlewares = append(r.middlewares, middlewares...)
	return r
}

// Consumer consume les routes présentes dans la structure interne pour les
// implanter avec le package http cette méthode est finale sur la structure, les
// modifications futures faites sur l'instance de la structure router ne pourra
// être consumé une seconde fois.
func (r *router) Consumer(prefix string) {
	if r.isConsummed {
		panic("error")
	}
	r.isConsummed = true

	for pattern, methods := range r.routes {
		func(pattern string, methods map[Method]Route) {
			handler := CORS(func(w http.ResponseWriter, rq *http.Request) {
				if next, ok := methods[Method(rq.Method)]; ok {
					// s'il n'y a pas de middleware, on retourne la fonction
					if len(r.middlewares) == 0 {
						next.HandlerFunc(w, rq)
						return
					}

					// sinon on fait une construction inverse des middlewares pour faire une imbrication
					// ex: m1(m2(m3(next.HandlerFunc(w, rq))))
					// pour m1, m2 et m3 des middlewares
					wrapped := next.HandlerFunc
					for i := len(r.middlewares) - 1; i >= 0; i-- {
						wrapped = r.middlewares[i](wrapped)
					}
					wrapped(w, rq)
				}
			})

			http.HandleFunc(prefix+pattern, handler)

			for method := range methods {
				fmt.Printf("[%-7s] %s\n", method, prefix+pattern)
			}

		}(prefix+pattern, methods)
	}
	fmt.Println(">> ---------------------------------")
	fmt.Println("Logger:")
}
