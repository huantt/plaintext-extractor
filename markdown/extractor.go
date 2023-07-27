package markdown

import (
	"bytes"
)

// Extractor represents a Markdown specific plain text extractor with custom tags support.
type Extractor struct {
	tags []Tag
}

// NewExtractor creates a new Markdown Extractor instance with optional custom tags.
func NewExtractor(customTags ...Tag) *Extractor {
	extractor := &Extractor{
		tags: tags,
	}
	extractor.tags = append(extractor.tags, customTags...)
	return extractor
}

// AddTag adds a custom tag to the Markdown Extractor instance.
func (e *Extractor) AddTag(tag Tag) *Extractor {
	e.tags = append(e.tags, tag)
	return e
}

// PlainText extracts plain text from the input Markdown string by removing the specified tags.
// It replaces the full tags and their contents with the plain text content inside those tags.
func (e *Extractor) PlainText(input string) (*string, error) {
	tmp := []byte(input)
	for _, tag := range e.tags {
		listFullTag := tag.FullRegex.FindAll(tmp, -1)
		for _, fullTag := range listFullTag {
			var plain = fullTag
			if tag.StartRegex != nil {
				plain = tag.StartRegex.ReplaceAll(plain, nil)
			}
			if tag.EndRegex != nil {
				plain = tag.EndRegex.ReplaceAll(plain, nil)
			}
			tmp = bytes.Replace(tmp, fullTag, plain, -1)
		}
	}
	output := string(tmp)
	return &output, nil
}
