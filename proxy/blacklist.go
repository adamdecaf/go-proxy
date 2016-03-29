package proxy

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
