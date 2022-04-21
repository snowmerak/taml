package strcase

import "strings"

// PascalToSnake converts a pascal string to snake case.
func PascalToSnake(t string) string {
	sb := strings.Builder{}
	for i := 0; i < len(t); i++ {
		v := t[i]
		if v >= 'A' && v <= 'Z' {
			if i > 0 {
				sb.WriteByte('_')
			}
			v = v - 'A' + 'a'
		}
		sb.WriteByte(v)
	}
	return sb.String()
}

// SnakeToPascal converts a snake case string to pascal case.
func SnakeToPascal(t string) string {
	sb := strings.Builder{}
	for i := 0; i < len(t); i++ {
		v := t[i]
		if i == 0 {
			v = v - 'a' + 'A'
		}
		if i > 0 && v == '_' {
			v = t[i+1] - 'a' + 'A'
			i++
		}
		sb.WriteByte(v)
	}
	return sb.String()
}
