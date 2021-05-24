package test

import (
	"fmt"
	"testing"
)

// TestStringAndInteger 11
func TestStringAndInteger(t *testing.T) {
	s1 := "111"
	s2 := 111
	r1 := fmt.Sprintf("%v", s1)
	r2 := fmt.Sprintf("%v", s2)
	if r1 == r2 {
		fmt.Println("equal--")
	}
	// fmt.Sprintf("%v", interface{}{})
	r3 := fmt.Sprintf("%v", nil)
	fmt.Println("--r3--", r3)
	r4 := map[string]int{r3: 1}
	fmt.Println(r4)
}
