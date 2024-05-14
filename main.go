package main

import (
	"github.com/Aleks335/pkg_test/test1"
	"github.com/go-chi/chi/v5"
	"net/http"
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
