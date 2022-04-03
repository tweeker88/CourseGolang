package main

import (
	"fmt"
	"sync"
	"time"
)

// Обычная горутина
func simpleGorutine(value chan bool) {
	fmt.Println("Hello! i'm simple gorutine")
	value <- true
}

func gorutineWithClose(value chan int) {
	value <- 10
	fmt.Println("I'm gorutine with close")
	close(value)
}

func gorutineWithCloseAndFor(value chan int) {

	for i := 0; i <= 4; i++ {
		value <- i
	}

	close(value)
}

type Money struct {
	value int
}

func process(money *Money, wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	money.value++
	m.Unlock()
	wg.Done() // WaitGroup --
}

func server1(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "from server1"
}

func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from server2"
}

func funcWithSelect() {
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)
	for {
		select {
		case res1 := <-output1:
			fmt.Println(res1)
			return
		case res2 := <-output2:
			fmt.Println(res2)
			return
		default:
			fmt.Println("Servers are not avaible")
			time.Sleep(1000 * time.Millisecond)
		}
	}
}

func main() {

	var zeroValue chan bool

	if zeroValue == nil {
		fmt.Println("Variable zeroValue is nil")
		zeroValue = make(chan bool, 0)

		go simpleGorutine(zeroValue) // вызов обычной горутины
	}

	first := <-zeroValue
	fmt.Println("First Read from zeroValue chan", first)

	calculate(456)

	// Пример deadlock'a
	// varForDeadlock := make(chan bool)
	// <-varForDeadlock (пытаемся прочитать, но НИКТО не записывает)

	valueForCloseGorutine := make(chan int)

	go gorutineWithClose(valueForCloseGorutine)
	resultFromValueForCloseGroutine := <-valueForCloseGorutine
	fmt.Println("Result from gorutine with close chan", resultFromValueForCloseGroutine)

	valueForCloseGorutineAndFor := make(chan int)
	go gorutineWithCloseAndFor(valueForCloseGorutineAndFor)
	for {
		res, ok := <-valueForCloseGorutineAndFor
		if ok == false {
			fmt.Println("Chan is close")
			break
		}
		fmt.Printf("Value: %v from circle in gorutine \n", res)
	}

	var wg sync.WaitGroup
	var m sync.Mutex
	money := Money{0}
	for i := 0; i < 1000; i++ {
		wg.Add(1) // WaitGroup ++
		go process(&money, &wg, &m)
	}
	wg.Wait()
	fmt.Println("Count money is", money.value)

	funcWithSelect()
	fmt.Println("I'm main gorutina and I'm end")

}

//1. Создадим чуть более полезную программу, которая будет делать следующие действия:
// берем число, например 456
// (4^2 + 5^2 + 6^2) + (4^3 + 5^3 + 6^3)
// Подсчитаем сумму квадартов цифр и сумму кубов, а затем сложим полученные результаты
//Делать будем следующим образом:
// * main gorutine получает число и вызывает 2 другие горутины, по итогу, получив от них результаты,
// она просто их сложит и выведет на консоль
// ** squaresGoRoutine - будет запущена main и подсчитает сумму квадратов всех цифр, результат положит в канал
// ** cubesGoRoutine - будет запущена main и подсчиатет сумму кубов всех цифр , результат полужит в канал

func calculate(variable int) {

	squareChan := make(chan int)
	cubeChan := make(chan int)

	go squaresGoRoutine(variable, squareChan)
	go cubesGoRoutine(variable, cubeChan)

	squareSum, cubeSum := <-squareChan, <-cubeChan

	fmt.Printf("squareSum: %v and cubeSum: %v for varialbe \n", squareSum, cubeSum)
}

func squaresGoRoutine(number int, squareChan chan int) {
	sum := 0
	dChan := make(chan int)
	go digit(number, dChan)

	for digit := range dChan {
		sum += digit * digit
	}

	squareChan <- sum
}

func cubesGoRoutine(number int, cubeChan chan int) {
	sum := 0
	dChan := make(chan int)
	go digit(number, dChan)
	for digit := range dChan {
		sum += digit * digit * digit
	}

	cubeChan <- sum
}

func digit(number int, digitChan chan int) {
	for number != 0 {
		digit := number % 10
		digitChan <- digit
		number /= 10
	}

	close(digitChan)
}
