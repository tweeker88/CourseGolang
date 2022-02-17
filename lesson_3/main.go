package main

import (
	"fmt"
	"strings"
)

func main() {
	// обычный цикл
	for i := 0; i <= 10; i++ {
		fmt.Println("Current value =", i)
	}

	// цикл только с условием
	var i int = 0
	for i <= 10 {
		fmt.Println("Current value:", i)
		i++
	}
	fmt.Printf("Value i = %d before cycle \n", i)

	// бесконечный цикл
	var password string
	for {
		fmt.Println("Insert password:")
		fmt.Scan(&password)
		if strings.Contains(password, "abc") {
			fmt.Println("Soryy! your password to simple")
			continue
		}
		fmt.Println("great!")
		break
	}

	//цикл с множественными переменными цикла
	for x, y := 0, 1; x <= 10 && y <= 12; x, y = x+1, y+2 {
		fmt.Printf("%d + %d = %d\n", x, y, x+y)
	}

	var name string
	fmt.Scan(&name)

	switch name {
	case "Maxim":
		fmt.Println("Hello me")
	case "Nick":
		fmt.Println("Hello my friend")
	case "Anna":
		fmt.Println("Mrrrr")
	default:
		fmt.Println("Who are you?")
	}

	switch name {
	case "Maxim":
		fmt.Println("Hello me and continue")
		fallthrough
	case "Nick":
		fmt.Println("It is last")
	case "Anna":
		fmt.Println("Mrrrr")
	default:
		fmt.Println("Who are you?")
	}
}
