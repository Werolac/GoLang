package main

import "fmt"

func main() {
	var i int = 42
	fmt.Println(i)

	firstName := "Laci"
	fmt.Println(firstName)

	c := complex(3, 6)
	fmt.Println(c)

	j, k := real(c), imag(c)
	fmt.Println(j, k)

	test := new(string)
	*test = "777"
	fmt.Println(test, *test)

	test2 := &c
	fmt.Println(test2, *test2)
}
