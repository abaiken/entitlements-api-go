package server

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"cloud.redhat.com/entitlements/controllers"
	"github.com/RedHatInsights/platform-go-middlewares/identity"
)

func DoRoutes() chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(identity.Identity)

	r.Route("/api/entitlements/v1", func(r chi.Router) {
		r.Route("/", controllers.LubDub)
		r.Route("/services", controllers.Subscriptions)
	})

	return r
}
