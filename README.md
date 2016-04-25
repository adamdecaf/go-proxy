# go-proxy

> A lightweight proxy in golang

## Usage

1. `go get github.com/adamdecaf/go-proxy`
1. `cd $GOPATH/src/github.com/adamdecaf/go-proxy`
1. `make build && go install`

Then you can run `go-proxy` (or `-p 8080` for some other port)

## Blacklist

Included are outgoing hostname and incoming source address blacklists. You can configure them under `proxy/blacklist.json`.

**Note** You do need to setup this proxy behind a load balancer (typically nginx or haproxy) as it expects to read `X-Forwarded-For` or other remote address / request forwarding headers. You can disable this behaviour be removing the blocked `sourceIPs` in the blacklist json file.
