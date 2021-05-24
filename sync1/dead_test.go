package sync1

import (
	"fmt"
	"sync"
	"testing"
)

type Counter struct {
	sync.Mutex
	Count int
}

func TestDeadStart(t *testing.T) {
	var c Counter
	c.Lock()
	defer c.Unlock()
	c.Count++
	f1(c) // 复制锁
}

func f1(c Counter) {
	c.Lock()
	defer c.Unlock()
	fmt.Println("in f1 func")
}
