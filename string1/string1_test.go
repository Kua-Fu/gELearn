package string1

import (
	"fmt"
	"strings"
	"testing"
)

func TestLower(t *testing.T) {
	s1 := "test OR 111 or 2222"
	s2 := strings.ToLower(s1)
	fmt.Println("s1--", s1, "--s2--", s2)
	s3 := strings.ReplaceAll(s2, "or", "OR")
	fmt.Println("s3--", s3)

}

func TestSlice(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6}
	for i, v := range s[2:4] {
		fmt.Println("--", i, "--", v)
	}
}

func TestMap(t *testing.T) {
	m := map[string]string{

		"1": "11",
	}
	ss, ok := m["2"]
	fmt.Println("---", ss, ok)

}

func TestSplit(t *testing.T) {
	s := "111-222"
	ss := strings.Split(s, "-")
	if len(ss) > 0 {
		fmt.Println("--ss--", ss)
		new := strings.Join(ss[1:], "-")
		fmt.Println("--new--", new)
	}

}
