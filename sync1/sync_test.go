package sync1

import (
	"fmt"
	"sync"
	"testing"
)

func TestCount(t *testing.T) {
	var count = 0
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				count++
			}
		}()
	}
	wg.Wait()
	fmt.Println(count)
}

func TestMutex(t *testing.T) {
	var (
		mu    sync.Mutex
		count = 0
		wg    sync.WaitGroup
	)
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				mu.Lock()
				count++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println(count)

}

func TestMissLock(t *testing.T) {
	var mu sync.Mutex
	defer mu.Unlock()
	fmt.Println("hello")
}
