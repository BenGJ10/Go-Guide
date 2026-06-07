package main

import "fmt"

func main() {

	nullSlice := []int{}
	fmt.Println("Null slice:", nullSlice)
	fmt.Println("Is the slice nil?", nullSlice == nil)

	// Creating a slice with a specific length and capacity
	dynamicSlice := make([]int, 5, 10)
	fmt.Println("\nDynamic slice:", dynamicSlice)
	fmt.Println("Length:", len(dynamicSlice))
	fmt.Println("Capacity:", cap(dynamicSlice))

	// Adding elements to the slice
	dynamicSlice = append(dynamicSlice, 6, 7, 8)
	fmt.Println("\nAfter appending elements:", dynamicSlice)
	fmt.Println("Length after appending:", len(dynamicSlice))
	fmt.Println("Capacity after appending:", cap(dynamicSlice))

	// Still the initial 5 elements are zero values, and the new elements are added after them
	// Let's override the initial elements to see the effect clearly
	for i := 0; i < 5; i++ {
		dynamicSlice[i] = i + 1
	}

	fmt.Println("\nAfter overriding initial elements:", dynamicSlice)
}
