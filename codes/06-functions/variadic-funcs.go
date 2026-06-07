package main

func main() {
	result := sum(1, 2, 3, 4, 5)
	println("The sum is:", result)

	sentence := concat("Hello", "world!", "How", "are", "you?")
	println("Concatenated string:", sentence)
}

func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func concat(strings ...string) string {
	result := ""
	for _, str := range strings {
		result += str
		if str != strings[len(strings)-1] {
			result += " " // Add a space between words, but not after the last one
		}
	}
	return result
}
