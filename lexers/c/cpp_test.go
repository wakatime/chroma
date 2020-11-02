package c_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers/c"
)

func TestCpp_AnalyseText(t *testing.T) {
	tests := map[string]struct {
		Filepath string
		Expected float32
	}{
		"include": {
			Filepath: "testdata/cpp_include.cpp",
			Expected: 0.2,
		},
		"namespace": {
			Filepath: "testdata/cpp_namespace.cpp",
			Expected: 0.4,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			data, err := ioutil.ReadFile(test.Filepath)
			assert.NoError(t, err)

			analyser, ok := c.CPP.(chroma.Analyser)
			assert.True(t, ok)

			assert.Equal(t, test.Expected, analyser.AnalyseText(string(data)))
		})
	}
}

func TestIssue290(t *testing.T) {
	input := `// 64-bit floats have 53 digits of precision, including the whole-number-part.
double a =     0011111110111001100110011001100110011001100110011001100110011010; // imperfect representation of 0.1
double b =     0011111111001001100110011001100110011001100110011001100110011010; // imperfect representation of 0.2
double c =     0011111111010011001100110011001100110011001100110011001100110011; // imperfect representation of 0.3
double a + b = 0011111111010011001100110011001100110011001100110011001100110100; // Note that this is not quite equal to the "canonical" 0.3!a
`
	it, err := c.CPP.Tokenise(nil, input)
	assert.NoError(t, err)
	for {
		token := it()
		if token == chroma.EOF {
			break
		}
	}
}
