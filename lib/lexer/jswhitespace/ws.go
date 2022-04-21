package jswhitespace

import (
	"errors"
)

func Find(data []byte) ([]byte, []byte, []byte, error) {
	s := -1
	e := -1
	for i := 0; i < len(data); i++ {
		if data[i] == ' ' || data[i] == '\t' || data[i] == '\n' || data[i] == '\r' {
			if s == -1 {
				s = i
			}
			e = i + 1
			continue
		}
		if s != -1 {
			break
		}
	}
	if s == e {
		return nil, nil, nil, errors.New("whitespace not found")
	}
	return data[:s], data[s:e], data[e:], nil
}
