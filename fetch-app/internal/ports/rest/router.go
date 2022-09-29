package rest

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Route("/fetch-app", func(r chi.Router) {
		r.Get("/ping", func(rw http.ResponseWriter, r *http.Request) {
			_, err := rw.Write([]byte("Hello, World!"))
			if err != nil {
				log.Println(err)
			}
		})
	})
	return router
}
