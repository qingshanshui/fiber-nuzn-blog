package utils

import "time"

// GetDay 返回当前的时间
func GetDay() string {
	//template := "2006-01-02"
	template := "2006:01:02 15:04:05" // 标准模板

	return time.Now().Format(template)
}

// GetUnixNano 返回当前纳秒
func GetUnixNano() int64 {
	return time.Now().UnixNano()
}

// GetUnix 返回当前时间戳
func GetUnix() int64 {
	return time.Now().Unix()
}

// UnixDate 时间戳转时间
func UnixDate(unix int64) string {
	template := "2006-01-02" // 标准模板
	return time.Unix(unix, 0).Format(template)
}
