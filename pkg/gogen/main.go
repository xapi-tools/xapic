package gogen

import (
	"context"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

type SpecMeta struct {
	Spec            *openapi3.T
	Code            strings.Builder
	SchemaGenerated map[string]struct{}
}

func ParseSpec(spec *openapi3.T) error {
	log.Println("Parsing spec ...")

	meta := SpecMeta{Spec: spec}
	meta.LogVersions()
	if err := meta.ParseSchemas(); err != nil {
		return fmt.Errorf("could not parse schemas: %v", err)
	}

	return WritePackage("pkg/sdk", meta.GetCode())
}

func ParseYamlPath(path string) error {
	log.Printf("Loading spec from %s ...\n", path)

	loader := openapi3.Loader{Context: context.Background()}
	spec, err := loader.LoadFromFile(path)
	if err != nil {
		return fmt.Errorf("could not load file %s: %v", path, err)
	}

	log.Println("Successfully loaded.")
	return ParseSpec(spec)
}

func WritePackage(dir string, code string) error {
	if err := os.RemoveAll(dir); err != nil {
		return fmt.Errorf("could not remove %s: %v", dir, err)
	}

	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("could not create %s: %v", dir, err)
	}

	_, name := path.Split(dir)
	path := path.Join(dir, name+".go")

	code = "package " + name + "\n\n" + code

	if err := os.WriteFile(path, []byte(code), 0644); err != nil {
		return fmt.Errorf("could not write file %s: %v", path, err)
	}

	return nil
}
