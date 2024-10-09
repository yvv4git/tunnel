package infrastructure

import "errors"

var (
	ErrInvalidHost     = errors.New("got invalid host")
	ErrInvalidPort     = errors.New("got invalid port")
	ErrInvalidPlatform = errors.New("got invalid platform")
)
