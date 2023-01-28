package stime

import (
	"fmt"
	"time"
)

type STime time.Time

func (t STime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

func (t *STime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+"2006-01-02 15:04:05"+`"`, string(data), time.Local)
	*t = STime(now)
	return
}

// 获取time.Time类型，方便拓展方法
func (t STime) Time() time.Time {
	return time.Time(t)
}

// 格式化
func (t STime) Format() string {
	return time.Time(t).Format("2006-01-02 15:04:05")
}

// 简单格式化
func (t STime) FormatSimple() string {
	return time.Time(t).Format("2006-01-02")
}
