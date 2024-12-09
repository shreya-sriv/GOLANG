package main

import "fmt"

func main() {
	var username string = "Shreya"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLog bool = true
	fmt.Println(isLog)
	fmt.Printf("Variable is of type: %T \n", isLog)

	//default values and aliases
	var another float32
	fmt.Println(another)
	fmt.Printf("Variable is of type: %T \n", another)

	//implicit type
	var name = "Palak"
	fmt.Println(name)
	names := "hehe"
	fmt.Println(names)
}
