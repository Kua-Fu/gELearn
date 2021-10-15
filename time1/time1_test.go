package time1

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestTime2(t *testing.T) {

	now := time.Now() //获取当前时间，放到now里面，要给next用

	now_t := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0, 0, now.Location()) //获取整点
	fmt.Println("--now--", now, now_t, now_t.Unix())

	v := now_t.UnixNano()
	createTime := time.Unix(v/1e9, 0)
	fmt.Println(createTime.Format("2006-01-02 15:04:05"))
	preTime := createTime.Add(time.Hour * -1)
	// 获取前一个小时的起点时间
	start := time.Date(
		preTime.Year(), preTime.Month(), preTime.Day(),
		preTime.Hour(), 0, 0, 0,
		preTime.Location(),
	)
	// 获取前一个小时的终点时间
	end := time.Date(
		createTime.Year(), createTime.Month(), createTime.Day(),
		createTime.Hour(), 0, 0, 0,
		createTime.Location(),
	)

	rumStart := time.Date(
		createTime.Year(), createTime.Month(), createTime.Day(),
		0, 0, 0, 0,
		createTime.Location(),
	)
	fmt.Println("--start--", start, "--end--", end, "--rumstart--", rumStart)

}

func TestTimeRange1(t *testing.T) {

	now := time.Now()
	// 8*60*60 也可用 int((8 * time.Hour).Seconds()) 表示
	cusZone := time.FixedZone("UTC+8", 8*60*60)
	// now = now.In(cusZone)
	// loc, _ := time.LoadLocation("Asia/Shanghai")
	v := time.Now().UnixNano()
	now = time.Unix(v/1e9, 0).In(cusZone) // 保证是北京时间

	// 获取前一天的时间
	preDate := now.AddDate(0, 0, -1)
	// 获取前一天的零点时间
	start := time.Date(
		preDate.Year(), preDate.Month(), preDate.Day(),
		0, 0, 0, 0,
		preDate.Location(),
	)
	// 获取前一天的24点时间
	end := time.Date(
		now.Year(), now.Month(), now.Day(),
		0, 0, 0, 0,
		now.Location(),
	)

	// 获取前一天的零点时间
	start1 := time.Date(
		preDate.Year(), preDate.Month(), preDate.Day(),
		0, 0, 0, 0,
		now.UTC().Location(),
	)
	// 获取前一天的24点时间
	end1 := time.Date(
		now.Year(), now.Month(), now.Day(),
		0, 0, 0, 0,
		now.UTC().Location(),
	)

	fmt.Println("--start--", start, start.Unix(), "--end--", end, end.Unix())
	fmt.Println("--start--", start1, start1.Unix(), "--end--", end1, end1.Unix())
}

func TestNow(t *testing.T) {
	now := time.Now()
	fmt.Println(now.Unix())
	nNow := now.Add(time.Minute)
	fmt.Println("new now--", nNow.Unix())
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
