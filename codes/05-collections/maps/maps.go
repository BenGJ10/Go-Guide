package main

import "fmt"

func main() {

	websites := map[string]string{
		"Google":              "https://www.google.com",
		"Amazon Web Services": "https://aws.amazon.com",
	}
	fmt.Println(websites)

	// Extracting a value from the map
	fmt.Println(websites["Google"])

	// Adding a new key-value pair to the map
	websites["GitHub"] = "https://github.com"
	fmt.Println(websites)

	// Deleting a key-value pair from the map
	delete(websites, "Amazon Web Services")
	fmt.Println(websites)

	// We can use a make() function to create an empty map and then add key-value pairs to it
	courseRatings := make(map[string]float64, 3)
	courseRatings["Go Programming"] = 4.5
	courseRatings["Python Programming"] = 4.7
	courseRatings["Java Programming"] = 4.3
	fmt.Println(courseRatings)

	for key, value := range courseRatings {
		fmt.Printf("Course: %s, Rating: %.1f\n", key, value)
	}
}
