package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 8

	fmt.Println("Map ", m)
	fmt.Println("K1 ", m["k1"])

	v1 := m["k2"]
	fmt.Println("V1 ", v1)
	fmt.Println("Length", len(m))

	delete(m, "k2")
	fmt.Println(m)

	_, prs := m["k2"]
	fmt.Println(prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map: ", n)
	fmt.Println("foo", n["foo"])
}
