package c

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// CBM BASIC V2 lexer.
var CbmBasicV2 = internal.Register(MustNewLexer(
	&Config{
		Name:      "CBM BASIC V2",
		Aliases:   []string{"cbmbas"},
		Filenames: []string{"*.bas"},
		MimeTypes: []string{},
	},
	Rules{
		"root": {},
	},
))
