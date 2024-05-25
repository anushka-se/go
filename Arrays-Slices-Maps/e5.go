package main

import "fmt"

func main() {
	sliceOne := []int{1, 2, 3}
	sliceTwo := make([]int, 2)
	
	// Destination, source
	copy(sliceTwo, sliceOne)
	fmt.Println(sliceOne, sliceTwo)
}
