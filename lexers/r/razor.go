package r

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Razor lexer. Lexer for Blazor's Razor files.
var Razor = internal.Register(MustNewLexer(
	&Config{
		Name:      "Razor",
		Aliases:   []string{"razor"},
		Filenames: []string{"*.razor"},
		MimeTypes: []string{"text/html"},
	},
	Rules{
		"root": {},
	},
))
