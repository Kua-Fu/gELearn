package time1

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestNow(t *testing.T) {
	now := time.Now()
	fmt.Println(now.Unix())
	fmt.Println(now.Location())
	s, _ := time.Parse("2006-01-02 15:04:05", "2021-06-06 00:00:00")
	fmt.Println(s.Unix())

}

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

func TestTimeStr(t *testing.T) {

	now := time.Now().Unix()

	strNow := fmt.Sprintln(now)
	fmt.Println(strNow)
	fmt.Println()
}

func checkAndGetSTime() string {
	// s := "01:00"
	// ns := config.C.CC.Time
	// if ns != "" {
	// 	hm := strings.Split(ns, ":")
	// 	if len(hm) == 2 {
	// 		hour, err := strconv.Atoi(hm[0])
	// 		if err != nil {
	// 			l.Error("get post time err--", err)
	// 			return s
	// 		}
	// 		minute, err := strconv.Atoi(hm[1])
	// 		if err != nil {
	// 			l.Error("get post time err--", err)
	// 			return s
	// 		}
	// 	}
	// }
	// return s
	return ""

}

func TestSTime(t *testing.T) {
	s := checkAndGetSTime()
	hm := strings.Split(s, ":")
	if len(hm) > 1 {

	}
	hour, err := strconv.Atoi(hm[0])
	minute, err := strconv.Atoi(hm[1])
	if err != nil {
		fmt.Println("-----error--", err)
	}
	hour = hour + 24
	eValue := hour*60 + minute
	now := time.Now()
	cHour := now.Hour()
	cMinute := now.Minute()
	cValue := cHour*60 + cMinute
	dValue := eValue - cValue
	fmt.Printf("eValue %d, cValue %d, dValue %d", eValue, cValue, dValue)
}

func TestUpper(t *testing.T) {

	s := "post"
	ss := strings.ToUpper("post")
	fmt.Println(s, ss)
}
