package repository

import (
	"context"
	model "github.com/brianshepanek/turnbull/testing/output/scaffold/domain/model"
	golangscribble "github.com/nanobox-io/golang-scribble"
)

type scribbleBarScaffoldRepository struct {
	driver     *golangscribble.Driver
	collection string
}

func NewScribbleBarScaffoldRepository(driver *golangscribble.Driver, collection string) *scribbleBarScaffoldRepository {
	return &scribbleBarScaffoldRepository{
		collection: collection,
		driver:     driver,
	}
}
func (r *scribbleBarScaffoldRepository) Add(ctx context.Context, req model.BarScaffoldInterface) (resp *model.BarScaffoldInterface, err error) {
	err = r.driver.Write(r.collection, req.Primary(), req)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *scribbleBarScaffoldRepository) FindAll(ctx context.Context, req interface{}) (resp *[]model.BarScaffoldInterface, err error) {
	return nil, nil
}

func (r *scribbleBarScaffoldRepository) FindOne(ctx context.Context, req interface{}) (resp *model.BarScaffoldInterface, err error) {
	return nil, nil
}

func (r *scribbleBarScaffoldRepository) CountEmUp(ctx context.Context, req interface{}) (resp *int, err error) {
	return nil, nil
}
