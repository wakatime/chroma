package q

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

// QML lexer.
var QML = internal.Register(MustNewLexer(
	&Config{
		Name:      "QML",
		Aliases:   []string{"qml", "qbs"},
		Filenames: []string{"*.qml", "*.qbs"},
		MimeTypes: []string{"application/x-qml", "application/x-qt.qbs+qml"},
	},
	Rules{
		"root": {},
	},
))
