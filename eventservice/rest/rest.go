package rest

import (
	"github.com/gorilla/mux"
	"net/http"
)

func serveAPI(endpoint string) error{

	handler := eventServiceHandler{}

	r := mux.NewRouter()
	eventsrouter := r.PathPrefix("/events").Subrouter()

	eventsrouter.Methods("GET").Path("/{SearchCriteria}/{search}").
		HandlerFunc(handler.FindEventHandler)
	eventsrouter.Methods(("GET")).Path("").
		HandlerFunc(handler.AllEventHandler)
	eventsrouter.Methods("POST").Path("").
		HandlerFunc(handler.NewEventHandler)

	return http.ListenAndServe(endpoint, r)
}