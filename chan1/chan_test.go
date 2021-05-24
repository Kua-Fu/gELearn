package chan1

import (
	"fmt"
	"testing"
)

func sum(s []int, c chan int) {

	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func TestChan(t *testing.T) {

	s := []int{1, 2, 3}
	c := make(chan int)
	go sum(s, c)
	go sum(s, c)

	x, y := <-c, <-c
	fmt.Println(x, y)
}
