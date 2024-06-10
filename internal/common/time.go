package common

import "time"

func FormatTime() string {
	t := time.Now()
	s2 := t.Format("2006-01-02 15:04:05")
	return s2
}

func FormatTimeByParam(t time.Time) string {
	s2 := t.Format("2006-01-02 15:04:05")
	return s2
}

func FormatFromTimeToStr(t *time.Time) string {
	s2 := t.Format("2006-01-02")
	return s2
}
