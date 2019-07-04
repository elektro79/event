package event

import (
	"testing"
)

const EVTEST = 1

var sum = 0

func TestEvent(t *testing.T) {
	st1 := &stTest{str: "1"}
	st2 := &stTest{str: "2"}
	st3 := &stTest{str: "3"}
	ev := NewEvent()
	ev.On(EVTEST, testEv)
	f2 := ev.On(EVTEST, testEv2)
	ev.On(EVTEST, st1.testEv)
	stf2 := ev.On(EVTEST, st2.testEv)
	stf3 := ev.On(EVTEST, st3.testEv)
	ev.Fire(EVTEST, t, "testit")
	ev.Off(f2)
	ev.Off(stf2)
	ev.Off(stf3)
	ev.Fire(EVTEST, t, "testit2")
	if sum != 7 {
		t.Fatal("Sum should be 7, is:", sum)
	}
}

func testEv(l ...interface{}) {
	l[0].(*testing.T).Log("testEv:", l[1])
	sum++
}

func testEv2(l ...interface{}) {
	l[0].(*testing.T).Log("testEv2:", l[1])
	sum++
}

type stTest struct {
	str string
}

func (s *stTest) testEv(l ...interface{}) {
	l[0].(*testing.T).Log("sTtestEv:", s.str, l[1])
	sum++
}
