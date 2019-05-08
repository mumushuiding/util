package util

import "time"

const (
	YYYY_MM_DD          = "2006-01-02"
	YYYY_MM_DD_HH_MM_SS = "2006-01-02 15:04:05"
)

// FormatDate 将日期转换成指定格式的字符串
func FormatDate(date time.Time, format string) string {
	return date.Format(format)
}
