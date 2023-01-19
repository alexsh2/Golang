package main

import (
	"fmt"
	"runtime"
)

func main() {

	/*numbers := [...]int{2, 5, 13, 7, 6, 4}
	size := 6
	sum := 0
	avg := 0
	index := 0

	for index < size {
		sum = sum + numbers[index]
		index++
	}

	avg = sum / size
	fmt.Println(avg)
	fmt.Printf("Type: %T Value: %#v\n", numbers, numbers)

	for ind, value := range numbers {
		fmt.Printf("Index: %d Value: %d\n", ind, value)
	}

	for _, value := range numbers {
		fmt.Printf("Value: %d\n", value)
	}

	srez := []string{"first", "second", "third"}

	for _, str := range srez {
		fmt.Printf("Value: %s\n", str)
	}

	fmt.Printf("Pointer: %p\n", &srez)*/

	intArr := [...]int{1, 2, 3, 4, 5}

	intSlice := intArr[:]
	intSlice = append(intSlice, 6)

	fmt.Printf("Type: %T Value: %#v\n", intSlice, intSlice)
	fmt.Printf("Length: %d Capacity: %d\n", len(intSlice), cap(intSlice))
	fmt.Printf("Pointer: %p\n\n", &intSlice)
	fmt.Printf("Type: %T Value: %#v\n", intArr, intArr)
	fmt.Printf("Pointer: %p\n", &intArr)

	intSlice[0] = 33
	fmt.Printf("Type: %T Value: %#v\n", intSlice, intSlice)
	fmt.Printf("Type: %T Value: %#v\n\n", intArr, intArr)

	fmt.Printf("Number of available CPUs: %v\n\n", runtime.NumCPU())

}
