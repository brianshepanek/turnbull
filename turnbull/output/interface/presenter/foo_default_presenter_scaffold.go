package presenter

import (
	"context"
	entity "github.com/brianshepanek/turnbull/turnbull/output/domain/entity"
)

type defaultFooScaffoldPresenter struct{}

func NewDefaultFooScaffoldPresenter() *defaultFooScaffoldPresenter {
	return &defaultFooScaffoldPresenter{}
}
func (r *defaultFooScaffoldPresenter) Browse(ctx context.Context, req entity.Foos) (entity.Foos, error) {
	return req, nil
}
func (r *defaultFooScaffoldPresenter) Read(ctx context.Context, req entity.Foo) (entity.Foo, error) {
	return req, nil
}
func (r *defaultFooScaffoldPresenter) Edit(ctx context.Context, req entity.Foo) (entity.Foo, error) {
	return req, nil
}
func (r *defaultFooScaffoldPresenter) Add(ctx context.Context, req entity.Foo) (entity.Foo, error) {
	return req, nil
}
func (r *defaultFooScaffoldPresenter) Delete(ctx context.Context, req entity.Foo) (entity.Foo, error) {
	return req, nil
}
