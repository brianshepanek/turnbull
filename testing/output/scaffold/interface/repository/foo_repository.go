package repository

import (
	"context"
	model "github.com/brianshepanek/turnbull/testing/output/scaffold/domain/model"
	golangscribble "github.com/nanobox-io/golang-scribble"
)

type scribbleFooScaffoldRepository struct {
	driver     *golangscribble.Driver
	collection string
}

func NewScribbleFooScaffoldRepository(driver *golangscribble.Driver, collection string) *scribbleFooScaffoldRepository {
	return &scribbleFooScaffoldRepository{
		collection: collection,
		driver:     driver,
	}
}
func (r *scribbleFooScaffoldRepository) Add(ctx context.Context, req model.FooScaffoldInterface) (err error) {
	return r.driver.Write(r.collection, req.Primary(), req)
}
