package main

import (
	"fmt"
)

type Student struct {
	firstName, middleName string
	classRoom             ClassRoom
	averageRating         int
}

type ClassRoom struct {
	symbol rune
}

type MySuperInt int

func (student *Student) replaceNameWithPointer(newName string) {
	student.firstName = newName
}

func (student Student) replaceNameWithoutPointer(newName string) {
	student.firstName = newName
}

func (classRoom *ClassRoom) WhereClassRoom() {
	fmt.Println("Your class with symbol:", string(classRoom.symbol))
}

func (myInt *MySuperInt) isEven() bool {
	return *myInt%2 == 0
}

func NewStudent(name string, middleName string, classRoom ClassRoom) *Student {
	newStudent := new(Student)

	newStudent.firstName = name
	newStudent.middleName = middleName
	newStudent.classRoom = classRoom
	newStudent.averageRating = 5

	return newStudent
}

type WorkInterface interface {
	CalcAverageRating() int
}

func (student *Student) CalcAverageRating() int {
	return student.averageRating / 2 * 3
}

func main() {
	student := Student{"Maxtim", "Teplov", ClassRoom{'B'}, 3}

	fmt.Println("Before call function", student)

	student.replaceNameWithoutPointer("New name") // у значения стркутура пытаемся вызвать метод со значением
	// изменений не будет, так как происходит локальное копирование
	fmt.Println("After call function without pointer", student)

	student.replaceNameWithPointer("Oleg") // у значения структуры пытаемся вызвать метод с поинтером
	// изменения будут
	fmt.Println("After call function with pointer", student)

	student.classRoom.WhereClassRoom()

	(&student).replaceNameWithoutPointer("New name") // у поинтера мы пытаемся вызвать метод со значением
	// ничего не поменяется
	fmt.Println("After call function without pointer", student)

	// собственный тип с новым методом
	var myInt MySuperInt = 10
	fmt.Println("Result function:", myInt.isEven())

	// создание структуры через конструктор
	classRoom := ClassRoom{'A'}
	newStudent := NewStudent("Ivan", "Ivanov", classRoom)
	fmt.Println("New student:", newStudent)

	// Вызов методов интерфейса у студента
	fmt.Println("First student:", student.CalcAverageRating(), "Second student:", student.CalcAverageRating())
	students := []WorkInterface{&student, newStudent}
	for id, student := range students {
		fmt.Printf("avarega rating for student №%v = %v \n", id, student.CalcAverageRating())
	}
}
