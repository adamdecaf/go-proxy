package proxy

import (
	"errors"
)

var (
	// errors
	HostnameBlacklisted = errors.New("This hostname has been blacklisted.")
	SourceAddressBlacklisted = errors.New("This hostname has been blacklisted.")
)

// Blacklist is ...
//
type Blacklist interface {
	Contains(string) bool
}

type FileBlacklist struct {
	Blacklist

	// Private
	items []string
}

func (b FileBlacklist) Contains(url string) bool {
	// todo: read from file
	return false
}

func BlacklistFromFile() (Blacklist, error) {
	return FileBlacklist{}, nil
}
