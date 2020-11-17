package b

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// BUGS lexer.
var Bugs = internal.Register(MustNewLexer(
	&Config{
		Name:      "BUGS",
		Aliases:   []string{"bugs", "winbugs", "openbugs"},
		Filenames: []string{"*.bug"},
		MimeTypes: []string{},
	},
	Rules{
		"root": {},
	},
))
