package markdown

import (
	"github.com/aerogo/aero"
	"github.com/microcosm-cc/bluemonday"
	blackfriday "gopkg.in/russross/blackfriday.v2"
)

var policy = bluemonday.UGCPolicy()

// Render converts the given markdown code to an HTML string.
func Render(code string) string {
	codeBytes := aero.StringToBytesUnsafe(code)
	codeBytes = policy.SanitizeBytes(codeBytes)
	codeBytes = blackfriday.Run(codeBytes)
	return aero.BytesToStringUnsafe(codeBytes)
}
