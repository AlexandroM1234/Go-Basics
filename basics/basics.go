package main

import (
	"fmt"
)

func main()  {
	// Variables
	// Long ways is var, then the name of the varible , the type , and then value
	var y int = 1
	// Short hand is variable name, colon equals the value 
	x := 1
	var sum int = x + y
	fmt.Println(sum)

	// if statements
	// note there are no parenthesis around the condition
	if sum > 10 {
		fmt.Println("More than 10")
	} else if sum > 5 {
		fmt.Println("More than 5")
	} else {
		fmt.Println("Less than 5")
	}

	// Arrays
	// Arrays are fixed in length and we can declase the contents of the array or not and if not the default values will all be 0
	// Note when declaring an array we have to set the length inside of the brackets
	a := [5]int{1,2,3,4,5}
	var b = [5]int {}
	fmt.Println(a)
	fmt.Println(b)

	// Slices
	// Slices are like arrays but they don't have a fixed length and we remove the number from the declration
	slice := []int{23,12,231}
	var slice2 = []int {0}

	fmt.Println(slice)
	fmt.Println(slice2)

	// With slices we can append to the end of the slice note it doesn't modify the original rather it creates a new one
	slice2 = append(slice2,1)
	fmt.Println(slice2)

	// Maps
	// when declaring a map you use the map keyword followed by the type of the keys and then the type of the values
	// then call the built in make function to create that new type
	hashMap := make(map[string]int)

	



	
}