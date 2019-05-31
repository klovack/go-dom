package godom

import (
	"io"

	class "github.com/klovack/go-dom/godom/class"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// Element contains HTML objects, object collections, and properties
type Element struct {
	ID        string
	Class     []string
	Attr      []html.Attribute
	ClassList class.List
	Children  []*Element

	node            *html.Node
	anchors         []*Element
	baseURI         string
	body            *Element
	docType         DocTypeDeclaration
	documentElement *Element
	forms           []*Element
	head            *Element
	links           []*Element
	images          []*Element
	scripts         []*Element
	title           *Element
	url             string
}

// ParseDocument parses the HTML and convert it to godom element
func ParseDocument(r io.Reader) *Element {
	doc, err := html.Parse(r)
	if err != nil {
		return &Element{}
	}

	newDoc := ParseElement(doc)
	return newDoc
}

// ParseElement converts the html.node to godom element
func ParseElement(n *html.Node) *Element {
	children := fetchChildren(n)

	return &Element{
		node:     n,
		Children: children,
	}
}

func (e *Element) Anchors() []*Element {
	if e.node == nil {
		return nil
	}

	if e.anchors != nil && len(e.anchors) > 0 {
		return e.anchors
	}

	var anchors []*Element
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.DataAtom == atom.A {
			el := ParseElement(n)
			anchors = append(anchors, el)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(e.node)

	return anchors
}

func (e *Element) BaseURI() string {
	return ""
}

func (e *Element) Body() *Element {
	return nil
}

func (e *Element) DocType() DocTypeDeclaration {
	return None
}

func (e *Element) DocumentElement() *Element {
	return nil
}

func (e *Element) Forms() []*Element {
	return nil
}

func (e *Element) Head() *Element {
	return nil
}

func (e *Element) Links() []*Element {
	return nil
}

func (e *Element) Images() []*Element {
	return nil
}

func (e *Element) Scripts() []*Element {
	return nil
}

func (e *Element) Title() []*Element {
	return nil
}

func (e *Element) URL() string {
	return ""
}

func GetElementById(id string) *Element {
	return nil
}
func GetElementsByClassName(classname string) []*Element {
	return nil
}
func GetElementsByName(name string) []*Element {
	return nil
}
func GetElementsByTagName(tagname string) []*Element {
	return nil
}
func GetElementsByTagNameNS(namespace string) []*Element {
	return nil
}

func QuerySelector(selector string) *Element {
	return nil
}
func QuerySelectorAll(selector string) []*Element {
	return nil
}
