package main

import "fmt"

type Student struct{
	name string
	age int
}

type Class struct{
	name string
	group [2] Student
}


func main()  {
	//student1 := Student{"Amos",12}

	groups := [...]Student{{"Hey",1},{"Hi",2}}

	classFour := Class{name:"Standard Four East",group: groups}

	fmt.Println(classFour)
}