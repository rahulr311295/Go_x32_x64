package main

import "fmt"

func funcHello() int {
	return 10
}
func funcPrint() {
	fmt.Println("Hello")
}

func funReturn() (int, string) {
	return 10, "hello"
}

func sum(a, b int) int {
	return a + b
}

func factorial(n int) (result int) {
	if n > 1 {
		result := n * factorial(n-1)
		return result
	}
	return 1
}

func main() {
	n := funcHello()
	fmt.Println("Hello ", n)

	funcPrint()

	a, b := funReturn()
	fmt.Println(b, a)

	sum := sum(10, 5)
	fmt.Println("Sum is", sum)

	fmt.Println(factorial(10))

}
