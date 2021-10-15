package math1

import (
	"fmt"
	"math"
	"testing"
)

func TestMathLog(t *testing.T) {
	a1 := math.Log(0)
	a2 := math.Log(2)
	a3 := math.Log(1) // 对数函数，默认的底数是自然对数e
	fmt.Println(a1, a2, a3)

}
