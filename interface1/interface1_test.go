package interface1

import (
	"fmt"
	"testing"
)

// 创建一个接口a
// 包含两个方法, a1, a2
type a interface {
	a1() string
	a2() string
}

// 创建一个接口b
type b interface {
	// a
	a1() string
	a2() string
}

type c interface {
	a
	b
}

func TestDuplicate(t *testing.T) {
	fmt.Println("111")
}
