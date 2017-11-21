package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shigwata/gae-go-sample/hello"
)

func init() {
	r := mux.NewRouter()
	r.HandleFunc("/", hello.SayHello)
	http.Handle("/", r)
}
