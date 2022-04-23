package generator

import (
	"reflect"
	"strings"

	"github.com/snowmerak/tson/lib/analyzer"
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
	sb.WriteString("\t\"github.com/snowmerak/tson/lib/strcase\"\n")
	sb.WriteString("\t\"github.com/valyala/fastjson\"")
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

	sb.WriteString("func (v *")
	sb.WriteString(name)
	sb.WriteString(") FromJSON(data []byte) error {\n")
	sb.WriteString("values, err := fastjson.ParseBytes(data)\n")
	sb.WriteString("return nil\n")
	sb.WriteString("}\n\n")

	return sb.String(), nil
}
