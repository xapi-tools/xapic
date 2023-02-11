package gogen

import (
	"fmt"
	"log"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/iancoleman/strcase"
)

func (m *SpecMeta) LogVersions() {
	log.Println("OpenAPI version:", m.Spec.OpenAPI)
	log.Println("API version:", m.Spec.Info.Version)
}

func (m *SpecMeta) ParseSchemas() error {
	log.Printf("Parsing schemas ...")
	if m.Spec.Components != nil {
		for name, schema := range m.Spec.Components.Schemas {
			if err := m.ParseSchema(name, schema); err != nil {
				return fmt.Errorf("could not parse schema %s: %v", name, err)
			}
		}
	}
	return nil
}

func (m *SpecMeta) ParseSchema(name string, schema *openapi3.SchemaRef) error {
	if _, ok := m.SchemaGenerated[schema.Ref]; ok {
		return nil
	}

	log.Printf("Generating Type %s ...\n", schema.Ref)
	var body strings.Builder
	var methods strings.Builder

	if schema.Value.Type == "object" {
		for pName, prop := range schema.Value.Properties {
			logPrefix := fmt.Sprintf("[%s.%s]", name, pName)
			log.Printf("%s Parsing property ...\n", logPrefix)

			// switch prop.Value.Type {
			// case "object":
			// 	m.ParseSchema()
			// }
			if pType := m.GetPrimitiveType(prop.Value.Type, prop.Value.Format); len(pType) != 0 {
				if _, err := body.WriteString(m.NewProperty(prop.Value.Description, strcase.ToCamel(pName), pType)); err != nil {
					return fmt.Errorf("could not write property: %v", err)
				}

				if _, err := methods.WriteString(m.NewPropertyGetter(strcase.ToCamel(name), strcase.ToCamel(pName), pType, prop.Value.Default)); err != nil {
					return fmt.Errorf("could not write getter: %v", err)
				}
				methods.WriteString("\n")
				if _, err := methods.WriteString(m.NewPropertySetter(strcase.ToCamel(name), strcase.ToCamel(pName), pType)); err != nil {
					return fmt.Errorf("could not write setter: %v", err)
				}
				methods.WriteString("\n")
			}
		}
	}

	m.Code.WriteString(m.NewStruct(schema.Value.Description, strcase.ToCamel(name), body.String()))
	m.Code.WriteString("\n")
	m.Code.WriteString(methods.String())
	return nil
}
