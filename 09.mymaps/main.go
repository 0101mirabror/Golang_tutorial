package main

import (
	"fmt"
	// "sort"
	// "bufio"
	// "os"
 )

func main(){
	fmt.Println("Maps in Golang")
	languages := make(map[string]string)
	languages["Js"] = "Javascript"
	languages["Rb"] = "Ruby"
	languages["Py"] = "Python"

	fmt.Println("List of all languages: ", languages)
	fmt.Println("Js shorts for: ", languages["Js"])
	
	delete(languages, "Rb")
	fmt.Println("List of all languages: ", languages)

	//loops are interesting in Golang

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
 
	// OK COMMA SYNTAX
	for _, value := range languages {
		fmt.Printf("For key v, value is %v\n",  value)
	}
 
 
}