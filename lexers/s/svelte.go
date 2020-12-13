package s

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Svelte lexer.
var Svelte = internal.Register(MustNewLexer(
	&Config{
		Name:      "Svelte",
		Aliases:   []string{"svelte"},
		Filenames: []string{"*.svelte"},
	},
	Rules{
		"root": {},
	},
))
