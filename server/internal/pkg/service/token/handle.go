package token

import (
	"net/http"
	"time"

	"github.com/Capucinoxx/vibrance/internal/pkg/common/router"
	"github.com/gocql/gocql"
)

type handler struct {
	repo Repository
}

func Handle(conn *gocql.Session, timeout time.Duration) router.Router {
	h := handler{repo: NewRepository(conn, timeout)}

	routes := router.Routes{
		{Method: http.MethodPost, Pattern: "/token_create", HandlerFunc: h.create()},
		{Method: http.MethodPost, Pattern: "/token_refresh", HandlerFunc: h.refresh()},
		{Method: http.MethodPost, Pattern: "/token_revoke", HandlerFunc: h.revoke()},
		{Method: http.MethodPost, Pattern: "/token_find", HandlerFunc: h.find()},
	}

	r := router.Consumer("/oauth", routes, nil)

	return r
}

func (h handler) create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("route token_create is not implemented"))
	}
}

func (h handler) refresh() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("route token_refresh is not implemented"))
	}
}

func (h handler) revoke() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("route token_revoke is not implemented"))
	}
}

func (h handler) find() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("route token_find is not implemented"))
	}
}
