package regexp1

import (
	"fmt"
	"regexp"
	"testing"
)

func TestExp(t *testing.T) {
	inputs := []string{
		"select f1 from t1",
		"select f11, f12 from t11;",
	}
	re := regexp.MustCompile("[\\s,;]+")
	for _, input := range inputs {
		l := re.Split(input, -1)
		var a []string
		for _, b := range l {
			if b != "" {
				a = append(a, b)
			}
		}
		fmt.Println(a)
	}
}

func TestCO(t *testing.T) {

	CObjectRegExp := "^(extra_).*"
	s1 := "extra_f1"
	match, err := regexp.MatchString(CObjectRegExp, s1)
	fmt.Println("match--", match, err)
}
