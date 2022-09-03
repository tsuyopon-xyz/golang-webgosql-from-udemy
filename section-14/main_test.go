package main

import (
	"fmt"
	"testing"
)

var Debug bool = false

func TestIsOne(t *testing.T) {
	i := 1
	if Debug {
		t.Skip("スキップ")
	}

	v := IsOne(i)
	fmt.Println(v)
	if !v {
		t.Errorf("%v != %v", i, 1)
	}
}
