package main

import (
	"fmt"
	"runtime"
	"bufio"
	"os"

)

func main() {
	num := runtime.NumCPU()
	fmt.Println(num)

	anvar := bufio.NewReader(os.Stdin)
	input, _ := anvar.ReadString('\n')
	fmt.Println("Thanks for ratings,", input)
}