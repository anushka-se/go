package main

import "fmt"

var scopeTest string = "Testing out the scope "

func main() {
	mainString := "Inside main function: "
	fmt.Println(mainString + scopeTest)
	f()
}

func f() {
	fString := "Inside the testing function: "
	fmt.Println(fString + scopeTest)
}
