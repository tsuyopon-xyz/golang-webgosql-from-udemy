package main

import "fmt"

func main() {
	var i int = -100
	var ui uint = 100
	var i8 int8 = 100

	fmt.Printf("%d %T %T %T\n", uint32(i), i, ui, i8)
}
