package main

import "fmt"

func main() {
	// var ptr *int
	// fmt.Println("value of pointer", ptr)

	myNumber := 23
	var ptr = &myNumber
	fmt.Println("Value of actual pointer is", ptr)
	fmt.Println("Value of actual pointer is", *ptr)
	*ptr = *ptr + 23
	fmt.Println(*ptr)

}
