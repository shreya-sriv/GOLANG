package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your rating")

	input, _ := reader.ReadString('\n')
	fmt.Println("Your rating is :", input)
}
