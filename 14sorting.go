package main

import (
	"fmt"
	"sort"
)

func main() {
	str := []string{"c", "b", "a"}
	sort.Strings(str)
	fmt.Println(str)

	sort.Sort(sort.Reverse(sort.StringSlice(str)))
	fmt.Println(str)

	numbers := []int{7, 5, 3, 2, 9}
	sort.Ints(numbers)

	fmt.Println(numbers)

	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
	fmt.Println(numbers)

	s := sort.IntsAreSorted(numbers)
	fmt.Println(s)
}
