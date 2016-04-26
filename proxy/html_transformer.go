package proxy

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/adamdecaf/go-proxy/codec"
	"golang.org/x/net/html"
	"io"
	"log"
)

type HTMLTransformer struct {
	Transformer
}

func (t HTMLTransformer) Transform(in Response) Response {
	doc, err := html.Parse(in.Reader)
	if err != nil {
		log.Printf("error parsing html document err=%s\n", err)
	}

	var f func(*html.Node)

	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for i, a := range n.Attr {
				if a.Key == "href" {
					// todo: need to resolve
					// relative (/about.html)
					// and
					// protcol dependent urls (//foo.com)
					a.Val = fmt.Sprintf("/url/%s", codec.ToBase64(a.Val))
				}
				n.Attr[i] = a
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(doc)

	out := new(bytes.Buffer)
	err = html.Render(out, doc)
	if err != nil {
		// log.Printf
		return in
	}

	buff := []io.Reader{out}
	combined := io.MultiReader(buff...)

	return Response{
		Reader: combined,
	}
}

// `encodeBase64` returns a UTF-8 string in it's base64 encoding
func encodeBase64(s string) string {
	bytes := []byte(s)
	return base64.StdEncoding.EncodeToString(bytes)
}

func copyNode(old *html.Node) *html.Node {
	n := new(html.Node)
	n.Type = old.Type
	n.DataAtom = old.DataAtom
	n.Data = old.Data
	n.Attr = make([]html.Attribute, len(old.Attr))
	copy(n.Attr, old.Attr)
	for c := old.FirstChild; c != nil; c = c.NextSibling {
		nc := copyNode(c)
		n.AppendChild(nc)
	}
	return n
}

func replaceNode(old, new *html.Node) *html.Node {
	parent := old.Parent

	// nil would mean the root element
	if parent != nil {
		parent.InsertBefore(copyNode(new), old)
		parent.RemoveChild(old)
	}

	return new
}

func NewHTMLTransformer() Transformer {
	return HTMLTransformer{}
}
