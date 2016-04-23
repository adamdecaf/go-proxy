package proxy

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	DefaultBlacklistJSONFile = "./proxy/blacklist.json"
)

// JSONBlacklist represents a blacklist that's read from a JSON blob from
// the local filesystem.
type JSONBlacklist struct {
	Blacklist

	// Private items
	data JSONBlacklistData
}

// todo: way better sorting -- prefix/radix tree
func (b JSONBlacklist) IsBlacklisted(req Request) *error {
	host := removePort(req.URL.Host)

	for i := range b.data.Hostnames {
		if host == b.data.Hostnames[i] {
			return &HostnameBlacklisted
		}
	}

	// todo: source addresses

	return nil
}

// JSONBlacklistData contains the json structure of the blacklist file to
// read and return inside the JSONBlacklist.
type JSONBlacklistData struct {
	Hostnames []string `json:"hostnames"`
	SourceIPs []string `json:"sourceIPs"`
}

// NewJSONBlacklist returns a defaulted instance of the `blacklist.json`
// file parsed.
func NewJSONBlacklist() (Blacklist, error) {
	return NewJSONBlacklistFromFile(DefaultBlacklistJSONFile)
}

// NewJSONBlacklistFromFile reads the given file and returns a `JSONBlacklist`
// instance.
func NewJSONBlacklistFromFile(file string) (Blacklist, error) {
	r, err := os.Open(file)
	if err != nil {
		log.Printf("error reading blacklist file '%s' (err=%s)\n", file, err)
	}

	bytes, err := ioutil.ReadAll(r)
	if err != nil {
		log.Printf("error reading json blacklist err=%s\n", err)
	}

	data := JSONBlacklistData{}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		log.Printf("error parsing blacklist json err=%s\n", err)
	}

	// create the blacklist
	bl := JSONBlacklist{
		data: data,
	}

	return bl, nil
}

func removePort(s string) string {
	loc := strings.IndexRune(s, ':')
	if loc > 0 {
		return s[:loc]
	}
	return s
}
