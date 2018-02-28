package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Println("Write ", i, " as ")

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// Switch with Day
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's weekend :)")
	default:
		fmt.Println("It's weekday :()")
	}

	// Switch with Time

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's forenoon")
	default:
		fmt.Println("It's afternoon")
	}

	// Type - Interface
	myType := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm bool")
		case int:
			fmt.Println("I'm int")
		default:
			fmt.Printf("I'm %T\n", t)
		}
	}

	myType(true)
	myType(10)
	myType("Hello")
}
