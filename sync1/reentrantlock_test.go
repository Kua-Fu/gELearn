package sync1

import (
	"fmt"
	"sync"
	"testing"
)

func TestReentrantLock(t *testing.T) {
	l := &sync.Mutex{}
	reentrantLockF1(l)

}

func reentrantLockF1(l sync.Locker) {

	fmt.Println("in f1")
	l.Lock()
	reentrantLockF2(l)
	l.Unlock()
}

func reentrantLockF2(l sync.Locker) {

	l.Lock()
	fmt.Println("in f2")
	l.Unlock()
}
