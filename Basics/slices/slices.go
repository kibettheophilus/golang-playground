package main

import "fmt"

func main()  {
	days := [...]string{"Monday","Tuesday","Wednesday","Thursday","Friday","Saturday","Sunday"}

	weekdays := days[0:5]

	fmt.Println(weekdays)
}