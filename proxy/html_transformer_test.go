package proxy

import (
	"fmt"
	"io/ioutil"
	"github.com/adamdecaf/go-proxy/codec"
	"strings"
	"testing"
)

func TestHTMLReplceLinks(t *testing.T) {
	str := `<html><p>links</p><ul><li><a href="foo">Foo</a><li><a href="/bar/baz">BarBaz</a></ul></html>`
	tr := NewHTMLTransformer()

	r := strings.NewReader(str)
	if r == nil {
		t.Fatalf("unable to create reader for str '%s'\n", str)
	}

	after := tr.Transform(Response{Reader: r})

	// read the response
	resp, err := ioutil.ReadAll(after.Reader)
	if err != nil {
		t.Fatalf("error reading transformed response err=%s\n", err)
	}

	answer := `<html><head></head><body><p>links</p><ul><li><a href="/url/aHR0cDovL2Zvbw==">Foo</a></li><li><a href="/url/aHR0cDovLy9iYXIvYmF6">BarBaz</a></li></ul></body></html>`

	res := string(resp)

	if res != answer {
		t.Fatalf("parsed response '%s' doens't match answer\n", res)
	}
}

func TestHTMLReplceLinkHrefs(t *testing.T) {
	str := `<html><head><link href="foo" /></head></html>`
	tr := NewHTMLTransformer()

	r := strings.NewReader(str)
	if r == nil {
		t.Fatalf("unable to create reader for str '%s'\n", str)
	}

	after := tr.Transform(Response{Reader: r})

	// read the response
	resp, err := ioutil.ReadAll(after.Reader)
	if err != nil {
		t.Fatalf("error reading transformed response err=%s\n", err)
	}

	answer := `<html><head><link href="/url/aHR0cDovL2Zvbw=="/></head><body></body></html>`

	res := string(resp)

	if res != answer {
		t.Fatalf("parsed response '%s' doens't match answer\n", res)
	}
}

func TestHTMLReplaceImages(t *testing.T) {
	str := `<html><img src="foo" /></html>`
	tr := NewHTMLTransformer()

	r := strings.NewReader(str)
	if r == nil {
		t.Fatalf("unable to create reader for str '%s'\n", str)
	}

	after := tr.Transform(Response{Reader: r})

	// read the response
	resp, err := ioutil.ReadAll(after.Reader)
	if err != nil {
		t.Fatalf("error reading transformed response err=%s\n", err)
	}

	answer := `<html><head></head><body><img src="/url/aHR0cDovL2Zvbw=="/></body></html>`

	res := string(resp)

	if res != answer {
		t.Fatalf("parsed response '%s' doens't match answer\n", res)
	}
}

func TestHTMLReplaceScript(t *testing.T) {
	str := `<html><head></head><body><script src="foo"></script></body></html>`
	tr := NewHTMLTransformer()

	r := strings.NewReader(str)
	if r == nil {
		t.Fatalf("unable to create reader for str '%s'\n", str)
	}

	after := tr.Transform(Response{Reader: r})

	// read the response
	resp, err := ioutil.ReadAll(after.Reader)
	if err != nil {
		t.Fatalf("error reading transformed response err=%s\n", err)
	}

	answer := `<html><head></head><body><script src="/url/aHR0cDovL2Zvbw=="></script></body></html>`

	res := string(resp)

	if res != answer {
		t.Fatalf("parsed response '%s' doens't match answer\n", res)
	}
}

func TestHTMLReplaceAllElements(t *testing.T) {
	str := `<html><head><link href="foo"/></head><body><img src="foo" /><script src="foo"></script><a href="foo">Foo</a></body></html>`
	tr := NewHTMLTransformer()

	r := strings.NewReader(str)
	if r == nil {
		t.Fatalf("unable to create reader for str '%s'\n", str)
	}

	after := tr.Transform(Response{Reader: r})

	// read the response
	resp, err := ioutil.ReadAll(after.Reader)
	if err != nil {
		t.Fatalf("error reading transformed response err=%s\n", err)
	}

	answer := `<html><head><link href="/url/aHR0cDovL2Zvbw=="/></head><body><img src="/url/aHR0cDovL2Zvbw=="/><script src="/url/aHR0cDovL2Zvbw=="></script><a href="/url/aHR0cDovL2Zvbw==">Foo</a></body></html>`
	res := string(resp)

	if res != answer {
		t.Fatalf("parsed response '%s' doens't match answer\n", res)
	}
}

func TestProxyableUrlCreation(t *testing.T) {
	res1 := createProxyableUrl("http://ashannon.us")
	if res1 != fmt.Sprintf("/url/%s", codec.ToBase64("http://ashannon.us")) {
		t.Fatalf("generated url doesn't match expected = '%s'", res1)
	}

	res2 := createProxyableUrl("ashannon.us")
	if res2 != fmt.Sprintf("/url/%s", codec.ToBase64("http://ashannon.us")) {
		t.Fatalf("generated url doesn't match expected = '%s'", res2)
	}
}
