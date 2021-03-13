package helper

import "time"

// TimeToMillis convert time to millisecond
func TimeToMillis(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

// GetNowMillis get millisecond at current time
func GetNowMillis() int64 {
	return TimeToMillis(time.Now())
}
