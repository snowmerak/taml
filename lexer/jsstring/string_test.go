package jsstring_test

import (
	"bytes"
	"testing"

	"github.com/snowmerak/tson/lexer/jsstring"
)

func TestStrings(t *testing.T) {
	inputs := [][]byte{[]byte("  \"hello, world!\" !!!!"), []byte("\n\"이거 하나 못하나\""), []byte("\"sdfsdfsdfsdfsfddsf\r\rdfgdfgfd\n\nsdfsdfsd\"")}
	outputs := [][]byte{[]byte("\"hello, world!\""), []byte("\"이거 하나 못하나\""), []byte("\"sdfsdfsdfsdfsfddsf\r\rdfgdfgfd\n\nsdfsdfsd\"")}
	for i, input := range inputs {
		_, v, _, err := jsstring.Find(input)
		if err != nil {
			t.Errorf("%d: %s", i, err)
		}
		if !bytes.Equal(v, outputs[i]) {
			t.Errorf("%d: %s != %s", i, v, outputs[i])
		}
	}
}
