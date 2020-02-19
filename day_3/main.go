package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type service struct{}

func hello(w http.ResponseWriter, r *http.Request) {
	sleep := rand.Intn(300)
	time.Sleep(time.Duration(sleep) * time.Millisecond)
	fmt.Fprintf(w, "%d ms", sleep)
}

func main() {
	http.HandleFunc("/hi", hello)
	fmt.Println("start listen")
	http.ListenAndServe(":1234", nil)
}
