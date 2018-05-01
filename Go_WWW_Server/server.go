package main

import (
	"net/http"
	"fmt"
	"log"
	"sync"
)

var count int
var mu sync.Mutex

func main() {
	http.HandleFunc("/", handler) //every request on this path calls handler method
	http.HandleFunc("/counter", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(writer http.ResponseWriter, request *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(writer, "%s %s %s\n", request.Method, request.URL, request.Proto)
	for k, v := range request.Header {
		fmt.Fprintf(writer, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(writer, "Host = %q\n", request.Host)
	fmt.Fprintf(writer, "RemoteAddr = %q\n", request.RemoteAddr)
	if err := request.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range request.Form {
		fmt.Fprintf(writer, "Form[%q] = %q\n", k, v)
	}
}

func counter(writer http.ResponseWriter, request *http.Request) {
	mu.Lock()
	fmt.Fprintf(writer, "Number of requests: %d\n", count)
	mu.Unlock()
}
