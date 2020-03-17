package repository

import (
	"context"
	"encoding/json"
	entity "github.com/brianshepanek/turnbull/turnbull/output/scaffold/domain/entity"
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
func (r *scribbleFooScaffoldRepository) Browse(ctx context.Context, query interface{}, req entity.FoosScaffoldInterface) error {
	records, err := r.driver.ReadAll(r.collection)
	if err != nil {
		return err
	}
	for _, record := range records {
		resp := entity.NewFooScaffoldStruct()
		err := json.Unmarshal([]byte(record), resp)
		if err != nil {
			return err
		}
		req.Append(resp)
	}
	return nil
}
func (r *scribbleFooScaffoldRepository) Read(ctx context.Context, query interface{}, req entity.FooScaffoldInterface) error {
	return r.driver.Read(r.collection, req.Id(), req)
}
func (r *scribbleFooScaffoldRepository) Edit(ctx context.Context, req entity.FooScaffoldInterface) error {
	return r.driver.Read(r.collection, req.Id(), req)
}
func (r *scribbleFooScaffoldRepository) Add(ctx context.Context, req entity.FooScaffoldInterface) error {
	return r.driver.Write(r.collection, req.Id(), req)
}
func (r *scribbleFooScaffoldRepository) Delete(ctx context.Context, req entity.FooScaffoldInterface) error {
	return r.driver.Delete(r.collection, req.Id())
}
