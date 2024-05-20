package main

import "fmt"

func main() {
	for i := 3; i <= 100; i += 3 {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}
