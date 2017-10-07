package markdown

import (
	"github.com/aerogo/aero"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

var strictPolicy = bluemonday.StrictPolicy()

// Render converts the given markdown code to an HTML string.
func Render(code string) string {
	codeBytes := aero.StringToBytesUnsafe(code)
	codeBytes = strictPolicy.SanitizeBytes(codeBytes)
	codeBytes = blackfriday.MarkdownCommon(codeBytes)
	return aero.BytesToStringUnsafe(codeBytes)
}
