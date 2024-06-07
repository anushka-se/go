package main

import "fmt"

/*
# command-line-arguments
./e3.go:6:14: undefined: x
./e3.go:10:2: x declared and not used
*/

func f() {
	fmt.Println(x)
}

func main() {
	x := 5
	f()
}