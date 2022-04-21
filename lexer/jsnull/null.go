package jsnull

import "errors"

func Find(data []byte) ([]byte, []byte, []byte, error) {
	null := []byte{'n', 'u', 'l', 'l'}
	for i := 0; i < len(data); i++ {
		if data[i] == 'n' {
			j := 0
			for ; j < 4 && i+j < len(data); j++ {
				if data[i+j] != null[j] {
					break
				}
			}
			if j == 4 {
				return data[:i], data[i : i+4], data[i+4:], nil
			}
		}
	}
	return nil, nil, nil, errors.New("null not found")
}
