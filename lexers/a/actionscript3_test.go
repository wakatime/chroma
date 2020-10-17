package a_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers/a"
)

func TestActionscript3_AnalyseText(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/actionscript3.as")
	assert.NoError(t, err)

	analyser, ok := a.Actionscript3.(chroma.Analyser)
	assert.True(t, ok)

	assert.Equal(t, 0.3, analyser.AnalyseText(string(data)))
}
