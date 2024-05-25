package main

import "fmt"

func main() {
	x := make([]float64, 5, 10)
	fmt.Println(x, len(x))

	arr := [5]float32{1,2,3,4,5}
	x1 := arr[0:5]
	fmt.Println(x1, len(x1))
}
