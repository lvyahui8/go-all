package main

import (
	"fmt"
	"log"
	"net/http"
)

type Hello struct {
}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("recv a request")
	fmt.Fprint(w, "Hello!")
}

func main() {
	var h Hello
	// curl 192.168.0.100:4000
	err := http.ListenAndServe("0.0.0.0:4000", h)

	if err != nil {
		log.Fatal(err)
	}
}
