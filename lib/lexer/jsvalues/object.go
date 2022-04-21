package jsvalues

import (
	"errors"

	"github.com/snowmerak/tson/lib/lexer/jswhitespace"
)

func FindObject(data []byte) ([]byte, []byte, []byte, error) {
	s := -1
	e := -1
	for i := 0; i < len(data); i++ {
		if data[i] == '{' {
			s = i
			if i+1 < len(data) && data[i+1] == '}' {
				e = i + 2
				return data[:s], data[s:e], data[e:], nil
			}
			break
		}
	}
	if s == -1 {
		return nil, nil, nil, errors.New("no object found")
	}
	nums := 0
	for i := s + 1; i < len(data); i++ {
		if nums > 0 {
			if data[i] == ',' {
				i++
				continue
			}
			return nil, nil, nil, errors.New("no object found")
		}
		prev, value, _, err := FindValue(data[i:])
		if err != nil {
			break
		}
		e = i + len(prev) + len(value)
		if e < len(data) && data[e] == ':' {
			e++
		} else {
			prev, value, _, err = jswhitespace.Find(data[e:])
			if err != nil {
				return nil, nil, nil, err
			}
			e += len(prev) + len(value)
		}
		prev, value, _, err = FindValue(data[e:])
		if err != nil {
			return nil, nil, nil, err
		}
		e += len(prev) + len(value)
		i = e
	}
	for i := e; i < len(data); i++ {
		if data[i] == '}' {
			e = i + 1
			break
		}
	}
	if s == -1 {
		return nil, nil, nil, errors.New("no object found")
	}
	return data[:s], data[s:e], data[e:], nil
}
