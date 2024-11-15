package application

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (a *App) loadRoutes() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/users", func(router chi.Router) {
		router.Handle("/*", a.forwardRequest(a.config.UserServiceAddress))
	})

	router.Route("/channels", func(router chi.Router) {
		router.Handle("/*", a.forwardRequest(a.config.ChannelServiceAddress))
	})

	router.Route("/messages", func(router chi.Router) {
		router.Handle("/*", a.forwardRequest(a.config.MessageServiceAddress))
	})

	router.Route("/typing", func(router chi.Router) {
		router.Handle("/*", a.forwardRequest(a.config.LiveTypingServiceAddress))
	})

	a.router = router
}

func (a *App) forwardRequest(serviceURL string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Forwarding request: %s %s\n", r.Method, r.URL.Path)

		serviceURLParsed, err := url.Parse(serviceURL)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to parse service URL: %v", err), http.StatusInternalServerError)
			return
		}

		proxy := httputil.NewSingleHostReverseProxy(serviceURLParsed)
		proxy.ServeHTTP(w, r)
	}
}
