package plaintext

import (
	"github.com/huantt/plaintext-extractor/html"
	"github.com/huantt/plaintext-extractor/markdown"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseHtml(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{`<div>This is a <a href="https://example.com">link</a></div>`, "This is a link"},
	}
	for _, test := range tests {
		extractor := NewHtmlExtractor()
		output, err := extractor.PlainText(test.input)
		assert.NoError(t, err)
		assert.NotNil(t, output)
		assert.Equal(t, test.expected, *output)
	}
}

func TestParseMarkdown(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"# H1 \n*italic* **bold** `code` `not code [link](https://example.com) ![image](https://image.com/image.png) ~~strikethrough~~", "H1 \nitalic bold code `not code link image strikethrough"},
	}
	for _, test := range tests {
		extractor := NewMarkdownExtractor()
		output, err := extractor.PlainText(test.input)
		assert.NoError(t, err)
		assert.NotNil(t, output)
		assert.Equal(t, test.expected, *output)
	}
}

func TestMultipleExtractors(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"<div> html </div> *markdown*", "html markdown"},
		{"<div> *markdown in html* </div>", "markdown in html"},
	}
	for _, test := range tests {
		markdownExtractor := markdown.NewExtractor()
		htmlExtractor := html.NewExtractor()
		extractor := NewExtractor(markdownExtractor.PlainText, htmlExtractor.PlainText)
		output, err := extractor.PlainText(test.input)
		assert.NoError(t, err)
		assert.NotNil(t, output)
		assert.Equal(t, test.expected, *output)
	}
}
