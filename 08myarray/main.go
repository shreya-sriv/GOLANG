package main

import "fmt"

func main() {
	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[3] = "Peach"
	fmt.Println(fruitList)

	var vegList = [3]string{"Potato", "beans", "mushroom"}
	fmt.Println(vegList)
}
