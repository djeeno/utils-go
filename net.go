package utils

import "net"

var Net = netT{
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
