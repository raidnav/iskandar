package helper

import "time"

func ToMillis(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}
