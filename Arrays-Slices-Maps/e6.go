package main

import "fmt"

func main() {
	// map of strings to int
	// var x map[string]int

	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x["key"])

	x2 := make(map[int]int)
	x2[1] = 11
	x2[2] = 12
	x2[3] = 13
	x2[4] = 14
	x2[5] = 15

	for i := 1; i < 6; i++ {
		fmt.Println(x2[i])
	}

	delete(x2, 1)
	fmt.Println("Post Deletion")
	for i := 1; i < 6; i++ {
		fmt.Println(x2[i])
	}
}
