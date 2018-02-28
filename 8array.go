package main

import "fmt"

func main() {
	var a [5]int

	fmt.Println("Empty array ", a)

	a[3] = 3
	a[1] = 1

	fmt.Println("Assignment ", a)

	fmt.Println("Single value ", a[3])

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	// TwoD array
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("TwoD Array ", twoD)
}
