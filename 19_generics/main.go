package main

import (
	"constraints"
	"fmt"
)

type NumberType interface {
	int | float32 | float64 | uint16 | uint64 | uint32 | int32
}

func addList[T constraints.Ordered](list []T) T {
	var sum T
	for _, item := range list {
		sum += item
	}
	return sum
}

// func printMe(thing interface{}) {
// 	fmt.Println(thing)
// }

func printMe(thing any) {
	fmt.Println(thing)
}


func addInts(list []int) int {
	var sum int
	for _, item := range list {
		sum += item
	}
	return sum
}

func addFloats(list []float32) float32 {
	var sum float32
	for _, item := range list {
		sum += item
	}
	return sum
}

func main() {

	fmt.Printf("sum of ints: %d\n", addInts([]int{1, 2, 3, 4, 5}))
	fmt.Printf("sum of floats: %.2f\n", addFloats([]float32{1.2, 2.3, 3.4}))

	fmt.Printf("sum of ints: %d\n", addList([]int{1, 2, 3, 4, 5}))
	fmt.Printf("sum of floats: %.2f\n", addList([]float32{1.2, 2.3, 3.4}))

	printMe("hi")
	printMe(1.2)

}