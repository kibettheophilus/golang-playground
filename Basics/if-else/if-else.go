package main

import "fmt"

func main()  {
	customerHeight := 140
	customerAge := 18

	if customerHeight >= 150 && customerAge >= 18 {
		fmt.Println("Old enoogh")
	} else if customerHeight >= 120 || customerAge >= 15 {
		fmt.Println("Come back later")
	} else{
		fmt.Println("Don't try coming near")
	}
}