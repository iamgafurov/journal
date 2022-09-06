package tools

import "time"

//TodayStart return start of current day in time
func TodayStart() time.Time {
	return time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, time.Now().Second(), 0, time.UTC)
}
