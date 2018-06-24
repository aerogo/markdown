package markdown

import (
	"github.com/aerogo/aero"
	blackfriday "gopkg.in/russross/blackfriday.v2"
)

// Render converts the given markdown code to an HTML string.
func Render(code string) string {
	codeBytes := aero.StringToBytesUnsafe(code)
	codeBytes = blackfriday.Run(codeBytes)
	return aero.BytesToStringUnsafe(codeBytes)
}
