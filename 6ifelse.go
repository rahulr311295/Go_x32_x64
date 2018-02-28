package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if num := 9; num < 0 {
		fmt.Println("Number is negative")
	} else if num < 10 {
		fmt.Println("Number is single digit")
	} else {
		fmt.Println("Number is greater than 10")
	}
}
