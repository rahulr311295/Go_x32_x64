package main

import "fmt"

func sum(n ...int) {
	fmt.Print(n, " ")
	total := 0
	for _, value := range n {
		total += value
	}
	fmt.Print(total)
	fmt.Println("")
}
func main() {
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4, 5, 6}

	sum(nums...)
}
