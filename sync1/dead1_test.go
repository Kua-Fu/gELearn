package sync1

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestDead1(t *testing.T) {

	var (
		mu1 sync.Mutex
		mu2 sync.Mutex
		wg  sync.WaitGroup
	)
	wg.Add(2)

	go func() {
		defer wg.Done()
		mu1.Lock()
		defer mu1.Unlock()

		time.Sleep(time.Second * 5)
		mu2.Lock()
		mu2.Unlock()
	}()

	go func() {
		defer wg.Done()
		mu2.Lock()
		defer mu2.Unlock()

		time.Sleep(time.Second * 5)
		mu1.Lock()
		mu1.Unlock()
	}()

	wg.Wait()
	fmt.Println("success!")
}
