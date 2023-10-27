package main


import "fmt"


func main(){
	fmt.Println("Welcome to defers in Golang")
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("\nTwo")
	fmt.Println("Hello")

	myDefer()
}
// 0, 1 ,2 ,3 ,4
//Welcome .../ hello/ 43210 / Two/ One/ World
func myDefer(){

	for i:=0; i < 5; i++{
		defer fmt.Print(i)
	}
}