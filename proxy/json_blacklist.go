package proxy

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"os"
	"reflect"
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

	// parsed internal data
	blockedIPs []net.IP
	blockedSubnets []net.IPNet
}

func (b JSONBlacklist) IsBlacklisted(req Request) *error {
	// check outbound hostnames
	host := removePort(req.URL.Host)
	for i := range b.data.Hostnames {
		if host == b.data.Hostnames[i] {
			return &HostnameBlacklisted
		}
	}

	// check incoming source ips
	for i := range b.blockedIPs {
		if reflect.DeepEqual(b.blockedIPs[i], req.SourceAddress) {
			log.Printf("rejecting request from source ip %s", req.SourceAddress)
			return &SourceAddressBlacklisted
		}
	}

	// check incoming source ip against blocked subnets
	for i := range b.blockedSubnets {
		if b.blockedSubnets[i].Contains(req.SourceAddress) {
			log.Printf("rejecting request from source ip against subnet %s", req.SourceAddress)
			return &SourceAddressBlacklisted
		}
	}

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

	// parse sourceIPs array
	blockedIPs, blockedSubnets := parseBlacklistedIPs(data.SourceIPs)

	// create the blacklist
	bl := JSONBlacklist{
		data: data,
		blockedIPs: blockedIPs,
		blockedSubnets: blockedSubnets,
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

// incoming ips can look like: 127.0.0.1 or 10.0.0.0/8
func parseBlacklistedIPs(ips []string) ([]net.IP, []net.IPNet) {
	var blocked_ips []net.IP
	var blocked_subnets []net.IPNet

	for i := range ips {
		// build up ip or subnet
		if strings.Contains(ips[i], "/") {
			_, cidrnet, err := net.ParseCIDR(ips[i])
			if err != nil {
				log.Printf("error parsing ip '%s', err=%s\n", ips[i], err)
			}
			if cidrnet != nil {
				blocked_subnets = append(blocked_subnets, *cidrnet)
			}
		} else {
			ip := net.ParseIP(ips[i])
			if ip != nil {
				blocked_ips = append(blocked_ips, ip)
			}
		}
	}

	return blocked_ips, blocked_subnets
}
