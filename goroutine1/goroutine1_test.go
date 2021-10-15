package goroutine1

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"testing"
	"time"
)

// --------chan default-----
func TestDefault(t *testing.T) {

	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {

		select {
		case <-tick:
			fmt.Println("tick")
		case <-boom:
			fmt.Println("boom")
			return
		default:
			fmt.Println("----")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

// --------chan select-------
func fibonacci1(c, quit chan int) {
	x, y := 0, 1
	for {

		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}

}

func TestSelect(t *testing.T) {

	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci1(c, quit)
}

// --------chan close---------
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)

}

func TestClose(t *testing.T) {

	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

// ---------buffer chan--------
func TestBuffer(t *testing.T) {

	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func say(s string) {
	for i := 0; i < 5; i++ {
		if s == "hello" {
			time.Sleep(1 * time.Second)
		} else {
			time.Sleep(100 * time.Millisecond)
		}
		fmt.Println(s)
	}

}

func TestSay(t *testing.T) {
	go say("world")
	say("hello")
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum = sum + v
	}
	c <- sum
}

func TestSum(t *testing.T) {

	s := []int{1, 2, 3, 4, 5, 6}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)

}

func GetGoid() int64 {
	var (
		buf [64]byte
		n   = runtime.Stack(buf[:], false)
		stk = strings.TrimPrefix(string(buf[:n]), "goroutine ")
	)

	idField := strings.Fields(stk)[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Errorf("can not get goroutine id: %v", err))
	}

	return int64(id)
}

func TestDefer(t *testing.T) {
	defer f2()
	fmt.Printf("pid: %d, goid: %d, main func \n", os.Getpid(), GetGoid())
	go f1()
}

func f1() {
	fmt.Printf("pid: %d, goid: %d, f1 func \n", os.Getpid(), GetGoid())
}

func f2() {

	fmt.Printf("pid: %d, goid: %d, f2 func \n", os.Getpid(), GetGoid())
}
