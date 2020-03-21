package presenter

import (
	"context"
	entity "github.com/brianshepanek/turnbull/turnbull/output/domain/entity"
)

type defaultFooPresenterStruct struct{}

func (r *defaultFooPresenterStruct) Browse(ctx context.Context, req entity.Foos) (entity.Foos, error) {
	return req, nil
}
func (r *defaultFooPresenterStruct) Read(ctx context.Context, req entity.Foo) (entity.Foo, error) {
	return req, nil
}
func (r *defaultFooPresenterStruct) Edit(ctx context.Context, req entity.Foo) (entity.Foo, error) {
	return req, nil
}
func (r *defaultFooPresenterStruct) Add(ctx context.Context, req entity.Foo) (entity.Foo, error) {
	return req, nil
}
func (r *defaultFooPresenterStruct) Delete(ctx context.Context, req entity.Foo) (entity.Foo, error) {
	return req, nil
}
