package main

import (
	"auto-test/middleware"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Method func(string, http.Handler)

func main() {

	r := chi.NewRouter()

	//this line can be removed(and the dependency too) once the test are generated through the request
	//GET localhost:8080/auto-generate-test endpoint
	r.Use(middleware.New("D:\\develop\\Golang\\auto-test\\autoTest.go").Handle)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	r.Group(func(r chi.Router) {
		r.Post("/more/tests", func(writer http.ResponseWriter, request *http.Request) {
			writer.WriteHeader(http.StatusOK)
			writer.Write([]byte("More tests"))
		})
	})

	http.ListenAndServe(":8080", r)
}
