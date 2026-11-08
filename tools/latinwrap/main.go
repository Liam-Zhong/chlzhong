package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)
var textProcessorRegex = regexp.MustCompile(`(?s)\$\$.*?\$\$|\$(?:[^\s\$][^\$]*[^\s\$]|[^\s\$])\$|[a-zA-Z0-9.\+_/:@#%&=-]+(?:\s+[a-zA-Z0-9.\+_/:@#%&=-]+)*`)

var skipTags = map[string]bool{
	"code":     true,
	"pre":      true,
	"script":   true,
	"style":    true,
	"svg":      true,
	"math":     true,
	"kbd":      true,
	"title":    true, 
	"textarea": true, 
	"option":   true, 
}

func main() {
	targetDir := "../../public"

	fmt.Println("Start Typography Post Processor...")

	err := filepath.Walk(targetDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".html") {
			processHTMLFile(path)
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Done")
	}
}

func processHTMLFile(path string) {
	content, err := os.ReadFile(path)
	if err != nil {
		return
	}

	doc, err := html.Parse(bytes.NewReader(content))
	if err != nil {
		return
	}

	traverse(doc, false)

	var buf bytes.Buffer
	html.Render(&buf, doc)
	os.WriteFile(path, buf.Bytes(), 0644)
}

func traverse(n *html.Node, skip bool) {
	if n.Type == html.ElementNode && skipTags[n.Data] {
		skip = true
	}

	if n.Type == html.TextNode && !skip {
		if textProcessorRegex.MatchString(n.Data) {
			wrapNodes(n)
		}
	}

	for c := n.FirstChild; c != nil; {
		next := c.NextSibling
		traverse(c, skip)
		c = next
	}
}

func wrapNodes(n *html.Node) {
	text := n.Data
	matches := textProcessorRegex.FindAllStringIndex(text, -1)
	if len(matches) == 0 {
		return
	}

	parent := n.Parent
	if parent == nil {
		return
	}

	lastIndex := 0

	for _, match := range matches {
		start, end := match[0], match[1]

		if start > lastIndex {
			prevText := text[lastIndex:start]
			parent.InsertBefore(&html.Node{Type: html.TextNode, Data: prevText}, n)
		}

		matchStr := text[start:end]

		if strings.HasPrefix(matchStr, "$") {
			parent.InsertBefore(&html.Node{Type: html.TextNode, Data: matchStr}, n)
		} else {
			span := &html.Node{
				Type: html.ElementNode,
				Data: "span",
				Attr: []html.Attribute{{Key: "class", Val: "latin"}},
			}
			span.AppendChild(&html.Node{Type: html.TextNode, Data: matchStr})
			parent.InsertBefore(span, n)
		}

		lastIndex = end
	}

	if lastIndex < len(text) {
		nextText := text[lastIndex:]
		parent.InsertBefore(&html.Node{Type: html.TextNode, Data: nextText}, n)
	}

	parent.RemoveChild(n)
}
