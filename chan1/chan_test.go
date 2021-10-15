package chan1

import (
	"fmt"
	"testing"
	"time"
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

func f1(queryChan chan string) {
	for {
		fmt.Println("----f1 enter--")
		v, ok := <-queryChan
		if ok {
			fmt.Println("----f1 v---", v, v == "")
		}
		time.Sleep(1 * time.Second)
	}
}

func f2(queryChan chan string) {

}

func TestChan1(t *testing.T) {
	queryChan := make(chan string) // 创建一个管道，保存是否可以查询

	go f1(queryChan)

	go f2(queryChan)
	for {

	}

}
