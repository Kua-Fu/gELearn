package assert1

import (
	"fmt"
	"testing"
)

func TestAssert1(t *testing.T) {
	var s interface{}
	s = 1609145234000
	rs, ok := s.(int)
	fmt.Println(rs, ok)
}
