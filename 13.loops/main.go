package main


import "fmt"

func main(){
	fmt.Println("Welcome to loops in Golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println( days)

	// for d:=0; d<len(days); d++{
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }


	// for index, day 	:= range days {
	// 	fmt.Printf("Index is %v, day is %v\n", index, day)
	// }

	// //vomma ok syntax
	// for _, day 	:= range days {
	// 	fmt.Printf("Index is v, day is %v\n", day)
	// }

	rougueValue := 1

		//break

	// for rougueValue < 10 {
	// 	if rougueValue == 5{
	// 		break
	// 	}
	// 	fmt.Println("Value is: ", rougueValue)
	// 	rougueValue++
	// }

		//continue

	// for rougueValue < 10 {
	// 	if rougueValue == 5{
	// 		rougueValue++
	// 		continue
	// 	}
	// 	fmt.Println("Value is: ", rougueValue)
	// 	rougueValue++
	// }

	//goto

	for rougueValue < 10 {
		if rougueValue == 2 {
			goto lco
		}
		if rougueValue == 5{
			rougueValue++
			continue
		}
		fmt.Println("Value is: ", rougueValue)
		rougueValue++
	}
	lco: 
		fmt.Println("Jump in LearnCodeOnline.in")


}