package main

import (
	"fmt"
	"sort"
)

func main(){
	// fmt.Println("Welcome to video on slices")
	var fruitList = []string{"Apple", "Tomato", "Peach"}
	// fmt.Printf("Type of fruitlist is %T \n", fruitList)
	// fmt.Println(fruitList)
	fruitList = append(fruitList, "Banana", "Mango")
	// fmt.Println(fruitList)
	fruitList = append(fruitList[1:])
	// fmt.Println(fruitList)
	fruitList = append(fruitList[1:3])
	// fmt.Println(fruitList)
	fruitList = append(fruitList[0:3])
	// fmt.Println(fruitList)

	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867


	highScores = append(highScores, 666, 555, 777)
	// fmt.Println(highScores)

	// fmt.Println(sort.IntsAreSorted(highScores))

	sort.Ints(highScores)
	// fmt.Println(highScores)
	// fmt.Println(sort.IntsAreSorted(highScores))


	//how to remove a value from slices based on index
	
	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:2], courses[index+1:]...)
	fmt.Println(courses)
 
}