package main

import(
	"fmt"
)

func main() {
	var myArray [2] string
	myArray[0] = "Hello"
	myArray[1] = "World"

	fmt.Println(myArray[0]+" "+myArray[1])
}