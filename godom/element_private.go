package godom

import (
	"golang.org/x/net/html"
)

func fetchChildren(n *html.Node) []*Element {
	var childrenList []*Element

	var f func(*html.Node)
	f = func(n *html.Node) {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
		childrenList = append(childrenList, ParseElement(n))
	}
	f(n)

	return childrenList
}
