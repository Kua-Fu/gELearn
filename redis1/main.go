package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis"
) //redis package
//connect redis
var client = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "abc123",
	DB:       0,
})

//lock
func lock(myfunc func()) {
	var lockKey = "mylockr"
	//lock
	lockSuccess, err := client.SetNX(lockKey, 1, time.Second*5).Result()
	if err != nil || !lockSuccess {
		fmt.Println("get lock fail")
		return
	}
	fmt.Println("get lock")

	//run func
	myfunc()
	//unlock
	_, err = client.Del(lockKey).Result()
	if err != nil {
		fmt.Println("unlock fail")
	} else {
		fmt.Println("unlock")
	}
}

//do action
var counter int64

func incr() {
	counter++
	fmt.Printf("after incr is %d\n", counter)
}

//5 goroutine compete lock
var wg sync.WaitGroup

func main() {
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			lock(incr)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("final counter is %d \n", counter)
	return
}
