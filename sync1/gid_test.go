package sync1

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"testing"

	"github.com/petermattis/goid"
)

func TestGetGoID(t *testing.T) {

	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("can not get goroutine id: %v", err))
	}
	fmt.Println("get goroutine id: ", id)
}

func TestGetGoID2(t *testing.T) {
	gid := goid.Get()
	fmt.Println("get goroutine id is ", gid)
}
