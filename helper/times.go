package helper

import "time"

func TimeToMillis(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}
