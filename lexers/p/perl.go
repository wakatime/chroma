package p

import (
	"regexp"
	"strings"

	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
	"github.com/alecthomas/chroma/pkg/shebang"
)

var perlAnalyzerRe = regexp.MustCompile(`(?:my|our)\s+[$@%(]`)

// Perl lexer.
var Perl = internal.Register(MustNewLazyLexer(
	&Config{
		Name:      "Perl",
		Aliases:   []string{"perl", "pl"},
		Filenames: []string{"*.pl", "*.pm", "*.t"},
		MimeTypes: []string{"text/x-perl", "application/x-perl"},
		DotAll:    true,
	},
	perlRules,
).SetAnalyser(func(text string) float32 {
	if matched, _ := shebang.MatchString(text, "perl"); matched {
		return 1.0
	}

	var result float32 = 0

	if perlAnalyzerRe.MatchString(text) {
		result += 0.9
	}

	if strings.Contains(text, ":=") {
		// := is not valid Perl, but it appears in unicon, so we should
		// become less confident if we think we found Perl with :=
		result /= 2
	}

	return result
}))

func perlRules() Rules {
	return Rules{
		"balanced-regex": {
			{`/(\\\\|\\[^\\]|[^\\/])*/[egimosx]*`, LiteralStringRegex, Pop(1)},
			{`!(\\\\|\\[^\\]|[^\\!])*![egimosx]*`, LiteralStringRegex, Pop(1)},
			{`\\(\\\\|[^\\])*\\[egimosx]*`, LiteralStringRegex, Pop(1)},
			{`\{(\\\\|\\[^\\]|[^\\}])*\}[egimosx]*`, LiteralStringRegex, Pop(1)},
			{`<(\\\\|\\[^\\]|[^\\>])*>[egimosx]*`, LiteralStringRegex, Pop(1)},
			{`\[(\\\\|\\[^\\]|[^\\\]])*\][egimosx]*`, LiteralStringRegex, Pop(1)},
			{`\((\\\\|\\[^\\]|[^\\)])*\)[egimosx]*`, LiteralStringRegex, Pop(1)},
			{`@(\\\\|\\[^\\]|[^\\@])*@[egimosx]*`, LiteralStringRegex, Pop(1)},
			{`%(\\\\|\\[^\\]|[^\\%])*%[egimosx]*`, LiteralStringRegex, Pop(1)},
			{`\$(\\\\|\\[^\\]|[^\\$])*\$[egimosx]*`, LiteralStringRegex, Pop(1)},
		},
		"root": {
			{`\A\#!.+?$`, CommentHashbang, nil},
			{`\#.*?$`, CommentSingle, nil},
			{`^=[a-zA-Z0-9]+\s+.*?\n=cut`, CommentMultiline, nil},
			{Words(``, `\b`, `case`, `continue`, `do`, `else`, `elsif`, `for`, `foreach`, `if`, `last`, `my`, `next`, `our`, `redo`, `reset`, `then`, `unless`, `until`, `while`, `print`, `new`, `BEGIN`, `CHECK`, `INIT`, `END`, `return`), Keyword, nil},
			{`(format)(\s+)(\w+)(\s*)(=)(\s*\n)`, ByGroups(Keyword, Text, Name, Text, Punctuation, Text), Push("format")},
			{`(eq|lt|gt|le|ge|ne|not|and|or|cmp)\b`, OperatorWord, nil},
			{`s/(\\\\|\\[^\\]|[^\\/])*/(\\\\|\\[^\\]|[^\\/])*/[egimosx]*`, LiteralStringRegex, nil},
			{`s!(\\\\|\\!|[^!])*!(\\\\|\\!|[^!])*![egimosx]*`, LiteralStringRegex, nil},
			{`s\\(\\\\|[^\\])*\\(\\\\|[^\\])*\\[egimosx]*`, LiteralStringRegex, nil},
			{`s@(\\\\|\\[^\\]|[^\\@])*@(\\\\|\\[^\\]|[^\\@])*@[egimosx]*`, LiteralStringRegex, nil},
			{`s%(\\\\|\\[^\\]|[^\\%])*%(\\\\|\\[^\\]|[^\\%])*%[egimosx]*`, LiteralStringRegex, nil},
			{`s\{(\\\\|\\[^\\]|[^\\}])*\}\s*`, LiteralStringRegex, Push("balanced-regex")},
			{`s<(\\\\|\\[^\\]|[^\\>])*>\s*`, LiteralStringRegex, Push("balanced-regex")},
			{`s\[(\\\\|\\[^\\]|[^\\\]])*\]\s*`, LiteralStringRegex, Push("balanced-regex")},
			{`s\((\\\\|\\[^\\]|[^\\)])*\)\s*`, LiteralStringRegex, Push("balanced-regex")},
			{`m?/(\\\\|\\[^\\]|[^\\/\n])*/[gcimosx]*`, LiteralStringRegex, nil},
			{`m(?=[/!\\{<\[(@%$])`, LiteralStringRegex, Push("balanced-regex")},
			{`((?<==~)|(?<=\())\s*/(\\\\|\\[^\\]|[^\\/])*/[gcimosx]*`, LiteralStringRegex, nil},
			{`\s+`, Text, nil},
			{Words(``, `\b`, `abs`, `accept`, `alarm`, `atan2`, `bind`, `binmode`, `bless`, `caller`, `chdir`, `chmod`, `chomp`, `chop`, `chown`, `chr`, `chroot`, `close`, `closedir`, `connect`, `continue`, `cos`, `crypt`, `dbmclose`, `dbmopen`, `defined`, `delete`, `die`, `dump`, `each`, `endgrent`, `endhostent`, `endnetent`, `endprotoent`, `endpwent`, `endservent`, `eof`, `eval`, `exec`, `exists`, `exit`, `exp`, `fcntl`, `fileno`, `flock`, `fork`, `format`, `formline`, `getc`, `getgrent`, `getgrgid`, `getgrnam`, `gethostbyaddr`, `gethostbyname`, `gethostent`, `getlogin`, `getnetbyaddr`, `getnetbyname`, `getnetent`, `getpeername`, `getpgrp`, `getppid`, `getpriority`, `getprotobyname`, `getprotobynumber`, `getprotoent`, `getpwent`, `getpwnam`, `getpwuid`, `getservbyname`, `getservbyport`, `getservent`, `getsockname`, `getsockopt`, `glob`, `gmtime`, `goto`, `grep`, `hex`, `import`, `index`, `int`, `ioctl`, `join`, `keys`, `kill`, `last`, `lc`, `lcfirst`, `length`, `link`, `listen`, `local`, `localtime`, `log`, `lstat`, `map`, `mkdir`, `msgctl`, `msgget`, `msgrcv`, `msgsnd`, `my`, `next`, `oct`, `open`, `opendir`, `ord`, `our`, `pack`, `pipe`, `pop`, `pos`, `printf`, `prototype`, `push`, `quotemeta`, `rand`, `read`, `readdir`, `readline`, `readlink`, `readpipe`, `recv`, `redo`, `ref`, `rename`, `reverse`, `rewinddir`, `rindex`, `rmdir`, `scalar`, `seek`, `seekdir`, `select`, `semctl`, `semget`, `semop`, `send`, `setgrent`, `sethostent`, `setnetent`, `setpgrp`, `setpriority`, `setprotoent`, `setpwent`, `setservent`, `setsockopt`, `shift`, `shmctl`, `shmget`, `shmread`, `shmwrite`, `shutdown`, `sin`, `sleep`, `socket`, `socketpair`, `sort`, `splice`, `split`, `sprintf`, `sqrt`, `srand`, `stat`, `study`, `substr`, `symlink`, `syscall`, `sysopen`, `sysread`, `sysseek`, `system`, `syswrite`, `tell`, `telldir`, `tie`, `tied`, `time`, `times`, `tr`, `truncate`, `uc`, `ucfirst`, `umask`, `undef`, `unlink`, `unpack`, `unshift`, `untie`, `utime`, `values`, `vec`, `wait`, `waitpid`, `wantarray`, `warn`, `write`), NameBuiltin, nil},
			{`((__(DATA|DIE|WARN)__)|(STD(IN|OUT|ERR)))\b`, NameBuiltinPseudo, nil},
			{`(<<)([\'"]?)([a-zA-Z_]\w*)(\2;?\n.*?\n)(\3)(\n)`, ByGroups(LiteralString, LiteralString, LiteralStringDelimiter, LiteralString, LiteralStringDelimiter, Text), nil},
			{`__END__`, CommentPreproc, Push("end-part")},
			{`\$\^[ADEFHILMOPSTWX]`, NameVariableGlobal, nil},
			{"\\$[\\\\\\\"\\[\\]'&`+*.,;=%~?@$!<>(^|/-](?!\\w)", NameVariableGlobal, nil},
			{`[$@%#]+`, NameVariable, Push("varname")},
			{`0_?[0-7]+(_[0-7]+)*`, LiteralNumberOct, nil},
			{`0x[0-9A-Fa-f]+(_[0-9A-Fa-f]+)*`, LiteralNumberHex, nil},
			{`0b[01]+(_[01]+)*`, LiteralNumberBin, nil},
			{`(?i)(\d*(_\d*)*\.\d+(_\d*)*|\d+(_\d*)*\.\d+(_\d*)*)(e[+-]?\d+)?`, LiteralNumberFloat, nil},
			{`(?i)\d+(_\d*)*e[+-]?\d+(_\d*)*`, LiteralNumberFloat, nil},
			{`\d+(_\d+)*`, LiteralNumberInteger, nil},
			{`'(\\\\|\\[^\\]|[^'\\])*'`, LiteralString, nil},
			{`"(\\\\|\\[^\\]|[^"\\])*"`, LiteralString, nil},
			{"`(\\\\\\\\|\\\\[^\\\\]|[^`\\\\])*`", LiteralStringBacktick, nil},
			{`<([^\s>]+)>`, LiteralStringRegex, nil},
			{`(q|qq|qw|qr|qx)\{`, LiteralStringOther, Push("cb-string")},
			{`(q|qq|qw|qr|qx)\(`, LiteralStringOther, Push("rb-string")},
			{`(q|qq|qw|qr|qx)\[`, LiteralStringOther, Push("sb-string")},
			{`(q|qq|qw|qr|qx)\<`, LiteralStringOther, Push("lt-string")},
			{`(q|qq|qw|qr|qx)([\W_])(.|\n)*?\2`, LiteralStringOther, nil},
			{`(package)(\s+)([a-zA-Z_]\w*(?:::[a-zA-Z_]\w*)*)`, ByGroups(Keyword, Text, NameNamespace), nil},
			{`(use|require|no)(\s+)([a-zA-Z_]\w*(?:::[a-zA-Z_]\w*)*)`, ByGroups(Keyword, Text, NameNamespace), nil},
			{`(sub)(\s+)`, ByGroups(Keyword, Text), Push("funcname")},
			{Words(``, `\b`, `no`, `package`, `require`, `use`), Keyword, nil},
			{`(\[\]|\*\*|::|<<|>>|>=|<=>|<=|={3}|!=|=~|!~|&&?|\|\||\.{1,3})`, Operator, nil},
			{`[-+/*%=<>&^|!\\~]=?`, Operator, nil},
			{`[()\[\]:;,<>/?{}]`, Punctuation, nil},
			{`(?=\w)`, Name, Push("name")},
		},
		"format": {
			{`\.\n`, LiteralStringInterpol, Pop(1)},
			{`[^\n]*\n`, LiteralStringInterpol, nil},
		},
		"varname": {
			{`\s+`, Text, nil},
			{`\{`, Punctuation, Pop(1)},
			{`\)|,`, Punctuation, Pop(1)},
			{`\w+::`, NameNamespace, nil},
			{`[\w:]+`, NameVariable, Pop(1)},
		},
		"name": {
			{`[a-zA-Z_]\w*(::[a-zA-Z_]\w*)*(::)?(?=\s*->)`, NameNamespace, Pop(1)},
			{`[a-zA-Z_]\w*(::[a-zA-Z_]\w*)*::`, NameNamespace, Pop(1)},
			{`[\w:]+`, Name, Pop(1)},
			{`[A-Z_]+(?=\W)`, NameConstant, Pop(1)},
			{`(?=\W)`, Text, Pop(1)},
		},
		"funcname": {
			{`[a-zA-Z_]\w*[!?]?`, NameFunction, nil},
			{`\s+`, Text, nil},
			{`(\([$@%]*\))(\s*)`, ByGroups(Punctuation, Text), nil},
			{`;`, Punctuation, Pop(1)},
			{`.*?\{`, Punctuation, Pop(1)},
		},
		"cb-string": {
			{`\\[{}\\]`, LiteralStringOther, nil},
			{`\\`, LiteralStringOther, nil},
			{`\{`, LiteralStringOther, Push("cb-string")},
			{`\}`, LiteralStringOther, Pop(1)},
			{`[^{}\\]+`, LiteralStringOther, nil},
		},
		"rb-string": {
			{`\\[()\\]`, LiteralStringOther, nil},
			{`\\`, LiteralStringOther, nil},
			{`\(`, LiteralStringOther, Push("rb-string")},
			{`\)`, LiteralStringOther, Pop(1)},
			{`[^()]+`, LiteralStringOther, nil},
		},
		"sb-string": {
			{`\\[\[\]\\]`, LiteralStringOther, nil},
			{`\\`, LiteralStringOther, nil},
			{`\[`, LiteralStringOther, Push("sb-string")},
			{`\]`, LiteralStringOther, Pop(1)},
			{`[^\[\]]+`, LiteralStringOther, nil},
		},
		"lt-string": {
			{`\\[<>\\]`, LiteralStringOther, nil},
			{`\\`, LiteralStringOther, nil},
			{`\<`, LiteralStringOther, Push("lt-string")},
			{`\>`, LiteralStringOther, Pop(1)},
			{`[^<>]+`, LiteralStringOther, nil},
		},
		"end-part": {
			{`.+`, CommentPreproc, Pop(1)},
		},
	}
}
