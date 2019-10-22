package utils

import "testing"

func TestTimeT_NowUnixTime(t *testing.T) {
	Time.NowUnixTime()
}

func TestTimeT_NowUTC(t *testing.T) {
	Time.NowUTC()
}

func TestTimeT_NowLocal(t *testing.T) {
	Time.NowLocal()
}

func TestTimeT_SleepRandomMicrosecond(t *testing.T) {
	Time.SleepRandomMicroseconds(1)
	Time.SleepRandomMicroseconds(0)
	Time.SleepRandomMilliseconds(1)
	Time.SleepRandomMilliseconds(0)
}
