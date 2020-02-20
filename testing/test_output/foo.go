package entities

type fooEntity struct {
	title    string
	subtitle []string
}
type FooEntity interface {
	Title() string
	Subtitle() []string
	SetTitle(title string)
	SetSubtitle(subtitle []string)
	MarshalJSON() ([]byte, error)
}

func (fooEntity *fooEntity) Title() string {
	return fooEntity.title
}
func (fooEntity *fooEntity) Subtitle() []string {
	return fooEntity.subtitle
}
func (fooEntity *fooEntity) SetTitle(title string) {
	fooEntity.title = title
}
func (fooEntity *fooEntity) SetSubtitle(subtitle []string) {
	fooEntity.subtitle = subtitle
}
