package main

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
	"log"
)

var json = jsoniter.ConfigFastest

func main() {

	log.Println("Starting server on 80")
	if err := fasthttp.ListenAndServe(":80", requestHandler); err != nil {
		log.Fatalf("Error in server: %s", err)
	}
}

