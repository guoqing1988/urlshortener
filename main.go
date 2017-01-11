package main

import (
	"log"
	"net/http"
	"runtime"
	"urlshortener/handler"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	http.HandleFunc("/shorten", handler.ShortenHandler)
	http.HandleFunc("/", handler.ExpandHandler)

	err := http.ListenAndServe(handler.ServerHost + ":" + handler.ServerPort, nil)
	if err != nil {
		log.Fatal(err)
	}
}
