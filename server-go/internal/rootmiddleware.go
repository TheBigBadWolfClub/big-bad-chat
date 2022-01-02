package internal

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"time"
)

type RouteRegister interface {
	register(route ChatRoute) *HTTPRouter
}

type HTTPRouter struct {
	Mux *chi.Mux
}

func NewHTTPRouter() *HTTPRouter {
	router := HTTPRouter{
		Mux: chi.NewRouter(),
	}
	return router.middleware().register(&WolfChatHandler{})
}

func (h *HTTPRouter) register(route ChatRoute) *HTTPRouter {
	h.Mux.Get(route.Path(), route.Handler)
	return h
}

func (h *HTTPRouter) middleware() *HTTPRouter {
	// A good base middleware stack
	h.Mux.Use(middleware.RequestID)
	h.Mux.Use(middleware.RealIP)
	h.Mux.Use(middleware.Logger)
	h.Mux.Use(middleware.Recoverer)
	h.Mux.Use(middleware.Timeout(3 * time.Minute))
	// https://github.com/go-chi/cors

	h.Mux.Use(allowCors)
	h.Mux.Use(MethodOptions)
	h.Mux.Use(startMiddleware)
	h.Mux.Use(endMiddleware)

	return h
}

func MethodOptions(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	})
}

func allowCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=ascii")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")

		next.ServeHTTP(w, r)
	})
}

func startMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "startTime", time.Now())
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
func endMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		value := r.Context().Value("startTime")
		diff := time.Now().Sub(value.(time.Time)).Microseconds()
		fmt.Println(diff, " Micros")
	})
}
