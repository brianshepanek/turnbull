package model

type Entity struct {
	Name string
	Interface bool
	Fields []Field
	Methods []Method
	Repositories []Repository
	Presenters []Presenter
	Controllers []Controller
}

type Field struct {
	Name string
	Primary bool
	Op string
	Package string
	Type string
	Slice bool
	Private bool
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
	Primary bool
}

type Presenter struct {
	Type string
	Primary bool
}

type Controller struct {
	Type string
	Primary bool
}