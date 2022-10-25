package spec

type Spec struct {
	Schemas Schemas `yaml:"schemas"`
}

type Schemas map[string]Schema

type Schema struct {
	Description string     `yaml:"description"`
	Type        string     `yaml:"type"`
	Properties  Properties `yaml:"properties"`
}

type Properties map[string]Property

type Property struct {
	Description string   `yaml:"description"`
	Type        string   `yaml:"type"`
	Format      string   `yaml:"format"`
	Default     string   `yaml:"default"`
	Enum        []string `yaml:"enum"`
}
