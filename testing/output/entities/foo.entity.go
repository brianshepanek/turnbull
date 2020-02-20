package entities

type fooEntity struct {
	title string
	bars  []Bar
}
type FooEntity interface {
	Title() string
	Bars() []Bar
	SetTitle(title string)
	SetBars(bars []Bar)
}

func (fooEntity *fooEntity) Title() string {
	return fooEntity.title
}
func (fooEntity *fooEntity) Bars() []Bar {
	return fooEntity.bars
}
func (fooEntity *fooEntity) SetTitle(title string) {
	fooEntity.title = title
}
func (fooEntity *fooEntity) SetBars(bars []Bar) {
	fooEntity.bars = bars
}
