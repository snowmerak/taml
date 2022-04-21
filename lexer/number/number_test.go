package number_test

import (
	"bytes"
	"testing"

	"github.com/snowmerak/tson/lexer/number"
)

func TestInteger(t *testing.T) {
	inputs := [][]byte{[]byte("a100 "), []byte("	-100()"), []byte("0"), []byte("-0"), []byte("99"), []byte("-99"), []byte("100."), []byte("34.123e")}
	outputs := [][]byte{[]byte("100"), []byte("-100"), []byte("0"), []byte("-0"), []byte("99"), []byte("-99")}
	for i, input := range inputs {
		_, v, _, err := number.Find(input)
		if i < len(outputs) && err != nil {
			t.Error(err)
		}
		if i < len(outputs) && bytes.Equal(v, outputs[i]) == false {
			t.Errorf("%s != %s", v, outputs[i])
		}
		if i >= len(outputs) && err == nil {
			t.Errorf("%s should be error", v)
		}
	}
}

func TestFloat(t *testing.T) {
	inputs := [][]byte{[]byte("	 123.2333 3434"), []byte("%&^%&sdf932.0   fdsgdfg"), []byte("dsfsd-2.2"), []byte("-0.9999asad"), []byte("0.003")}
	outputs := [][]byte{[]byte("123.2333"), []byte("932.0"), []byte("-2.2"), []byte("-0.9999"), []byte("0.003")}
	for i, input := range inputs {
		_, v, _, err := number.Find(input)
		if i < len(outputs) && err != nil {
			t.Error(err)
		}
		if i < len(outputs) && bytes.Equal(v, outputs[i]) == false {
			t.Errorf("%s != %s", v, outputs[i])
		}
		if i >= len(outputs) && err == nil {
			t.Errorf("%s should be error", v)
		}
	}
}

func TestExponent(t *testing.T) {
	inputs := [][]byte{[]byte("aaaa!!!!!0.000001e5				"), []byte("sdfsdfwer....0.000001e-5.......e"), []byte("$%&^ vfa0.000001e+5asdd"), []byte("0.000001e-5sdfdfw"), []byte("0.000001e+-5.5"), []byte("0.000001e+-5.5e+-5.5")}
	outputs := [][]byte{[]byte("0.000001e5"), []byte("0.000001e-5"), []byte("0.000001e+5"), []byte("0.000001e-5")}
	for i, input := range inputs {
		_, v, _, err := number.Find(input)
		if i < len(outputs) && err != nil {
			t.Error(err)
		}
		if i < len(outputs) && bytes.Equal(v, outputs[i]) == false {
			t.Errorf("%s != %s", v, outputs[i])
		}
		if i >= len(outputs) && err == nil {
			t.Errorf("%s should be error", v)
		}
	}
}
