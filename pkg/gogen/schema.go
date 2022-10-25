package gogen

import (
	"fmt"
	"log"
	"strings"

	"github.com/ashutshkumr/openapiart/pkg/spec"
	"github.com/iancoleman/strcase"
)

func (m *SpecMeta) ParseSchema(name string, schema spec.Schema) error {
	log.Printf("Parsing schema %s ...\n", name)
	var body strings.Builder
	var methods strings.Builder

	for pName, prop := range schema.Properties {
		log.Printf("Parsing property %s.%s ...\n", name, pName)

		if pType := m.GetPrimitiveType(prop.Type, prop.Format); len(pType) != 0 {
			if _, err := body.WriteString(m.NewProperty(prop.Description, strcase.ToCamel(pName), pType)); err != nil {
				return fmt.Errorf("could not write property: %v", err)
			}

			if _, err := methods.WriteString(m.NewPropertyGetter(strcase.ToCamel(name), strcase.ToCamel(pName), pType, prop.Default)); err != nil {
				return fmt.Errorf("could not write getter: %v", err)
			}
			methods.WriteString("\n")
			if _, err := methods.WriteString(m.NewPropertySetter(strcase.ToCamel(name), strcase.ToCamel(pName), pType)); err != nil {
				return fmt.Errorf("could not write setter: %v", err)
			}
			methods.WriteString("\n")
		}
	}

	m.Code.WriteString(m.NewStruct(schema.Description, strcase.ToCamel(name), body.String()))
	m.Code.WriteString("\n")
	m.Code.WriteString(methods.String())
	return nil
}
