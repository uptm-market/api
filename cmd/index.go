package main

import "net/http"

type getIndexHandlerFunc func(w http.ResponseWriter, r *http.Request)

func (f getIndexHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(w, r)
}
