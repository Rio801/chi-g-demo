package main

import (
	"chi_demo2/routes"
	"net/http"
)

func main() {
	router := routes.NewServer()

	http.ListenAndServe(":3000", router.Router)
}
