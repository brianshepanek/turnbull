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
func (r *scribbleFooScaffoldRepository) Add(ctx context.Context, req model.FooScaffoldInterface) (resp *model.FooScaffoldInterface, err error) {
	err = r.driver.Write(r.collection, req.Primary(), req)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
