package scraper

import (
	"errors"
	"net"
	"net/url"
	"slices"
)

func ValidateURL(raw string) (*url.URL, error) {
	parsed, err := url.ParseRequestURI(raw)
	if err != nil {
		return nil, err
	}

	if parsed.Scheme != "http" && parsed.Scheme != "https" {
		return nil, errors.New("the provided URL points to a restricted internal network")
	}

	host := parsed.Hostname()
	ips, err := net.LookupIP(host)

	if err != nil {
		return nil, errors.New("DNS Resolution failed")
	}

	if slices.ContainsFunc(ips, isRestrictedIP) {
		return nil, errors.New("Invalid Resolved IP")
	}

	return parsed, nil
}

// A URL is restricted if it is:
// - A loopback (127.0.0.1)
// - A private network (192.168.x.x, etc.)
// - Link-local (fe80::) - often used for internal cloud metadata services
// - NOT Global Unicast (meaning it's multicast or reserved)

func isRestrictedIP(ip net.IP) bool {
	return ip.IsPrivate() || ip.IsLinkLocalUnicast() || ip.IsLoopback()
}
