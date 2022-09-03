package main

import (
	"fmt"
	"myapp/alib"
)

func IsOne(n int) bool {
	return n == 1
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(alib.Average(s))
}
