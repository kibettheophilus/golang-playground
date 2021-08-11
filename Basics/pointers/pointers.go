package main

import "fmt"

func main()  {
	// & is used to reference a location of a variable in memory
	// * is used to reference the value of a variable

	x := 12
	y := &x
	
	fmt.Println(x,y)

	*y = 20

	fmt.Println(x,y)
}