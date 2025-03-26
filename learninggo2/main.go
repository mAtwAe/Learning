package main

import "fmt"

func main() {
	x := 10
	p := x         // p is a pointer to x
	fmt.Println(p) // prints 10 (dereferencing)
}
