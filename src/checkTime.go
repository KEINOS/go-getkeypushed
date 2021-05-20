package gkp

import "time"

// checkTime sets the value of keyDefault to returnChan if the timer exceeds it's wait time.
func checkTime(returnChan chan<- string, timer *time.Timer, keyDefault string) {
	<-timer.C // Continue the below lines if the timer exceeds

	returnChan <- keyDefault
}
