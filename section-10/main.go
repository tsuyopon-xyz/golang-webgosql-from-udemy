package main

import "fmt"

func Double(i *int) {
	*i = *i * 2
}

func main() {
	n := 100
	Double(&n)

	fmt.Println(n)
}
