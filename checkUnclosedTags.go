package markdown

import (
	"bytes"
	"fmt"
)

var selfClosingTags = map[string]bool{
	"area":    true,
	"base":    true,
	"br":      true,
	"col":     true,
	"command": true,
	"embed":   true,
	"hr":      true,
	"img":     true,
	"input":   true,
	"link":    true,
	"meta":    true,
	"param":   true,
	"source":  true,
}

func checkUnclosedTags(codeBytes []byte) error {
	startTag := false
	isClosingTag := false
	var tagName []byte
	stack := [][]byte{}

	for _, b := range codeBytes {
		if b == '<' {
			startTag = true
			continue
		}

		if startTag && b == '/' {
			isClosingTag = true
			continue
		}

		if b == ' ' || b == '>' {
			if tagName != nil {
				if isClosingTag {
					lastTagName := stack[len(stack)-1]

					if !bytes.Equal(lastTagName, tagName) {
						return fmt.Errorf("Unclosed HTML tag: %s", lastTagName)
					}

					stack = stack[:len(stack)-1]
				} else if !selfClosingTags[string(tagName)] {
					stack = append(stack, tagName)
				}
			}

			startTag = false
			isClosingTag = false
			tagName = nil
			continue
		}

		if startTag {
			tagName = append(tagName, b)
		}
	}

	return nil
}
