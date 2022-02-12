package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
	"unsafe"
)

func main() {
	var (
		truth bool = true
		lying bool = false
	)

	fmt.Println("AND: ", truth && lying)
	fmt.Println("OR: ", truth || lying)
	fmt.Println("NOT: ", !truth)
	///////////////////////////////////////////////

	var (
		one  int   = 24
		two  uint  = 34
		thee int32 = 32
	)

	fmt.Println("One:", one, "Two:", two, "Three:", thee, "Sum:", one+int(two)+int(thee))
	// Происходит приведение типов так как мы не можем сложить разный по разрядности int
	// Наш int 64 разрядности, НО нельзя int + int64

	//Вывод типа через %T форматирование
	fmt.Printf("Type is %T\n", one)
	//Узнаем, сколько байт занимает переемнная типа int
	fmt.Printf("Type %T size of %d bytes\n", one, unsafe.Sizeof(one))

	///////////////////////////////////////////////
	var (
		four float32 = 32
		five float64 = 64
	)

	fmt.Println("Four:", four, "Five:", five)

	///////////////////////////////////////////////
	c1 := complex(5, 7)
	c2 := 12 + 32i
	fmt.Println(c1 + c2)
	fmt.Println(c1 * c2)

	///////////////////////////////////////////////
	var (
		name       string = "Максим"
		middleName string = "Теплов"
	)

	fmt.Println("Name:", name, "Middle:", middleName, "Full:", name+" "+middleName)
	fmt.Println("Length of string :", name, len(name)) // Функция len() возвращает количество элементов в наборе
	fmt.Println("Amount of chars:", name, utf8.RuneCountInString(name))

	var (
		totalString string = "My name is Maxim"
		subString   string = "Maxim"
	)

	fmt.Print(strings.Contains(totalString, subString), "\n")

	///////////////////////////////////////////////
	var simpleRune rune = 'Q'
	var numberRune rune = 234
	fmt.Println("Simple rune", simpleRune)
	fmt.Println("Number rune", numberRune)

	fmt.Println(strings.Compare(totalString, subString)) // 1 or 0 or -1

	///////////////////////////////////////////////
	var aByte byte = 8
	fmt.Println("Byte:", aByte)

	///////////////////////////////////////////////
	var userAge int
	fmt.Scan(&userAge)
	if userAge < 18 {
		fmt.Println("Малолетка, какое тебе пиво")
	} else {
		fmt.Println("Добрый вечерок")
	}

	var height int = 100

	if name := "Квадрат"; height > 40 {
		fmt.Println("Ты не", name)

		return
	}

	fmt.Println("Ладно, ты ", name)
}
