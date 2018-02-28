package main

import "fmt"

func main() {
	// While
	var i = 0
	for i < 3 {
		fmt.Println("First Loop ", i)
		i = i + 1
	}

	// For Loop

	for j := 1; j < 5; j++ {
		fmt.Println("Second Loop ", j)
	}

	// Break

	for {
		fmt.Println("Infinite Loop")
		break
	}

	// Continue
	i = 0
	for i < 3 {
		if i == 0 {
			fmt.Println("Continue Loop ", i)
			i++
			continue
		} else {
			fmt.Println("Break Loop ", i)
			i += 1
			break
		}
	}

}
