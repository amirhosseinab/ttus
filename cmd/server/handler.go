package main

import (
	"net/http"
	"strings"
)

type handler struct{}

func newHandler() *handler {
	return &handler{}
}

func (handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	uri := strings.Split(strings.TrimPrefix(r.RequestURI, "/"), "/")
	if len(uri) != 1 {
		http.NotFound(w, r)
		return
	}
	if uri[0] == "" {
		root(w, r)
		return
	} else {
		redirectTo(w, r)
		return
	}
	http.NotFound(w, r)
}

func root(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hi"))
}
func redirectTo(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}
