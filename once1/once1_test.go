package once1

import (
	"fmt"
	"sync"
	"testing"
)

// once常常用于初始化单例资源，
// 或者并发访问只需要初始化一次的共享资源，
// 或者在测试的时候初始化一次测试资源
func TestOnce(t *testing.T) {
	var once sync.Once

	// 第一个初始化函数
	f1 := func() {
		fmt.Println("in f1")
	}

	once.Do(f1) //执行f1

	// 第二个初始化函数
	f2 := func() {
		fmt.Println("in f2")
	}

	once.Do(f2) // 不会被执行
}

type Once struct {
	done uint32
	m    sync.Mutex
}
