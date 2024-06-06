package common

import "time"

func FormatTime() string {
	t := time.Now()
	s2 := t.Format("2006-01-02 00:00:00")
	return s2
}
