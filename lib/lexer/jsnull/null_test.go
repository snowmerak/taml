package jsnull_test

import (
	"bytes"
	"testing"

	"github.com/snowmerak/tson/lib/lexer/jsnull"
)

func TestNull(t *testing.T) {
	inputs := [][]byte{[]byte("  null !!!!"), []byte("\nnull"), []byte("null")}
	outputs := [][]byte{[]byte("null"), []byte("null"), []byte("null")}
	for i, input := range inputs {
		_, v, _, err := jsnull.Find(input)
		if err != nil {
			t.Errorf("%d: %s", i, err)
		}
		if !bytes.Equal(v, outputs[i]) {
			t.Errorf("%d: %s != %s", i, v, outputs[i])
		}
	}
}
