package main

import "fmt"

func main()  {
	//for storing key-value pair
	studentsMarks := map[string]int{
		"Amos" : 200,
		"Jane" : 230,
		"Jayden" : 79,
		"Ian" : 138,
	}
	fmt.Println(studentsMarks["Amos"])
}