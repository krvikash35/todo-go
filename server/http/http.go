package http

import (
	"log"
	"net/http"
	"todo/postgres"

	"github.com/gorilla/mux"
)

func StartServer() {

	db, err := postgres.New()
	if err != nil {
		log.Fatal("http.startServer: could not initialize postgres due to error ", err)
	}

	h := mux.NewRouter()

	h.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	}).Methods(http.MethodGet)

	h.HandleFunc("/todo", func(w http.ResponseWriter, r *http.Request) {
		db.GetTodos()
	}).Methods(http.MethodPost)

	h.HandleFunc("/todo", func(w http.ResponseWriter, r *http.Request) {
		db.CreateTodo("sample")
	}).Methods(http.MethodPost)

	s := http.Server{
		Addr:    "0.0.0.0:8081",
		Handler: h,
	}
	s.ListenAndServe()

}
