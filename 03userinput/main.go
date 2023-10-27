package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to userinput"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza:")
	//comma ok // comma err
		input, _ := reader.ReadString('\n')
		fmt.Println("Thanks for ratings,", input)
	fmt.Printf("Type of this rating is %T ", input)

}
