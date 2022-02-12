package main // package name

import "fmt" // импорт пакета

func main() {
	fmt.Println("Hello world!")         // вывод строки с переносом
	fmt.Print("Hello second world! \n") // вывод строки без переноса

	var a int = 10         // иниациализации переменной со значением
	var b int              // декларирование переменной без значения (по дефолту 0)
	var c, d = 5, "string" // инициализация неск.переменных

	fmt.Printf("a = %d, b = %d, c = %d, d = %s \n", a, b, c, d)

	e, f := 1, 2 // короткая инициализация
	// e, f := 3, 3 не будет работать
	fmt.Printf("Short e = %d, f = %d \n", e, f)
	p, e, f := 4, 4, 4 // будет, так как слева новая переменная

	fmt.Printf("New short e = %d, f = %d p = %d \n", e, f, p)

	var (
		personName string = "Max"
		personAge  int    = 25
	)

	fmt.Printf("My name is %s \nMy I'm %d \n", personName, personAge)

	fmt.Scan(&personName, &personAge) // ввод из терминала
	// fmt.Fscan(os.Stdin, &age)
	fmt.Printf("My new name is %s \nMy I'm %d", personName, personAge)
}
