package main

import "fmt"

func moreSlices() {

	initialMarks := []int{85, 90, 78, 88}
	fmt.Println("Initial marks:", initialMarks)

	finalMarks := make([]int, len(initialMarks))
	copy(finalMarks, initialMarks)
	finalMarks = append(finalMarks, 92)

	fmt.Println("Final marks:", finalMarks)
	fmt.Println("Capacity of final marks:", cap(finalMarks))

	finalMarks = make([]int, len(finalMarks), len(finalMarks)*2)
	fmt.Println("Capacity of final marks after resizing:", cap(finalMarks))

	averageMarks := calculateAverage(finalMarks...)
	fmt.Println("Average marks:", averageMarks)

	fmt.Print("Enter numbers (press Ctrl+D to stop): ")
	numbers := getNNumbersAsInput()
	fmt.Println("Numbers entered:", numbers)
}

func calculateAverage(marks ...int) float64 {
	if len(marks) == 0 {
		return 0
	}
	total := 0
	for _, mark := range marks {
		total += mark
	}
	return float64(total) / float64(len(marks))
}

// Function to get N numbers as input (till a break condition is met) and return them as a slice
func getNNumbersAsInput() []int {
	numbers := []int{}
	for {
		var num int
		_, err := fmt.Scan(&num)
		if err != nil {
			break
		}
		numbers = append(numbers, num)
	}
	return numbers
}
