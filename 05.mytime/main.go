package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println("Welcome to time study of golang")
	presentTime := time.Now()	
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
	created_date := time.Date(2001, time.January, 24, 0, 17, 43, 0, time.UTC)
	fmt.Println(created_date)
	fmt.Println(created_date.Format("01-02-2006 Monday"))
}
