package proxy

import (
	"errors"
)

var (
	// defaults
	DefaultBlacklist = NewBlacklist()

	// errors
	HostnameBlacklisted = errors.New("This hostname has been blacklisted.")
	SourceAddressBlacklisted = errors.New("This source address has been blacklisted.")
)

// Blacklist is a type which handles marking if a `Request` should not
// be processed. This most often happens based on hostname or source ip.
//
// See `DefaultBacklist` for a shared instance
type Blacklist interface {
	 IsBlacklisted(req Request) *error
}

func NewBlacklist() Blacklist {
	b, err := NewJSONBlacklist()
	if err != nil {
		return NewEmptyBlacklist()
	}
	return b
}
