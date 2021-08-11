package main

import "fmt"

//NUMBERS,BOOLEANS AND STRINGS


func main(){
	//integers can be signed or unsigned from 8 to 64

	//signed ints can hold zero,negative and positive values
	 var x int16 = -22222
	fmt.Println(x, x + 3, x - 3)

	//unsigned ints can only hold zeros and positive values
	var y uint8 = 225
	fmt.Println(y, y + 3, y - 3)

	//you can declare a datatype or leave for the compiler to decide

	/*
	CONVERSION OF TYPES
	*/

	var oranges uint8 = 12
	var bananas uint16 = 20

	//you have to convert them to be of the same types before any operations
	var fruits = oranges + uint8(bananas)

	fmt.Println(fruits)

	
}