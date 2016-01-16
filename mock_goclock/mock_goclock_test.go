package mock_goclock

import "testing"
import "github.com/golang/mock/gomock"
import "time"

func TestSleep(t *testing.T) {
	var ctrl = gomock.NewController(t)
	var clock = NewMockClock(ctrl)
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
func TestAfter(t *testing.T) {
	var ctrl = gomock.NewController(t)
	var clock = NewMockClock(ctrl)
	var done = clock.After(time.Second * 1000)
	clock.Force(time.Second * 1100)
	select {
	case <-time.After(time.Nanosecond):
		t.Error("Timeout")
	case <-done:

	}
	//
	done = clock.After(time.Second * 1000)
	clock.Force(time.Second * 900)
	select {
	case <-time.After(time.Nanosecond):

	case <-done:
		t.Error("After don't be complete")
	}
}
