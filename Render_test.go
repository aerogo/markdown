package markdown_test

import (
	"strings"
	"testing"

	"github.com/aerogo/markdown"
)

func TestRender(t *testing.T) {
	example := "# Header\n\nThis is markdown."
	expect := "<h1>Header</h1><p>This is markdown.</p>"
	html := markdown.Render(example)
	html = strings.ReplaceAll(html, "\n", "")

	if html != expect {
		t.Logf("Expected output differs:\n1.)\n%s\n2.)\n%s", html, expect)
		t.FailNow()
	}
}

func TestNoClosingTag(t *testing.T) {
	example := "Hello<div>"
	expect := "Unclosed HTML tag: div"
	html := markdown.Render(example)
	html = strings.ReplaceAll(html, "\n", "")

	if html != expect {
		t.Logf("Expected output differs:\n1.)\n%s\n2.)\n%s", html, expect)
		t.FailNow()
	}
}

func TestClosingWithoutOpeningTag(t *testing.T) {
	example := "Hello</p>"
	expect := "Closing tag without opening tag: p"
	html := markdown.Render(example)
	html = strings.ReplaceAll(html, "\n", "")

	if html != expect {
		t.Logf("Expected output differs:\n1.)\n%s\n2.)\n%s", html, expect)
		t.FailNow()
	}
}
