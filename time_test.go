package utils

import "testing"

var testTime = Time

func TestTimeT_NowUnixTime(t *testing.T) {
	t.Helper()
	testTime.NowUnixTime()
}

func TestTimeT_NowUTC(t *testing.T) {
	t.Helper()
	testTime.NowUTC()
}

func TestTimeT_NowLocal(t *testing.T) {
	t.Helper()
	testTime.NowLocal()
}

func TestTimeT_SleepRandomMicrosecond(t *testing.T) {
	t.Helper()
	testTime.SleepRandomMicroseconds(1)
	testTime.SleepRandomMicroseconds(0)
	testTime.SleepRandomMilliseconds(1)
	testTime.SleepRandomMilliseconds(0)
}
