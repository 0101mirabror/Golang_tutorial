package main


import ("os"
		"fmt"
		"io"
		"io/ioutil"
		
	)

func main(){
	fmt.Println("Welcome to Files in golang")

	content := "This needs to go in  a file - LearnCodeOnline.com"
	file, err := os.Create("./mylcofile.txt")
	checkNilErr(err)
	// if err != nil {
	// 	panic(err)

	// }
	length, err := io.WriteString(file, content)
	checkNilErr(err)
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Println("Length of string is",length)
	defer file.Close()

	readFile("./mylcofile.txt")
}


func readFile(filename string){
	databyte, err := ioutil.ReadFile(filename)
	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)
	fmt.Println("Text data inside the file is \n ", string(databyte))

}

//errorHandling


func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}