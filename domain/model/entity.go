package model

type Entity struct {
	Name string
	Fields []Field
	JSON bool
	BSON bool
	Methods []Method
	Repositories []Repository
}

type Field struct {
	Name string
	Primary bool
	Op string
	Package string
	Type string
	Slice bool
}

type Method struct {
	Name string
	Type string
	Callbacks []Callback
}

type Callback struct {
	Name string
	Type string
}

type Repository struct {
	Type string
}