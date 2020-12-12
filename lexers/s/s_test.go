package s_test

import (
	"io/ioutil"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers/s"
)

func TestS_AnalyseText(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/s.S")
	assert.NoError(t, err)

	analyser, ok := s.S.(chroma.Analyser)
	assert.True(t, ok)

	assert.Equal(t, float32(0.11), analyser.AnalyseText(string(data)))
}
