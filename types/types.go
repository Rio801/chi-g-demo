package models

import "github.com/go-chi/chi/v5"

type Server struct {
	Router *chi.Mux
}
