package main

import "fmt"

func main(){
	greeter()
	fmt.Println("Welcome to functions in golang")

	// function ichida function yozish mumkin emas

	// greeterTwo()
	// func greeterTwo(){
	// 	fmt.Println("Another greeting method")
	// }

	// result := adder(3,5)

	// fmt.Println("Result is: ", result)

	// values :=  []int{1, 5, 6, 8}
	result_pro, myMessage := proAdder(1,5,6,8,8,8)
	fmt.Println("Pro result is ",result_pro)
	fmt.Println("My message is ", myMessage )
}
// func greeterTwo(){
// 	fmt.Println("Another greeting method")
// }
func adder(num1 int, num2 int)int{
		return num1 + num2
}

func proAdder( values ... int) (int, string) {
	total := 0
	for _, value := range values {
		total = total + value 
	}
	return total, "Golang is best language for learning"
}


func greeter(){
	fmt.Println("Salom from Golang")
}