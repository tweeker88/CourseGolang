package main

import (
	"fmt"
)

// Обычное объявление структуры
type Student struct {
	firstName, middleName string
	age                   int
}

// Структура с анонимными свойствами
type UnknownStudent struct {
	string
	int
}

// Вложенная структура
type Class struct {
	text    string
	symbol  rune
	student Student
}

func main() {
	student := Student{
		"Maxim",
		"Teplov",
		25,
	}

	unknown := UnknownStudent{
		"Undefined",
		0,
	}

	printInfo(student)

	fmt.Println("Undefined student")
	fmt.Println(unknown)

	// анонимная структура
	anon := struct {
		age  int
		text string
	}{
		age:  21,
		text: "bla bla bla",
	}

	fmt.Println(anon)

	// работа со структурой по ссылке
	renameStudentFromPointer(&student)
	fmt.Println("Student after rename", student)

	// объявление вложенной структуры
	class := Class{
		"Some class",
		'R',
		student,
	}

	fmt.Println("_____Class_____")
	fmt.Println("Class name:", class.text)
	fmt.Println("Class symbol:", class.symbol)
	fmt.Println("Student middle name in class:", class.student.middleName)

}

func printInfo(student Student) {
	fmt.Println("My name is", student.firstName)
	fmt.Println("My middleName is", student.middleName)
	fmt.Println("My age is", student.age)
}

func renameStudentFromPointer(student *Student) *Student {
	student.firstName = "New Max"
	fmt.Println("Student in function", student)

	return student
}
