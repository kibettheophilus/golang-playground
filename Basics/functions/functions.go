package main

import "fmt"

func add(firstName string,lastName string)(string, error)  {

	return firstName + "" + lastName, nil

}

func main()  {
	fullName,err := add("Go","Lang") 

	if err != nil {
	fmt.Println("Error found")
	}

	fmt.Println(fullName)
}