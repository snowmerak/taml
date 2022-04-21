package generator

import (
	"errors"
	"strings"

	"github.com/snowmerak/tson/lib/analyzer"
	"github.com/snowmerak/tson/lib/strcase"
)

func GoVar2JSON(m analyzer.Member) (string, error) {
	name := m.Name
	if strings.Contains(m.Name, ".") {
		names := strings.Split(m.Name, ".")
		name = names[len(names)-1]
	}
	sb := strings.Builder{}
	sb.WriteString("sb.WriteString(\"\\\"")
	sb.WriteString(strcase.PascalToSnake(name))
	sb.WriteString("\\\" : \")\n")
	switch m.Type {
	case "string":
		sb.WriteString("sb.WriteString(string(v.")
		sb.WriteString(m.Name)
		sb.WriteString("))\n")
	case "bool":
		sb.WriteString("sb.WriteString(strconv.FormatBool(v.")
		sb.WriteString(m.Name)
		sb.WriteString("))\n")
	case "int", "int8", "int16", "int32", "int64":
		sb.WriteString("sb.WriteString(strconv.FormatInt(int64(v.")
		sb.WriteString(m.Name)
		sb.WriteString("), 10))\n")
	case "uint", "uint8", "uint16", "uint32", "uint64":
		sb.WriteString("sb.WriteString(strconv.FormatUint(uint64(v.")
		sb.WriteString(m.Name)
		sb.WriteString("), 10))\n")
	case "float32", "float64":
		sb.WriteString("sb.WriteString(strconv.FormatFloat(float64(v.")
		sb.WriteString(m.Name)
		sb.WriteString("), 'f', -1, 64))\n")
	case "[]string":
		sb.WriteString("sb.WriteString(\"[\")\n")
		sb.WriteString("for i, x := range v.")
		sb.WriteString(m.Name)
		sb.WriteString(" {\n")
		sb.WriteString("sb.WriteString(\"\\\"\")\n")
		sb.WriteString("sb.WriteString(string(x))\n")
		sb.WriteString("sb.WriteString(\"\\\"\")\n")
		sb.WriteString("if i < len(v.")
		sb.WriteString(m.Name)
		sb.WriteString(")-1 {\n")
		sb.WriteString("sb.WriteString(\", \")\n")
		sb.WriteString("}\n")
		sb.WriteString("}\n")
		sb.WriteString("sb.WriteString(\"]\")\n")
	case "[]int8", "[]int16", "[]int32", "[]int64", "[]int":
		sb.WriteString("sb.WriteString(\"[\")\n")
		sb.WriteString("for i, x := range v.")
		sb.WriteString(m.Name)
		sb.WriteString(" {\n")
		sb.WriteString("sb.WriteString(strconv.FormatInt(int64(x), 10))\n")
		sb.WriteString("if i < len(v.")
		sb.WriteString(m.Name)
		sb.WriteString(")-1 {\n")
		sb.WriteString("sb.WriteString(\", \")\n")
		sb.WriteString("}\n")
		sb.WriteString("}\n")
		sb.WriteString("sb.WriteString(\"]\")\n")
	case "[]uint8", "[]uint16", "[]uint32", "[]uint64", "[]uint":
		sb.WriteString("sb.WriteString(\"[\")\n")
		sb.WriteString("for i, x := range v.")
		sb.WriteString(m.Name)
		sb.WriteString(" {\n")
		sb.WriteString("sb.WriteString(strconv.FormatUint(uint64(x), 10))\n")
		sb.WriteString("if i < len(v.")
		sb.WriteString(m.Name)
		sb.WriteString(")-1 {\n")
		sb.WriteString("sb.WriteString(\", \")\n")
		sb.WriteString("}\n")
		sb.WriteString("}\n")
		sb.WriteString("sb.WriteString(\"]\")\n")
	case "[]float32", "[]float64":
		sb.WriteString("sb.WriteString(\"[\")\n")
		sb.WriteString("for i, x := range v.")
		sb.WriteString(m.Name)
		sb.WriteString(" {\n")
		sb.WriteString("sb.WriteString(strconv.FormatFloat(float64(x), 'f', -1, 64))\n")
		sb.WriteString("if i < len(v.")
		sb.WriteString(m.Name)
		sb.WriteString(")-1 {\n")
		sb.WriteString("sb.WriteString(\", \")\n")
		sb.WriteString("}\n")
		sb.WriteString("}\n")
		sb.WriteString("sb.WriteString(\"]\")\n")
	case "[]bool":
		sb.WriteString("sb.WriteString(\"[\")\n")
		sb.WriteString("for i, x := range v.")
		sb.WriteString(m.Name)
		sb.WriteString(" {\n")
		sb.WriteString("sb.WriteString(strconv.FormatBool(x))\n")
		sb.WriteString("if i < len(v.")
		sb.WriteString(m.Name)
		sb.WriteString(")-1 {\n")
		sb.WriteString("sb.WriteString(\", \")\n")
		sb.WriteString("}\n")
		sb.WriteString("}\n")
		sb.WriteString("sb.WriteString(\"]\")\n")
	case "map[string]string":
		sb.WriteString("sb.WriteString(\"{ \")\n")
		sb.WriteString("l := 0\n")
		sb.WriteString("for k, x := range v.")
		sb.WriteString(m.Name)
		sb.WriteString(" {\n")
		sb.WriteString("sb.WriteString(\"\\\"\")\n")
		sb.WriteString("sb.WriteString(k)\n")
		sb.WriteString("sb.WriteString(\"\\\" : \")\n")
		sb.WriteString("sb.WriteString(\"\\\"\")\n")
		sb.WriteString("sb.WriteString(x)\n")
		sb.WriteString("sb.WriteString(\"\\\"\")\n")
		sb.WriteString("if l < len(v.")
		sb.WriteString(m.Name)
		sb.WriteString(") - 1 {\n")
		sb.WriteString("sb.WriteString(\", \")\n")
		sb.WriteString("l++\n")
		sb.WriteString("}\n")
		sb.WriteString("}\n")
		sb.WriteString("sb.WriteString(\" }\")\n")
	case "map[string]int", "map[string]int8", "map[string]int16", "map[string]int32", "map[string]int64":
		sb.WriteString("sb.WriteString(\"{ \")\n")
		sb.WriteString("l := 0\n")
		sb.WriteString("for k, x := range v.")
		sb.WriteString(m.Name)
		sb.WriteString(" {\n")
		sb.WriteString("sb.WriteString(\"\\\"\")\n")
		sb.WriteString("sb.WriteString(k)\n")
		sb.WriteString("sb.WriteString(\"\\\" : \")\n")
		sb.WriteString("sb.WriteString(strconv.FormatInt(int64(x), 10))\n")
		sb.WriteString("if l < len(v.")
		sb.WriteString(m.Name)
		sb.WriteString(") - 1 {\n")
		sb.WriteString("sb.WriteString(\", \")\n")
		sb.WriteString("l++\n")
		sb.WriteString("}\n")
		sb.WriteString("}\n")
		sb.WriteString("sb.WriteString(\" }\")\n")
	case "map[string]uint", "map[string]uint8", "map[string]uint16", "map[string]uint32", "map[string]uint64":
		sb.WriteString("sb.WriteString(\"{ \")\n")
		sb.WriteString("l := 0\n")
		sb.WriteString("for k, x := range v.")
		sb.WriteString(m.Name)
		sb.WriteString(" {\n")
		sb.WriteString("sb.WriteString(\"\\\"\")\n")
		sb.WriteString("sb.WriteString(k)\n")
		sb.WriteString("sb.WriteString(\"\\\" : \")\n")
		sb.WriteString("sb.WriteString(strconv.FormatUint(uint64(x), 10))\n")
		sb.WriteString("if l < len(v.")
		sb.WriteString(m.Name)
		sb.WriteString(") - 1 {\n")
		sb.WriteString("sb.WriteString(\", \")\n")
		sb.WriteString("l++\n")
		sb.WriteString("}\n")
		sb.WriteString("}\n")
		sb.WriteString("sb.WriteString(\" }\")\n")
	case "map[string]float32", "map[string]float64":
		sb.WriteString("sb.WriteString(\"{ \")\n")
		sb.WriteString("l := 0\n")
		sb.WriteString("for k, x := range v.")
		sb.WriteString(m.Name)
		sb.WriteString(" {\n")
		sb.WriteString("sb.WriteString(\"\\\"\")\n")
		sb.WriteString("sb.WriteString(k)\n")
		sb.WriteString("sb.WriteString(\"\\\" : \")\n")
		sb.WriteString("sb.WriteString(strconv.FormatFloat(float64(x, 'f', -1, 64)))\n")
		sb.WriteString("), 'f', -1, 64))\n")
		sb.WriteString("if l < len(v.")
		sb.WriteString(m.Name)
		sb.WriteString(") - 1 {\n")
		sb.WriteString("sb.WriteString(\", \")\n")
		sb.WriteString("l++\n")
		sb.WriteString("}\n")
		sb.WriteString("}\n")
		sb.WriteString("sb.WriteString(\" }\")\n")
	case "map[string]bool":
		sb.WriteString("sb.WriteString(\"{ \")\n")
		sb.WriteString("l := 0\n")
		sb.WriteString("for k, x := range v.")
		sb.WriteString(m.Name)
		sb.WriteString(" {\n")
		sb.WriteString("sb.WriteString(\"\\\"\")\n")
		sb.WriteString("sb.WriteString(k)\n")
		sb.WriteString("sb.WriteString(\"\\\" : \")\n")
		sb.WriteString("sb.WriteString(strconv.FormatBool(x))\n")
		sb.WriteString("if l < len(v.")
		sb.WriteString(m.Name)
		sb.WriteString(") - 1 {\n")
		sb.WriteString("sb.WriteString(\", \")\n")
		sb.WriteString("l++\n")
		sb.WriteString("}\n")
		sb.WriteString("}\n")
		sb.WriteString("sb.WriteString(\" }\")\n")
	default:
		if len(m.SubMmbers) == 0 {
			return "", errors.New("unknown type: " + m.Type)
		}
		sb.WriteString("sb.WriteString(\"{ \")\n")
		for i, v := range m.SubMmbers {
			if i > 0 {
				sb.WriteString("sb.WriteString(\", \")\n")
			}
			v.Name = m.Name + "." + v.Name
			s, err := GoVar2JSON(v)
			if err != nil {
				return "", err
			}
			sb.WriteString(s)
		}
		sb.WriteString("sb.WriteString(\" }\")\n")
	}

	return sb.String(), nil
}
