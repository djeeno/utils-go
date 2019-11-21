package utils

import "testing"

func TestTime(t *testing.T) {
	_ = Time()
}

func TestTimeT_NowUnixTime(t *testing.T) {
	t.Helper()
	testTime := _time

	testTime.NowUnixTime()
}

func TestTimeT_NowUTC(t *testing.T) {
	t.Helper()
	testTime := _time

	testTime.NowUTC()
}

func TestTimeT_NowLocal(t *testing.T) {
	t.Helper()
	testTime := _time

	testTime.NowLocal()
}

func TestTimeT_SleepRandomMicrosecond(t *testing.T) {
	t.Helper()
	testTime := _time

	testTime.SleepRandomMicroseconds(1)
	testTime.SleepRandomMicroseconds(0)
	testTime.SleepRandomMilliseconds(1)
	testTime.SleepRandomMilliseconds(0)
}
