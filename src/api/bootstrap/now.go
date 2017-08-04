package bootstrap

import "time"

//StartTime float64
var StartTime float64

// Now func to get unixtimestamp
func Now() float64 {
	myTime := float64(time.Now().UnixNano())
	if StartTime < 0.000001 {
		StartTime = myTime
	}
	return myTime - StartTime
}
