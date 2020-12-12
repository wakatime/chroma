package s

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// S lexer.
var S = internal.Register(MustNewLexer(
	&Config{
		Name:      "S",
		Aliases:   []string{"splus", "s", "r"},
		Filenames: []string{"*.S", "*.R", ".Rhistory", ".Rprofile", ".Renviron"},
		MimeTypes: []string{"text/S-plus", "text/S", "text/x-r-source", "text/x-r",
			"text/x-R", "text/x-r-history", "text/x-r-profile"},
	},
	Rules{
		"root": {},
	},
))
