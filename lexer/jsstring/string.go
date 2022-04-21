package jsstring

import "errors"

func Find(data []byte) ([]byte, []byte, []byte, error) {
	s := -1
	e := -1
	for i := 0; i < len(data); i++ {
		if data[i] == '"' {
			if s == -1 {
				s = i
			} else {
				e = i + 1
				break
			}
		}
		if data[i] == '\\' {
			if data[i] == '"' || data[i] == '\\' || data[i] == '/' || data[i] == 'b' || data[i] == 'f' || data[i] == 'n' || data[i] == 'r' || data[i] == 't' || data[i] == 'u' {
				i++
				e++
			} else {
				return nil, nil, nil, nil
			}
		}
	}
	if s == -1 || e == -1 {
		return nil, nil, nil, errors.New("no string found")
	}
	return data[:s], data[s:e], data[e:], nil
}
