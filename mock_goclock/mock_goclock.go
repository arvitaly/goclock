// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/arvitaly/goclock (interfaces: Clock)

package mock_goclock

import (
	//	"log"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

type scheduled struct {
	ready     chan bool
	duration  time.Duration
	startTime time.Time
}

// Mock of Clock interface
type MockClock struct {
	scheduleds []scheduled
	now        time.Time
	ctrl       *gomock.Controller
	recorder   *_MockClockRecorder
}

// Recorder for MockClock (not exported)
type _MockClockRecorder struct {
	mock *MockClock
}

func NewMockClock(ctrl *gomock.Controller) *MockClock {
	mock := &MockClock{ctrl: ctrl}
	mock.recorder = &_MockClockRecorder{mock}
	return mock
}
func (_m *MockClock) Force(d time.Duration) {
	_m.now = _m.now.Add(d)
	_m.callScheduleds()
}
func (_m *MockClock) callScheduleds() {
	var newScheduleds = []scheduled{}

	for _, v := range _m.scheduleds {

		if _m.now.After(v.startTime.Add(v.duration)) {

			v.ready <- true
		} else {
			newScheduleds = append(newScheduleds, v)
		}
	}
	_m.scheduleds = newScheduleds
}

func (_m *MockClock) EXPECT() *_MockClockRecorder {
	return _m.recorder
}

func (_m *MockClock) After(_param0 time.Duration) <-chan time.Time {
	//	ret := _m.ctrl.Call(_m, "After", _param0)
	//	ret0, _ := ret[0].(<-chan time.Time)
	//	return ret0
	var t = make(chan time.Time)
	var ready = make(chan bool)
	_m.scheduleds = append(_m.scheduleds, scheduled{startTime: _m.now, duration: _param0, ready: ready})
	go func() {
		<-ready
		t <- time.Now()
	}()
	return t
}

func (_mr *_MockClockRecorder) After(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "After", arg0)
}

func (_m *MockClock) AfterFunc(_param0 time.Duration, _param1 func()) *time.Timer {
	ret := _m.ctrl.Call(_m, "AfterFunc", _param0, _param1)
	ret0, _ := ret[0].(*time.Timer)
	return ret0
}

func (_mr *_MockClockRecorder) AfterFunc(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AfterFunc", arg0, arg1)
}

func (_m *MockClock) Date(_param0 int, _param1 time.Month, _param2 int, _param3 int, _param4 int, _param5 int, _param6 int, _param7 *time.Location) time.Time {
	ret := _m.ctrl.Call(_m, "Date", _param0, _param1, _param2, _param3, _param4, _param5, _param6, _param7)
	ret0, _ := ret[0].(time.Time)
	return ret0
}

func (_mr *_MockClockRecorder) Date(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Date", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

func (_m *MockClock) FixedZone(_param0 string, _param1 int) *time.Location {
	ret := _m.ctrl.Call(_m, "FixedZone", _param0, _param1)
	ret0, _ := ret[0].(*time.Location)
	return ret0
}

func (_mr *_MockClockRecorder) FixedZone(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FixedZone", arg0, arg1)
}

func (_m *MockClock) LoadLocation(_param0 string) (*time.Location, error) {
	ret := _m.ctrl.Call(_m, "LoadLocation", _param0)
	ret0, _ := ret[0].(*time.Location)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClockRecorder) LoadLocation(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LoadLocation", arg0)
}

func (_m *MockClock) NewTicker(_param0 time.Duration) *time.Ticker {
	ret := _m.ctrl.Call(_m, "NewTicker", _param0)
	ret0, _ := ret[0].(*time.Ticker)
	return ret0
}

func (_mr *_MockClockRecorder) NewTicker(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NewTicker", arg0)
}

func (_m *MockClock) NewTimer(_param0 time.Duration) *time.Timer {
	ret := _m.ctrl.Call(_m, "NewTimer", _param0)
	ret0, _ := ret[0].(*time.Timer)
	return ret0
}

func (_mr *_MockClockRecorder) NewTimer(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NewTimer", arg0)
}

func (_m *MockClock) Now() time.Time {
	ret := _m.ctrl.Call(_m, "Now")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

func (_mr *_MockClockRecorder) Now() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Now")
}

func (_m *MockClock) Parse(_param0 string, _param1 string) (time.Time, error) {
	ret := _m.ctrl.Call(_m, "Parse", _param0, _param1)
	ret0, _ := ret[0].(time.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClockRecorder) Parse(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Parse", arg0, arg1)
}

func (_m *MockClock) ParseDuration(_param0 string) (time.Duration, error) {
	ret := _m.ctrl.Call(_m, "ParseDuration", _param0)
	ret0, _ := ret[0].(time.Duration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClockRecorder) ParseDuration(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ParseDuration", arg0)
}

func (_m *MockClock) ParseInLocation(_param0 string, _param1 string, _param2 *time.Location) (time.Time, error) {
	ret := _m.ctrl.Call(_m, "ParseInLocation", _param0, _param1, _param2)
	ret0, _ := ret[0].(time.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockClockRecorder) ParseInLocation(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ParseInLocation", arg0, arg1, arg2)
}

func (_m *MockClock) Since(_param0 time.Time) time.Duration {
	ret := _m.ctrl.Call(_m, "Since", _param0)
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

func (_mr *_MockClockRecorder) Since(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Since", arg0)
}

func (_m *MockClock) Sleep(_param0 time.Duration) {
	//_m.ctrl.Call(_m, "Sleep", _param0)
	var ready = make(chan bool)
	_m.scheduleds = append(_m.scheduleds, scheduled{startTime: _m.now, duration: _param0, ready: ready})
	<-ready
}

func (_mr *_MockClockRecorder) Sleep(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Sleep", arg0)
}

func (_m *MockClock) Tick(_param0 time.Duration) <-chan time.Time {
	ret := _m.ctrl.Call(_m, "Tick", _param0)
	ret0, _ := ret[0].(<-chan time.Time)
	return ret0
}

func (_mr *_MockClockRecorder) Tick(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Tick", arg0)
}

func (_m *MockClock) Unix(_param0 int64, _param1 int64) time.Time {
	ret := _m.ctrl.Call(_m, "Unix", _param0, _param1)
	ret0, _ := ret[0].(time.Time)
	return ret0
}

func (_mr *_MockClockRecorder) Unix(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Unix", arg0, arg1)
}