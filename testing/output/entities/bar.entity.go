package entities

type barEntity struct {
	title string
}
type BarEntity interface {
	Title() string
	SetTitle(title string)
}

func (barEntity *barEntity) Title() string {
	return barEntity.title
}
func (barEntity *barEntity) SetTitle(title string) {
	barEntity.title = title
}
