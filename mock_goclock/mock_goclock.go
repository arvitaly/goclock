package mock_goclock

import (
	"time"

	"github.com/arvitaly/goclock"
)

type MockClock struct {
	goclock.Clock
	now    time.Time
	timers []*MockTimer
}
type MockTimer struct {
	start    time.Time
	duration time.Duration
	ready    chan time.Time
}
type MockTick struct {
}

//Create
func NewMockClock() *MockClock {
	var m = &MockClock{Clock: goclock.NewClock()}
	return m
}
func NewMockTimer(clock *MockClock, d time.Duration) *MockTimer {
	var t = new(MockTimer)
	t.duration = d
	t.start = clock.now
	t.ready = make(chan time.Time)
	clock.timers = append(clock.timers, t)
	return t
}

func (m *MockClock) Force(d time.Duration) {
	m.now = m.now.Add(d)
	var timers2 = []*MockTimer{}
	for _, timer := range m.timers {
		if m.now.After(timer.start.Add(timer.duration)) {
			timer.ready <- m.now
		} else {
			timers2 = append(timers2, timer)
		}
	}
	m.timers = timers2
}
func (m *MockClock) Sleep(d time.Duration) {
	var timer = NewMockTimer(m, d)
	<-timer.ready
}
func (m *MockClock) AfterFunc(d time.Duration, f func()) goclock.Timer {
	var timer = NewMockTimer(m, d)
	go func() {
		//TODO Add timer stop
		<-timer.ready
		f()
	}()
	return timer
}
func (m *MockClock) After(d time.Duration) <-chan time.Time {
	var timer = NewMockTimer(m, d)
	var timeReady = make(chan time.Time)
	go func() {
		timeReady <- <-timer.ready
	}()
	return timeReady
}
