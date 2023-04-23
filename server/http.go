package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func StartHttp() {

	h := mux.NewRouter()

	h.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	}).Methods(http.MethodGet)

	s := http.Server{
		Addr:    "0.0.0.0:8081",
		Handler: h,
	}
	s.ListenAndServe()

}
