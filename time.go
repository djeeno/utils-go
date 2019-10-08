package utils

import "time"

var Time timeT

type timeT struct{}

// NowUnixTime returns now Unix Time and nanosecond.
func (timeT) NowUnixTime() (unixTime int64, nanosecond int64) {
	now := time.Now()
	return now.Unix(), int64(now.Nanosecond())
}

// NowUTC returns now UTC time zone date string formatted RFC3339Nano.
func (timeT) NowUTC() (utcRFC3339Nano string) {
	return time.Now().UTC().Format(time.RFC3339Nano)
}

// NowLocal returns now local time zone date string formatted RFC3339Nano.
func (timeT) NowLocal() (localRFC3339Nano string) {
	return time.Now().In(time.Local).Format(time.RFC3339Nano)
}
