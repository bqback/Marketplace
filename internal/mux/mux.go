package mux

import (
	"net/http"

	"marketplace/internal/auth"
	"marketplace/internal/config"
	"marketplace/internal/handlers"
	"marketplace/internal/logging"
	"marketplace/internal/mux/middleware"

	_ "marketplace/docs"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

func SetupMux(handlers *handlers.Handlers, config *config.Config, logger *logging.LogrusLogger) http.Handler {

	mux := chi.NewRouter()

	baseUrl := config.API.BaseUrl

	manager := auth.NewManager(config.JWT)
	authMW := middleware.NewAuthMiddleware(manager)

	mux.Use(middleware.RequestID, middleware.NewLogger(logger), middleware.PanicRecovery, middleware.JsonHeader)

	mux.Route(baseUrl, func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			r.Post("/login/", handlers.AuthHandler.Login)
			r.Post("/signup/", handlers.AuthHandler.Signup)
		})
		r.Route("/listing", func(r chi.Router) {
			r.Use(authMW.Auth)
			r.Post("/", handlers.ListingHandler.Create)
		})
		r.Route("/feed", func(r chi.Router) {
			r.Use(authMW.OptionalAuth)
			r.Use(middleware.ExtractFeedParams)
			r.Get("/?sortby={sortby}&order={order}&page={page}&per={perpage}&minprice={minprice}&maxprice={maxprice}", handlers.ListingHandler.Feed)
		})
	})
	mux.Route("/swagger/", func(r chi.Router) {
		r.Get("/*", httpSwagger.Handler(
			httpSwagger.URL("swagger/doc.json"),
		))
	})

	return mux
}
