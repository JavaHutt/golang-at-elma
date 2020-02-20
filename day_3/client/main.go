package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

const (
	concurrency = 80
	requests    = 500
)

func doWork() string {
	res, err := http.Post("http://localhost:1234/hi", "text/plain", bytes.NewBufferString("test"))
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(body)
}

func main() {
	works := make(chan struct{})
	wg := new(sync.WaitGroup)

	for i := 0; i < concurrency; i++ {
		wg.Add(concurrency)
		go func(j int) {
			for range works {
				fmt.Println("work on goroutine", j, "ends with response", doWork())
			}
			wg.Done()
		}(i)
	}
	for i := 0; i < requests; i++ {
		works <- struct{}{}
	}
	close(works)
	wg.Wait()
}
