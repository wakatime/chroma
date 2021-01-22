package n_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers/n"
)

func TestNotmuch_AnalyseText(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/notmuch")
	assert.NoError(t, err)

	analyser, ok := n.Notmuch.(chroma.Analyser)
	assert.True(t, ok)

	assert.Equal(t, float32(1.0), analyser.AnalyseText(string(data)))
}
