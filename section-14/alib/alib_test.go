package alib

import "testing"

var Dubug bool = false

func TestAverage(t *testing.T) {
	if Dubug {
		t.Skip("スキップ")
	}

	s := []int{1, 2, 3, 4, 5}
	avr := Average(s)

	if avr != 3 {
		t.Errorf("%v != %v", avr, 3)
	}
}
