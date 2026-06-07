package main

import "fmt"

func functions() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Initial numbers: ", numbers)

	// We can pass a function as an argument to another function
	doubledNumbers := transformNumbers(numbers, double)
	fmt.Println("Numbers after doubling: ", doubledNumbers)

	// Usage of an anonymous function to triple the numbers
	tripledNumbers := transformNumbers(numbers, func(num int) int {
		return num * 3
	})
	fmt.Println("Numbers after tripling: ", tripledNumbers)
}

func transformNumbers(nums []int, transform func(int) int) []int {
	result := []int{}
	for _, num := range nums {
		result = append(result, transform(num))
	}
	return result
}

func double(num int) int {
	return num * 2
}
