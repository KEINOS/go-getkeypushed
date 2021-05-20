package gkp

import "time"

// startTimer returns the pointer to the timer obj which will trigger in timeSec seconds.
func startTimer(timeSec int) *time.Timer {
	return time.NewTimer(time.Duration(timeSec) * time.Second)
}
