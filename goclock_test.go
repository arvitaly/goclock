package goclock

import "time"
import "testing"

func TestCreate(t *testing.T) {
	var time1 = time.Now()
	var clock = NewClock()
	clock.Sleep(time.Nanosecond)
	if !clock.Now().After(time1) {
		t.Error("Invalid time")
	}
}
