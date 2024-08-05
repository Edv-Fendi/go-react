package api

import (
	"github.com/rocketseat-education/semana-tech-go-react-server/internal/store/pgstore"

	"github.com/go-chi/chi/v5"
)

type apiHandler struct {
	q *pgstore.Queries
	r *chi.Mux
}

func (h apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	h.r.ServeHTTP(w, r)

}

func NewHandler(q *pgstore.Queries) hpp.Handler {
	a := apiHandler{
		q: q,
	}
	r := chi.newRouter()
	a.r = r
	return a
}