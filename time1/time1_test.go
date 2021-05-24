package time1

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeSub(t *testing.T) {
	now := time.Now()
	t1 := now.UnixNano() / int64(time.Millisecond)
	beforeTime, _ := time.ParseDuration("-5m")
	t2 := now.Add(beforeTime)
	t3 := t2.UnixNano() / int64(time.Microsecond)
	t4 := t2.Second()
	sub := now.Sub(t2)
	fmt.Println("sub ", sub.Seconds())
	fmt.Println("t1 ", t1)
	fmt.Println("t2 ", t2)
	fmt.Println("t3 ", t3)
	fmt.Println("t4 ", t4)
}
