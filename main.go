package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"pkg_test/test1"
)

func main() {

	//- отвечает за логику которой можно подедиться в не проекта
	//(поделится внешне)

	test1.Start1()

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	http.ListenAndServe(":3000", r)

}
