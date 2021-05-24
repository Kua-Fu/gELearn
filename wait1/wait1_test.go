package wait1

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type Counter struct {
	mu    sync.Mutex
	count uint64
}

func (c *Counter) Incr() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

// 获取当前的计数
func (c *Counter) Count() uint64 {

	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func worker(c *Counter, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second)
	c.Incr()
}

func TestCount(t *testing.T) {

	var counter Counter
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {

		go worker(&counter, &wg)
	}

	// wait
	wg.Wait()
	fmt.Println(counter.Count())
}
