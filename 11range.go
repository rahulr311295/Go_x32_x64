package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, num := range nums {
		sum += num
	}

	for index, num := range nums {
		fmt.Println("Index ", index, " Value ", num)
	}

	maps := map[string]string{"a": "apple", "b": "banana"}

	for index, value := range maps {
		fmt.Println(index, " -> ", value)
	}
}
