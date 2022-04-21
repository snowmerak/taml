package jsbool_test

import (
	"bytes"
	"testing"

	"github.com/snowmerak/tson/lib/lexer/jsbool"
)

func TestBool(t *testing.T) {
	inputs := [][]byte{[]byte("  true !!!!"), []byte("\ntrue"), []byte("true"), []byte("false"), []byte("false !!!!"), []byte("\nfalse"), []byte("false")}
	outputs := [][]byte{[]byte("true"), []byte("true"), []byte("true"), []byte("false"), []byte("false"), []byte("false"), []byte("false")}
	for i, input := range inputs {
		_, v, _, err := jsbool.Find(input)
		if err != nil {
			t.Errorf("%d: %s", i, err)
		}
		if !bytes.Equal(v, outputs[i]) {
			t.Errorf("%d: %s != %s", i, v, outputs[i])
		}
	}
}
