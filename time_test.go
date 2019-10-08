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
