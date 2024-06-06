package main

import (
	"chi_test_second/postgres"
	"chi_test_second/web"
	"log"
	"net/http"
)

func main() {
	store, err := postgres.NewStore("postgres://postgres:" +
		"12345@localhost/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	h := web.NewHandler(store)
	http.ListenAndServe(":8080", h)
}
