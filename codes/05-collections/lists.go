package main

import "fmt"

func lists() {
	prices := [5]float64{19.99, 29.99, 39.99, 49.99, 59.99}
	for i := 0; i < len(prices); i++ {
		fmt.Printf("Price %d: $%.2f\n", i+1, prices[i])
	}

	contents := []string{"Go is great!", "I love programming.", "Go is fast and efficient."}
	fmt.Println("Slicing contents: ", contents[:3])

	// Slicing doesn't create a copy, it references the same underlying array
	selectedContents := contents[1:3]
	selectedContents[0] = "I enjoy coding."

	fmt.Println("Selected contents: ", selectedContents)
	fmt.Println("Original contents: ", contents)

}
