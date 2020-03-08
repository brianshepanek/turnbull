package model

type FooScaffoldStruct struct {
	id     string
	title  string
	things []string
}

type FoosScaffoldStruct []FooScaffoldInterface

type FooScaffoldInterface interface {
	Id() string
	Primary() string
	Title() string
	Things() []string
	SetId(id string)
	SetTitle(Title string)
	SetThings(Things []string)
	SetAll(req FooScaffoldInterface)
}
type FoosScaffoldInterface interface {
	Len() int
	Append(req FooScaffoldInterface)
	Elements() []FooScaffoldInterface
}

func NewFooScaffoldStruct() FooScaffoldInterface {
	return &FooScaffoldStruct{}
}

func NewFoosScaffoldStruct() FoosScaffoldInterface {
	return &FoosScaffoldStruct{}
}

func (m *FoosScaffoldStruct) Len() int {
	if m != nil {
		return len(*m)
	}
	return 0
}
func (m *FoosScaffoldStruct) Append(req FooScaffoldInterface) {
	if m != nil {
		*m = append(*m, req)
	}
}
func (m *FoosScaffoldStruct) Elements() []FooScaffoldInterface {
	return *m
}
func (m *FooScaffoldStruct) Id() string {
	return m.id
}

func (m *FooScaffoldStruct) Primary() string {
	return m.id
}

func (m *FooScaffoldStruct) Title() string {
	return m.title
}

func (m *FooScaffoldStruct) Things() []string {
	return m.things
}

func (m *FooScaffoldStruct) SetId(id string) {
	m.id = id
}

func (m *FooScaffoldStruct) SetTitle(title string) {
	m.title = title
}

func (m *FooScaffoldStruct) SetThings(things []string) {
	m.things = things
}

func (m *FooScaffoldStruct) SetAll(req FooScaffoldInterface) {
	m.SetId(req.Id())
	m.SetTitle(req.Title())
	m.SetThings(req.Things())
}
