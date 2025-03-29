package link

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	// dfs(doc, "")
	nodes := linkNodes(doc)

	var links []Link
	for _, node := range nodes {
		// fmt.Println(node.Attr[0].Val)
		// fmt.Println(node.LastChild.Data)
		// link := Link{
		// 	Href: node.Attr[0].Val,
		// 	Text: node.LastChild.Data,
		// }
		var link Link
		for _, attr := range node.Attr {
			if attr.Key == "href" {
				link.Href = attr.Val
				break
			}
		}

		link.Text = text(node)

		links = append(links, link)
	}

	return links, nil
}

func text(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}

	if n.Type != html.ElementNode {
		return ""
	}

	var t string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		t += text(c) + " "
	}

	return strings.Join(strings.Fields(t), " ")
}

func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}

	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, linkNodes(c)...)
	}

	return ret
}

// func dfs(n *html.Node, padding string) {
// 	fmt.Println(padding, n.Data)

// 	for c:=n.FirstChild; c != nil; c = c.NextSibling {
// 		fmt.Println(c.Data)
// 		dfs(c,padding + "  ")
// 	}
// }