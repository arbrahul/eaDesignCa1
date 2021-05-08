package main

import (
	"fmt"
	"log"
	"net/http"
	"math/rand"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Newsfetcher: No news! Is good news ;-)")
}

func main() {

	if rand.Intn(90) < 30 {
		panic("flaky error occurred ")
    } 
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8888", nil))
}
