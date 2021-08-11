package main

import "fmt"

func main()  {
	ages := []int {23,30,18,17,27}

	for i := 0; i < len(ages); i++ {
		fmt.Println(ages[i])
	} 
}