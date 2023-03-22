package protogen

type Proto struct {
	Messages []Message
}

type Message struct {
	Description string
	Name        string
	Fields      []MessageField
	Enums       []Enum
	Messages    []Message
}

type MessageField struct {
	Description string
	Id          uint
	Name        string
	Type        string
	Optional    bool
	Repeated    bool
}

type Enum struct {
	Description string
	Constants   []EnumConstant
}

type EnumConstant struct {
	Description string
	Name        string
	Value       string
}
