package util

import "time"

func formatDateForDay(t time.Time) string {
	return t.Format("2006-01-02")
}

func FormatDate(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func TurnTime(ti int64) string {
	return FormatDate(time.Unix(ti, 0))
}

func TurnMysqlTime(timestamp int64) string {
	return FormatDate(time.Unix(0, timestamp*int64(time.Millisecond)))
}
