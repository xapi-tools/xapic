package gogen

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
)

func (m *SpecMeta) GetPrimitiveType(typ string, format string) string {
	switch typ {
	case "integer":
		switch format {
		case "int32":
			return "int32"
		case "int64":
			return "int64"
		default:
			return "int"
		}
	case "number":
		switch format {
		case "float":
			return "float32"
		case "double":
			return "float64"
		default:
			return "float64"
		}
	case "string":
		return "string"
	default:
		return ""
	}
}

func (m *SpecMeta) GetDefaultValue(val string, typ string) string {
	if len(strings.TrimSpace(val)) == 0 {
		return ""
	}

	switch typ {
	case "string":
		return fmt.Sprintf("\"%s\"", val)
	default:
		return val
	}
}

func (m *SpecMeta) GetCode() string {
	return m.Code.String()
}

func (m *SpecMeta) GetDoc(doc string) string {
	if len(strings.TrimSpace(doc)) == 0 {
		return "TODO: Add comment"
	}

	return doc
}

func (m *SpecMeta) NewStruct(doc string, name string, body string) string {
	return fmt.Sprintf(
		"// %s\ntype %s struct {\n%s}\n",
		m.GetDoc(doc), name, body,
	)
}

func (m *SpecMeta) NewMethod(doc string, objName string, objType string, name string, inArgs []string, outTypes []string, body string) string {
	out := ""
	if len(outTypes) == 1 {
		out = outTypes[0]
	} else if len(outTypes) > 1 {
		out = "(" + strings.Join(outTypes, ",") + ")"
	}
	return fmt.Sprintf(
		"// %s\nfunc (%s %s) %s(%s) %s {\n%s}\n",
		m.GetDoc(doc), objName, objType, name, strings.Join(inArgs, ","), out, body,
	)
}

func (m *SpecMeta) NewProperty(doc string, name string, typ string) string {
	return fmt.Sprintf(
		"    // %s\n    %s *%s `yaml:\"%s\"`\n",
		m.GetDoc(doc), name, typ, strcase.ToSnake(name),
	)
}

func (m *SpecMeta) GetterName(propName string) string {
	return "Get" + strcase.ToCamel(propName)
}

func (m *SpecMeta) SetterName(propName string) string {
	return "Set" + strcase.ToCamel(propName)
}

func (m *SpecMeta) NewPropertyGetter(objType string, propName string, propType string, defaultVal string) string {
	getter := m.GetterName(propName)
	setter := m.SetterName(propName)
	doc := fmt.Sprintf("%s returns the value of %s", getter, propName)

	body := fmt.Sprintf("    return *o.%s\n", propName)

	if defVal := m.GetDefaultValue(defaultVal, propType); len(defVal) == 0 {

	} else {
		body = fmt.Sprintf(
			"    if o.%s == nil {\n        o.%s(%s)\n    }\n\n",
			propName, setter, defVal,
		) + body
	}

	return m.NewMethod(doc, "o", "*"+objType, getter, []string{}, []string{propType}, body)
}

func (m *SpecMeta) NewPropertySetter(objType string, propName string, propType string) string {
	setter := m.SetterName(propName)
	doc := fmt.Sprintf("%s sets the value of %s", setter, propName)

	body := fmt.Sprintf(
		"    if o.%s == nil {\n        o.%s = new(%s)\n    }\n    *o.%s = val\n\n    return o\n",
		propName, propName, propType, propName,
	)

	return m.NewMethod(doc, "o", "*"+objType, setter, []string{"val " + propType}, []string{"*" + objType}, body)
}
