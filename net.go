package utils

import "net"

func Net() *netT {
	return &_net
}

var _net = netT{
	IP: &ipT{},
}

type netT struct {
	IP *ipT
}

type ipT struct{}

// UnmarshalText is wrapper of *net.IP.UnmarshalText()
func (ipT) UnmarshalText(text []byte) (*net.IP, error) {
	ip := &net.IP{}
	return ip, ip.UnmarshalText(text)
}
