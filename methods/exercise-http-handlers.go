package main

import (
	"fmt"
	"log"
	"net/http"
)

type String string
type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}

func (s *Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>%s %s %s</h1>", s.Greeting, s.Punct, s.Who)
}

func main() {
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"hello", ":", "Grophers!"})
	log.Fatal(http.ListenAndServe("0.0.0.0:4000", nil))
}
