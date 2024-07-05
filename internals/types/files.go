package types

type DtoFile struct {
	NameSpace string
	Imports   []string
	ClassName string
	Fields    []Field
}

type BasicJsonController struct {
	NameSpace string
	ClassName string
	Route     string
	RouteName string
}

type AdvancedJsonController struct {
	NameSpace    string
	ClassName    string
	Route        string
	RouteName    string
	ServiceNS    string
	RepositoryNS string
}

type File struct {
	NameSpace string
	ClassName string
}

type ServiceFile struct {
	NameSpace    string
	ClassName    string
	RepositoryNS string
	MapperNS     string
}

type Field struct {
	Type string
	Name string
}
