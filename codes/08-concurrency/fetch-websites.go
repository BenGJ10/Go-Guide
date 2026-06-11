package main

import (
	"fmt"
	"net/http"
	"sync"
)

var waitG sync.WaitGroup // pointer to a WaitGroup struct
var mutex sync.Mutex     // pointer to a Mutex struct

var signal = []string{"https://amazon.com"}

func main() {
	websites := []string{
		"https://www.google.com",
		"https://www.github.com",
		"https://bengregoryjohn.in",
	}

	for _, website := range websites {
		// Increment the WaitGroup counter for each goroutine
		go getStatusCode(website)
		waitG.Add(1)
	}

	// Wait for all the goroutines to finish
	waitG.Wait()
}

func getStatusCode(endpoint string) {
	// Wait for the goroutine to finish before decrementing the WaitGroup counter
	defer waitG.Done()

	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("Could not fetch the endpoint!")
	} else {
		mutex.Lock() // Lock the mutex before accessing the shared resource
		signal = append(signal, endpoint)
		mutex.Unlock() // Unlock the mutex after accessing the shared resource

		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
}
