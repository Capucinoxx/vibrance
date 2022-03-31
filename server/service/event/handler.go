package event

import (
	"encoding/json"
	"net/http"

	"github.com/Capucinoxx/vibrance/server/model"
	"github.com/Capucinoxx/vibrance/server/utils/response"
	"github.com/Capucinoxx/vibrance/server/utils/router"
)

type handler struct {
	r Repository
}

func Handle() router.Routes {
	h := handler{r: NewRepository()}

	return router.Routes{
		{Method: http.MethodPost, Pattern: "/events", HandlerFunc: h.listEvent()},
	}
}

func (h handler) listEvent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		defer func() {
			if err := r.Body.Close(); err != nil {
				response.Write(w, http.StatusBadRequest, nil)
				return
			}
		}()

		var in model.ListEventsInput
		if err := decoder.Decode(&in); err != nil {
			response.Write(w, http.StatusBadRequest, nil)
			return
		}

		events, err := h.r.Events(in.From, in.To)
		if err != nil {
			response.Write(w, http.StatusBadRequest, err)
			return
		}

		response.Write(w, http.StatusOK, response.Response{Data: events})
	}
}
