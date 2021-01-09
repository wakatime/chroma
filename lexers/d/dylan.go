package d

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Dylan lexer.
var Dylan = internal.Register(MustNewLexer(
	&Config{
		Name:      "Dylan",
		Aliases:   []string{"dylan"},
		Filenames: []string{"*.dylan", "*.dyl", "*.intr"},
		MimeTypes: []string{"text/x-dylan"},
	},
	Rules{
		"root": {},
	},
))
