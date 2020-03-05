package model

type FooScaffoldStruct struct {
	id     string
	title  string
	things []string
}
type FooScaffoldInterface interface {
	Id() string
	Primary() string
	Title() string
	Things() []string
	SetId(id string)
	SetTitle(Title string)
	SetThings(Things []string)
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
