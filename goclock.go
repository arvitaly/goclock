package goclock

import "time"

type Ticker interface {
	Stop()
}

type Timer interface {
	/*Reset(d time.Duration) bool
	Stop() bool*/
}

type Clock interface {
	After(d time.Duration) <-chan time.Time
	Sleep(d time.Duration)
	Tick(d time.Duration) <-chan time.Time

	ParseDuration(s string) (time.Duration, error)
	Since(t time.Time) time.Duration

	FixedZone(name string, offset int) *time.Location
	LoadLocation(name string) (*time.Location, error)

	NewTicker(d time.Duration) Ticker

	Date(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) time.Time
	Parse(layout, value string) (time.Time, error)
	ParseInLocation(layout, value string, loc *time.Location) (time.Time, error)
	Unix(sec int64, nsec int64) time.Time

	AfterFunc(d time.Duration, f func()) Timer
	NewTimer(d time.Duration) Timer

	Now() time.Time
}
type _Clock struct {
}

func NewClock() Clock {
	return new(_Clock)
}

func (c _Clock) After(d time.Duration) <-chan time.Time {
	return time.After(d)
}
func (c _Clock) Sleep(d time.Duration) {
	time.Sleep(d)
}
func (c _Clock) Tick(d time.Duration) <-chan time.Time {
	return time.Tick(d)
}

func (c _Clock) ParseDuration(s string) (time.Duration, error) {
	return time.ParseDuration(s)
}
func (c _Clock) Since(t time.Time) time.Duration {
	return time.Since(t)
}

func (c _Clock) FixedZone(name string, offset int) *time.Location {
	return time.FixedZone(name, offset)
}
func (c _Clock) LoadLocation(name string) (*time.Location, error) {
	return time.LoadLocation(name)
}

func (c _Clock) NewTicker(d time.Duration) Ticker {
	return time.NewTicker(d)
}

func (c _Clock) Date(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) time.Time {
	return time.Date(year, month, day, hour, min, sec, nsec, loc)
}
func (c _Clock) Parse(layout, value string) (time.Time, error) {
	return time.Parse(layout, value)
}
func (c _Clock) ParseInLocation(layout, value string, loc *time.Location) (time.Time, error) {
	return time.ParseInLocation(layout, value, loc)
}
func (c _Clock) Unix(sec int64, nsec int64) time.Time {
	return time.Unix(sec, nsec)
}

func (c _Clock) AfterFunc(d time.Duration, f func()) Timer {
	return time.AfterFunc(d, f)
}
func (c _Clock) NewTimer(d time.Duration) Timer {
	return time.NewTimer(d)
}

func (c _Clock) Now() time.Time {
	return time.Now()
}
