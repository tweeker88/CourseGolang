package pointer

import "fmt"

const (
	PORT = 8000
)

func main() {
	fmt.Println("Port:", PORT)

	// Указатели

	var varibale = 3
	var pointer *int = &varibale // инициализируем ссылку на переменную (хранится адрес)

	fmt.Printf("Variable is %v and type is %T \n", varibale, varibale)
	fmt.Printf("Pointer is %v and type is %T \n", pointer, pointer)

	// получение значение из переменной с ссылкой
	fmt.Println("Get value from pointer", *pointer)

	// создание поинтера на 0
	zeroValuePointer := new(int)
	fmt.Println("value in zeroValuePointer after:", *zeroValuePointer)
	*zeroValuePointer += 50
	fmt.Println("value in zeroValuePointer before:", *zeroValuePointer)

	pointerFromFunc := returnPointer()
	fmt.Println("Get pointer from function")
	fmt.Printf("Place in memory: %v Value is: %v and type is %T \n", pointerFromFunc, *pointerFromFunc, pointerFromFunc)

	// одно и тоже значение, так как мы обращаемся к одному и тому же месту в памяти
	pointerToFunc := incrementWithPointer(pointerFromFunc)
	fmt.Printf("Value is: %v and type is %T \n", pointerToFunc, pointerFromFunc)
	fmt.Printf("Place in memory: %v Value is: %v and type is %T \n", pointerFromFunc, *pointerFromFunc, pointerFromFunc)

}

// Определение функции, которая возвращает указатель
func returnPointer() *int {
	var value int = 100

	return &value
}

// Определение функции, которая принимает указатель
func incrementWithPointer(pointer *int) int {
	*pointer++

	return *pointer
}
