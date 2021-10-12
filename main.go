package main // this is written because this the main package we are creating and will be helpful in manging
// different packages

import (
	"fmt"
	"reflect"
)

// imported packages

func main() { // this is the main function

	fmt.Println("Hello, playground")
	var myName string = "Puran"
	fmt.Println(myName)

	var arrayName [5]int
	fmt.Println(reflect.TypeOf(arrayName))

	// var array2 []int = arrayName[1:3]
	// fmt.Println(array2)
}

// Reverse string
// sum of an array
// star pattern
// find the armstring number
