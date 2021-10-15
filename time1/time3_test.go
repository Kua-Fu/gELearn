package time1

import (
	"fmt"
	"testing"
	"time"
)

func TestDate(t *testing.T) {
	// 当前时间
	now := time.Now()
	// 获取前一天的时间
	// preDate := now.AddDate(0, 0, -1)
	preDate := now
	// 获取前一天的零点时间
	start := time.Date(
		preDate.Year(), preDate.Month(), preDate.Day(),
		0, 0, 0, 0,
		preDate.Location(),
	)
	// 获取前一天的24点时间
	end := time.Date(
		preDate.Year(), preDate.Month(), preDate.Day(),
		23, 59, 59, 0,
		preDate.Location(),
	)

	fmt.Println(start.Format("2006/01/02 15:04:05"))
	fmt.Println(end.Format("2006/01/02 15:04:05"))

	startTimestamp := start.Unix()
	endTimeStamp := end.Unix()
	fmt.Println(startTimestamp, endTimeStamp)

	fmt.Println(startTimestamp-8*60*60, endTimeStamp-8*60*60)

}

func TestRum(t *testing.T) {
	createTime := time.Unix(1623056400, 0)

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

	fmt.Println(start.Unix(), end.Unix())
}
