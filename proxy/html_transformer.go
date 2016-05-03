package proxy

import (
	"bytes"
	"fmt"
	"github.com/adamdecaf/go-proxy/codec"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/url"
)

type HTMLTransformer struct {
	Transformer
}

func (t HTMLTransformer) Transform(orig_url url.URL, in Response) Response {
	doc, err := html.Parse(in.Reader)
	if err != nil {
		log.Printf("error parsing html document err=%s\n", err)
	}

	replaceAHrefs(orig_url, doc)
	replaceImgSrcs(orig_url, doc)
	replaceScriptSrcs(orig_url, doc)
	replaceLinkHrefs(orig_url, doc)

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

// needs the `proxy.Request` object
func createProxyableUrl(orig_url url.URL, s string) string {
	p1, err := url.Parse(s)
	if err != nil {
		return s
	}

	// if there isn't a scheme append 'http'
	if p1.Scheme == "" {
		p1.Scheme = "http"
	}

	// reparse the url
	// this is because of urls like 'example.com' are parsed as a
	// Path, but it's under the Host
	p2, err := url.Parse(p1.String())
	if err != nil {
		return s
	}

	// when a url like 'example.com' is parsed by `url.Parse` it's Host is
	// blank and it's Path is set to 'example.com'
	if p2.Host == "" {
		if orig_url.Host == "" {
			p2.Host = orig_url.Path
		} else {
			p2.Host = orig_url.Host
		}
	}

	// use the scheme from the original url (if there is one)
	if orig_url.Scheme != "" {
		p2.Scheme = orig_url.Scheme
	}

	// append proxy specific url path prefix
	return fmt.Sprintf("/url/%s", codec.ToBase64(p2.String()))
}

// `replaceAHrefs` is a greedy depth-first search and replace for
// `href` attributes in `a` elements.
func replaceAHrefs(orig_url url.URL, n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for i, a := range n.Attr {
			if a.Key == "href" {
				a.Val = createProxyableUrl(orig_url, a.Val)
			}
			n.Attr[i] = a
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		replaceAHrefs(orig_url, c)
	}
}

// `replaceLinkHrefs` is a greedy depth-first search and replace for
// `href` attributes in `link` elements.
func replaceLinkHrefs(orig_url url.URL, n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "link" {
		for i, a := range n.Attr {
			if a.Key == "href" {
				a.Val = createProxyableUrl(orig_url, a.Val)
			}
			n.Attr[i] = a
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		replaceLinkHrefs(orig_url, c)
	}
}

// replaceImgSrcs is a greedy depth-first search and replace for
// `src` attributes in `img` elements.
func replaceImgSrcs(orig_url url.URL, n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "img" {
		for i, a := range n.Attr {
			if a.Key == "src" {
				a.Val = createProxyableUrl(orig_url, a.Val)
			}
			n.Attr[i] = a
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		replaceImgSrcs(orig_url, c)
	}
}

// replaceScriptSrcs is a greedy depth-first search and replace for
// `src` attributes in `script` elements.
func replaceScriptSrcs(orig_url url.URL, n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "script" {
		for i, a := range n.Attr {
			if a.Key == "src" {
				a.Val = createProxyableUrl(orig_url, a.Val)
			}
			n.Attr[i] = a
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		replaceScriptSrcs(orig_url, c)
	}
}

func NewHTMLTransformer() Transformer {
	return HTMLTransformer{}
}
