package repository

import (
	"context"
	"encoding/json"
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
func (r *scribbleBarScaffoldRepository) Add(ctx context.Context, req model.BarScaffoldInterface) (err error) {
	return r.driver.Write(r.collection, req.Primary(), req)
}

func (r *scribbleBarScaffoldRepository) FindAll(ctx context.Context, req interface{}, resp model.BarsScaffoldInterface) (err error) {
	records, err := r.driver.ReadAll(r.collection)
	if err != nil {
		return err
	}
	for _, record := range records {
		rec := model.NewBarScaffoldStruct()
		err := json.Unmarshal([]byte(record), rec)
		if err != nil {
			return err
		}
		resp.Append(rec)
	}
	return nil
}

func (r *scribbleBarScaffoldRepository) FindOne(ctx context.Context, req model.BarScaffoldInterface) (err error) {
	return r.driver.Read(r.collection, req.Primary(), req)
}

func (r *scribbleBarScaffoldRepository) CountEmUp(ctx context.Context, req interface{}) (resp *int, err error) {
	return nil, nil
}

func (r *scribbleBarScaffoldRepository) ByeBye(ctx context.Context, req model.BarScaffoldInterface) (err error) {
	return r.driver.Delete(r.collection, req.Primary())
}
