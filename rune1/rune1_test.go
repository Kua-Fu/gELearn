package rune1

import (
	"fmt"
	"testing"
)

func TestRune(t *testing.T) {

	s := ".?+*|"
	r := []rune(s)
	for _, i := range r {
		fmt.Printf("%v\n", i)
		ns := []rune{i}
		fmt.Println(string(ns))
	}
}
