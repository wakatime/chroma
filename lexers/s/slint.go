package s

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Slash lexer. Lexer for the Slash programming language.
var Slash = internal.Register(MustNewLexer(
	&Config{
		Name:      "Slint",
		Aliases:   []string{"slint"},
		Filenames: []string{"*.slint"},
	},
	Rules{
		"root": {},
	},
))
