package utils

import (
	"math/rand"
	"time"
)

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

// SleepRandomMicroseconds sleeps for random microseconds.
func (timeT) SleepRandomMicroseconds(seed int64) {
	if seed == 0 {
		seed = int64(time.Now().Nanosecond())
	}
	rand.Seed(seed)
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Microsecond)
}

// SleepRandomMilliseconds sleeps for random milliseconds.
func (timeT) SleepRandomMilliseconds(seed int64) {
	if seed == 0 {
		seed = int64(time.Now().Nanosecond())
	}
	rand.Seed(seed)
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
}
