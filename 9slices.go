package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("Empty Slice ", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("Assignment ", s)

	s = append(s, "d")
	fmt.Println("Append ", s)

	s = append(s, "e", "f", "g", "h")
	fmt.Println("Appending multiple values ", s)

	c := make([]string, len(s))
	copy(c, s)

	l := s[2:5]
	fmt.Println(l)

	l = s[:5]
	fmt.Println(s)

	l = s[2:]
	fmt.Println(s)

	t := []string{"a", "b", "c", "d"}
	fmt.Println("Declaration ", t)

	// TwoD Array
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		twoD[i] = make([]int, i+1)
		for j := 0; j < i+1; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("TwoD Array ", twoD)

}
