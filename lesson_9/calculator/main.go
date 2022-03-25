package main

import (
	"calculator/rectangle"
	"errors"
	"fmt"
	"runtime/debug"
)

func init() {
	fmt.Println("Hello! I'm init function in main package")
	fmt.Printf("My vars: int - %v, string - %s \n", testInteger, testString)
}

var (
	testString  string = "Some string"
	testInteger int8   = 127
)

func PrintFinishInfo(variable int) {
	fmt.Println("Func is done", variable)
}

func FuncWithError(myInt int) (string, error) {
	if myInt%2 == 0 {
		return "Even", nil
	}

	return "Odd", errors.New("Error")
}

func FuncWithPanic() {
	defer PanicRecover()
	panic("AAA I'm panic!!")
}

func PanicRecover() {
	if r := recover(); r != nil {
		fmt.Println("Panic satisfied:", r)
		debug.PrintStack()
	}
}

func main() {
	var variableForPrint = 12

	defer PrintFinishInfo(variableForPrint)
	variableForPrint++
	defer PrintFinishInfo(variableForPrint)

	resSub := Sub(1, 2)
	fmt.Println(resSub)

	r := rectangle.New(1, 2, "green")
	fmt.Println(r)

	res, ok := FuncWithError(20)

	if ok != nil {
		fmt.Println("Oops", ok)
		return
	} else {
		fmt.Println("Good", res)
	}

	FuncWithPanic()
}
