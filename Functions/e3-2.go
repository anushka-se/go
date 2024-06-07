package main

import "fmt"

// We need to either do this:
func f(x int) {
	fmt.Println(x)
}

func main() {
	x := 5
	f(x)
}
