package time2

import "time"

// NowUTCUnix 获取当时UTC时间戳
func NowUTCUnix() int64 {
	return time.Now().UTC().Unix()
}
