package s

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Singularity lexer.
var Singularity = internal.Register(MustNewLexer(
	&Config{
		Name:      "Singularity",
		Aliases:   []string{"singularity"},
		Filenames: []string{"*.def", "Singularity"},
		MimeTypes: []string{},
	},
	Rules{
		"root": {},
	},
))
