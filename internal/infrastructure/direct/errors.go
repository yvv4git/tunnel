package direct

import "errors"

var (
	ErrInvalidPlatform    = errors.New("got invalid platform")
	ErrInvalidChannelType = errors.New("got invalid channel type")
)
