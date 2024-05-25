package main

import "fmt"

func main() {
	sliceOne := []int{1, 2, 3}
	sliceTwo := append(sliceOne, 4, 5)

	fmt.Println(sliceOne, sliceTwo)
}
