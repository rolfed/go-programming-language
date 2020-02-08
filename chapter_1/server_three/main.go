package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
) 

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// Handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "%s %s %s/n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprint(w, "Header[%q] = %q/n", k, v)
	}
	fmt.Fprint(w, "Host = %q/n", r.Host)
	fmt.Fprint(w, "RemoteAddr = %q/n", r.RemoteAddr)

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q/n", k, v)
	}

}

// Counter echoes the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d/n", count)
	mu.Unlock()
}
