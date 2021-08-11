package main

import "fmt"

func main(){
	//declaring empty array
	var months [4]string

	months[0] = "January"
	months[1] = "February"
	months[2] = "March"
	months[3] = "April"

	fmt.Println(months)

	//declaring array with elements
	days := []string{"Monday","Tuesday","Wednesday","Thursday","Friday","Saturday","Sunday"}

	fmt.Println(days)
	fmt.Println(len(days))
}