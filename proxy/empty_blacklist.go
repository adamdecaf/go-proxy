package proxy

// EmptyBlacklist allows all requests through. This isn't recommended for most
// deployments because it offers no prevention of clients or requests.
//
// However, right now it is used when there is an error loading another blacklist.
type EmptyBlacklist struct {
	Blacklist
}

func (b EmptyBlacklist) IsBlacklisted(req Request) *error {
	return nil
}

// NewEmptyBlacklist returns an empty instance of EmptyBlacklist
func NewEmptyBlacklist() Blacklist {
	return EmptyBlacklist{}
}
