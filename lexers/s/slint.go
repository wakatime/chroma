package s

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Slint lexer. Lexer for the Slint programming language.
var Slint = internal.Register(MustNewLexer(
	&Config{
		Name:      "Slint",
		Aliases:   []string{"slint"},
		Filenames: []string{"*.slint"},
	},
	Rules{
		"root": {},
	},
))
