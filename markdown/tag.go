package markdown

import "regexp"

type Tag struct {
	Name       string
	FullRegex  *regexp.Regexp
	StartRegex *regexp.Regexp
	EndRegex   *regexp.Regexp
}

var tags = []Tag{
	{
		Name:       "Header",
		FullRegex:  regexp.MustCompile(`^#{1,6}\s+(.*)`),
		StartRegex: regexp.MustCompile(`^#{1,6}\s+`),
		EndRegex:   nil,
	},
	{
		Name:       "Bold",
		FullRegex:  regexp.MustCompile(`\*\*(.*?)\*\*|__(.*?)__`),
		StartRegex: regexp.MustCompile(`\*\*`),
		EndRegex:   regexp.MustCompile(`\*\*`),
	},
	{
		Name:       "Italic",
		FullRegex:  regexp.MustCompile(`\*(.*?)\*|_(.*?)_`),
		StartRegex: regexp.MustCompile(`\*`),
		EndRegex:   regexp.MustCompile(`\*`),
	},
	{
		Name:       "Strikethrough",
		FullRegex:  regexp.MustCompile(`~~(.*?)~~`),
		StartRegex: regexp.MustCompile(`^~~`),
		EndRegex:   regexp.MustCompile(`~~$`),
	},
	{
		Name:       "InlineCode",
		FullRegex:  regexp.MustCompile("\\`(.+?)\\`"),
		StartRegex: regexp.MustCompile("\\`"),
		EndRegex:   regexp.MustCompile("\\`"),
	},
	{
		Name:       "Image",
		FullRegex:  regexp.MustCompile(`\!\[(.*?)\]\((.*?)\)`),
		StartRegex: regexp.MustCompile(`\!\[`),
		EndRegex:   regexp.MustCompile(`\]\((.*?)\)`),
	},
	{
		Name:       "Link",
		FullRegex:  regexp.MustCompile(`\[(.*?)\]\((.*?)\)`),
		StartRegex: regexp.MustCompile(`\[`),
		EndRegex:   regexp.MustCompile(`\]\((.*?)\)`),
	},
}
