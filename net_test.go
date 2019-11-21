package utils

import (
	"net"
	"reflect"
	"testing"
)

const (
	localhostV4 = "127.0.0.1"
)

func TestNet(t *testing.T) {
	_ = Net()
}

func TestIpT_UnmarshalText(t *testing.T) {
	t.Helper()
	testNet := _net

	ip, err := testNet.IP.UnmarshalText([]byte(localhostV4))
	if err != nil {
		t.Errorf("test: TestIpT_UnmarshalText(): Net.IP.UnmarshalText()")
	}
	// check
	ip2 := &net.IP{}
	_ = ip2.UnmarshalText([]byte(localhostV4))
	if !reflect.DeepEqual(ip, ip2) {
		t.Errorf("test: TestIpT_UnmarshalText(): %#v not equal %#v", ip, ip2)
	}
}
