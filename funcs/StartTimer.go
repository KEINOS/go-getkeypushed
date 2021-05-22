package funcs

import "time"

// startTimer returns the pointer to the timer obj which will trigger in timeSec seconds.
func StartTimer(timeSec int) *time.Timer {
	return time.NewTimer(time.Duration(timeSec) * time.Second)
}
