package main

import (
	"fmt"
)

func q3() {
	x := [6]string{"a", "b", "c", "d", "e", "f"}

	fmt.Println(x[2:5])
}

func q4() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	smallest := x[1]
	for i := 0; i < len(x); i++ {
		if x[i] < smallest {
			smallest = x[i]
		}
	}

	fmt.Println(x)
	fmt.Println(smallest)
}

func main() {
	// q3()
	q4()
}
