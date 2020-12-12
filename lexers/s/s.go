package s

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
	"github.com/dlclark/regexp2"
)

var sAnalyserRe = regexp2.MustCompile(`[a-z0-9_\])\s]<-(?!-)`, regexp2.None)

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
).SetAnalyser(func(text string) float32 {
	if matched, _ := sAnalyserRe.MatchString(text); matched {
		return 0.11
	}

	return 0
}))
