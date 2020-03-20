package repository

import (
	"context"
	"encoding/json"
	entity "github.com/brianshepanek/turnbull/turnbull/output/domain/entity"
	golangscribble "github.com/nanobox-io/golang-scribble"
)

type scribbleFooRepositoryStruct struct {
	driver     *golangscribble.Driver
	collection string
}

func (r *scribbleFooRepositoryStruct) Browse(ctx context.Context, query interface{}, req entity.Foos) error {
	records, err := r.driver.ReadAll(r.collection)
	if err != nil {
		return err
	}
	for _, record := range records {
		resp := entity.NewFoo()
		err := json.Unmarshal([]byte(record), resp)
		if err != nil {
			return err
		}
		req.Append(resp)
	}
	return nil
}
func (r *scribbleFooRepositoryStruct) Read(ctx context.Context, query interface{}, req entity.Foo) error {
	return r.driver.Read(r.collection, req.Id(), req)
}
func (r *scribbleFooRepositoryStruct) Edit(ctx context.Context, req entity.Foo) error {
	return r.driver.Read(r.collection, req.Id(), req)
}
func (r *scribbleFooRepositoryStruct) Add(ctx context.Context, req entity.Foo) error {
	return r.driver.Write(r.collection, req.Id(), req)
}
func (r *scribbleFooRepositoryStruct) Delete(ctx context.Context, req entity.Foo) error {
	return r.driver.Delete(r.collection, req.Id())
}
