package jsnumber

import "errors"

func Find(data []byte) ([]byte, []byte, []byte, error) {
	s := -1
	e := -1
	for i := 0; i < len(data); i++ {
		if data[i] == '-' || data[i] == '0' || data[i] == '1' || data[i] == '2' || data[i] == '3' || data[i] == '4' || data[i] == '5' || data[i] == '6' || data[i] == '7' || data[i] == '8' || data[i] == '9' {
			s = i
			e = s + 1
			break
		}
	}
	if s == -1 {
		return nil, nil, nil, errors.New("no number found")
	}
	for i := e; i < len(data); i++ {
		if data[i] == '0' || data[i] == '1' || data[i] == '2' || data[i] == '3' || data[i] == '4' || data[i] == '5' || data[i] == '6' || data[i] == '7' || data[i] == '8' || data[i] == '9' {
			e++
			continue
		}
		break
	}
	if e < len(data) && data[e] == '.' {
		e++
	}
	for i := e; i < len(data); i++ {
		if data[i] == '0' || data[i] == '1' || data[i] == '2' || data[i] == '3' || data[i] == '4' || data[i] == '5' || data[i] == '6' || data[i] == '7' || data[i] == '8' || data[i] == '9' {
			e++
			continue
		}
		break
	}
	if e < len(data) && (data[e] == 'e' || data[e] == 'E') {
		e++
		if e < len(data) && (data[e] == '+' || data[e] == '-') {
			e++
		}
		for i := e; i < len(data); i++ {
			if data[i] == '0' || data[i] == '1' || data[i] == '2' || data[i] == '3' || data[i] == '4' || data[i] == '5' || data[i] == '6' || data[i] == '7' || data[i] == '8' || data[i] == '9' {
				e++
				continue
			}
			break
		}
	}
	if s == -1 || data[e-1] == '.' || data[e-1] == 'e' || data[e-1] == 'E' || data[e-1] == '+' || data[e-1] == '-' {
		return nil, nil, nil, errors.New("number not found")
	}
	return data[:s], data[s:e], data[e:], nil
}
