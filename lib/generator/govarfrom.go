package generator

import (
	"strings"

	"github.com/snowmerak/tson/lib/analyzer"
	"github.com/snowmerak/tson/lib/strcase"
)

func GoVarFromJSON(m analyzer.Member) (string, error) {
	name := strcase.SnakeToPascal(m.Name)
	sb := strings.Builder{}
	switch m.Type {
	case "string":
		sb.WriteString("\tv.")
		sb.WriteString(name)
		sb.WriteString(" = ")
		sb.WriteString(`values.Get("` + m.Name + `").String()`)
	case "int", "int8", "int16", "int32", "int64":
		sb.WriteString("\tv.")
		sb.WriteString(name)
		sb.WriteString(" = ")
		sb.WriteString(m.Type + `(` + `values.GetInt64("` + m.Name + `")`)
	case "uint", "uint8", "uint16", "uint32", "uint64":
		sb.WriteString("\tv.")
		sb.WriteString(name)
		sb.WriteString(" = ")
		sb.WriteString(m.Type + `(` + `values.GetUint64("` + m.Name + `")`)
	case "float32", "float64":
		sb.WriteString("\tv.")
		sb.WriteString(name)
		sb.WriteString(" = ")
		sb.WriteString(m.Type + `(` + `values.GetFloat64("` + m.Name + `")`)
	case "bool":
		sb.WriteString("\tv.")
		sb.WriteString(name)
		sb.WriteString(" = ")
		sb.WriteString(`values.GetBool("` + m.Name + `")`)
	case "[]string":
		sb.WriteString(`{\narr := values.GetArray("` + m.Name + `")`)
		sb.WriteString(`for _, x := range arr {`)
		sb.WriteString(`\tv.`)
		sb.WriteString(name)
		sb.WriteString(` = append(v.`)
		sb.WriteString(name)
		sb.WriteString(`, x.String())`)
		sb.WriteString(`}\n}\n`)
	case "[]int", "[]int8", "[]int16", "[]int32", "[]int64":
		sb.WriteString(`{\narr := values.GetArray("` + m.Name + `")`)
		sb.WriteString(`for _, x := range arr {`)
		sb.WriteString(`\tv.`)
		sb.WriteString(name)
		sb.WriteString(` = append(v.`)
		sb.WriteString(name)
		sb.WriteString(`, `)
		sb.WriteString(m.Type[2:])
		sb.WriteString(`(x.GetInt64()))`)
		sb.WriteString(`}\n}\n`)
	case "[]uint", "[]uint8", "[]uint16", "[]uint32", "[]uint64":
		sb.WriteString(`{\narr := values.GetArray("` + m.Name + `")`)
		sb.WriteString(`for _, x := range arr {`)
		sb.WriteString(`\tv.`)
		sb.WriteString(name)
		sb.WriteString(` = append(v.`)
		sb.WriteString(name)
		sb.WriteString(`, `)
		sb.WriteString(m.Type[2:])
		sb.WriteString(`(x.GetUint64()))`)
		sb.WriteString(`}\n}\n`)
	case "[]float32", "[]float64":
		sb.WriteString(`{\narr := values.GetArray("` + m.Name + `")`)
		sb.WriteString(`for _, x := range arr {`)
		sb.WriteString(`\tv.`)
		sb.WriteString(name)
		sb.WriteString(` = append(v.`)
		sb.WriteString(name)
		sb.WriteString(`, `)
		sb.WriteString(m.Type[2:])
		sb.WriteString(`(x.GetFloat64()))`)
		sb.WriteString(`}\n}\n`)
	case "[]bool":
		sb.WriteString(`{\narr := values.GetArray("` + m.Name + `")`)
		sb.WriteString(`for _, x := range arr {`)
		sb.WriteString(`\tv.`)
		sb.WriteString(name)
		sb.WriteString(` = append(v.`)
		sb.WriteString(name)
		sb.WriteString(`, x.GetBool())`)
		sb.WriteString(`}\n}\n`)
	case "map[string]string":
		sb.WriteString(`values.GetObject("` + m.Name + `").Visit(func(key []byte, v *fastjson.Value) {\n`)
		sb.WriteString(`\tv.`)
		sb.WriteString(name)
		sb.WriteString(`[string(key)] = v.String()\n`)
		sb.WriteString(`})\n`)
	case "map[string]int", "map[string]int8", "map[string]int16", "map[string]int32", "map[string]int64":
		sb.WriteString(`values.GetObject("` + m.Name + `").Visit(func(key []byte, v *fastjson.Value) {\n`)
		sb.WriteString(`\tv.`)
		sb.WriteString(name)
		sb.WriteString(`[string(key)] = `)
		sb.WriteString(m.Type[11:])
		sb.WriteString(`(v.GetInt64())`)
		sb.WriteString(`\n})\n`)
	case "map[string]uint", "map[string]uint8", "map[string]uint16", "map[string]uint32", "map[string]uint64":
		sb.WriteString(`values.GetObject("` + m.Name + `").Visit(func(key []byte, v *fastjson.Value) {\n`)
		sb.WriteString(`\tv.`)
		sb.WriteString(name)
		sb.WriteString(`[string(key)] = `)
		sb.WriteString(m.Type[11:])
		sb.WriteString(`(v.GetUint64())`)
		sb.WriteString(`\n})\n`)
	case "map[string]float32", "map[string]float64":
		sb.WriteString(`values.GetObject("` + m.Name + `").Visit(func(key []byte, v *fastjson.Value) {\n`)
		sb.WriteString(`\tv.`)
		sb.WriteString(name)
		sb.WriteString(`[string(key)] = `)
		sb.WriteString(m.Type[11:])
		sb.WriteString(`(v.GetFloat64())`)
		sb.WriteString(`\n})\n`)
	case "map[string]bool":
		sb.WriteString(`values.GetObject("` + m.Name + `").Visit(func(key []byte, v *fastjson.Value) {\n`)
		sb.WriteString(`\tv.`)
		sb.WriteString(name)
		sb.WriteString(`[string(key)] = v.GetBool()`)
		sb.WriteString(`\n})\n`)
	}
	return sb.String(), nil
}
