# Plain Text Extractor
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/huantt/plaintext-extractor)](https://goreportcard.com/report/github.com/huantt/plaintext-extractor)

Plain Text Extractor is a Golang library that helps you extract plain text from `HTML` and `Markdown`. 

It provides a flexible and extensible interface for extracting the plain text content using both the predefined extraction methods and your own custom extraction requirements.

## Features
- Parse HTML and Markdown documents into plain text.
- Support for custom extraction functions.
- Easy-to-use API to convert complex documents to simple plain text.

## Installation
```shell
go get github.com/huantt/plaintext-extractror
```

## Usage
### Markdown extractor
```go
markdownContent := "# H1 \n*italic* **bold** `code` `not code [link](https://example.com) ![image](https://image.com/image.png) ~~strikethrough~~"
extractor := NewMarkdownExtractor()
output, err := extractor.PlainText(markdownContent)
if err != nil {
    panic(err)
}
fmt.Println(output)
// Output: H1 \nitalic bold code `not code link image strikethrough
```

### Custom Markdown Tag
```go
markdownContent := "This is {color:#0A84FF}red{color}"

customTag := markdown.Tag{
    Name:       "color-custom-tag",
    FullRegex:  regexp.MustCompile("{color:[a-zA-Z0-9#]+}(.*?){color}"),
    StartRegex: regexp.MustCompile("{color:[a-zA-Z0-9#]+}"),
    EndRegex:   regexp.MustCompile("{color}"),
}

markdownExtractor := NewMarkdownExtractor(customTag)
plaintextExtractor := plaintext.NewExtractor(markdownExtractor.PlainText)
plaintext, err := plaintextExtractor.PlainText(markdownContent)
if err != nil{
    panic(nil)
}
fmt.Println(plaintext)
// Output: This is red
```

### HTML Extractor
```go
html := `<div>This is a <a href="https://example.com">link</a></div>`
extractor := NewHtmlExtractor()
output, err := extractor.PlainText(html)
if err != nil {
    panic(err)
}
fmt.Println(output)
// Output: This is a link
```

### Multiple extractors
```go
input := `<div> html </div> *markdown*`
markdownExtractor := markdown.NewExtractor()
htmlExtractor := html.NewExtractor()
extractor := NewExtractor(markdownExtractor.PlainText, htmlExtractor.PlainText)
output, err := extractor.PlainText(input)
if err != nil {
    panic(err)
}
fmt.Println(output)
// Output: html markdown
```

## Contribution
Contributions to the Plain Text Parser project are welcome!
If you find any issues or want to add new features, please feel free to open an issue or submit a pull request. 
Please see the [CONTRIBUTING.md](./CONTRIBUTING.md) for more information.

## License
This project released under the MIT License, refer [LICENSE](./LICENSE.md) file.