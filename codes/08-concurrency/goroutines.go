package main

import (
	"fmt"
	"time"
)

func greet(str string, done chan bool) {
	fmt.Println(str)
	done <- true
}

func slowGreet(str string, done chan bool) {
	time.Sleep(2 * time.Second)
	fmt.Println("Last greet:", str)
	done <- true
	close(done)
}

func goroutineSample() {
	done := make(chan bool)
	go greet("Hello, World!", done)
	go slowGreet("How are you?", done)
	go greet("My name is Go.", done)

	for range done {
		// Wait for all goroutines to finish
	}
}
