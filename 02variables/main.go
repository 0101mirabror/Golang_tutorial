package main

import "fmt"

// jwtToken := 300000
const LoginToken =  "sksjdk44sd55s7dfsd57sd5sd5s" //public

func main() {
	// fmt.Println("Variables")
	var username string = "hitesh"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)
	
	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)


	var smallFloat float64 = 255.4545478797545454
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	//default values and some alies
	var anotherVariable string
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	//implicit type
	var website = "onlinelearncode.in"
	fmt.Println(website)

	//no var style

	numberOfUser := 300000
	fmt.Println(numberOfUser)

	//public const
	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}