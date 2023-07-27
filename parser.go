package plaintext

import (
	"github.com/huantt/plaintext-extractor/html"
	"github.com/huantt/plaintext-extractor/markdown"
)

// Extractor represents a plain text extractor that can parse input strings using multiple extract functions (for example html or markdown).
type Extractor struct {
	extractFuncs []ExtractFunc
}

// NewExtractor creates a new Extractor instance with the given extract function.
func NewExtractor(extractFunc ExtractFunc, moreFuncs ...ExtractFunc) *Extractor {
	extractor := &Extractor{
		extractFuncs: []ExtractFunc{extractFunc},
	}
	extractor.extractFuncs = append(extractor.extractFuncs, moreFuncs...)
	return extractor
}

func NewMarkdownExtractor(customTags ...markdown.Tag) *Extractor {
	return NewExtractor(markdown.NewExtractor(customTags...).PlainText)
}

func NewHtmlExtractor(blockTags ...string) *Extractor {
	return NewExtractor(html.NewExtractor(blockTags...).PlainText)
}

// AddExtractor adds an extract function to the Extractor instance.
func (p *Extractor) AddExtractor(extractor ExtractFunc) *Extractor {
	p.extractFuncs = append(p.extractFuncs, extractor)
	return p
}

// ExtractFunc is the function signature for extracting plain text from a given input string.
// Implement this function to extend availability of extracting plain text by passing into Extractor.AddExtractor function.
type ExtractFunc func(input string) (*string, error)

// PlainText extracts plain text from the input string using registered extract functions.
// It iterates over all extract functions, applying them in sequence, and returns the final plain text.
func (p *Extractor) PlainText(input string) (plainText *string, err error) {
	plainText = &input
	for _, extractFunc := range p.extractFuncs {
		plainText, err = extractFunc(*plainText)
		if err != nil {
			return nil, err
		}
	}
	return plainText, nil
}
