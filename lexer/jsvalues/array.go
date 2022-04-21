package jsvalues

import "errors"

func FindArray(data []byte) ([]byte, []byte, []byte, error) {
	s := -1
	e := -1
	for i := 0; i < len(data); i++ {
		if data[i] == '[' {
			s = i
			if i+1 < len(data) && data[i+1] == ']' {
				e = i + 2
				return data[:s], data[s:e], data[e:], nil
			}
			break
		}
	}
	if s == -1 {
		return nil, nil, nil, errors.New("no array found")
	}
	prev, value, _, err := FindValue(data[s+1:])
	if err != nil {
		return nil, nil, nil, err
	}
	e = s + 1 + len(prev) + len(value)
	for i := e; i < len(data); i++ {
		if data[i] == ',' {
			prev, value, _, err := FindValue(data[i+1:])
			if err != nil {
				return nil, nil, nil, err
			}
			e = i + 1 + len(prev) + len(value)
			i = e
		}
	}
	closed := false
	for i := e; i < len(data); i++ {
		if data[i] == ']' {
			e = i + 1
			closed = true
			break
		}
	}
	if !closed {
		return nil, nil, nil, errors.New("array not closed")
	}
	return data[:s], data[s:e], data[e:], nil
}
