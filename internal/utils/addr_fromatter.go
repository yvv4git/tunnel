package utils

import "fmt"

func FormatAddrString(host string, port uint16) (string, error) {
	if host == "" {
		return "", ErrInvalidHost
	}

	if port == 0 {
		return "", ErrInvalidPort
	}

	return fmt.Sprintf("%s:%d", host, port), nil
}
