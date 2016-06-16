package main

import "fmt"

func main() {
	fmt.Println(42)
	//binary
	fmt.Printf("%d - %b \n", 42, 42)
	//hexadecimal
	fmt.Printf("%d \t %b \t %#X \n", 42, 42, 42)
	//loop
	for i := 1000000; i < 1000100; i++ {
		fmt.Printf("%d \t %b \t %x \n", i, i, i)
	}
	//encoding

	for i := 60; i < 122; i++ {
		fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
	}
}
