package entity

type luckyStruct struct {
	thing  *string
	thing2 *int64
}

func newLucky() *lucky {
	return &lucky{luckyStruct: newLuckyStruct()}
}

func newLuckyStruct() *luckyStruct {
	return &luckyStruct{}
}

type luckiesStruct []luckyInterface

type luckyInterface interface {
	Thing() *string
	Thing2() *int64
	SetThing(thing *string)
	SetThing2(thing2 *int64)
}
type luckiesInterface interface {
	Len() int
	Append(req luckyInterface)
	Elements() []luckyInterface
}

func (m *luckiesStruct) Len() int {
	if m != nil {
		return len(*m)
	}
	return 0
}
func (m *luckiesStruct) Append(req luckyInterface) {
	if m != nil {
		*m = append(*m, req)
	}
}
func (m *luckiesStruct) Elements() []luckyInterface {
	return *m
}
func (m *luckyStruct) Thing() *string {
	return m.thing
}

func (m *luckyStruct) Thing2() *int64 {
	return m.thing2
}

func (m *luckyStruct) SetThing(thing *string) {
	m.thing = thing
}

func (m *luckyStruct) SetThing2(thing2 *int64) {
	m.thing2 = thing2
}
