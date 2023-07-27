package html

import (
	"golang.org/x/net/html"
	"strings"
)

// HTMLExtractor represents an HTML-specific plain text extractor.
type Extractor struct {
	blockTags map[string]bool
}

// NewExtractor creates a new HTMLExtractor instance.
func NewExtractor(otherBlockTags ...string) *Extractor {
	uniqueBlockTags := map[string]bool{}
	for _, tag := range blockTags {
		uniqueBlockTags[tag] = true
	}
	for _, tag := range otherBlockTags {
		uniqueBlockTags[tag] = true
	}

	return &Extractor{blockTags: uniqueBlockTags}
}

// PlainText extracts plain text from the input HTML string.
func (e *Extractor) PlainText(input string) (*string, error) {
	doc, err := html.Parse(strings.NewReader(input))
	if err != nil {
		return nil, err
	}

	var plainText strings.Builder
	e.extractText(&plainText, doc)

	output := plainText.String()
	output = strings.ReplaceAll(output, "\n ", "\n")
	return &output, nil
}

// Recursively extract plain text from the HTML nodes.
func (e *Extractor) extractText(plainText *strings.Builder, node *html.Node) {
	if node.Type == html.TextNode {
		// Trim and append the text content
		text := strings.TrimSpace(node.Data)
		if text != "" {
			if plainText.Len() > 0 {
				if found := e.blockTags[node.Parent.DataAtom.String()]; found {
					plainText.WriteString("\n")
				} else {
					plainText.WriteString(" ")
				}
			}
			plainText.WriteString(text)
		}
	}
	if node.DataAtom.String() == "br" {
		plainText.WriteString("\n")
		return
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		e.extractText(plainText, child)
	}
}
