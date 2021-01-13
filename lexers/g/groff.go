package g

import (
	"regexp"
	"unicode"

	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

var groffAlphanumericRe = regexp.MustCompile(`^[a-zA-Z0-9]+$`)

// Groff lexer.
var Groff = internal.Register(MustNewLexer(
	&Config{
		Name:      "Groff",
		Aliases:   []string{"groff", "nroff", "man"},
		Filenames: []string{"*.[1234567]", "*.man"},
		MimeTypes: []string{"application/x-troff", "text/troff"},
	},
	Rules{
		"root": {},
	},
).SetAnalyser(func(text string) float32 {
	if text[:1] != "." {
		return 0
	}

	if text[:3] == `.\"` {
		return 1.0
	}

	if text[:4] == ".TH " {
		return 1.0
	}

	if groffAlphanumericRe.MatchString(text[1:3]) && unicode.IsSpace(rune(text[3])) {
		return 0.9
	}

	return 0
}))
