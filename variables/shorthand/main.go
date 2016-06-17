package main

import "fmt"

func main() {
	var b string
	a := 10
	d, f := true, "a"

	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", d)

	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", d)
	fmt.Printf("%v \n", b)
	fmt.Println(a, b, d, f)

}
