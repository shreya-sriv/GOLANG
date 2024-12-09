package main

import (
	"fmt"
	"sort"
)

func main() {

	var vegList = []string{"Potato", "beans", "mushroom"}
	fmt.Println(vegList)
	vegList = append(vegList, "mango", "banana")
	fmt.Println(vegList)

	vegList = append(vegList[1:])
	fmt.Println(vegList)

	highScores := make([]int, 4)
	highScores[0] = 231
	highScores[1] = 321
	highScores[2] = 455
	highScores[3] = 768
	fmt.Println(highScores)

	highScores = append(highScores, 555, 666, 888)
	fmt.Println(highScores)
	sort.Ints(highScores)
	fmt.Println(highScores)

	var courses = []string{"Java", "Python", "Rust", "JavaScript", "GoLang"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
