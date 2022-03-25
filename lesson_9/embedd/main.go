package main

import "fmt"

// наследование интерфейсов

type Student struct {
	FirstName, MiddleName string
}

type Worker interface {
	Work()
}

type Studer interface {
	Stude()
}

type Studenter interface {
	Worker
	Studer
}

func (student *Student) Work() {
	fmt.Println("Yes, i work")
}

func (studend *Student) Stude() {
	fmt.Println("Yes, i stude")
}

func main() {

	// Инициализируем переменные с типом интерфейсов
	var studer Studer
	var worker Worker
	var studenter Studenter

	student := Student{"Maxim", "Teplov"}

	studer = &student
	worker = &student
	studenter = &student

	studer.Stude()
	worker.Work()

	studenter.Stude()
	studenter.Work()

}
