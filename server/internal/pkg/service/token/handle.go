package token

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Capucinoxx/vibrance/internal/pkg/common/oauth"
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

type CreateInput struct {
	Key string `json:"key"`
}

func (h handler) create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		defer func() {
			if err := r.Body.Close(); err != nil {
				fmt.Print(err)
				return
			}
		}()

		var in CreateInput
		if err := decoder.Decode(&in); err != nil {
			fmt.Print(err)
			return
		}

		accessToken := oauth.GenerateToken(in.Key, oauth.TokenTTLAccess, nil)
		refreshToken := oauth.GenerateToken(in.Key, oauth.TokenTTLRefresh, nil)

		if err := h.repo.Create(r.Context(), accessToken, refreshToken); err != nil {
			fmt.Printf("create token: %s", err)
			return
		}

		fmt.Println("create token: ok")
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
