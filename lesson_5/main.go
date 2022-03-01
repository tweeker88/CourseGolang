package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	word := "Тестовая строка"

	fmt.Printf("Bytes: ")
	for i := 0; i < len(word); i++ {
		fmt.Printf("%x ", word[i]) //%x - формат представления 16-ти ричного байта
	}
	fmt.Println()

	fmt.Printf("Characters: ")
	for i := 0; i < len(word); i++ {
		fmt.Printf("%c", word[i]) //%c - формат представления символа
	}
	fmt.Println()

	runeSlice := []rune(word)
	fmt.Printf("Runes: ")
	for i := 0; i < len(runeSlice); i++ {
		fmt.Printf("%c", runeSlice[i])
	}
	fmt.Println()

	for key, rune := range word {
		fmt.Printf("%c starts at postion %d\n", rune, key)
	}
	fmt.Println()

	myBytes := []byte{0x40, 0x41, 0x42, 0x43}
	myStr := string(myBytes)
	fmt.Println(myStr)

	// Длина и емкость строки
	engWord := "One"
	rusWord := "Три"

	fmt.Printf("eng string lenght: %d \n", len(engWord))                         // будет 3, так как 1 символ занимает 1 ячейку
	fmt.Printf("rus string lenght: %d \n", len(rusWord))                         // будет 6, так как 1 символ занимает 2 ячейки
	fmt.Printf("rus string real lenght: %d \n", utf8.RuneCountInString(rusWord)) // будет 3, так как мы тут считаем Руны

	myFullName := "Teplov Maxim"

	// myFullName[0] = "G"  --- не будет работать
	myFullNameInSlice := []rune(myFullName)
	myFullNameInSlice[0] = 'G'

	fmt.Printf("my fullName is: %s \n", string(myFullNameInSlice))

	// Мапы
	mapper := make(map[string]int)
	fmt.Println("Empty map:", mapper)

	var nilMap map[string]int // ТАК НЕЛЬЗЯ ОБЪЯВЛЯТЬ
	if nilMap == nil {
		fmt.Println("Zero value map")
	}

	mapper["Maxim"] = 25
	fmt.Println("Map with first value:", mapper)

	newMap := map[string]int{
		"Maxim": 100,
		"Oleg":  200,
	}
	fmt.Println("Map with first value:", newMap)

	employee := map[string]int{
		"Max":  0,
		"Oleg": 0,
	}

	// employee2 := map[string]int{
	// 	"Max":  0,
	// 	"Oleg": 0,
	// }

	// if employee == employee2 {
	// 	fmt.Println("yes")
	// }							 НЕЛЬЗЯ СРАВНИВАТЬ, ТАК КАК ССЫЛОЧНЫЙ ТИП

	if value, ok := employee["Max"]; ok {
		fmt.Println("Yes, Max is employee, his value: ", value)
	}
}
