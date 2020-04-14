package model

type Entity struct {
	Name string
	Interface bool
	Fields []Field
	Methods []Method
	Repositories []Repository
	Presenters []Presenter
	Controllers []Controller
	Interactors []string
}

type Field struct {
	Name string
	Primary bool
	Op string
	Package string
	Type string
	Slice bool
	Private bool
	Embedded bool
	Entity Entity
}

type Method struct {
	Name string
	Type string
	Callbacks []Callback
	Repository RepositoryMethod
	Presenter PresenterMethod
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

type RepositoryMethod struct {
	Name string
	Arguments []Field
	ReturnValues []Field
}

type PresenterMethod struct {
	Name string
	Arguments []Field
	ReturnValues []Field
}