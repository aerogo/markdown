package markdown

import (
	"github.com/akyoto/stringutils/unsafe"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

var policy = bluemonday.UGCPolicy()

// Render converts the given markdown code to an HTML string.
func Render(code string) string {
	codeBytes := unsafe.StringToBytes(code)
	codeBytes = blackfriday.Run(codeBytes)
	codeBytes = policy.SanitizeBytes(codeBytes)
	err := checkUnclosedTags(codeBytes)

	if err != nil {
		return err.Error()
	}

	return unsafe.BytesToString(codeBytes)
}
