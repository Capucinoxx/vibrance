package client

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
		{Method: http.MethodPost, Pattern: "/client_create", HandlerFunc: h.create()},
		{Method: http.MethodPost, Pattern: "/client_find", HandlerFunc: h.find()},
		{Method: http.MethodPost, Pattern: "/client_updateSecret", HandlerFunc: h.updateSecret()},
		{Method: http.MethodPost, Pattern: "/client_softDelete", HandlerFunc: h.softDelete()},
		{Method: http.MethodPost, Pattern: "/client_delete", HandlerFunc: h.delete()},
	}

	r := router.Consumer("/oauth", routes, nil)

	return r
}

func (h handler) create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("route client_create is not implemented"))
	}
}

func (h handler) find() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("route client_find is not implemented"))
	}
}

func (h handler) updateSecret() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("route client_updateSecret is not implemented"))
	}
}

func (h handler) softDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("route client_softDelete is not implemented"))
	}
}

func (h handler) delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("route client_delete is not implemented"))
	}
}
