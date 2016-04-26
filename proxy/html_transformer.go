package proxy

import (
	"bytes"
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
		replaceAHrefs(n)
		replaceImgSrcs(n)
	}

	f(doc)

	out := new(bytes.Buffer)
	err = html.Render(out, doc)
	if err != nil {
		return in
	}

	buff := []io.Reader{out}
	combined := io.MultiReader(buff...)

	return Response{
		Reader: combined,
	}
}

// `replaceAHrefs` is a greedy depth-first search and replace for
// `href` attributes in `a` elements.
func replaceAHrefs(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for i, a := range n.Attr {
			if a.Key == "href" {
				a.Val = fmt.Sprintf("/url/%s", codec.ToBase64(a.Val))
			}
			n.Attr[i] = a
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		replaceAHrefs(c)
	}
}

// replaceImgSrcs is a greedy depth-first search and replace for
// `src` attributes in `img` elements.
func replaceImgSrcs(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "img" {
		for i, a := range n.Attr {
			if a.Key == "src" {
				a.Val = fmt.Sprintf("/url/%s", codec.ToBase64(a.Val))
			}
			n.Attr[i] = a
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		replaceImgSrcs(c)
	}
}

func NewHTMLTransformer() Transformer {
	return HTMLTransformer{}
}
