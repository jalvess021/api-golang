package web

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/jalvess021/kartka/api/internal/infra/web/routes"
)

func SetupRoutes(r *chi.Mux, db *sql.DB) {
	r.Route("/api", func(r chi.Router) {
		//Subrotas de produtos
		r.Mount("/products", routes.SetupProductRoutes(db))
	})
}
