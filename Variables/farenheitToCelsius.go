package main

import "fmt"

func main() {
	var far float32 = 212
	var cel float32 = 5 * (far - 32) / 9
	fmt.Printf("%v\n", cel)
}
