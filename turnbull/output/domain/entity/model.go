package entity

type model struct {
	*modelStruct
}

type models struct {
	modelsStruct
}

type Model interface {
	modelInterface
}
type Models interface {
	modelsInterface
}

func NewModel() Model {
	return newModel()
}

func NewModels() Models {
	return &models{}
}
