package controller

import (
	"github.com/go-chi/chi"
	"net/http"
)


func MountHealthCheckRouters(r *chi.Mux) {
	r.Get("/health-check", handlerHealthCheck)
}

func handlerHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}
