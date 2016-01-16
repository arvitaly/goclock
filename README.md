Goclock
Adapter for package "time" by interface with mocks

[![Build Status](https://travis-ci.org/arvitaly/goclock.svg?branch=master)](https://travis-ci.org/arvitaly/goclock)

Install:

	go get github.com/arvitaly/goclock

Lib contains two packages: goclock for real-code and mock_goclock for testing

Example usage of goclock:

	package main

	import (
		"time"

		"github.com/arvitaly/goclock"
	)

	type A struct {
	}

	func NewA(clock goclock.Clock) {
		clock.Sleep(time.Second)
	}

	func main() {
		var clock = goclock.New()
		NewA(clock)
	}

For mock_goclock has special method for shift time:

	Force(d time.Duration)

Example usage of mock_goclock:

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
