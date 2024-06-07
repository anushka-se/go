package main

import "fmt"

func factorial(x int) int {
	if x == 1 || x == 0 {
		return 1
	} else {
		return factorial(x-1) * x
	}
}

func main() {
	ans := factorial(6)
	fmt.Println(ans)
}
