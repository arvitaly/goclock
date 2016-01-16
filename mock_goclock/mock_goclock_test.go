package mock_goclock

import "testing"
import "time"

func TestSleep(t *testing.T) {
	var clock = NewMockClock()
	var done = make(chan bool)
	go func() {
		done <- true
		clock.Sleep(time.Second * 10)
		done <- true
	}()
	<-done
	clock.Force(time.Second * 11)
	select {
	case <-time.After(time.Second):
		t.Error("Timeout")
	case <-done:

	}
}
func TestAfterCompleted(t *testing.T) {
	var clock = NewMockClock()

	//After should be completed
	var done = clock.After(time.Second * 1000)
	//Shift a time
	clock.Force(time.Second * 1100)
	select {
	case <-time.After(time.Nanosecond):
		t.Error("Timeout")
	case <-done:

	}
}
func TestAfterNotCompleted(t *testing.T) {
	var clock = NewMockClock()
	//After should not be completed
	var done = clock.After(time.Second * 1000)
	//Shift a time
	clock.Force(time.Second * 900)
	select {
	case <-time.After(time.Nanosecond):

	case <-done:
		t.Error("After don't be complete")
	}
}
