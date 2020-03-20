package repository

import golangscribble "github.com/nanobox-io/golang-scribble"

type scribbleFooRepository struct {
	scribbleFooRepositoryStruct
}

func NewScribbleFooRepository(driver *golangscribble.Driver, collection string) *scribbleFooRepository {
	return &scribbleFooRepository{scribbleFooRepositoryStruct{
		collection: collection,
		driver:     driver,
	}}
}
