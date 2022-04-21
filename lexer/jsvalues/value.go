package jsvalues

import (
	"errors"

	"github.com/snowmerak/tson/lexer/jsbool"
	"github.com/snowmerak/tson/lexer/jsnull"
	"github.com/snowmerak/tson/lexer/jsnumber"
	"github.com/snowmerak/tson/lexer/jsstring"
)

func FindValue(data []byte) ([]byte, []byte, []byte, error) {
	prevs := [][]byte{}
	values := [][]byte{}
	nexts := [][]byte{}
	prev, value, next, err := FindObject(data)
	if err == nil {
		prevs = append(prevs, prev)
		values = append(values, value)
		nexts = append(nexts, next)
	}
	prev, value, next, err = FindArray(data)
	if err == nil {
		prevs = append(prevs, prev)
		values = append(values, value)
		nexts = append(nexts, next)
	}
	prev, value, next, err = jsstring.Find(data)
	if err == nil {
		prevs = append(prevs, prev)
		values = append(values, value)
		nexts = append(nexts, next)
	}
	prev, value, next, err = jsnumber.Find(data)
	if err == nil {
		prevs = append(prevs, prev)
		values = append(values, value)
		nexts = append(nexts, next)
	}
	prev, value, next, err = jsbool.Find(data)
	if err == nil {
		prevs = append(prevs, prev)
		values = append(values, value)
		nexts = append(nexts, next)
	}
	prev, value, next, err = jsnull.Find(data)
	if err == nil {
		prevs = append(prevs, prev)
		values = append(values, value)
		nexts = append(nexts, next)
	}
	if len(prevs) == 0 {
		return nil, nil, nil, errors.New("no value found")
	}
	min := 0
	for i := 1; i < len(prevs); i++ {
		if len(prevs[i]) < len(prevs[min]) {
			min = i
		}
	}
	return prevs[min], values[min], nexts[min], nil
}
