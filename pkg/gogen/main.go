package gogen

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/ashutshkumr/openapiart/pkg/spec"

	"gopkg.in/yaml.v2"
)

type SpecMeta struct {
	Spec spec.Spec
	Code strings.Builder
}

func ParseSpec(spec spec.Spec) error {
	log.Println("Parsing schemas ...")

	meta := SpecMeta{
		Spec: spec,
	}

	for name, schema := range spec.Schemas {
		if err := meta.ParseSchema(name, schema); err != nil {
			return fmt.Errorf("could not parse schema %s: %v", name, err)
		}
	}

	return WritePackage("pkg/sdk", meta.GetCode())
}

func ParseYamlBytes(bytes []byte) error {
	spec := spec.Spec{}
	if err := yaml.Unmarshal(bytes, &spec); err != nil {
		return fmt.Errorf("could not unmarshal bytes: %v", err)
	}

	return ParseSpec(spec)
}

func ParseYamlPath(path string) error {
	log.Printf("Parsing spec from %s ...\n", path)
	b, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("could not read file %s: %v", path, err)
	}

	return ParseYamlBytes(b)
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
