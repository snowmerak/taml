package jsbool

import "errors"

func Find(data []byte) ([]byte, []byte, []byte, error) {
	trues := [4]byte{'t', 'r', 'u', 'e'}
	falses := [5]byte{'f', 'a', 'l', 's', 'e'}
	for i := 0; i < len(data); i++ {
		if data[i] == 't' {
			j := 0
			for ; j < 4 && i+j < len(data); j++ {
				if data[i+j] != trues[j] {
					break
				}
			}
			if j == 4 {
				return data[:i], data[i : i+4], data[i+4:], nil
			}
		}
		if data[i] == 'f' {
			j := 0
			for ; j < 5 && i+j < len(data); j++ {
				if data[i+j] != falses[j] {
					break
				}
			}
			if j == 5 {
				return data[:i], data[i : i+5], data[i+5:], nil
			}
		}
	}
	return nil, nil, nil, errors.New("no bool found")
}
