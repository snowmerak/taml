package generator

import (
	"errors"
	"reflect"
	"strings"

	"github.com/snowmerak/tson/lib/analyzer"
	"github.com/snowmerak/tson/lib/strcase"
)

func GoCode(name string, structure interface{}) (string, error) {

	members, err := analyzer.MembersOf(reflect.TypeOf(structure))
	if err != nil {
		return "", err
	}

	sb := strings.Builder{}

	sb.WriteString("package ")
	sb.WriteString(name)
	sb.WriteString("\n\n")
	sb.WriteString("import (\n")
	sb.WriteString("\t\"github.com/snowmerak/tson/lib/lexer/jsvalues\"\n")
	sb.WriteString("\t\"github.com/snowmerak/tson/lib/lexer/jsnumber\"\n")
	sb.WriteString("\t\"github.com/snowmerak/tson/lib/lexer/jsstring\"\n")
	sb.WriteString("\t\"github.com/snowmerak/tson/lib/lexer/jsbool\"\n")
	sb.WriteString("\t\"github.com/snowmerak/tson/lib/lexer/jsnull\"\n")
	sb.WriteString("\t\"strings\"\n")
	sb.WriteString("\t\"strconv\"\n")
	sb.WriteString("\t\"errors\"\n")
	sb.WriteString(")\n\n")

	sb.WriteString("type ")
	sb.WriteString(name)
	sb.WriteString(" struct {\n")
	for _, m := range members {
		sb.WriteString("\t")
		sb.WriteString(m.Name)
		sb.WriteString(" ")
		sb.WriteString(m.Type)
		sb.WriteString("\n")
	}
	sb.WriteString("}\n\n")

	sb.WriteString("func New(")
	for i, m := range members {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(m.Name)
		sb.WriteString(" ")
		sb.WriteString(m.Type)
	}
	sb.WriteString(") ")
	sb.WriteString(name)
	sb.WriteString(" {\n")
	sb.WriteString("\treturn ")
	sb.WriteString(name)
	sb.WriteString("{")
	for i, m := range members {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(m.Name)
	}
	sb.WriteString("}\n")
	sb.WriteString("}\n\n")

	sb.WriteString("func (v *")
	sb.WriteString(name)
	sb.WriteString(") ToJSON() string {\n")
	sb.WriteString("sb := strings.Builder{}\n")
	sb.WriteString("sb.WriteString(\"{ \")\n")
	for i, m := range members {
		if i > 0 {
			sb.WriteString("sb.WriteString(\", \")\n")
		}
		c, err := GoVar2JSON(m)
		if err != nil {
			return "", err
		}
		sb.WriteString(c)
	}
	sb.WriteString("sb.WriteString(\" }\")\n")
	sb.WriteString("return sb.String()\n")
	sb.WriteString("}\n\n")

	return sb.String(), nil
}

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
	case "[]byte", "[]string":
		sb.WriteString("sb.WriteString(\"[\")\n")
		sb.WriteString("for i, v := range v.")
		sb.WriteString(m.Name)
		sb.WriteString(" {\n")
		sb.WriteString("sb.WriteString(\"\\\"\")\n")
		sb.WriteString("sb.WriteString(string(v))\n")
		sb.WriteString("sb.WriteString(\"\\\"\")\n")
		sb.WriteString("if i < len(v) - 1 {\n")
		sb.WriteString("sb.WriteString(\", \")\n")
		sb.WriteString("}\n")
		sb.WriteString("}\n")
		sb.WriteString("sb.WriteString(\"]\")\n")
	case "[]int8", "[]int16", "[]int32", "[]int64":
		sb.WriteString("sb.WriteString(\"[\")\n")
		sb.WriteString("for i, v := range v.")
		sb.WriteString(m.Name)
		sb.WriteString(" {\n")
		sb.WriteString("sb.WriteString(strconv.FormatInt(int64(v), 10))\n")
		sb.WriteString("if i < len(v)-1 {\n")
		sb.WriteString("sb.WriteString(\", \")\n")
		sb.WriteString("}\n")
		sb.WriteString("}\n")
		sb.WriteString("sb.WriteString(\"]\")\n")
	case "[]uint8", "[]uint16", "[]uint32", "[]uint64", "[]uint":
		sb.WriteString("sb.WriteString(\"[\")\n")
		sb.WriteString("for i, v := range v.")
		sb.WriteString(m.Name)
		sb.WriteString(" {\n")
		sb.WriteString("sb.WriteString(strconv.FormatUint(uint64(v), 10))\n")
		sb.WriteString("if i < len(v)-1 {\n")
		sb.WriteString("sb.WriteString(\", \")\n")
		sb.WriteString("}\n")
		sb.WriteString("}\n")
		sb.WriteString("sb.WriteString(\"]\")\n")
	case "[]float32", "[]float64":
		sb.WriteString("sb.WriteString(\"[\")\n")
		sb.WriteString("for i, v := range v.")
		sb.WriteString(m.Name)
		sb.WriteString(" {\n")
		sb.WriteString("sb.WriteString(strconv.FormatFloat(float64(v), 'f', -1, 64))\n")
		sb.WriteString("if i < len(v)-1 {\n")
		sb.WriteString("sb.WriteString(\", \")\n")
		sb.WriteString("}\n")
		sb.WriteString("}\n")
		sb.WriteString("sb.WriteString(\"]\")\n")
	case "[]bool":
		sb.WriteString("sb.WriteString(\"[\")\n")
		sb.WriteString("for i, v := range v.")
		sb.WriteString(m.Name)
		sb.WriteString(" {\n")
		sb.WriteString("sb.WriteString(strconv.FormatBool(v))\n")
		sb.WriteString("if i < len(v)-1 {\n")
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
