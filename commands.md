# tutorial3

- `go mod init hello`
- control + l terminalni tozalash
- agar go fayldagi kodda biror error bo'lsa, quyidagicha ` If any of the dependencies are missing, the ⚠️ Analysis Tools Missing warning is displayed. Click on the warning to download dependencies.`
- lexers
-

# tutorial 6

- variables
- variable yaratish uchun `var` dan foydalaniladi.
- `fp` bu `fmt.Println` uchun shortcut
- `fmt.Printf` stringni formatlash uchun ishlatiladi, `\n` yangi qator, `\d` tab uchun. `%T` o'zidagi o'zgaruvchini tipini ko'rsatadi
- `:=` walrus operator metod ichida ishlatiladi, metoddan tashqarida ishlatish mumkine emas.

# user input

- `pkg.go.dev` pypi kabi golang uchun packagelar yuklanadigan sayt.

# inputdan ma'lumotni o'qib uni variablega tenglashtirish

```
packages (
    "os"
    "buffio"
)
func main(){
    son = buffio.NewReader(os.Stdin)
}
son input orqali kiritilgan o'zgaruvchiga teng bo'ladi
```

`walrus operator << := >>`

# time

```
func main(){
	fmt.Println("Welcome to time study of golang")
	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

}
```

- Format ichida aynan '01-02-2006' yoki Monday bo'lishi kerak aks holda hato sana chiqaradi.
- aniq sana bo'lishi uchun faqat Monday berilishi kerak
- aniq soat uchun aynan "15:04:05" bo'lishi kerak

# Aynan bir vaqtni input orqali kiritish

```
created_date := time.Date(2020, time.July, 18, 0, 17, 43, 0, time.UTC)
```

# Pointer

```
package main


import (
	"fmt"


)
func main(){
	fmt.Println("Welcome to class on pointers")
	// var ptr *int
	// fmt.Println("Value of this pointer is, ", ptr)
	myNumber := 23
	var ptr = &myNumber

	fmt.Println("Value of actual pointer is ", ptr)
	fmt.Println("Value of actual pointer is ", *ptr)

	*ptr = *ptr*2
	fmt.Println("New value is ", myNumber)
}
```

- my numberga qiymat olindi
- & belgisi my number oldiga qo'yilib, my number manzili ptrga tenglandi
- _ptr da _ belgi bilan ptr ichidagi qiymatni oldik

# Arrays

```
package main

import (
	"fmt"
)

func  main(){

	fmt.Println("Welcome to array in Golang")
	var fruitList [4] string
	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"
	// fruitList[] = "Apple"
	fmt.Println("Fruit list is:",  fruitList)
	fmt.Println("Fruit list length is:",  len(fruitList))

	var vegList = [3] string {"potato", "beans", "mushroom"}
	fmt.Println("Vegetables list is", vegList)
	fmt.Println("Vegetables list length is", len(vegList))

}
```

- array elon qilinganda datatype va uzunligi ko'rsatilishi kerak
- ikki hil elon qilish mumkin

```
var fruitList [4] string
fruitList[0] = "Apple"

yoki

var vegList = [3] string {"potato", "beans", "mushroom"}
```

# Slices

```
package main

import (
	"fmt"
	"sort"
)

func main(){
	// fmt.Println("Welcome to video on slices")
	var fruitList = []string{"Apple", "Tomato", "Peach"}
	// fmt.Printf("Type of fruitlist is %T \n", fruitList)
	// fmt.Println(fruitList)
	fruitList = append(fruitList, "Banana", "Mango")
	// fmt.Println(fruitList)
	fruitList = append(fruitList[1:])
	// fmt.Println(fruitList)
	fruitList = append(fruitList[1:3])
	// fmt.Println(fruitList)
	fruitList = append(fruitList[0:3])
	// fmt.Println(fruitList)

	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867


	highScores = append(highScores, 666, 555, 777)
	// fmt.Println(highScores)

	// fmt.Println(sort.IntsAreSorted(highScores))

	sort.Ints(highScores)
	// fmt.Println(highScores)
	// fmt.Println(sort.IntsAreSorted(highScores))


	//how to remove a value from slices based on index

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:2], courses[index+1:]...)
	fmt.Println(courses)

}
```

- append orqali slicedan element o'shirish yoki qo'shish mumkin
- slices pythondagi listga o'xshaydi yana golangdagi arrayga ham
- make va new funksiyalari orqali variables yaratish mumkin. new yangidan yaratadi, make yaratganda qiymat bo'ladi
- make bilan slice yaratish

```
highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
```

- slicedan qiymatni o'chirish

```
var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:2], courses[index+1:]...)
	fmt.Println(courses)

```

- slicelarni sortlash(buning uchun "sort" packageni import qilish kerak)

```
sort.Ints(highScores)
// fmt.Println(highScores)
// fmt.Println(sort.IntsAreSorted(highScores))
```

# maps

```
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
```

- maps pythondagi dictionaryiga o'xshash unda ham key value bo'ladi
- make() orqali map yaratish

```
	languages := make(map[string]string)
	languages["Js"] = "Javascript"
	languages["Rb"] = "Ruby"
	languages["Py"] = "Python"
```

- mapdan qiymat o'chirish

```
delete(languages, "Rb")
fmt.Println("List of all languages: ", languages)
```

- %v bilan mapdan key valueni olish

```
for key, value := range languages {
	fmt.Printf("For key %v, value is %v\n", key, value)
	}

```

# struct

```

package main

import "fmt"
func main(){

	fmt.Println("Structs in golang")
	//no inheritance in golang: No super or parent

	hitesh := User{"Hitesh", "hitesh@go.dev", true, 16}
	fmt.Println(hitesh)
	fmt.Printf("hitesh details are: %+v\n", hitesh)
	fmt.Printf("Name is %v and email is %v:", hitesh.Name, hitesh.Email)
}

type User struct {
	Name string
	Email string
	Status bool
	Age int

}
```

- struct golangda bo'ladi, golangda klass bo'lmaydi
- structda struct nomi va undagi o'zgaruvchilar doimo bosh harfda bo'lishi kerak
- main funksiyasidan keyin struct yoziladi
- structda hamma elementlarni olish(bunda formatlaganda %+v dan foydalaniladi)

```

	hitesh := User{"Hitesh", "hitesh@go.dev", true, 16}
	fmt.Println(hitesh)
	fmt.Printf("hitesh details are: %+v\n", hitesh)
```

- structda aynan bir fieldni olish(%vdan foydalanamiz)

```

hitesh := User{"Hitesh", "hitesh@go.dev", true, 16}

	fmt.Printf("Name is %v and email is %v:", hitesh.Name, hitesh.Email)
```

# ifelse

```

package main

import "fmt"

func main(){
	fmt.Println("If else in golang")

	loginCount := 9

	var result string

	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = "Exactly 10 login count"
	}
	fmt.Println(result)

	if 9%2 == 0{
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")

	} else {
		fmt.Println("Num is not less than 10")
	}
}

```

# switchcase

```
package main


import (
	"fmt"
	"time"
	"math/rand"
)

// oltita nuqtadali kubli qimor o'yini

func main(){
	fmt.Println("Switch and case in Golang")
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is one and you can open")
	case 2:
		fmt.Println("You can move 2 spot")
	case 3:
		fmt.Println("You can move 3 spot")
		fallthrough
	case 4:
		fmt.Println("You can move 4 spot")
		fallthrough
	case 5:
		fmt.Println("You can move 5 spot")
	case 6:
		fmt.Println("You can move 6 spot and roll dice again")
	default:
		fmt.Println("What was that!")

	}

}

```

# loop for

```
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
```

# function

```

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
```

# methods

```
package main


import "fmt"

func main(){
	fmt.Println("Welcome to methods in Golang")

	//no inheritance in golang: No super or parent

	hitesh := User{"Hitesh", "hitesh@go.dev", true, 16}
	fmt.Println(hitesh)
	fmt.Printf("hitesh details are: %+v\n", hitesh)
	fmt.Printf("Name is %v and email is %v\n", hitesh.Name, hitesh.Email)
	hitesh.GetStatus()

	hitesh.NewMail()
	fmt.Println(hitesh.Email)
}

type User struct {
	Name string
	Email string
	Status bool
	Age int

}

func (u User) GetStatus(){
	fmt.Println("Is user active:", u.Status)
}

func (u User) NewMail(){
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is ", u.Email)
}
```

# defer

```
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
```

# files

```
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
```

# webrequests

```
package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

const url = "https://lco.dev"
func main(){
	fmt.Println("Welcome to web requests in Golang")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response of type is: %T\n", response)
	defer response.Body.Close() // caller's responsibility to close the connection

	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)
}

```

# handling url in golang

```

package main

import (
	"fmt"
	"net/url"

)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=gbhj456ghb"

func main(){
	fmt.Println("Handling URLsin Golang")
	fmt.Println(myurl)

	//parsing

	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query params are  \n %T\n", qparams)
	fmt.Println(qparams["coursename"])

	for  _, val := range qparams{
		fmt.Println("Param is ", val)
	}

	partsOfUrl := &url.URL{
		Scheme:"https",
		Host: "lco.dev",
		Path:"/tutcss",
		RawPath: "user=hitesh",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)


}
```

# webreqverbs

```
package main


import (
	"fmt"
	"net/http"
	"io/ioutil"
	"strings"
)


func main(){
	fmt.Println("Welcome to web verb video")
	PerformGetRequest()
}

func PerformGetRequest(){
	const myurl = "http://localhost:8000/get"
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length is ", response.ContentLength  )

	/* 2 - usul*/
	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("ByteCount is:", byteCount)
	fmt.Println(responseString.String())



	/* 1 - usul*/

	// content, _ := ioutil.ReadAll(response.Body)
	// fmt.Println(content)
	// fmt.Println(string(content))
}
```

# making POST request in Golang

```
package main


import (
	"fmt"
	"net/http"
	"io/ioutil"
	"strings"
)


func main(){
	fmt.Println("Welcome to web verb video")
	// PerformGetRequest()
	PerformPostJsonRequest()
}

func PerformPostJsonRequest(){
	const myurl = "http://localhost:8000/post"

	//fake json payload

	requestBody := strings.NewReader(`
		{
			"coursename": "Let's go with Golang",
			"price": 0,
			"platform": "learncodeonline.in"
		}
	`)
	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
```

# json

```
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
	EncodeJSON()
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

```

# JSON datani qabul qilish

```
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
```
