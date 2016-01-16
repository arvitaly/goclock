Goclock
Adapter for package "time" by interface with mocks

Example:

func TestAfter(t *testing.T) {

	var ctrl = gomock.NewController(t)
	var clock = NewMockClock(ctrl)

	//After should be completed
	var done = clock.After(time.Second * 1000)
	//Shift a time
	clock.Force(time.Second * 1100)
	select {
	case <-time.After(time.Second):
		t.Error("Timeout")
	case <-done:

	}

	//After should not be completed
	done = clock.After(time.Second * 1000)
	clock.Force(time.Second * 900)
	select {
	case <-time.After(time.Second):

	case <-done:
		t.Error("After don't be complete")
	}

}
