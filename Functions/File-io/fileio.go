package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	fmt.Println("Read File as String\n")

	//Pass the file name to the ReadFile() function from
	//the ioutil package to get the content of the file.

	content, error := ioutil.ReadFile("test.txt")

	// Check whether the 'error' is nil or not. If it
	//is not nil, then print the error and exit.
	if error != nil {

		log.Fatal(error)
	}

	// convert the content into a string
	str := string(content)

	//Print the string str
	fmt.Println(str)
}
