package m

import (
	"regexp"
	"strings"

	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

var (
	matlabAnalyserCommentRe   = regexp.MustCompile(`^\s*%`)
	matlabAnalyserSystemCMDRe = regexp.MustCompile(`^!\w+`)
)

// Matlab lexer.
var Matlab = internal.Register(MustNewLazyLexer(
	&Config{
		Name:      "Matlab",
		Aliases:   []string{"matlab"},
		Filenames: []string{"*.m", "*.matlab"},
		MimeTypes: []string{"text/matlab"},
	},
	matlabRules,
).SetAnalyser(func(text string) float32 {
	lines := strings.Split(strings.Replace(text, "\r\n", "\n", -1), "\n")

	var firstNonComment string
	for _, line := range lines {
		if !matlabAnalyserCommentRe.MatchString(line) {
			firstNonComment = strings.TrimSpace(line)
			break
		}
	}

	// function declaration
	if strings.HasPrefix(firstNonComment, "function") && !strings.Contains(firstNonComment, "{") {
		return 1.0
	}

	// comment
	for _, line := range lines {
		if matlabAnalyserCommentRe.MatchString(line) {
			return 0.2
		}
	}

	// system cmd
	for _, line := range lines {
		if matlabAnalyserSystemCMDRe.MatchString(line) {
			return 0.2
		}
	}

	return 0
}))

func matlabRules() Rules {
	return Rules{
		"root": {
			{`\n`, Text, nil},
			{`^!.*`, LiteralStringOther, nil},
			{`%\{\s*\n`, CommentMultiline, Push("blockcomment")},
			{`%.*$`, Comment, nil},
			{`^\s*function`, Keyword, Push("deffunc")},
			{Words(``, `\b`, `break`, `case`, `catch`, `classdef`, `continue`, `else`, `elseif`, `end`, `enumerated`, `events`, `for`, `function`, `global`, `if`, `methods`, `otherwise`, `parfor`, `persistent`, `properties`, `return`, `spmd`, `switch`, `try`, `while`), Keyword, nil},
			{`(sin|sind|sinh|asin|asind|asinh|cos|cosd|cosh|acos|acosd|acosh|tan|tand|tanh|atan|atand|atan2|atanh|sec|secd|sech|asec|asecd|asech|csc|cscd|csch|acsc|acscd|acsch|cot|cotd|coth|acot|acotd|acoth|hypot|exp|expm1|log|log1p|log10|log2|pow2|realpow|reallog|realsqrt|sqrt|nthroot|nextpow2|abs|angle|complex|conj|imag|real|unwrap|isreal|cplxpair|fix|floor|ceil|round|mod|rem|sign|airy|besselj|bessely|besselh|besseli|besselk|beta|betainc|betaln|ellipj|ellipke|erf|erfc|erfcx|erfinv|expint|gamma|gammainc|gammaln|psi|legendre|cross|dot|factor|isprime|primes|gcd|lcm|rat|rats|perms|nchoosek|factorial|cart2sph|cart2pol|pol2cart|sph2cart|hsv2rgb|rgb2hsv|zeros|ones|eye|repmat|rand|randn|linspace|logspace|freqspace|meshgrid|accumarray|size|length|ndims|numel|disp|isempty|isequal|isequalwithequalnans|cat|reshape|diag|blkdiag|tril|triu|fliplr|flipud|flipdim|rot90|find|end|sub2ind|ind2sub|bsxfun|ndgrid|permute|ipermute|shiftdim|circshift|squeeze|isscalar|isvector|ans|eps|realmax|realmin|pi|i|inf|nan|isnan|isinf|isfinite|j|why|compan|gallery|hadamard|hankel|hilb|invhilb|magic|pascal|rosser|toeplitz|vander|wilkinson)\b`, NameBuiltin, nil},
			{`\.\.\..*$`, Comment, nil},
			{`-|==|~=|<|>|<=|>=|&&|&|~|\|\|?`, Operator, nil},
			{`\.\*|\*|\+|\.\^|\.\\|\.\/|\/|\\`, Operator, nil},
			{`\[|\]|\(|\)|\{|\}|:|@|\.|,`, Punctuation, nil},
			{`=|:|;`, Punctuation, nil},
			{`(?<=[\w)\].])\'+`, Operator, nil},
			{`(\d+\.\d*|\d*\.\d+)([eEf][+-]?[0-9]+)?`, LiteralNumberFloat, nil},
			{`\d+[eEf][+-]?[0-9]+`, LiteralNumberFloat, nil},
			{`\d+`, LiteralNumberInteger, nil},
			{`(?<![\w)\].])\'`, LiteralString, Push("string")},
			{`[a-zA-Z_]\w*`, Name, nil},
			{`.`, Text, nil},
		},
		"string": {
			{`[^\']*\'`, LiteralString, Pop(1)},
		},
		"blockcomment": {
			{`^\s*%\}`, CommentMultiline, Pop(1)},
			{`^.*\n`, CommentMultiline, nil},
			{`.`, CommentMultiline, nil},
		},
		"deffunc": {
			{`(\s*)(?:(.+)(\s*)(=)(\s*))?(.+)(\()(.*)(\))(\s*)`, ByGroups(TextWhitespace, Text, TextWhitespace, Punctuation, TextWhitespace, NameFunction, Punctuation, Text, Punctuation, TextWhitespace), Pop(1)},
			{`(\s*)([a-zA-Z_]\w*)`, ByGroups(Text, NameFunction), Pop(1)},
		},
	}
}
