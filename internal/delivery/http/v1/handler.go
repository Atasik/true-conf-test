package v1

import (
	"refactoring/internal/service"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{services: services}
}

func (h *Handler) Init(api chi.Router) {
	api.Route("/v1", func(r chi.Router) {
		h.initUserRoutes(r)
	})
}
