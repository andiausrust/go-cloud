package rest

import "net/http"

type eventServiceHandler struct {}

func (eh *eventServiceHandler) FindEventHandler(w http.ResponseWriter, r *http.Request) {}
func (eh *eventServiceHandler) AllEventHandler(w http.ResponseWriter, r *http.Request) {}
func (eh *eventServiceHandler) NewEventHandler(w http.ResponseWriter, r *http.Request) {}
