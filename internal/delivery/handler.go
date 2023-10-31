package delivery

import (
	"net/http"
	_ "refactoring/docs"
	v1 "refactoring/internal/delivery/http/v1"
	"refactoring/internal/service"
	"time"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{services: services}
}

func (h *Handler) Init() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID,
		middleware.RealIP,
		middleware.Logger,
		middleware.Timeout(60*time.Second),
	)
	//r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(time.Now().String()))
	})

	r.Get("/swagger/*", httpSwagger.WrapHandler)

	h.initAPI(r)

	return r
}

func (h *Handler) initAPI(router *chi.Mux) {
	handlerV1 := v1.NewHandler(h.services)
	router.Route("/api", func(r chi.Router) {
		handlerV1.Init(r)
	})
}
