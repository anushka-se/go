package main

import "fmt"

func first() {
	fmt.Println("First")
}

func second() {
	fmt.Println("second")
}

// func three() {
// 	fmt.Println("three")
// }

func main() {
	defer second()
	// defer three()
	first()
}
