package entity

type lucky struct {
	*luckyStruct
}

type luckies struct {
	luckiesStruct
}

type Lucky interface {
	luckyInterface
}
type Luckies interface {
	luckiesInterface
}

func NewLucky() Lucky {
	return newLucky()
}

func NewLuckies() Luckies {
	return &luckies{}
}
