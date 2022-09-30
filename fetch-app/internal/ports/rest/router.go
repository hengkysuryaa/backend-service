package rest

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/hengkysuryaa/backend-service/fetch-app/internal/ports/rest/middleware"
)

func NewRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Route("/fetch-app", func(r chi.Router) {
		r.Get("/ping", middleware.AuthorizeAll(
			func(rw http.ResponseWriter, r *http.Request) {
				_, err := rw.Write([]byte("Hello, World!"))
				if err != nil {
					log.Println(err)
				}
			}),
		)
	})
	return router
}
