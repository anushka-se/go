package main

import "fmt"

// func main() {
// 	panic("PANIC")
// 	str := recover()
// 	fmt.Println(str)
// }

func main() {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}
