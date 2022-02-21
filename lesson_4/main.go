package main

import "fmt"

func main() {
	// init
	var myArray [5]int
	fmt.Println("My array:", myArray)
	myArray[0] = 1
	myArray[2] = 3
	fmt.Println("My array:", myArray)

	secondArray := [3]int{10, 20, 30}
	fmt.Println("Second array:", secondArray)

	endlessArray := [...]int{5, 4, 3, 2, 1}
	fmt.Println("Endless array:", endlessArray, len(endlessArray))

	// copying
	originalArray := [3]int{1, 2, 3}
	coppiedArray := originalArray

	fmt.Println("Origin", originalArray, "coppied:", coppiedArray)

	coppiedArray[0] = 100
	fmt.Println("Origin", originalArray, "coppied:", coppiedArray)

	// for
	for i := 0; i < len(coppiedArray); i++ {
		fmt.Println("i:", i, "array element:", coppiedArray[i])
	}

	// range
	for id, val := range coppiedArray {
		fmt.Println("id:", id, "val:", val)
	}

	for _, val := range coppiedArray {
		fmt.Println("only val:", val)
	}

	// slice
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Simple slice:", slice)

	sliceSlice := slice[0:2]
	fmt.Println("Simple slice:", slice, "Slice slice:", sliceSlice)
	sliceSlice[0] = 100
	fmt.Println("Simple slice:", slice, "Slice slice:", sliceSlice)

	secondSliceSlice := sliceSlice[:]
	secondSliceSlice[0]++
	fmt.Println("Simple slice:", slice, "Slice slice:", sliceSlice, "Second slice slice", secondSliceSlice)

	// capacity
	numerics := []int{1, 2}
	for i := 0; i < 200; i++ {
		if i%5 == 0 {
			fmt.Println("Current len:", len(numerics), "Current cap:", cap(numerics))
		}
		numerics = append(numerics, i)
	}

	// make
	makeSlice := make([]string, 5, 10)
	fmt.Println(makeSlice)
	makeSlice[0] = "Maxim"
	fmt.Println(makeSlice)
}
