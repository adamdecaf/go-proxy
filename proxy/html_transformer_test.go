package proxy

import (
	"io/ioutil"
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

	answer := `<html><head></head><body><p>links</p><ul><li><a href="/url/Zm9v">Foo</a></li><li><a href="/url/L2Jhci9iYXo=">BarBaz</a></li></ul></body></html>`

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

	answer := `<html><head></head><body><img src="/url/Zm9v"/></body></html>`

	res := string(resp)

	if res != answer {
		t.Fatalf("parsed response '%s' doens't match answer\n", res)
	}
}
