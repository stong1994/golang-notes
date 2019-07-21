package main

import (
	"golang-learning/analyse/pprof/data"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	go func() {
		for {
			log.Println(data.Add("hello world"))
			time.Sleep(time.Second)
		}
	}()

	http.ListenAndServe("0.0.0.0:8080", nil)
}
