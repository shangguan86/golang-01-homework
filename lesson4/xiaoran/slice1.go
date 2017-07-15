package main

import "fmt"

func main() {
	names := [4]string{
		"a",
		"b",
		"c",
		"d",
	}
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
	c := a[1:2]
	c[0] = "YYY"
}
