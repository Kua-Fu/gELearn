package time1

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestTicker1(t *testing.T) {

	ticker := time.NewTicker(3 * time.Second)
	fmt.Println("当前时间是: ", time.Now())

	go func() {
		for {
			// 从定时器中获取时间
			t := <-ticker.C
			fmt.Println("新的时间是: ", t)
		}
	}()

	// for {
	// 	time.Sleep(time.Second * 1)
	// }
}

func TestTimeRange(t *testing.T) {
	t1 := int64(1622542154889753000)
	t2 := t1 / 1e9
	t3 := time.Unix(t2, 0)
	fmt.Println(t3.Format("2006-01-02 15:04:05"))
	t0 := t3.Add(time.Hour * -1)
	// 获取前一个小时的起点时间
	start := time.Date(
		t0.Year(), t0.Month(), t0.Day(),
		t0.Hour(), 0, 0, 0,
		t0.Location(),
	)
	// 获取前一个小时的终点时间
	end := time.Date(
		t0.Year(), t0.Month(), t0.Day(),
		t0.Hour(), 59, 59, 0,
		t0.Location(),
	)

	startV := start.Unix()
	endV := end.Unix()
	fmt.Println(end.Format("2006-01-02 15:04:05"), endV)
	fmt.Println(start.Format("2006-01-02 15:04:05"), startV)

}

func TestTime3(t *testing.T) {
	t1 := "2021-05-25"
	t11, err := time.Parse("2006-01-02", t1)
	if err != nil {
		fmt.Println("err--", err)
	}
	t2 := "10d"
	if t2[len(t2)-1:len(t2)] == "d" {
		vv, _ := strconv.Atoi(t2[:len(t2)-1])
		vv = vv * 24 // 改成小时
		t2 = fmt.Sprintf("%dh", vv)
		fmt.Println("天转换为小时", t2)
	}
	t22, err := time.ParseDuration(t2)
	if err != nil {
		fmt.Println("err---", err)
	}
	v := t11.Add(t22)
	v1 := v.Unix() // 过期时间

	now := time.Now()
	n1 := time.Date(
		now.Year(), now.Month(), now.Day(),
		0, 0, 0, 0,
		now.Location(),
	)
	v2 := n1.Unix()
	if v1 > v2 {
		fmt.Println("没有过期", v1, v2)
	} else {
		fmt.Println("过期了", v1, v2)
	}
}

func TestFirstSleepTime(t *testing.T) {
	s := checkAndGetSTime1()
	hm := strings.Split(s, ":")
	hour, _ := strconv.Atoi(hm[0])
	minute, _ := strconv.Atoi(hm[1])
	hour = hour + 24
	eValue := hour*60 + minute
	now := time.Now()
	cHour := now.Hour()
	cMinute := now.Minute()
	cValue := cHour*60 + cMinute
	dValue := eValue - cValue
	fmt.Printf("eValue %d, cValue %d, dValue %d", eValue, cValue, dValue)
	// return time.Minute * time.Duration(dValue)
}

func checkAndGetSTime1() string {
	s := "01:00"
	ns := "01:00"
	if ns != "" {
		hm := strings.Split(ns, ":")
		if len(hm) == 2 {
			hour, err := strconv.Atoi(hm[0])
			if err != nil || hour > 24 || hour < 0 {

				return s
			}
			minute, err := strconv.Atoi(hm[1])
			if err != nil || minute > 60 || minute < 0 {

				return s
			}
		}
	}
	return s

}
