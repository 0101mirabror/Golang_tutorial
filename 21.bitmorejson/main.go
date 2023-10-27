package main

import (
	"fmt"
	"encoding/json"
)

type course struct{
	Name string `json:"coursename"`
	Price int
	Platform string `json:"website"`
	Password string	`json:"-"` /*-(dash) qo'yilsa bu maydonni aidan olib tashlaydi*/
	Tags []string	`json:"tags,omitempty"` /*comma bilan so'zlar orasida bo'sh joy bo'lmasligi kerak*/
}

func  main(){
	fmt.Println("Welcome to JSON video in Golang")
	// EncodeJSON()
	DecodeJson()
}

func EncodeJSON(){
	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "LearnCodeOnline.in", "bcd123", []string{"fullstack", "js"}},
		{"Angular Bootcamp", 99, "LearnCodeOnline.in", "hit123", nil},
	}

	//package this data ad JSON data

	// finalJson, err := json.Marshal(lcoCourses)
	finalJson, err := json.MarshalIndent(lcoCourses,"", "\t")

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson(){
	jsonDataFromWeb := []byte(`
	{	"coursename": "ReactJS Bootcamp",
		"Price": 299,
		"website": "LearnCodeOnline.in",
		"tags": ["web-dev", "js"]
	}
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	// some cases where you just want to add data to key value pair  	 


	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for key, value := range myOnlineData{
		fmt.Printf("Key is %v and value is %v and type is  %T\n", key, value, value)
	}

}
