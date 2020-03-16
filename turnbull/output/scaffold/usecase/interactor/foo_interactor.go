package interactor

import (
	"context"
	entity "github.com/brianshepanek/turnbull/turnbull/output/scaffold/domain/entity"
	presenter "github.com/brianshepanek/turnbull/turnbull/output/scaffold/usecase/presenter"
	repository "github.com/brianshepanek/turnbull/turnbull/output/scaffold/usecase/repository"
)

type fooScaffoldInteractor struct {
	repository repository.FooScaffoldRepository
	presenter  presenter.FooScaffoldPresenter
}
type FooScaffoldInteractor interface {
	Count(ctx context.Context, query interface{}, req *int) (*int, error)
	Browse(ctx context.Context, query interface{}, req entity.FoosScaffoldInterface) (*entity.FoosScaffoldInterface, error)
	Read(ctx context.Context, query interface{}, req entity.FooScaffoldInterface) (*entity.FooScaffoldInterface, error)
	Edit(ctx context.Context, req entity.FooScaffoldInterface) (*entity.FooScaffoldInterface, error)
	Add(ctx context.Context, req entity.FooScaffoldInterface) (*entity.FooScaffoldInterface, error)
	Delete(ctx context.Context, req entity.FooScaffoldInterface) (*entity.FooScaffoldInterface, error)
}

func NewFooScaffoldInteractor(r repository.FooScaffoldRepository, p presenter.FooScaffoldPresenter) FooScaffoldInteractor {
	return &fooScaffoldInteractor{r, p}
}
func (i *fooScaffoldInteractor) Count(ctx context.Context, query interface{}, req *int) (*int, error) {
	err := i.repository.Count(ctx, query, req)
	if err != nil {
		return nil, err
	}
	return i.presenter.Count(ctx, req)
}
func (i *fooScaffoldInteractor) Browse(ctx context.Context, query interface{}, req entity.FoosScaffoldInterface) (*entity.FoosScaffoldInterface, error) {
	err := i.repository.Browse(ctx, query, req)
	if err != nil {
		return nil, err
	}
	return i.presenter.Browse(ctx, req)
}
func (i *fooScaffoldInteractor) Read(ctx context.Context, query interface{}, req entity.FooScaffoldInterface) (*entity.FooScaffoldInterface, error) {
	err := i.repository.Read(ctx, query, req)
	if err != nil {
		return nil, err
	}
	return i.presenter.Read(ctx, req)
}
func (i *fooScaffoldInteractor) Edit(ctx context.Context, req entity.FooScaffoldInterface) (*entity.FooScaffoldInterface, error) {
	err := i.repository.Edit(ctx, req)
	if err != nil {
		return nil, err
	}
	return i.presenter.Edit(ctx, req)
}
func (i *fooScaffoldInteractor) Add(ctx context.Context, req entity.FooScaffoldInterface) (*entity.FooScaffoldInterface, error) {
	err := i.repository.Add(ctx, req)
	if err != nil {
		return nil, err
	}
	return i.presenter.Add(ctx, req)
}
func (i *fooScaffoldInteractor) Delete(ctx context.Context, req entity.FooScaffoldInterface) (*entity.FooScaffoldInterface, error) {
	err := i.repository.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return i.presenter.Delete(ctx, req)
}
