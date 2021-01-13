package g

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

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
))
