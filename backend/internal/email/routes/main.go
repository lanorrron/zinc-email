package routes

import (
	"awesomeProject/internal/email/handler"
	"github.com/go-chi/chi/v5"
)

func InitializeMailRoutes(r *chi.Mux, handler *handler.EmailHandler) {

	r.Post("/search", handler.SearchEmailInZinc)
	r.Get("/list", handler.ListIndex)
	r.Delete("/{index_name}", handler.DeleteIndex)

}
