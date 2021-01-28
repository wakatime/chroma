package p

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// Python2 lexer.
var Python2 = internal.Register(MustNewLexer(
	&Config{
		Name:      "Python 2.x",
		Aliases:   []string{"python2", "py2"},
		MimeTypes: []string{"text/x-python2", "application/x-python2"},
	},
	Rules{
		"root": {},
	},
))
