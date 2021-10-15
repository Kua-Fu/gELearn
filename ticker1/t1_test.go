package ticker1

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker1(t *testing.T) {
	// 使用time.Ticker
	var ticker *time.Ticker = time.NewTicker(1 * time.Second)

	for t := range ticker.C {
		fmt.Println("Tick at", t)
	}

	// time.Sleep(time.Second * 5)
	// ticker.Stop()
	// fmt.Println("Ticker stopped")

}
