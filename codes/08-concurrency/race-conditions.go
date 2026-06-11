package main

import (
	"fmt"
	"sync"
)

func raceConditions() {
	fmt.Println("Demo on Race Condition...")

	waitG := &sync.WaitGroup{}
	mutex := &sync.Mutex{}

	var scores = []int{0}

	waitG.Add(3)
	go func(waitG *sync.WaitGroup, mutex *sync.Mutex) {
		fmt.Println("Goroutine 1")
		mutex.Lock()
		scores = append(scores, 1)
		mutex.Unlock()
		waitG.Done()
	}(waitG, mutex)

	go func(waitG *sync.WaitGroup, mutex *sync.Mutex) {
		fmt.Println("Goroutine 2")
		mutex.Lock()
		scores = append(scores, 2)
		mutex.Unlock()
		waitG.Done()
	}(waitG, mutex)

	go func(waitG *sync.WaitGroup, mutex *sync.Mutex) {
		fmt.Println("Goroutine 3")
		mutex.Lock()
		scores = append(scores, 3)
		mutex.Unlock()
		waitG.Done()
	}(waitG, mutex)

	waitG.Wait()
	fmt.Println("Final scores:", scores)
}
