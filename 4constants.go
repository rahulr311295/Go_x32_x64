package main

import (
	"fmt"
	"math"
)

const s = "String"

func main() {
	fmt.Println("This is a ", s)

	const n = 500000
	const d = 5e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
