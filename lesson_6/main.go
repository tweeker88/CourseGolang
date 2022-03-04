package main

import "fmt"

func main() {
	sum := calc(2, 4)
	fmt.Println("Sym: ", sum)

	words := []string{"Hello", "my", "name", "is", "Maxim"}

	fullString := cancat(words...)

	fmt.Println(fullString)
}

func calc(a, b int) int {
	return a + b
}

func cancat(words ...string) string {
	var result string
	for _, word := range words {
		result += word + " "
	}

	return result
}
