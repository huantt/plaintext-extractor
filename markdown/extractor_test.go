package markdown

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestExtract(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"~~strikethrough~~", "strikethrough"},
		{"# H1 \n*italic* **bold** `code` `not code [link](https://example.com) ![image](https://image.com/image.png) ~~strikethrough~~", "H1 \nitalic bold code `not code link image strikethrough"},
		{"# H1 \n new line", "H1 \n new line"},
		{"*italic*", "italic"},
		{"**bold**", "bold"},
		{"`code`", "code"},
		{"`node code", "`node code"},
		{"[link](https://example.com)", "link"},
		{"[image](https://image.com/image.png)", "image"},
	}

	markdownExtractor := NewExtractor()
	for _, test := range tests {
		output, err := markdownExtractor.PlainText(test.input)
		assert.NoError(t, err)
		assert.NotNil(t, output)
		assert.Equal(t, test.expected, *output)
	}
	fmt.Println()
}

func TestCustomTag(t *testing.T) {
	customTag := Tag{
		Name:       "color-custom-tag",
		FullRegex:  regexp.MustCompile("{color:[a-zA-Z0-9#]+}(.*?){color}"),
		StartRegex: regexp.MustCompile("{color:[a-zA-Z0-9#]+}"),
		EndRegex:   regexp.MustCompile("{color}"),
	}
	markdownExtractor := NewExtractor(customTag)
	tests := []struct {
		input    string
		expected string
	}{
		{"This is {color:#0A84FF}red{color}", "This is red"},
	}
	for _, test := range tests {
		output, err := markdownExtractor.PlainText(test.input)
		assert.NoError(t, err)
		assert.NotNil(t, output)
		assert.Equal(t, test.expected, *output)
	}
}
