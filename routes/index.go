package routes

import (
	models "chi_demo2/types"
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func NewServer() *models.Server {
	server := &models.Server{
		Router: chi.NewRouter(),
	}
	server.Router.Use(middleware.Logger)
	server.Router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://localhost:3000", "http://localhost:3000"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	server.Router.Handle("/js/*", http.StripPrefix("/js", http.FileServer(http.Dir("./tmpls/js"))))

	server.Router.Group(func(r chi.Router) {
		r.Get("/", indexHandler)
	})

	return server
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	msg := make(map[string]string, 0)

	msg["Message"] = "Welcome to the index page"

	idxTmpl, err := template.ParseFiles("tmpls/base.html", "tmpls/index.html")

	if err != nil {
		http.Error(w, " Internal Server Error", http.StatusInternalServerError)
	}
	idxTmpl.Execute(w, msg)

}
