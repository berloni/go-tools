package tools

import (
	"errors"
	"net"
	"net/http"
	"strings"
)

// GetRequestIP gets the ip address of an incoming http request
func GetRequestIP(r *http.Request) (string, error) {
	ip := r.Header.Get("X-REAL-IP")
	netIP := net.ParseIP(ip)
	if netIP != nil {
		return ip, nil
	}

	ips := r.Header.Get("X-FORWARDED-FOR")
	splitIPs := strings.Split(ips, ",")
	for _, ip := range splitIPs {
		netIP := net.ParseIP(ip)
		if netIP != nil {
			return ip, nil
		}
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", nil
	}
	netIP = net.ParseIP(ip)
	if netIP != nil {
		return ip, nil
	}

	return "", errors.New("ip: no valid ip found")
}
